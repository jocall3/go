```go
// Package risk provides concrete risk controls, such as velocity and frequency
// limiters, that are consulted by the execution engine during command processing.
// These controls are designed to make the system 'slow to attackers' by
// enforcing deterministic, fail-closed limits on activity.
package risk

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// Limiter defines the interface for a generic risk control mechanism.
// Limiters are consulted to determine if an action should be allowed based on
// historical activity. They are designed to be "fail-closed", meaning any
// ambiguity or breach results in denial.
type Limiter interface {
	// Check assesses if a proposed value (e.g., transaction amount, or 1 for an event count)
	// is within the limiter's constraints without consuming the capacity.
	// It returns true if the action is permissible, and false with a reason if not.
	Check(value float64) (bool, string)

	// Consume records the value against the limit, assuming a prior Check passed.
	// This should only be called after the associated action has been successfully
	// committed to prevent consuming limits for failed operations.
	Consume(value float64)

	// Name returns the unique identifier for the limiter.
	Name() string

	// CurrentUsage returns the current total value within the active window.
	CurrentUsage() float64
}

// event represents a single occurrence tracked by a limiter.
type event struct {
	timestamp time.Time
	value     float64
}

// SlidingWindowLimiter implements a sliding window rate limiter. It can be used
// to control both frequency (number of events) and velocity (total value of events).
// It is safe for concurrent use.
type SlidingWindowLimiter struct {
	mu     sync.Mutex
	name   string
	limit  float64
	window time.Duration
	events []event

	// timeSource allows for injecting a custom time function, primarily for testing.
	// If nil, time.Now is used.
	timeSource func() time.Time
}

// NewVelocityLimiter creates a new limiter that tracks the cumulative value of events
// over a time window. For example, limit $1,000,000 per hour.
func NewVelocityLimiter(name string, limit float64, window time.Duration) *SlidingWindowLimiter {
	return newSlidingWindowLimiter(name, limit, window)
}

// NewFrequencyLimiter creates a new limiter that tracks the number of events
// over a time window. For example, limit 100 API calls per minute.
// The limit is an int for convenience, but stored as a float64 internally.
func NewFrequencyLimiter(name string, limit int, window time.Duration) *SlidingWindowLimiter {
	return newSlidingWindowLimiter(name, float64(limit), window)
}

// newSlidingWindowLimiter is the internal constructor.
func newSlidingWindowLimiter(name string, limit float64, window time.Duration) *SlidingWindowLimiter {
	if window <= 0 {
		// A non-positive window is a logical error, panic to prevent misconfiguration.
		panic("limiter window must be positive")
	}
	return &SlidingWindowLimiter{
		name:       name,
		limit:      limit,
		window:     window,
		events:     make([]event, 0),
		timeSource: time.Now,
	}
}

// now returns the current time, using the timeSource if available.
func (l *SlidingWindowLimiter) now() time.Time {
	if l.timeSource != nil {
		return l.timeSource()
	}
	return time.Now()
}

// prune removes events that are older than the current window.
// This must be called within a locked section.
func (l *SlidingWindowLimiter) prune(now time.Time) {
	cutoff := now.Add(-l.window)

	// Find the index of the first event that is within the window.
	// sort.Search is efficient as events are sorted by timestamp.
	firstValidIndex := sort.Search(len(l.events), func(i int) bool {
		// We are looking for the first event that is NOT before the cutoff time.
		return !l.events[i].timestamp.Before(cutoff)
	})

	if firstValidIndex > 0 {
		// copy is used to be memory-safe and prevent leaks from the underlying array.
		remaining := len(l.events) - firstValidIndex
		copy(l.events, l.events[firstValidIndex:])
		l.events = l.events[:remaining]
	}
}

// currentUsage calculates the total value of events within the current window.
// This must be called within a locked section after pruning.
func (l *SlidingWindowLimiter) currentUsage() float64 {
	var total float64
	for _, e := range l.events {
		total += e.value
	}
	return total
}

// Check assesses if a proposed value is within the limiter's constraints.
// For frequency limiters, the value should typically be 1.0.
func (l *SlidingWindowLimiter) Check(value float64) (bool, string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	l.prune(now)

	currentUsage := l.currentUsage()
	if currentUsage+value > l.limit {
		reason := fmt.Sprintf(
			"limiter '%s' breached: proposed value %.2f + current usage %.2f > limit %.2f in window %s",
			l.name, value, currentUsage, l.limit, l.window,
		)
		return false, reason
	}

	return true, ""
}

// Consume records the value against the limit.
// For frequency limiters, the value should typically be 1.0.
func (l *SlidingWindowLimiter) Consume(value float64) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	// Prune before adding to keep the slice size manageable.
	l.prune(now)

	l.events = append(l.events, event{
		timestamp: now,
		value:     value,
	})
}

// Name returns the unique identifier for the limiter.
func (l *SlidingWindowLimiter) Name() string {
	return l.name
}

// CurrentUsage returns the current total value within the active window.
// It is safe for concurrent use and useful for monitoring.
func (l *SlidingWindowLimiter) CurrentUsage() float64 {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.now()
	l.prune(now)
	return l.currentUsage()
}

```