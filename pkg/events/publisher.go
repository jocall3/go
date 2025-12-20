```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"context"
)

// Publisher defines the contract for broadcasting events to subscribers.
// This interface abstracts the underlying event bus or message queue, allowing
// for a decoupled architecture where components can react to state changes
// without direct dependencies.
//
// The design of the Publisher is critical for achieving the system's goals:
// - **Speed:** Implementations can be optimized for low-latency, in-memory
//   dispatch for co-located services, or high-throughput message brokers for
//   distributed systems.
// - **Stability:** A robust implementation (e.g., backed by a persistent queue)
//   ensures that events are not lost during service outages, supporting
//   at-least-once delivery semantics crucial for financial transactions.
// - **Auditability:** Every event published can be logged and stored, creating
//   an immutable, replayable history of system state changes. This is a
//   cornerstone for regulatory compliance and incident analysis.
//
// Implementations might range from a simple in-memory fan-out dispatcher to
// a sophisticated client for a distributed messaging system like Kafka, NATS,
// or a cloud-native service like Google Cloud Pub/Sub.
type Publisher interface {
	// Publish sends an event to the event stream or topic.
	//
	// The `ctx` parameter is used for request-scoped values like trace IDs,
	// and for managing deadlines and cancellation. A robust implementation
	// should propagate this context to downstream systems.
	//
	// The `event` parameter is the payload to be broadcast. It must conform
	// to the Event interface, ensuring it carries essential metadata like
	// type, timestamp, and a unique ID.
	//
	// An error is returned if the event cannot be published. The exact
	// semantics of this error (e.g., whether it's retryable) depend on the
	// specific implementation. For critical financial events, callers should
	// implement retry logic or use patterns like the transactional outbox
	// to guarantee delivery.
	Publish(ctx context.Context, event Event) error

	// Shutdown gracefully closes the publisher, ensuring that any buffered
	// events are flushed and underlying resources (like network connections)
	// are released.
	//
	// The `ctx` parameter allows for a timeout on the shutdown process,
	// preventing the application from hanging indefinitely.
	Shutdown(ctx context.Context) error
}

```