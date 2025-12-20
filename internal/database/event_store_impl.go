```go
package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.comcom/your-project-name/internal/domain/events" // Assuming events package is in internal/domain
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	// pgUniqueViolation is the error code for a unique constraint violation in PostgreSQL.
	// This is critical for detecting concurrency conflicts at the database level.
	pgUniqueViolation = "23505"
)

// SQL queries for the event store.
// Using constants for queries improves readability and maintainability.
const (
	insertEventSQL = `
		INSERT INTO events (event_id, aggregate_id, aggregate_type, event_type, payload, version, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	selectLatestVersionSQL = `
		SELECT COALESCE(MAX(version), 0)
		FROM events
		WHERE aggregate_id = $1;
	`

	selectEventsByAggregateIDSQL = `
		SELECT event_id, aggregate_id, aggregate_type, event_type, payload, version, created_at
		FROM events
		WHERE aggregate_id = $1
		ORDER BY version ASC;
	`
)

// PostgresEventStore is a PostgreSQL-based implementation of the events.Store interface.
// It provides an append-only, atomic, and concurrency-safe way to store and retrieve domain events.
type PostgresEventStore struct {
	db *sql.DB
}

// NewPostgresEventStore creates a new instance of PostgresEventStore.
// It requires a connected *sql.DB instance.
func NewPostgresEventStore(db *sql.DB) *PostgresEventStore {
	if db == nil {
		panic("database connection cannot be nil")
	}
	return &PostgresEventStore{
		db: db,
	}
}

// Save appends a slice of events to the store for a given aggregate.
// It enforces optimistic concurrency control by checking the expectedVersion against
// the aggregate's current version in the database. The entire operation is performed
// within a single atomic transaction.
// If the expected version does not match the stored version, it returns events.ErrConcurrency.
func (s *PostgresEventStore) Save(
	ctx context.Context,
	aggregateID string,
	aggregateType events.AggregateType,
	expectedVersion int,
	newEvents []events.Event,
) error {
	if len(newEvents) == 0 {
		return nil // Nothing to save
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Rollback is a no-op if the transaction is committed.

	// 1. Get current version and enforce optimistic concurrency lock.
	currentVersion, err := s.getCurrentVersion(ctx, tx, aggregateID)
	if err != nil {
		return fmt.Errorf("failed to get current version for aggregate %s: %w", aggregateID, err)
	}

	if currentVersion != expectedVersion {
		return fmt.Errorf("concurrency error for aggregate %s: expected version %d, but got %d: %w",
			aggregateID, expectedVersion, currentVersion, events.ErrConcurrency)
	}

	// 2. Prepare the insert statement once.
	stmt, err := tx.PrepareContext(ctx, insertEventSQL)
	if err != nil {
		return fmt.Errorf("failed to prepare event insert statement: %w", err)
	}
	defer stmt.Close()

	// 3. Insert all new events with incrementing versions.
	nextVersion := currentVersion + 1
	for _, event := range newEvents {
		// Ensure event metadata is consistent with the save request.
		event.AggregateID = aggregateID
		event.AggregateType = aggregateType
		event.Version = nextVersion
		event.CreatedAt = time.Now().UTC()
		if event.ID == uuid.Nil {
			event.ID = uuid.New()
		}

		_, err := stmt.ExecContext(
			ctx,
			event.ID,
			event.AggregateID,
			event.AggregateType,
			event.EventType,
			event.Payload,
			event.Version,
			event.CreatedAt,
		)

		if err != nil {
			// Check if the error is a unique constraint violation on (aggregate_id, version).
			// This is the database-level safeguard against race conditions.
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation {
				return fmt.Errorf("database concurrency conflict for aggregate %s at version %d: %w",
					aggregateID, event.Version, events.ErrConcurrency)
			}
			return fmt.Errorf("failed to execute event insert for aggregate %s: %w", aggregateID, err)
		}
		nextVersion++
	}

	// 4. Commit the transaction.
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit event store transaction for aggregate %s: %w", aggregateID, err)
	}

	return nil
}

// Load retrieves all events for a given aggregate ID, ordered by version.
// If no events are found for the aggregate, it returns an empty slice and no error.
func (s *PostgresEventStore) Load(ctx context.Context, aggregateID string) ([]events.Event, error) {
	rows, err := s.db.QueryContext(ctx, selectEventsByAggregateIDSQL, aggregateID)
	if err != nil {
		return nil, fmt.Errorf("failed to query events for aggregate %s: %w", aggregateID, err)
	}
	defer rows.Close()

	var loadedEvents []events.Event
	for rows.Next() {
		var event events.Event
		if err := rows.Scan(
			&event.ID,
			&event.AggregateID,
			&event.AggregateType,
			&event.EventType,
			&event.Payload,
			&event.Version,
			&event.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan event row for aggregate %s: %w", aggregateID, err)
		}
		loadedEvents = append(loadedEvents, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating event rows for aggregate %s: %w", aggregateID, err)
	}

	// It's not an error to find no events; it simply means the aggregate is new.
	// The domain layer is responsible for interpreting this.
	return loadedEvents, nil
}

// getCurrentVersion fetches the latest version number for an aggregate within a transaction.
// It uses the transaction to ensure a consistent read (read-committed isolation level is sufficient).
// If no events exist for the aggregate, it returns 0.
func (s *PostgresEventStore) getCurrentVersion(ctx context.Context, tx *sql.Tx, aggregateID string) (int, error) {
	var version int
	err := tx.QueryRowContext(ctx, selectLatestVersionSQL, aggregateID).Scan(&version)
	if err != nil {
		// It's possible sql.ErrNoRows could be returned if the table is empty,
		// but COALESCE should prevent this. We handle it just in case.
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("failed to scan current version: %w", err)
	}
	return version, nil
}

```