```go
// Copyright (c) 2024. The Bridge Project Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package observability

import (
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
)

// ReadinessProbe manages the readiness state of the application.
// It allows different components (e.g., database connections, message queue clients)
// to register their status and provides an HTTP handler to expose the overall readiness
// to orchestration systems like Kubernetes.
type ReadinessProbe struct {
	// mu protects the components map during writes.
	mu sync.RWMutex
	// components holds the readiness status of individual registered components.
	components map[string]bool
	// ready is an atomic boolean that reflects the overall readiness status.
	// It is updated whenever a component's status changes. Reading this value
	// is lock-free, making the readiness check endpoint highly performant.
	ready atomic.Bool
}

// NewReadinessProbe creates and initializes a new ReadinessProbe.
// The probe is initially not ready. It becomes ready only when all registered
// components have explicitly signaled that they are ready.
func NewReadinessProbe() *ReadinessProbe {
	p := &ReadinessProbe{
		components: make(map[string]bool),
	}
	// The system is not ready by default. This enforces fail-closed semantics.
	p.ready.Store(false)
	return p
}

// Register adds a new component to be tracked for readiness.
// Components are initially considered not ready upon registration.
// This function is safe for concurrent use.
func (p *ReadinessProbe) Register(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.components[name]; exists {
		// Avoid re-registration and potential state clobbering.
		return
	}
	p.components[name] = false
	p.recalculate()
}

// SetReady marks a component as ready. If the component was not ready before,
// the overall readiness state is recalculated.
// This function is safe for concurrent use.
func (p *ReadinessProbe) SetReady(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if ready, exists := p.components[name]; !exists || ready {
		// Do nothing if the component is not registered or is already ready.
		return
	}
	p.components[name] = true
	p.recalculate()
}

// SetNotReady marks a component as not ready. If the component was ready before,
// the overall readiness state is recalculated.
// This function is safe for concurrent use.
func (p *ReadinessProbe) SetNotReady(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if ready, exists := p.components[name]; !exists || !ready {
		// Do nothing if the component is not registered or is already not ready.
		return
	}
	p.components[name] = false
	p.recalculate()
}

// recalculate updates the overall readiness status based on the state of all components.
// It assumes all registered components must be ready for the application to be considered ready.
// This method must be called with the write lock held.
func (p *ReadinessProbe) recalculate() {
	// If there are no registered components, we are not ready. This forces
	// critical components to be explicitly registered and report as ready.
	if len(p.components) == 0 {
		p.ready.Store(false)
		return
	}

	for _, isReady := range p.components {
		if !isReady {
			// If any component is not ready, the entire system is not ready.
			p.ready.Store(false)
			return
		}
	}

	// All components are ready.
	p.ready.Store(true)
}

// Handler returns an http.HandlerFunc for the readiness probe (`/readyz`).
// It checks the overall readiness state and responds with HTTP 200 OK if ready,
// or HTTP 503 Service Unavailable if not. The response body contains a detailed
// breakdown of component statuses for operational visibility.
func (p *ReadinessProbe) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := p.currentStatus()
		responseBody, err := json.Marshal(status)
		if err != nil {
			// This should ideally not happen with our controlled map structure.
			http.Error(w, "Failed to serialize readiness status", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if p.ready.Load() {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		w.Write(responseBody)
	}
}

// currentStatus returns a snapshot of the current component statuses.
// This is useful for providing detailed debug information in the probe's response.
func (p *ReadinessProbe) currentStatus() map[string]any {
	p.mu.RLock()
	defer p.mu.RUnlock()

	componentStatus := make(map[string]string)
	for name, isReady := range p.components {
		if isReady {
			componentStatus[name] = "ready"
		} else {
			componentStatus[name] = "not_ready"
		}
	}

	overallStatus := "not_ready"
	if p.ready.Load() {
		overallStatus = "ready"
	}

	return map[string]any{
		"status":     overallStatus,
		"components": componentStatus,
	}
}

// LivenessHandler is a simple health check handler for the liveness probe (`/healthz`).
// It always returns HTTP 200 OK to indicate that the process is running and has not
// deadlocked. It should have no external dependencies.
func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	// A simple, constant response is sufficient.
	w.Write([]byte(`{"status":"ok"}`))
}

// NewHealthMux creates a new http.ServeMux and registers the standard health
// and readiness endpoints. It returns the mux and the readiness probe, allowing
// other parts of the application to register components and update their status.
func NewHealthMux() (*http.ServeMux, *ReadinessProbe) {
	mux := http.NewServeMux()
	probe := NewReadinessProbe()

	mux.HandleFunc("/healthz", LivenessHandler)
	mux.HandleFunc("/readyz", probe.Handler())

	return mux, probe
}

```