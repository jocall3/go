package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds the entire application configuration.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	AI       AIConfig
}

// ServerConfig holds configuration for the HTTP server.
type ServerConfig struct {
	Port            string
	Host            string
	Environment     string // development, staging, production
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	AllowedOrigins  []string
	Debug           bool
}

// DatabaseConfig holds configuration for the PostgreSQL database.
type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Name         string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}

// RedisConfig holds configuration for the Redis cache/queue.
type RedisConfig struct {
	Address      string
	Password     string
	DB           int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
}

// AIConfig holds configuration for AI services (LLMs).
type AIConfig struct {
	OpenAIAPIKey      string
	AnthropicAPIKey   string
	DefaultModel      string
	Temperature       float64
	MaxTokens         int
	RequestTimeout    time.Duration
	RetryCount        int
	EmbeddingModel    string
}

// Load initializes the configuration from environment variables.
// It sets reasonable defaults where environment variables are missing.
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port:            getEnv("SERVER_PORT", "8080"),
			Host:            getEnv("SERVER_HOST", "0.0.0.0"),
			Environment:     getEnv("APP_ENV", "development"),
			ReadTimeout:     getEnvAsDuration("SERVER_READ_TIMEOUT", 15*time.Second),
			WriteTimeout:    getEnvAsDuration("SERVER_WRITE_TIMEOUT", 15*time.Second),
			ShutdownTimeout: getEnvAsDuration("SERVER_SHUTDOWN_TIMEOUT", 10*time.Second),
			AllowedOrigins:  getEnvAsSlice("CORS_ALLOWED_ORIGINS", []string{"*"}, ","),
			Debug:           getEnvAsBool("APP_DEBUG", true),
		},
		Database: DatabaseConfig{
			Host:         getEnv("DB_HOST", "localhost"),
			Port:         getEnvAsInt("DB_PORT", 5432),
			User:         getEnv("DB_USER", "postgres"),
			Password:     getEnv("DB_PASSWORD", "postgres"),
			Name:         getEnv("DB_NAME", "app_db"),
			SSLMode:      getEnv("DB_SSLMODE", "disable"),
			MaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 25),
			MaxLifetime:  getEnvAsDuration("DB_MAX_LIFETIME", 5*time.Minute),
		},
		Redis: RedisConfig{
			Address:      getEnv("REDIS_ADDR", "localhost:6379"),
			Password:     getEnv("REDIS_PASSWORD", ""),
			DB:           getEnvAsInt("REDIS_DB", 0),
			ReadTimeout:  getEnvAsDuration("REDIS_READ_TIMEOUT", 3*time.Second),
			WriteTimeout: getEnvAsDuration("REDIS_WRITE_TIMEOUT", 3*time.Second),
			PoolSize:     getEnvAsInt("REDIS_POOL_SIZE", 10),
		},
		AI: AIConfig{
			OpenAIAPIKey:    getEnv("OPENAI_API_KEY", ""),
			AnthropicAPIKey: getEnv("ANTHROPIC_API_KEY", ""),
			DefaultModel:    getEnv("AI_DEFAULT_MODEL", "gpt-4-turbo"),
			Temperature:     getEnvAsFloat("AI_TEMPERATURE", 0.7),
			MaxTokens:       getEnvAsInt("AI_MAX_TOKENS", 2048),
			RequestTimeout:  getEnvAsDuration("AI_REQUEST_TIMEOUT", 60*time.Second),
			RetryCount:      getEnvAsInt("AI_RETRY_COUNT", 3),
			EmbeddingModel:  getEnv("AI_EMBEDDING_MODEL", "text-embedding-3-small"),
		},
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// DSN returns the PostgreSQL Data Source Name.
func (db *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode)
}

// validate performs basic validation on the configuration.
func (c *Config) validate() error {
	if c.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	if c.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}
	// Add more validation rules as necessary
	return nil
}

// Helper functions to read environment variables

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsFloat(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string, sep string) []string {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	return strings.Split(valueStr, sep)
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultValue
}