```go
package server

import (
	"context"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
)

// contextKey is a custom type for context keys to avoid collisions.
type contextKey string

const (
	// requestIDKey is the key used to store the request ID in the context.
	requestIDKey = contextKey("requestID")
	// principalKey is the key used to store the authenticated principal in the context.
	principalKey = contextKey("principal")
)

// Principal represents the authenticated entity making the request.
// It's placed in the request context by the authentication middleware.
type Principal struct {
	ID      string
	Roles   []string
	IsAdmin bool
}

// responseWriter is a wrapper around http.ResponseWriter that captures the status code
// and other metrics for logging.
type responseWriter struct {
	http.ResponseWriter
	statusCode    int
	headerWritten bool
}

// newResponseWriter creates a new responseWriter.
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	// Default to 200 OK, as http.ResponseWriter does.
	return &responseWriter{w, http.StatusOK, false}
}

// WriteHeader captures the status code and calls the underlying ResponseWriter's WriteHeader.
func (rw *responseWriter) WriteHeader(code int) {
	if rw.headerWritten {
		return
	}
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
	rw.headerWritten = true
}

// Write calls the underlying ResponseWriter's Write and ensures WriteHeader is called first.
func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.headerWritten {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

// requestID is a middleware that injects a unique request ID into each request's context.
// If a request already has an X-Request-ID header, it is honored.
// This should be one of the first middleware in the chain.
func (s *Server) requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		w.Header().Set("X-Request-ID", requestID)
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// logging is a middleware that provides structured logging for each request.
// It logs the request details, response status, and duration.
// It should be placed after requestID to include the ID in logs.
func (s *Server) logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID, _ := r.Context().Value(requestIDKey).(string)

		// Wrap the response writer to capture the status code.
		rw := newResponseWriter(w)

		// The logger with request-specific fields.
		requestLogger := s.logger.With(
			slog.String("request_id", requestID),
			slog.String("http.method", r.Method),
			slog.String("http.url", r.URL.RequestURI()),
			slog.String("remote_addr", r.RemoteAddr),
			slog.String("user_agent", r.UserAgent()),
		)

		requestLogger.Info("HTTP request started")

		// Pass our wrapped response writer to the next handler.
		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		requestLogger.Info("HTTP request completed",
			slog.Int("http.status_code", rw.statusCode),
			slog.Duration("duration", duration),
		)
	})
}

// recoverPanic is a middleware that recovers from panics, logs the error with a stack trace,
// and returns a 500 Internal Server Error. This helps prevent the server from crashing.
// It should be placed after logging to ensure panics in handlers are caught and logged correctly,
// and to correctly handle cases where response headers have already been written.
func (s *Server) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				requestID, _ := r.Context().Value(requestIDKey).(string)

				s.logger.Error("panic recovered",
					slog.String("request_id", requestID),
					slog.Any("panic_error", err),
					slog.String("stack_trace", string(debug.Stack())),
				)

				// Check if the response header has already been written.
				// This relies on our custom responseWriter from the logging middleware.
				if rw, ok := w.(*responseWriter); ok {
					if rw.headerWritten {
						// Headers already sent, we can't send a 500.
						// The server will close the connection.
						return
					}
				}

				// If headers are not written, send a generic 500 error.
				// Avoid leaking internal details in the error message.
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"error": "internal server error"}`))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// authenticate is a middleware that enforces authentication.
// It checks for a valid API key and populates the request context with a Principal.
// It's a fail-closed implementation: any failure in validation results in a 401 Unauthorized.
func (s *Server) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

		if apiKey == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error": "authentication required: missing X-API-Key header"}`))
			return
		}

		// In a production system, these keys would be loaded from a secure vault
		// and the validation logic would be more robust (e.g., constant-time comparison).
		// TODO: Replace with a proper authentication mechanism (e.g., JWT, mTLS, OAuth2).
		var principal Principal
		isValid := false

		// This is a placeholder for a real authentication system.
		// A real implementation would look up the key in a database or cache.
		if apiKey == s.config.Auth.AdminAPIKey {
			principal = Principal{ID: "admin-principal", Roles: []string{"admin", "user"}, IsAdmin: true}
			isValid = true
		} else if apiKey == s.config.Auth.UserAPIKey {
			principal = Principal{ID: "user-principal", Roles: []string{"user"}, IsAdmin: false}
			isValid = true
		}

		if !isValid {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error": "invalid credentials"}`))
			return
		}

		// Add the principal to the context for downstream handlers.
		ctx := context.WithValue(r.Context(), principalKey, principal)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getPrincipalFromContext retrieves the authenticated principal from the request context.
// It returns the principal and true if found, otherwise an empty principal and false.
func getPrincipalFromContext(ctx context.Context) (Principal, bool) {
	p, ok := ctx.Value(principalKey).(Principal)
	return p, ok
}
### END_OF_FILE_COMPLETED ###
```