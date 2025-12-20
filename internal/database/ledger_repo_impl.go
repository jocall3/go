```go
package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/your-org/your-project/internal/ledger"
)

// Assumed Database Schema for LedgerRepositoryImpl
//
// CREATE TABLE ledger_snapshots (
//     id UUID PRIMARY KEY,
//     ledger_id UUID NOT NULL,
//     version BIGINT NOT NULL,
//     state_data JSONB NOT NULL,
//     created_at TIMESTAMPTZ NOT NULL,
//     CONSTRAINT uq_ledger_snapshot_version UNIQUE (ledger_id, version)
// );
//
// CREATE INDEX idx_ledger_snapshots_lookup ON ledger_snapshots (ledger_id, version DESC);
//
// CREATE TABLE ledger_transactions (
//     id UUID PRIMARY KEY,
//     ledger_id UUID NOT NULL,
//     version BIGINT NOT NULL,
//     transaction_data JSONB NOT NULL,
//     created_at TIMESTAMPTZ NOT NULL,
//     CONSTRAINT uq_ledger_tx_version UNIQUE (ledger_id, version)
// );
//
// CREATE INDEX idx_ledger_transactions_lookup ON ledger_transactions (ledger_id, version ASC);

// Compile-time check to ensure LedgerRepositoryImpl implements ledger.Repository.
var _ ledger.Repository = (*LedgerRepositoryImpl)(nil)

const (
	// snapshotFrequency determines how often a new snapshot is created (e.g., every 1000 versions).
	// A lower number means faster reads but more storage and slower writes.
	// A higher number means slower reads but less storage and faster writes.
	// This value is a trade-off between write performance/storage and read (reconstruction) performance.
	snapshotFrequency = 1000
)

// LedgerRepositoryImpl provides a concrete implementation of the ledger.Repository
// interface using a PostgreSQL database. It is optimized for fast state reconstruction
// by using a combination of periodic state snapshots and a transaction log.
type LedgerRepositoryImpl struct {
	pool *pgxpool.Pool
}

// NewLedgerRepository creates a new instance of LedgerRepositoryImpl.
// It requires a configured pgxpool.Pool for database connections.
func NewLedgerRepository(pool *pgxpool.Pool) *LedgerRepositoryImpl {
	return &LedgerRepositoryImpl{
		pool: pool,
	}
}

// SaveState atomically saves a new ledger state. This involves persisting the
// transactions that led to this state and, if a certain version threshold is met,
// creating a new state snapshot. This entire operation is performed within a
// single database transaction to ensure consistency and auditability.
func (r *LedgerRepositoryImpl) SaveState(ctx context.Context, state *ledger.State, transactions []ledger.Transaction) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin database transaction: %w", err)
	}
	defer tx.Rollback(ctx) // Rollback is a no-op if the transaction is committed.

	// 1. Save all transactions that compose this state change.
	if err := r.saveTransactionsInTx(ctx, tx, transactions); err != nil {
		return err // Error is already wrapped
	}

	// 2. Decide whether to save a new snapshot.
	// We save a snapshot at version 0 (genesis) and then periodically
	// based on the configured frequency to optimize reconstruction time.
	if state.Version == 0 || state.Version%snapshotFrequency == 0 {
		if err := r.saveSnapshotInTx(ctx, tx, state); err != nil {
			return err // Error is already wrapped
		}
	}

	// 3. Commit the transaction.
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit ledger state transaction: %w", err)
	}

	return nil
}

// GetLatestState reconstructs and returns the latest state of a given ledger.
// It finds the most recent snapshot and replays all subsequent transactions
// to build the current state, ensuring a consistent and up-to-date view.
func (r *LedgerRepositoryImpl) GetLatestState(ctx context.Context, ledgerID uuid.UUID) (*ledger.State, error) {
	var maxVersion uint64
	// First, determine the latest known version for the ledger from the transaction log.
	// COALESCE handles the case where a ledger has a genesis snapshot but no transactions yet.
	err := r.pool.QueryRow(ctx, "SELECT COALESCE(MAX(version), 0) FROM ledger_transactions WHERE ledger_id = $1", ledgerID).Scan(&maxVersion)
	if err != nil {
		// If no rows are found, it might mean the ledger doesn't exist.
		// We can check for a genesis snapshot at version 0 as a fallback.
		if errors.Is(err, pgx.ErrNoRows) {
			return r.GetStateAt(ctx, ledgerID, 0)
		}
		return nil, fmt.Errorf("failed to query max version for ledger %s: %w", ledgerID, err)
	}

	return r.GetStateAt(ctx, ledgerID, maxVersion)
}

// GetStateAt reconstructs and returns the state of a ledger at a specific version.
// This method is central to providing full auditability and time-travel queries.
// It implements the core reconstruction logic:
// 1. Find the most recent snapshot at or before the target version.
// 2. Load that snapshot's state.
// 3. Fetch all transactions between the snapshot version and the target version.
// 4. Apply each transaction sequentially to the state.
func (r *LedgerRepositoryImpl) GetStateAt(ctx context.Context, ledgerID uuid.UUID, version uint64) (*ledger.State, error) {
	// 1. Find the most recent snapshot at or before the target version.
	snapshot, err := r.getLatestSnapshotBeforeOrAt(ctx, ledgerID, version)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// This case implies no state (not even a genesis snapshot) exists for this ledger
			// at or before the requested version.
			return nil, fmt.Errorf("no snapshot found for ledger %s at or before version %d: %w", ledgerID, version, ledger.ErrLedgerNotFound)
		}
		return nil, fmt.Errorf("failed to get latest snapshot: %w", err)
	}

	// 2. Deserialize the snapshot into a base state. This is our starting point.
	state, err := ledger.StateFromSnapshot(snapshot)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize snapshot %s: %w", snapshot.ID, err)
	}

	// If the requested version is exactly the snapshot version, we are done.
	if state.Version == version {
		return state, nil
	}

	// 3. Fetch transactions that occurred after the snapshot's version up to the target version.
	transactions, err := r.getTransactionsBetween(ctx, ledgerID, state.Version, version)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions after version %d: %w", state.Version, err)
	}

	// 4. Apply transactions to the state to reconstruct the final state.
	// The ApplyTransaction logic is part of the core ledger domain.
	for _, tx := range transactions {
		if err := state.ApplyTransaction(&tx); err != nil {
			return nil, fmt.Errorf("failed to apply transaction %s during state reconstruction: %w", tx.ID, err)
		}
	}

	// Final invariant check: ensure the reconstructed state matches the requested version.
	// This protects against data inconsistencies (e.g., missing transaction records).
	if state.Version != version {
		return nil, fmt.Errorf("state reconstruction failed for ledger %s: expected version %d, got %d", ledgerID, version, state.Version)
	}

	return state, nil
}

// saveTransactionsInTx saves a batch of transactions within a given database transaction.
func (r *LedgerRepositoryImpl) saveTransactionsInTx(ctx context.Context, tx pgx.Tx, transactions []ledger.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	// Use pgx's batching capabilities for efficient bulk inserts.
	batch := &pgx.Batch{}
	stmt := `
		INSERT INTO ledger_transactions (id, ledger_id, version, transaction_data, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	for _, t := range transactions {
		batch.Queue(stmt, t.ID, t.LedgerID, t.Version, t.TransactionData, t.CreatedAt)
	}

	br := tx.SendBatch(ctx, batch)
	defer br.Close()

	// Check the result of each queued operation to ensure all succeeded.
	for i := 0; i < len(transactions); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("failed to execute transaction insert in batch (index %d): %w", i, err)
		}
	}

	return br.Close()
}

// saveSnapshotInTx serializes and saves a ledger state snapshot within a given database transaction.
func (r *LedgerRepositoryImpl) saveSnapshotInTx(ctx context.Context, tx pgx.Tx, state *ledger.State) error {
	stateData, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to serialize state for snapshot (version %d): %w", state.Version, err)
	}

	snapshot := ledger.Snapshot{
		ID:        uuid.New(),
		LedgerID:  state.LedgerID,
		Version:   state.Version,
		StateData: stateData,
		CreatedAt: state.CreatedAt,
	}

	stmt := `
		INSERT INTO ledger_snapshots (id, ledger_id, version, state_data, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = tx.Exec(ctx, stmt, snapshot.ID, snapshot.LedgerID, snapshot.Version, snapshot.StateData, snapshot.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert ledger snapshot (version %d): %w", snapshot.Version, err)
	}

	return nil
}

// getLatestSnapshotBeforeOrAt retrieves the most recent snapshot for a ledger
// at or before a specific version. This is the starting point for state reconstruction.
func (r *LedgerRepositoryImpl) getLatestSnapshotBeforeOrAt(ctx context.Context, ledgerID uuid.UUID, version uint64) (*ledger.Snapshot, error) {
	snapshot := &ledger.Snapshot{}
	query := `
		SELECT id, ledger_id, version, state_data, created_at
		FROM ledger_snapshots
		WHERE ledger_id = $1 AND version <= $2
		ORDER BY version DESC
		LIMIT 1
	`
	err := r.pool.QueryRow(ctx, query, ledgerID, version).Scan(
		&snapshot.ID,
		&snapshot.LedgerID,
		&snapshot.Version,
		&snapshot.StateData,
		&snapshot.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, pgx.ErrNoRows // Propagate specific error for clear handling upstream
		}
		return nil, fmt.Errorf("db query for latest snapshot failed: %w", err)
	}
	return snapshot, nil
}

// getTransactionsBetween retrieves all transactions for a ledger between a starting
// version (exclusive) and an ending version (inclusive), ordered by version.
func (r *LedgerRepositoryImpl) getTransactionsBetween(ctx context.Context, ledgerID uuid.UUID, startVersion, endVersion uint64) ([]ledger.Transaction, error) {
	query := `
		SELECT id, ledger_id, version, transaction_data, created_at
		FROM ledger_transactions
		WHERE ledger_id = $1 AND version > $2 AND version <= $3
		ORDER BY version ASC
	`
	rows, err := r.pool.Query(ctx, query, ledgerID, startVersion, endVersion)
	if err != nil {
		return nil, fmt.Errorf("db query for transactions failed: %w", err)
	}
	defer rows.Close()

	var transactions []ledger.Transaction
	for rows.Next() {
		var t ledger.Transaction
		if err := rows.Scan(&t.ID, &t.LedgerID, &t.Version, &t.TransactionData, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan transaction row: %w", err)
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transaction rows: %w", err)
	}

	return transactions, nil
}

```