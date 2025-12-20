```go
// Copyright (c) 2024 The Bridge
//
// This file is part of The Bridge.
// The Bridge is a production-grade financial infrastructure that bridges
// the gap between traditional banking and modern digital asset ecosystems.
//
// The Bridge is free software: you can redistribute it and/or modify
// it under the terms of the MIT License.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package platform

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config holds the application's configuration.
// It is populated from environment variables at startup.
// The "fail-closed" principle is enforced by making critical
// configuration parameters required, causing the application to exit
// if they are not provided. This ensures the system always starts in a
// known, valid state.
type Config struct {
	Log      LogConfig
	Server   ServerConfig
	Database DatabaseConfig
	Risk     RiskConfig
	// Add other configuration sections as the system grows, e.g.,
	// Settlement SettlementConfig
	// Governance GovernanceConfig
}

// LogConfig contains configuration for the logging system.
type LogConfig struct {
	Level  string `envconfig:"LOG_LEVEL" default:"info"`
	Format string `envconfig:"LOG_FORMAT" default:"json"` // "json" or "text"
}

// ServerConfig contains configuration for the HTTP server.
type ServerConfig struct {
	Host            string        `envconfig:"SERVER_HOST" default:"0.0.0.0"`
	Port            int           `envconfig:"SERVER_PORT" default:"8080"`
	ReadTimeout     time.Duration `envconfig:"SERVER_READ_TIMEOUT" default:"5s"`
	WriteTimeout    time.Duration `envconfig:"SERVER_WRITE_TIMEOUT" default:"10s"`
	IdleTimeout     time.Duration `envconfig:"SERVER_IDLE_TIMEOUT" default:"120s"`
	ShutdownTimeout time.Duration `envconfig:"SERVER_SHUTDOWN_TIMEOUT" default:"30s"`
}

// DatabaseConfig contains configuration for connecting to the primary database.
// We favor explicit connection parameters over a single DSN string to make
// configuration clearer and less error-prone.
type DatabaseConfig struct {
	Host            string        `envconfig:"DB_HOST" required:"true"`
	Port            int           `envconfig:"DB_PORT" required:"true"`
	User            string        `envconfig:"DB_USER" required:"true"`
	Password        string        `envconfig:"DB_PASSWORD" required:"true"`
	Name            string        `envconfig:"DB_NAME" required:"true"`
	SSLMode         string        `envconfig:"DB_SSL_MODE" default:"disable"` // e.g., "disable", "require", "verify-full"
	MaxOpenConns    int           `envconfig:"DB_MAX_OPEN_CONNS" default:"25"`
	MaxIdleConns    int           `envconfig:"DB_MAX_IDLE_CONNS" default:"25"`
	ConnMaxLifetime time.Duration `envconfig:"DB_CONN_MAX_LIFETIME" default:"5m"`
}

// RiskConfig contains parameters for the core risk management engine.
// These values are critical invariants for capital safety.
type RiskConfig struct {
	// MaxGlobalPositionSize defines the maximum total position size the system can hold.
	MaxGlobalPositionSize float64 `envconfig:"RISK_MAX_GLOBAL_POSITION_SIZE" default:"100000000.00"`
	// MaxCounterpartyExposure defines the maximum exposure to a single counterparty.
	MaxCounterpartyExposure float64 `envconfig:"RISK_MAX_COUNTERPARTY_EXPOSURE" default:"5000000.00"`
	// CircuitBreakerThreshold is the percentage loss (e.g., 0.05 for 5%) that triggers a system-wide halt.
	CircuitBreakerThreshold float64 `envconfig:"RISK_CIRCUIT_BREAKER_THRESHOLD" default:"0.05"`
}

// Load reads configuration from environment variables and validates it.
// It returns a fully populated and validated Config struct, or an error
// if loading or validation fails. This function is designed to be called
// once at application startup.
func Load() (*Config, error) {
	var cfg Config

	// The "FIN" prefix is used to namespace all environment variables for this application.
	// This avoids collisions with other variables in the execution environment.
	if err := envconfig.Process("fin", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process configuration from environment: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("configuration validation failed: %w", err)
	}

	return &cfg, nil
}

// validate performs semantic validation on the loaded configuration.
// This ensures that the system starts in a known, valid state.
func (c *Config) validate() error {
	// Validate LogConfig
	logLevel := c.Log.Level
	if logLevel != "debug" && logLevel != "info" && logLevel != "warn" && logLevel != "error" {
		return fmt.Errorf("invalid log level: %q, must be one of 'debug', 'info', 'warn', 'error'", logLevel)
	}
	logFormat := c.Log.Format
	if logFormat != "json" && logFormat != "text" {
		return fmt.Errorf("invalid log format: %q, must be one of 'json', 'text'", logFormat)
	}

	// Validate ServerConfig
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d, must be between 1 and 65535", c.Server.Port)
	}
	if c.Server.ShutdownTimeout < 0 {
		return fmt.Errorf("server shutdown timeout cannot be negative: %v", c.Server.ShutdownTimeout)
	}

	// Validate DatabaseConfig
	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		return fmt.Errorf("invalid database port: %d, must be between 1 and 65535", c.Database.Port)
	}
	if c.Database.MaxOpenConns <= 0 {
		return fmt.Errorf("database max open connections must be positive: %d", c.Database.MaxOpenConns)
	}
	if c.Database.MaxIdleConns <= 0 {
		return fmt.Errorf("database max idle connections must be positive: %d", c.Database.MaxIdleConns)
	}

	// Validate RiskConfig
	if c.Risk.MaxGlobalPositionSize <= 0 {
		return fmt.Errorf("risk max global position size must be positive: %f", c.Risk.MaxGlobalPositionSize)
	}
	if c.Risk.MaxCounterpartyExposure <= 0 {
		return fmt.Errorf("risk max counterparty exposure must be positive: %f", c.Risk.MaxCounterpartyExposure)
	}
	if c.Risk.CircuitBreakerThreshold <= 0 || c.Risk.CircuitBreakerThreshold >= 1.0 {
		return fmt.Errorf("risk circuit breaker threshold must be between 0 and 1 (exclusive): %f", c.Risk.CircuitBreakerThreshold)
	}

	return nil
}

// DSN returns the Data Source Name for connecting to the database.
// This is a convenience method to construct the connection string
// required by most Go database drivers (e.g., for PostgreSQL).
func (db *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode)
}

// Addr returns the network address for the HTTP server.
func (s *ServerConfig) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

```