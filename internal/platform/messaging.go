```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package platform

import (
	"context"
	"errors"
)

// Standard errors for the messaging platform. These provide a consistent error
// vocabulary for consumers, regardless of the underlying message bus implementation.
var (
	// ErrNotConnected is returned when an operation is attempted on a disconnected client.
	ErrNotConnected = errors.New("messaging: not connected to the message bus")

	// ErrPublishFailed is returned when a message fails to be published.
	// The underlying error from the specific driver should be wrapped for more context.
	ErrPublishFailed = errors.New("messaging: failed to publish message")

	// ErrSubscribeFailed is returned when a subscription attempt fails.
	ErrSubscribeFailed = errors.New("messaging: failed to subscribe to subject")

	// ErrAckFailed is returned when a message acknowledgment fails.
	ErrAckFailed = errors.New("messaging: failed to acknowledge message")

	// ErrNackFailed is returned when a message negative-acknowledgment fails.
	ErrNackFailed = errors.New("messaging: failed to nack message")

	// ErrInvalidSubject is returned for an invalid subject or topic name.
	ErrInvalidSubject = errors.New("messaging: invalid subject name")

	// ErrShutdown is returned when an operation is attempted during or after shutdown.
	ErrShutdown = errors.New("messaging: client is shut down")
)

// Publisher defines the interface for publishing messages to an external message bus.
// This abstraction allows for different underlying implementations (e.g., NATS, Kafka)
// to be used interchangeably by the core application logic. It is designed for
// fire-and-forget event publication to external consumers.
type Publisher interface {
	// Publish sends a message to the specified subject.
	// The operation should be resilient and handle transient connection issues
	// according to its implementation's configuration (e.g., retries).
	// A context is used for deadlines, cancellation, and propagating tracing information.
	Publish(ctx context.Context, subject string, data []byte) error

	// Close gracefully terminates the connection to the message bus.
	// It should attempt to flush any buffered messages before closing to prevent data loss.
	// Calling Close on an already closed publisher should not panic.
	Close() error
}

// Message represents a message received from the message bus.
// It provides methods to access its data and to explicitly manage its lifecycle
// through acknowledgments. This explicit control is crucial for financial systems
// to ensure that every message is processed correctly and durably.
type Message interface {
	// Data returns the raw payload of the message.
	Data() []byte

	// Subject returns the subject the message was received on.
	Subject() string

	// Ack acknowledges the message, indicating successful processing.
	// This typically removes the message from the queue, preventing redelivery.
	// This is a critical step in ensuring "at-least-once" or "exactly-once" processing.
	Ack(ctx context.Context) error

	// Nack negatively acknowledges the message, indicating a processing failure.
	// The message bus may attempt to redeliver it or move it to a dead-letter queue,
	// depending on its configuration. This prevents message loss on transient failures.
	Nack(ctx context.Context) error
}

// MessageHandler is a callback function that processes messages received from a subscription.
// The handler is responsible for calling Ack or Nack on the message to signal its
// processing status. This explicit responsibility on the handler ensures that the
// business logic has full control over the message processing lifecycle.
type MessageHandler func(ctx context.Context, msg Message)

// Subscriber defines the interface for subscribing to subjects on a message bus.
// While the primary goal is publishing, a complete abstraction includes subscribing
// for potential future use cases or auxiliary services within the ecosystem.
type Subscriber interface {
	// Subscribe creates a durable subscription to a given subject.
	// Messages received on the subject will be passed to the provided MessageHandler.
	// The subscription will run in the background until the Subscriber is closed
	// or the provided context is cancelled.
	Subscribe(ctx context.Context, subject string, handler MessageHandler) error

	// Close gracefully unsubscribes from all subjects and terminates the connection.
	// It should ensure that all in-flight messages are processed before returning.
	Close() error
}

// NoopPublisher is a mock implementation of the Publisher interface that does nothing.
// It is useful for testing or for environments where messaging is disabled, allowing
// the application to run without a live message bus connection.
type NoopPublisher struct{}

// NewNoopPublisher creates a new no-op publisher.
func NewNoopPublisher() *NoopPublisher {
	return &NoopPublisher{}
}

// Publish does nothing and always returns nil.
func (p *NoopPublisher) Publish(_ context.Context, _ string, _ []byte) error {
	return nil
}

// Close does nothing and always returns nil.
func (p *NoopPublisher) Close() error {
	return nil
}

// NoopSubscriber is a mock implementation of the Subscriber interface that does nothing.
// It is useful for testing components that require a Subscriber dependency.
type NoopSubscriber struct{}

// NewNoopSubscriber creates a new no-op subscriber.
func NewNoopSubscriber() *NoopSubscriber {
	return &NoopSubscriber{}
}

// Subscribe does nothing and always returns nil. It will not call the handler.
func (s *NoopSubscriber) Subscribe(_ context.Context, _ string, _ MessageHandler) error {
	return nil
}

// Close does nothing and always returns nil.
func (s *NoopSubscriber) Close() error {
	return nil
}

### END_OF_FILE_COMPLETED ###
```