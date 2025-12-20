```go
package platform

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// ShutdownHook is a function that is called during the graceful shutdown process.
// It receives a context that is canceled when the shutdown timeout is reached.
type ShutdownHook func(ctx context.Context) error

// Manager handles the graceful shutdown of the application.
// It listens for OS signals and coordinates the shutdown of registered components,
// ensuring that in-flight operations are completed.
type Manager struct {
	// wg is used to wait for all in-flight operations to complete.
	wg sync.WaitGroup

	// mu protects access to the hooks slice.
	mu sync.Mutex

	// hooks contains the functions to be executed on shutdown.
	// They are executed in LIFO (Last-In, First-Out) order.
	hooks []ShutdownHook

	// shutdownSignal is a channel that receives OS signals for shutdown.
	shutdownSignal chan os.Signal

	// shutdownInitiated is closed when shutdown begins, either by signal or programmatically.
	shutdownInitiated chan struct{}

	// timeout is the maximum duration to wait for a graceful shutdown.
	timeout time.Duration
}

// NewManager creates a new shutdown manager.
// It takes a timeout duration that specifies the maximum time to wait for
// all hooks and in-flight operations to complete.
func NewManager(timeout time.Duration) *Manager {
	m := &Manager{
		timeout:           timeout,
		shutdownInitiated: make(chan struct{}),
		shutdownSignal:    make(chan os.Signal, 1),
	}

	// Register the OS signals we want to listen for.
	signal.Notify(m.shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	return m
}

// Register adds one or more shutdown hooks to the manager.
// Hooks are executed in reverse order of registration (LIFO).
// This is useful for ensuring dependencies are shut down in the correct order
// (e.g., shut down the server before the database connection it uses).
func (m *Manager) Register(hooks ...ShutdownHook) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.hooks = append(m.hooks, hooks...)
}

// Track adds a delta of 1 to the manager's WaitGroup counter.
// This should be called at the beginning of an operation that needs to
// complete before the application can shut down (e.g., an HTTP request handler).
func (m *Manager) Track() {
	m.wg.Add(1)
}

// Done decrements the manager's WaitGroup counter.
// This should be called at the end of an operation being tracked, typically in a defer statement.
func (m *Manager) Done() {
	m.wg.Done()
}

// WaitForShutdown blocks until a shutdown signal is received or shutdown is
// initiated programmatically. It then orchestrates the graceful shutdown process.
// This function should be called from the main goroutine as the final blocking call.
func (m *Manager) WaitForShutdown() {
	// Block until a signal is received or shutdown is triggered programmatically.
	select {
	case s := <-m.shutdownSignal:
		log.Printf("Shutdown signal received: %s. Starting graceful shutdown...", s)
	case <-m.shutdownInitiated:
		log.Printf("Programmatic shutdown initiated. Starting graceful shutdown...")
	}

	// Create a context with a timeout for the entire shutdown process.
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	// Perform the shutdown.
	if err := m.shutdown(ctx); err != nil {
		log.Printf("Graceful shutdown failed: %v", err)
		os.Exit(1)
	}

	log.Println("Graceful shutdown completed successfully.")
}

// InitiateShutdown programmatically starts the shutdown process.
// This is useful for triggering shutdown from within the application,
// for example, after a critical, unrecoverable error. This call is non-blocking
// and idempotent.
func (m *Manager) InitiateShutdown() {
	// Use a non-blocking close to signal shutdown.
	// Closing a channel is an idempotent broadcast mechanism.
	select {
	case <-m.shutdownInitiated:
		// Already closed, do nothing.
	default:
		close(m.shutdownInitiated)
	}
}

// shutdown orchestrates the actual shutdown logic.
func (m *Manager) shutdown(ctx context.Context) error {
	// Phase 1: Wait for all in-flight operations to complete.
	log.Println("Waiting for in-flight operations to complete...")

	waitGroupDone := make(chan struct{})
	go func() {
		m.wg.Wait()
		close(waitGroupDone)
	}()

	select {
	case <-waitGroupDone:
		log.Println("All in-flight operations completed.")
	case <-ctx.Done():
		return fmt.Errorf("shutdown timed out while waiting for in-flight operations: %w", ctx.Err())
	}

	// Phase 2: Execute all registered shutdown hooks in LIFO order.
	log.Println("Executing shutdown hooks...")
	m.mu.Lock()
	defer m.mu.Unlock()

	var firstErr error

	// Iterate backwards to execute in LIFO order.
	for i := len(m.hooks) - 1; i >= 0; i-- {
		hook := m.hooks[i]
		if err := hook(ctx); err != nil {
			// We continue executing other hooks even if one fails,
			// but we will return the first error encountered.
			// This ensures a best-effort cleanup.
			log.Printf("Shutdown hook failed: %v", err)
			if firstErr == nil {
				firstErr = err
			}
		}
	}

	if firstErr != nil {
		return fmt.Errorf("one or more shutdown hooks failed: %w", firstErr)
	}

	return nil
}
### END_OF_FILE_COMPLETED ###
```