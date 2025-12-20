```go
// Package server defines the main HTTP server for the application.
// It is responsible for initializing the router, applying middleware,
// and managing the server's lifecycle, including graceful shutdown.
package server

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

	"github.com/your-org/your-project/internal/config"
	"github.com/your-org/your-project/internal/router"
	"github.com/your-org/your-project/internal/services"
)

// Server represents the main HTTP server.
// It encapsulates the http.Server and its dependencies.
type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
	config     *config.ServerConfig
}

// New creates and configures a new Server instance.
// It requires configuration, a structured logger, and a container for all
// API dependencies (like services and repositories) to be injected into handlers.
// This approach promotes loose coupling and testability.
func New(cfg *config.Config, logger *slog.Logger, deps *services.APIDependencies) (*Server, error) {
	if cfg == nil {
		return nil, errors.New("config cannot be nil")
	}
	if logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	if deps == nil {
		return nil, errors.New("api dependencies cannot be nil")
	}

	// Initialize the main router, which encapsulates all routes and middleware.
	mainRouter := router.New(logger, deps)

	// Configure the http.Server with production-grade settings.
	// Timeouts are critical to prevent resource exhaustion from slow or malicious clients.
	// - ReadTimeout: Protects against slowloris attacks.
	// - WriteTimeout: Prevents connections from being held open by slow clients.
	// - IdleTimeout: Closes keep-alive connections that are not being used.
	// - ErrorLog: Redirects the server's internal errors to our structured logger.
	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      mainRouter,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	return &Server{
		httpServer: httpServer,
		logger:     logger.With(slog.String("component", "server")),
		config:     &cfg.Server,
	}, nil
}

// Start runs the HTTP server and sets up a graceful shutdown mechanism.
// This method blocks until the server is shut down by an OS signal
// (SIGINT, SIGTERM) or a fatal error occurs.
func (s *Server) Start() error {
	// shutdownError channel is used to propagate errors from the shutdown goroutine.
	shutdownError := make(chan error)

	// This goroutine listens for OS signals to trigger a graceful shutdown.
	// This is a standard pattern for robust Go services.
	go func() {
		// Wait for an interrupt or terminate signal.
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		s.logger.Info("shutting down server", "signal", sig.String())

		// Create a context with a timeout to allow in-flight requests to complete.
		// This timeout ensures the shutdown process doesn't hang indefinitely.
		ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()

		// Disable keep-alives to prevent new connections during shutdown.
		s.httpServer.SetKeepAlivesEnabled(false)

		// Attempt to gracefully shut down the server.
		// Shutdown() will block until all active connections are handled or the context times out.
		if err := s.httpServer.Shutdown(ctx); err != nil {
			shutdownError <- fmt.Errorf("graceful shutdown failed: %w", err)
		}

		s.logger.Info("server shutdown complete")
		close(shutdownError)
	}()

	s.logger.Info("starting server", "addr", s.httpServer.Addr)

	// Start listening for requests.
	// ListenAndServe will block until the server is shut down.
	// If it returns an error, we check if it's ErrServerClosed, which is expected
	// during a graceful shutdown. Any other error is considered fatal.
	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server failed to start: %w", err)
	}

	// Wait for the shutdown goroutine to complete and return any error it produced.
	if err := <-shutdownError; err != nil {
		return err
	}

	return nil
}

// Shutdown provides a way to programmatically shut down the server,
// for example, during integration tests.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("initiating programmatic shutdown")
	return s.httpServer.Shutdown(ctx)
}

// Addr returns the network address the server is listening on.
func (s *Server) Addr() string {
	return s.httpServer.Addr
}

```