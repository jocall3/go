package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// These imports assume a standard project layout.
	// Replace "github.com/yourusername/yourproject" with your actual module name defined in go.mod.
	"github.com/yourusername/yourproject/internal/config"
	"github.com/yourusername/yourproject/internal/database"
	"github.com/yourusername/yourproject/internal/server"
)

func main() {
	// 1. Initialize Structured Logging
	// Using slog (Go 1.21+) for high-performance, structured JSON logging suitable for production.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// 2. Load Configuration
	// Loads environment variables and defaults.
	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	logger.Info("starting server", "env", cfg.Environment, "port", cfg.Port)

	// 3. Initialize Database Connection
	// Establishes a connection pool with context awareness.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("failed to close database connection", "error", err)
		}
	}()

	logger.Info("database connection established")

	// 4. Initialize API Server
	// Inject dependencies (Config, DB, Logger) into the server instance.
	srv := server.New(cfg, db, logger)

	// 5. Configure HTTP Server
	// Set timeouts to prevent slowloris attacks and resource leaks.
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      srv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// 6. Start Server with Graceful Shutdown
	// Use a buffered channel to listen for OS signals (SIGINT, SIGTERM).
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server failed", "error", err)
			os.Exit(1)
		}
	}()

	logger.Info("server is ready to handle requests", "address", httpServer.Addr)

	// Block until a signal is received
	sig := <-shutdownChan
	logger.Info("shutdown signal received", "signal", sig.String())

	// Create a deadline for the shutdown process
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Attempt graceful shutdown
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("server exited gracefully")
}