```go
// Package platform contains foundational code for logging, configuration,
// and other cross-cutting concerns that support the application's core logic.
package platform

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// contextKey is a private type to prevent collisions with other packages' context keys.
// It's used to store and retrieve values from a context.Context.
type contextKey string

// Define the keys for context values. Using a custom type ensures that these keys
// do not conflict with keys from other packages.
const (
	// requestIDKey is the context key for the request ID.
	requestIDKey contextKey = "requestID"
	// loggerKey is the context key for the logger instance.
	loggerKey contextKey = "logger"
)

// ContextHandler is a slog.Handler wrapper that enriches log records with values
// extracted from the context.Context. This allows for automatic inclusion of
// metadata like request IDs in every log message within a specific request's lifecycle.
type ContextHandler struct {
	slog.Handler
}

// NewContextHandler creates a new ContextHandler that wraps the given handler.
// This is the core of our context-aware logging setup.
func NewContextHandler(h slog.Handler) *ContextHandler {
	return &ContextHandler{Handler: h}
}

// Handle processes the log record. Before passing the record to the wrapped
// handler, it inspects the context for well-known keys (like requestIDKey)
// and adds their values as attributes to the log record.
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	// Add request_id to the log record if it exists in the context.
	if id, ok := RequestIDFromContext(ctx); ok {
		r.AddAttrs(slog.String("request_id", id))
	}

	// This is extensible. Other context-based values can be added here.
	// For example, user ID, tenant ID, etc.

	return h.Handler.Handle(ctx, r)
}

// WithAttrs returns a new ContextHandler whose attributes consist of the
// wrapped handler's attributes followed by the provided attrs. This ensures
// that the context-aware behavior is preserved when new attributes are added
// to the logger.
func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewContextHandler(h.Handler.WithAttrs(attrs))
}

// WithGroup returns a new ContextHandler with the given group name. This
// preserves the context-aware behavior when a new logging group is created.
func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return NewContextHandler(h.Handler.WithGroup(name))
}

// NewLogger initializes a new structured logger for the application.
// It uses a JSON handler for machine-readable output, which is ideal for
// production environments and log aggregation systems. The logger is configured
// to be context-aware, automatically including request IDs.
//
// Parameters:
//   - w: The io.Writer to write logs to. Defaults to os.Stdout if nil.
//   - level: The minimum log level to record (e.g., slog.LevelInfo, slog.LevelDebug).
func NewLogger(w io.Writer, level slog.Level) *slog.Logger {
	if w == nil {
		w = os.Stdout
	}

	opts := &slog.HandlerOptions{
		// AddSource adds the source file and line number to the log record.
		// This is invaluable for debugging but has a minor performance cost.
		// It can be disabled in high-performance scenarios if needed.
		AddSource: true,
		Level:     level,
	}

	// The logging pipeline is composed:
	// 1. A base handler (slog.NewJSONHandler) formats logs as JSON.
	// 2. Our custom ContextHandler wraps the base handler to add context values.
	baseHandler := slog.NewJSONHandler(w, opts)
	contextHandler := NewContextHandler(baseHandler)

	logger := slog.New(contextHandler)

	return logger
}

// WithRequestID returns a new context that contains the given request ID.
// This should be called by the entry point of a request (e.g., in HTTP middleware)
// to tag the entire request lifecycle.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// RequestIDFromContext extracts the request ID from the context, if present.
// It returns the ID and a boolean indicating whether the ID was found.
func RequestIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIDKey).(string)
	return id, ok
}

// ToContext stores the logger in the context. This allows passing a logger
// instance (potentially with request-specific attributes already added)
// through the call stack.
func ToContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext retrieves the logger from the context.
// If no logger is found in the context, it returns the default logger configured
// by slog.SetDefault(). This provides a safe fallback.
func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
### END_OF_FILE_COMPLETED ###
```