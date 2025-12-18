package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// RouterConfig holds configuration for the HTTP router.
type RouterConfig struct {
	// Timeout is the maximum duration for reading the entire request, including the body.
	Timeout time.Duration
	// CorsAllowedOrigins is a list of origins a cross-domain request can be executed from.
	CorsAllowedOrigins []string
	// EnablePprof enables the standard library pprof debugging endpoints.
	EnablePprof bool
}

// DefaultRouterConfig returns a default configuration for the router.
func DefaultRouterConfig() RouterConfig {
	return RouterConfig{
		Timeout:            60 * time.Second,
		CorsAllowedOrigins: []string{"*"},
		EnablePprof:        false,
	}
}

// NewRouter initializes a new chi.Router with production-ready middleware and settings.
// It accepts a variadic list of routeHandler functions to register domain-specific routes.
func NewRouter(cfg RouterConfig, routeHandlers ...func(chi.Router)) http.Handler {
	r := chi.NewRouter()

	// -------------------------------------------------------------------------
	// Middleware Setup
	// -------------------------------------------------------------------------

	// A good base middleware stack
	r.Use(middleware.RequestID)               // Injects a request ID into the context of each request
	r.Use(middleware.RealIP)                  // Sets the request's RemoteAddr to either X-Real-IP or X-Forwarded-For
	r.Use(middleware.Logger)                  // Logs the start and end of each request with the elapsed time
	r.Use(middleware.Recoverer)               // Recovers from panics without crashing the server
	r.Use(middleware.Timeout(cfg.Timeout))    // Stop processing after a specified duration
	r.Use(middleware.Compress(5))             // Compress responses with gzip
	r.Use(middleware.URLFormat)               // Parse extension from url and put it on request context
	r.Use(render.SetContentType(render.ContentTypeJSON)) // Set content-type to json by default

	// CORS Configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.CorsAllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// -------------------------------------------------------------------------
	// System Routes
	// -------------------------------------------------------------------------

	// Health Check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Optional: Profiling routes for debugging
	if cfg.EnablePprof {
		r.Mount("/debug", middleware.Profiler())
	}

	// -------------------------------------------------------------------------
	// Domain Routes
	// -------------------------------------------------------------------------

	// Mount all injected route handlers
	// This allows the router to be agnostic of the specific business logic
	r.Group(func(r chi.Router) {
		for _, handler := range routeHandlers {
			handler(r)
		}
	})

	// -------------------------------------------------------------------------
	// Error Handling
	// -------------------------------------------------------------------------

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"resource not found"}`))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"method not allowed"}`))
	})

	return r
}