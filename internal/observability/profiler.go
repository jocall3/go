```go
// Copyright 2023 The Bridge Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package observability

import (
	"context"
	"log"
	"net/http"
	"net/http/pprof" // Import for side-effect of registering pprof handlers
)

// Profiler provides an HTTP server for runtime profiling data.
// It integrates with Go's built-in pprof tool to expose endpoints for
// CPU, memory, and other runtime profiling information. This is a crucial
// tool for identifying performance bottlenecks and optimizing the system
// in a production environment.
//
// The profiler is designed to be started as a non-blocking background service.
// It should typically be exposed on a private network or localhost for security,
// as the data it exposes can be sensitive.
type Profiler struct {
	listenAddr string
	server     *http.Server
}

// NewProfiler creates and configures a new Profiler instance.
// The listenAddr should be in the format "host:port", e.g., "localhost:6060".
func NewProfiler(listenAddr string) *Profiler {
	if listenAddr == "" {
		// Provide a default if none is specified to avoid silent failure.
		listenAddr = "localhost:6060"
		log.Printf("WARN: Profiler listen address not specified, defaulting to %s", listenAddr)
	}
	return &Profiler{
		listenAddr: listenAddr,
	}
}

// Start initializes and starts the pprof HTTP server in a separate goroutine.
// It does not block. Any errors during server startup will be logged.
// The server exposes the standard pprof endpoints under the /debug/pprof/ path.
func (p *Profiler) Start() {
	// pprof registers its handlers with the DefaultServeMux. For clarity and
	// to avoid potential conflicts in a complex system, we create a new mux
	// and explicitly register the handlers. This adheres to the principle of
	// explicit state and dependencies, preventing side-effects from other
	// packages that might also use the DefaultServeMux.
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	p.server = &http.Server{
		Addr:    p.listenAddr,
		Handler: mux,
	}

	go func() {
		log.Printf("INFO: Profiler server starting on %s", p.listenAddr)
		log.Println("INFO: Access pprof at http://" + p.listenAddr + "/debug/pprof/")
		log.Println("INFO: To profile CPU for 30s: go tool pprof http://" + p.listenAddr + "/debug/pprof/profile?seconds=30")
		log.Println("INFO: To get a heap profile: go tool pprof http://" + p.listenAddr + "/debug/pprof/heap")

		if err := p.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("ERROR: Profiler server failed: %v", err)
		}
	}()
}

// Stop gracefully shuts down the pprof HTTP server.
// It waits for active connections to finish up to the timeout specified
// in the provided context. This ensures a clean shutdown, which is critical
// for a production-grade system.
func (p *Profiler) Stop(ctx context.Context) error {
	if p.server == nil {
		log.Println("WARN: Profiler server was not running, nothing to stop.")
		return nil
	}

	log.Println("INFO: Shutting down profiler server...")
	if err := p.server.Shutdown(ctx); err != nil {
		log.Printf("ERROR: Profiler server graceful shutdown failed: %v", err)
		return err
	}
	log.Println("INFO: Profiler server stopped.")
	return nil
}

```