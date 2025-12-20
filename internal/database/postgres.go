```go
package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Config holds all the configuration parameters for the PostgreSQL database connection.
// This explicit structure ensures that all necessary parameters are provided and
// makes the connection setup deterministic and easy to audit.
type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxConns        int32
	MinConns        int32
	MaxConnIdleTime time.Duration
	MaxConnLifetime time.Duration
	ConnectTimeout  time.Duration
}

// New creates a new PostgreSQL connection pool with robust retry logic.
// It ensures the application starts only with a valid and verified database connection.
// This function embodies the "fail-closed" principle: if a connection cannot be
// established after a set number of retries, the function will return an error,
// preventing the application from starting in an unstable or non-operational state.
// This is critical in a financial system where operating without a database could
// lead to data loss, inconsistent state, or capital safety violations.
func New(ctx context.Context, cfg Config, logger *slog.Logger) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres config: %w", err)
	}

	// Apply connection pool settings from the configuration. These parameters are
	// crucial for managing database resources efficiently and preventing connection
	// exhaustion under load.
	poolConfig.MaxConns = cfg.MaxConns
	poolConfig.MinConns = cfg.MinConns
	poolConfig.MaxConnIdleTime = cfg.MaxConnIdleTime
	poolConfig.MaxConnLifetime = cfg.MaxConnLifetime
	poolConfig.ConnConfig.ConnectTimeout = cfg.ConnectTimeout

	var pool *pgxpool.Pool
	const maxRetries = 5
	const retryInterval = 5 * time.Second

	logger.Info("Attempting to connect to PostgreSQL database...", "host", cfg.Host, "db", cfg.DBName)

	for i := 0; i < maxRetries; i++ {
		// Use a context with a timeout for each connection attempt to avoid indefinite hangs.
		connectCtx, cancel := context.WithTimeout(ctx, cfg.ConnectTimeout)

		pool, err = pgxpool.NewWithConfig(connectCtx, poolConfig)
		cancel() // Release context resources immediately after the attempt.

		if err == nil {
			// Connection successful, now verify with a ping to ensure the database is responsive.
			pingCtx, pingCancel := context.WithTimeout(ctx, 5*time.Second)
			pingErr := pool.Ping(pingCtx)
			pingCancel()

			if pingErr == nil {
				logger.Info("Successfully connected and verified PostgreSQL database connection", "host", cfg.Host, "db", cfg.DBName)
				return pool, nil
			}
			// Ping failed, close the newly created pool and retry. This handles cases
			// where a connection is accepted but the DB is not yet ready to serve queries.
			logger.Warn("Database connection established, but ping failed", "error", pingErr, "attempt", i+1)
			pool.Close()
			err = pingErr // Set err to the ping error for logging below.
		}

		if ctx.Err() != nil {
			return nil, fmt.Errorf("database connection cancelled by parent context: %w", ctx.Err())
		}

		logger.Warn(
			"Failed to connect to database, retrying...",
			"attempt", i+1,
			"max_retries", maxRetries,
			"retry_in", retryInterval.String(),
			"error", err,
		)

		select {
		case <-time.After(retryInterval):
			// Wait before the next retry.
		case <-ctx.Done():
			return nil, fmt.Errorf("database connection context cancelled during retry wait: %w", ctx.Err())
		}
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}

// Close gracefully closes the database connection pool.
// It's a simple wrapper around the pool's Close method for API consistency.
// The application should call this on shutdown to ensure all database connections
// are terminated cleanly.
func Close(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
	}
}
### END_OF_FILE_COMPLETED ###
```