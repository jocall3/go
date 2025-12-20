```go
// Package platform contains shared dependency injection providers for foundational
// application components like logging, database connections, and configuration.
// These providers are used by Wire (in cmd/server/wire.go) to construct the
// application's dependency graph, ensuring a consistent and robust startup process.
package platform

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	// In a real project, this would be a dedicated package.
	// "github.com/your-org/your-project/internal/config"
)

// Config is a placeholder for the application's configuration structure.
// In a real-world application, this would be defined in its own `internal/config` package
// and loaded from a file (e.g., YAML) and/or environment variables.
type Config struct {
	Log      LogConfig
	Database DatabaseConfig
}

// LogConfig holds configuration for the logger.
type LogConfig struct {
	Level string `mapstructure:"level"` // e.g., "debug", "info", "warn", "error"
}

// DatabaseConfig holds configuration for the database connection pool.
type DatabaseConfig struct {
	URL             string        `mapstructure:"url"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

// ProvideLogger creates a new structured logger (slog) based on the application configuration.
// It sets the logging level and format. This is a foundational component for observability,
// auditability, and making the system "boring to regulators".
func ProvideLogger(cfg *Config) (*slog.Logger, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(cfg.Log.Level)); err != nil {
		slog.Warn("invalid log level in config, defaulting to INFO", "level", cfg.Log.Level)
		level = slog.LevelInfo
	}

	// Using a JSON handler for structured, machine-readable logs.
	// This is crucial for auditability and automated analysis in production environments.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
		// AddSource adds file and line number, useful for debugging.
		AddSource: true,
	}))

	return logger, nil
}

// ProvideDatabasePool creates a new PostgreSQL connection pool using pgxpool.
// It uses the configuration to set up connection parameters and pool settings.
// It returns the pool and a cleanup function to be called on application shutdown.
// This ensures that database connections are properly closed, adhering to fail-closed semantics.
func ProvideDatabasePool(ctx context.Context, cfg *Config, logger *slog.Logger) (*pgxpool.Pool, func(), error) {
	dbConfig, err := pgxpool.ParseConfig(cfg.Database.URL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Configure the connection pool settings. These are critical for performance and stability.
	dbConfig.MaxConns = int32(cfg.Database.MaxOpenConns)
	dbConfig.MinConns = 2 // Ensure a minimum number of connections are available.
	dbConfig.MaxConnLifetime = cfg.Database.ConnMaxLifetime
	dbConfig.MaxConnIdleTime = cfg.Database.ConnMaxIdleTime

	// Health check pings the database to ensure it's available before returning the pool.
	// This prevents the application from starting in a broken state.
	dbConfig.HealthCheckPeriod = 1 * time.Minute

	logger.Info("connecting to database...")

	pool, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create database connection pool: %w", err)
	}

	// Ping the database to verify the connection. This is a critical startup check.
	// It enforces the "halt safely on uncertainty" principle.
	if err := pool.Ping(ctx); err != nil {
		pool.Close() // Close the pool if ping fails
		return nil, nil, fmt.Errorf("failed to ping database on startup: %w", err)
	}

	logger.Info("database connection pool established")

	cleanup := func() {
		logger.Info("closing database connection pool")
		pool.Close()
	}

	return pool, cleanup, nil
}

// ShutdownContext provides a context that is canceled when a shutdown signal (SIGINT, SIGTERM)
// is received. This is the primary mechanism for triggering a graceful shutdown of the application.
// It allows long-running tasks to finish their work before the application exits.
func ShutdownContext(logger *slog.Logger) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		shutdownSignals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, shutdownSignals...)

		sig := <-sigChan
		logger.Info("received shutdown signal, initiating graceful shutdown", "signal", sig.String())
		cancel()

		// If another signal is received, force exit.
		// This prevents the application from hanging during a problematic shutdown.
		<-sigChan
		logger.Warn("received second shutdown signal, forcing exit")
		os.Exit(1)
	}()

	return ctx
}

// DBTX is an interface that abstracts database operations.
// It is satisfied by both *pgxpool.Pool and pgx.Tx, allowing repository methods
// to be executed within or outside of a transaction without changing the code.
// This is crucial for maintaining deterministic state transitions and auditability.
type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

// PlatformSet is a Wire provider set that includes all the foundational components
// of the application platform, such as logging, database connections, and shutdown handling.
// Grouping these providers makes it easy to include them in the main application injector.
var PlatformSet = wire.NewSet(
	ProvideLogger,
	ProvideDatabasePool,
	ShutdownContext,
	// Note: We do not include a provider for Config here.
	// Configuration is typically loaded first in the `main` function and then
	// passed as a parameter to the Wire-generated injector. This ensures that
	// the application fails fast if configuration is missing or invalid,
	// before any other components are initialized.
)
### END_OF_FILE_COMPLETED ###
```