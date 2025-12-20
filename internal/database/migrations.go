```go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"time"
)

// Migration defines the structure for a single database migration.
// The Up function applies the migration, and the Down function reverts it.
// Both operations are expected to be transactional.
type Migration struct {
	Version     int
	Description string
	Up          func(tx *sql.Tx) error
	Down        func(tx *sql.Tx) error
}

// Migrations is a slice of Migration, which can be sorted by version.
type Migrations []Migration

func (m Migrations) Len() int           { return len(m) }
func (m Migrations) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Migrations) Less(i, j int) bool { return m[i].Version < m[j].Version }

// allMigrations holds all the defined database migrations for the application.
// New migrations should be added to this slice in ascending order of Version.
var allMigrations = Migrations{
	{
		Version:     1,
		Description: "Create initial schema: ledgers, accounts, transfers, entries tables",
		Up: func(tx *sql.Tx) error {
			// Using UUIDs for primary keys is good practice for distributed systems
			// and prevents information leakage from sequential IDs.
			// Using NUMERIC type for financial values is crucial to avoid floating-point errors.
			// Timestamps are stored with time zone information.

			// ledgers table: A ledger is a high-level container for a set of accounts,
			// often representing a distinct business domain or currency type.
			// e.g., "customer_assets", "internal_revenue", "operational_expenses"
			_, err := tx.Exec(`
				CREATE TABLE IF NOT EXISTS ledgers (
					id UUID PRIMARY KEY,
					name VARCHAR(255) NOT NULL UNIQUE,
					created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
				);
			`)
			if err != nil {
				return fmt.Errorf("failed to create ledgers table: %w", err)
			}

			// accounts table: Represents a store of value for a specific currency within a ledger.
			// The balance is a critical invariant that must be maintained by the system's logic.
			_, err = tx.Exec(`
				CREATE TABLE IF NOT EXISTS accounts (
					id UUID PRIMARY KEY,
					ledger_id UUID NOT NULL REFERENCES ledgers(id),
					currency VARCHAR(10) NOT NULL,
					balance NUMERIC(36, 18) NOT NULL DEFAULT 0,
					status VARCHAR(50) NOT NULL DEFAULT 'active', -- e.g., active, frozen, closed
					created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					UNIQUE (ledger_id, currency) -- An account is typically unique per ledger and currency
				);
				CREATE INDEX IF NOT EXISTS idx_accounts_ledger_id ON accounts(ledger_id);
			`)
			if err != nil {
				return fmt.Errorf("failed to create accounts table: %w", err)
			}

			// transfers table: Represents an intent to move value between accounts.
			// A transfer results in one or more 'entries' that must balance.
			_, err = tx.Exec(`
				CREATE TABLE IF NOT EXISTS transfers (
					id UUID PRIMARY KEY,
					transaction_id UUID NOT NULL, -- Groups multiple transfers in a single atomic financial transaction
					status VARCHAR(50) NOT NULL, -- e.g., pending, posted, failed
					metadata JSONB,
					created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
				);
				CREATE INDEX IF NOT EXISTS idx_transfers_transaction_id ON transfers(transaction_id);
			`)
			if err != nil {
				return fmt.Errorf("failed to create transfers table: %w", err)
			}

			// entries table: The immutable, append-only log of all value movements.
			// This is the core of a double-entry accounting system.
			// A credit to one account must be balanced by a debit from another within a transfer.
			_, err = tx.Exec(`
				CREATE TABLE IF NOT EXISTS entries (
					id BIGSERIAL PRIMARY KEY,
					account_id UUID NOT NULL REFERENCES accounts(id),
					transfer_id UUID NOT NULL REFERENCES transfers(id),
					amount NUMERIC(36, 18) NOT NULL, -- Positive for credit, negative for debit
					created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
				);
				CREATE INDEX IF NOT EXISTS idx_entries_account_id ON entries(account_id);
				CREATE INDEX IF NOT EXISTS idx_entries_transfer_id ON entries(transfer_id);
			`)
			if err != nil {
				return fmt.Errorf("failed to create entries table: %w", err)
			}

			return nil
		},
		Down: func(tx *sql.Tx) error {
			// Drop tables in reverse order of creation to respect foreign key constraints.
			_, err := tx.Exec(`DROP TABLE IF EXISTS entries;`)
			if err != nil {
				return err
			}
			_, err = tx.Exec(`DROP TABLE IF EXISTS transfers;`)
			if err != nil {
				return err
			}
			_, err = tx.Exec(`DROP TABLE IF EXISTS accounts;`)
			if err != nil {
				return err
			}
			_, err = tx.Exec(`DROP TABLE IF EXISTS ledgers;`)
			if err != nil {
				return err
			}
			return nil
		},
	},
	// --- ADD NEW MIGRATIONS HERE ---
	// {
	// 	Version:     2,
	// 	Description: "Add users table and link to accounts",
	// 	Up: func(tx *sql.Tx) error { ... },
	// 	Down: func(tx *sql.Tx) error { ... },
	// },
}

// ensureMigrationsTable creates the schema_migrations table if it does not exist.
// This table is used to track which migrations have been applied.
func ensureMigrationsTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}
	return nil
}

// getAppliedMigrations retrieves the versions of all migrations that have already been applied.
func getAppliedMigrations(db *sql.DB) (map[int]bool, error) {
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to query schema_migrations: %w", err)
	}
	defer rows.Close()

	applied := make(map[int]bool)
	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return nil, fmt.Errorf("failed to scan migration version: %w", err)
		}
		applied[version] = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over applied migrations: %w", err)
	}

	return applied, nil
}

// ApplyMigrations runs all pending migrations on the database.
// It ensures that migrations are applied in order, transactionally, and exactly once.
// This function is designed to be idempotent and safe to run on application startup.
func ApplyMigrations(db *sql.DB) error {
	if err := ensureMigrationsTable(db); err != nil {
		return err
	}

	applied, err := getAppliedMigrations(db)
	if err != nil {
		return err
	}

	// Sort migrations to ensure they are always applied in the correct order.
	sort.Sort(allMigrations)

	log.Println("Starting database migrations...")

	for _, m := range allMigrations {
		if applied[m.Version] {
			continue
		}

		log.Printf("Applying migration version %d: '%s'...", m.Version, m.Description)

		// Each migration is run in its own transaction.
		// For databases that support DDL in transactions (like PostgreSQL),
		// this ensures that a failed migration doesn't leave the schema in a partially-applied state.
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("failed to begin transaction for migration %d: %w", m.Version, err)
		}

		// Defer a rollback. If the transaction is committed, this is a no-op.
		// If it fails for any reason, the transaction is rolled back. This is a fail-closed semantic.
		defer tx.Rollback()

		if err := m.Up(tx); err != nil {
			return fmt.Errorf("migration version %d failed during 'Up' execution: %w", m.Version, err)
		}

		// Record the migration version in the schema_migrations table within the same transaction.
		_, err = tx.Exec("INSERT INTO schema_migrations (version, applied_at) VALUES ($1, $2)", m.Version, time.Now().UTC())
		if err != nil {
			return fmt.Errorf("failed to record migration version %d: %w", m.Version, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction for migration %d: %w", m.Version, err)
		}

		log.Printf("Successfully applied migration version %d.", m.Version)
	}

	log.Println("Database migrations finished successfully.")
	return nil
}

```