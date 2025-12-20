```go
// Copyright 2024 The Bridge Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"time"

	"github.com/google/uuid"
)

// EventType is a string alias for the type of a domain event.
// Using a dedicated type improves type safety and clarity over a raw string.
type EventType string

// Event is the fundamental interface for all domain events in the system.
// It defines the contract that all events must adhere to, ensuring they
// carry essential metadata for processing, auditing, and replay. An event
// represents a fact that has occurred in the past and is immutable.
type Event interface {
	// Header returns the common metadata for the event.
	Header() *EventHeader

	// EventType returns the specific type of the event.
	EventType() EventType

	// AggregateID returns the identifier of the aggregate root this event belongs to.
	// For example, an AccountID or an OrderID.
	AggregateID() string

	// Timestamp returns the UTC time at which the event occurred.
	Timestamp() time.Time
}

// EventHeader contains common metadata for all domain events.
// It is intended to be embedded in concrete event structs.
// This composition ensures that all events share a consistent structure
// for identity, causality, and temporal ordering, which is critical for
// a deterministic and auditable system.
type EventHeader struct {
	// ID is the unique identifier for this specific event instance (UUID v4).
	ID string `json:"id"`

	// Type is the specific type of the event (e.g., "AccountCreated", "OrderPlaced").
	Type EventType `json:"type"`

	// AggregateID is the identifier of the aggregate root this event applies to.
	AggregateID string `json:"aggregate_id"`

	// Timestamp is the UTC time at which the event was created. It must be
	// deterministic and consistent across all nodes processing the event.
	Timestamp time.Time `json:"timestamp"`

	// Version is the version of the aggregate after this event is applied.
	// It is used for optimistic concurrency control and to ensure event order.
	// The first event for an aggregate has version 1.
	Version uint64 `json:"version"`

	// Source identifies the component, service, or system that generated the event.
	// Useful for tracing and debugging in a distributed system.
	Source string `json:"source"`

	// CorrelationID links events that are part of the same logical operation or request.
	// This helps in tracing a flow through the system.
	CorrelationID string `json:"correlation_id"`

	// CausationID is the ID of the command or event that caused this event to be generated.
	// This builds a causal chain, crucial for understanding system behavior and for debugging.
	CausationID string `json:"causation_id"`
}

// NewEventHeader creates and initializes a new EventHeader.
// It automatically generates a new UUID for the event ID and sets the timestamp to the current UTC time.
// The caller is responsible for providing the context (aggregate, version, source, correlation, causation).
func NewEventHeader(aggregateID string, eventType EventType, version uint64, source, correlationID, causationID string) EventHeader {
	return EventHeader{
		ID:            uuid.NewString(),
		Type:          eventType,
		AggregateID:   aggregateID,
		Timestamp:     time.Now().UTC(),
		Version:       version,
		Source:        source,
		CorrelationID: correlationID,
		CausationID:   causationID,
	}
}

// Header returns the EventHeader itself. This satisfies the Event interface.
func (h *EventHeader) Header() *EventHeader {
	return h
}

// EventType returns the specific type of the event. This satisfies the Event interface.
func (h *EventHeader) EventType() EventType {
	return h.Type
}

// AggregateID returns the identifier of the aggregate root. This satisfies the Event interface.
func (h *EventHeader) AggregateID() string {
	return h.AggregateID
}

// Timestamp returns the time at which the event occurred. This satisfies the Event interface.
func (h *EventHeader) Timestamp() time.Time {
	return h.Timestamp
}

// Ensure EventHeader correctly implements the Event interface at compile time.
// This is a standard Go pattern to verify interface compliance.
var _ Event = (*EventHeader)(nil)

```