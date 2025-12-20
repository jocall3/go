```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"context"
)

// Subscriber defines the interface for any component that needs to react to
// events published on the event bus. Implementations of this interface, such
// as the risk engine or settlement projector, process events to update their
// internal state or trigger further actions.
//
// The event handling logic must be idempotent, as events may be redelivered
// in certain failure scenarios. The system's event bus guarantees at-least-once
// delivery. A well-designed subscriber can handle the same event multiple times
// without causing incorrect state changes.
type Subscriber interface {
	// HandleEvent processes a single event from the event stream.
	//
	// The context.Context can be used for managing deadlines, cancellation signals,
	// and passing request-scoped values across API boundaries and between processes.
	//
	// The event parameter is the specific event to be processed. Implementations
	// should use a type switch to handle different concrete event types they are
	// interested in.
	//
	// Returning a non-nil error indicates that the event could not be processed
	// successfully. The event bus may attempt to redeliver the event based on its
	// retry policy. A permanent failure (e.g., an unrecoverable data validation
	// error) should be indicated by a specific error type that the bus can
	// recognize to prevent infinite retries and move the event to a dead-letter queue.
	// If the error is nil, the event is considered successfully processed and will be
	// acknowledged by the event bus.
	HandleEvent(ctx context.Context, event Event) error

	// SubscriberID returns the unique identifier for this subscriber. This ID is
	// crucial for the event bus to track the processing offset for each subscriber,
	// ensuring that no events are missed and that the system can resume from the
	// last known position after a restart or failure.
	SubscriberID() SubscriberID
}

// SubscriberID is a unique, human-readable identifier for a subscriber.
// Example: "risk-engine", "settlement-projector".
type SubscriberID string

// String returns the string representation of the SubscriberID.
func (s SubscriberID) String() string {
	return string(s)
}


```