```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package events

import (
	"context"
	"fmt"
	"sync"
)

// MemoryStore is an in-memory implementation of the Store interface.
// It is designed for fast and isolated unit/integration tests, embodying the
// principle of making the system 'fast to builders'. It is safe for concurrent use.
type MemoryStore struct {
	mu     sync.RWMutex
	events map[ID]Event // Primary store: event ID -> event

	// Indexes for efficient lookups
	byStream    map[string][]Event // stream ID -> sorted list of events by version
	byAggregate map[string][]Event // aggregate key -> sorted list of events by version
}

// NewMemoryStore creates and initializes a new MemoryStore.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		events:      make(map[ID]Event),
		byStream:    make(map[string][]Event),
		byAggregate: make(map[string][]Event),
	}
}

// A compile-time check to ensure MemoryStore implements the Store interface.
var _ Store = (*MemoryStore)(nil)

// Append adds a new event to the store.
// It enforces stream version consistency to prevent race conditions and ensure ordering,
// which is critical for deterministic state transitions.
// If an event with the same ID already exists, it returns ErrDuplicateEvent.
// If the event's stream version is not exactly one greater than the latest
// version in the stream, it returns ErrVersionConflict.
func (s *MemoryStore) Append(ctx context.Context, event Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check for duplicate event ID
	if _, exists := s.events[event.ID]; exists {
		return fmt.Errorf("%w: id %s", ErrDuplicateEvent, event.ID)
	}

	// Check for stream version conflict
	streamEvents := s.byStream[event.StreamID]
	currentVersion := uint64(0)
	if len(streamEvents) > 0 {
		currentVersion = streamEvents[len(streamEvents)-1].StreamVersion
	}

	if event.StreamVersion != currentVersion+1 {
		return fmt.Errorf("%w: expected version %d, got %d for stream %s",
			ErrVersionConflict, currentVersion+1, event.StreamVersion, event.StreamID)
	}

	// Add to primary store
	s.events[event.ID] = event

	// Add to stream index
	s.byStream[event.StreamID] = append(s.byStream[event.StreamID], event)

	// Add to aggregate index
	aggregateKey := s.getAggregateKey(event.AggregateType, event.AggregateID)
	s.byAggregate[aggregateKey] = append(s.byAggregate[aggregateKey], event)

	return nil
}

// Get retrieves a single event by its unique ID.
// If no event is found, it returns ErrEventNotFound.
func (s *MemoryStore) Get(ctx context.Context, id ID) (Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	event, ok := s.events[id]
	if !ok {
		return Event{}, fmt.Errorf("%w: id %s", ErrEventNotFound, id)
	}
	return event, nil
}

// GetByStream retrieves all events for a given stream ID starting after a specific version.
// The returned events are ordered by their stream version.
// If fromVersion is 0, it retrieves all events for the stream.
// If the stream does not exist, it returns a nil slice and no error.
func (s *MemoryStore) GetByStream(ctx context.Context, streamID string, fromVersion uint64) ([]Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	streamEvents, ok := s.byStream[streamID]
	if !ok {
		return nil, nil
	}

	// A linear scan is simple and sufficient for a test implementation.
	// For very long event streams in a production store, a more performant
	// lookup (e.g., binary search or a database query) would be used.
	var results []Event
	for _, event := range streamEvents {
		if event.StreamVersion > fromVersion {
			results = append(results, event)
		}
	}

	return results, nil
}

// GetByAggregate retrieves all events for a given aggregate type and ID.
// The returned events are ordered by their stream version.
// If the aggregate does not exist, it returns a nil slice and no error.
// It returns a copy of the event slice to prevent mutation of the store's internal state.
func (s *MemoryStore) GetByAggregate(ctx context.Context, aggregateType string, aggregateID string) ([]Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	aggregateKey := s.getAggregateKey(aggregateType, aggregateID)
	aggregateEvents, ok := s.byAggregate[aggregateKey]
	if !ok {
		return nil, nil
	}

	// Create a copy to prevent external modification of the internal slice
	result := make([]Event, len(aggregateEvents))
	copy(result, aggregateEvents)

	return result, nil
}

// getAggregateKey generates a consistent key for the aggregate index map.
func (s *MemoryStore) getAggregateKey(aggregateType, aggregateID string) string {
	return fmt.Sprintf("%s:%s", aggregateType, aggregateID)
}

```