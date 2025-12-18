package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config holds the configuration for the PostgreSQL database connection.
type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	Name         string
	SSLMode      string
	Timezone     string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
	Debug        bool
}

// Open initializes a new database connection using GORM with the provided configuration.
// It sets up connection pooling and returns the GORM DB instance.
func Open(cfg Config) (*gorm.DB, error) {
	// Set default SSLMode if not provided
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}
	// Set default Timezone if not provided
	if cfg.Timezone == "" {
		cfg.Timezone = "UTC"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s timezone=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
		cfg.SSLMode,
		cfg.Timezone,
	)

	logLevel := logger.Silent
	if cfg.Debug {
		logLevel = logger.Info
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		// PrepareStmt creates a prepared statement when executing any SQL and caches them to speed up future calls
		PrepareStmt: true,
		// DisableForeignKeyConstraintWhenMigrating can be useful in certain CI/CD pipelines, keeping false for integrity
		DisableForeignKeyConstraintWhenMigrating: false,
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// StatusCheck returns nil if the database is responding, or an error if it is not.
func StatusCheck(ctx context.Context, db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB for status check: %w", err)
	}
	return sqlDB.PingContext(ctx)
}

// Close closes the underlying database connection.
// Note: GORM does not require explicit closing usually, but this is useful for graceful shutdowns.
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB for close: %w", err)
	}
	return sqlDB.Close()
}