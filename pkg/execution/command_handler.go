```go
// Copyright (c) 2023-2024 The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package execution

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Command represents an intent to change the system state. It is the fundamental
// unit of work for the execution engine. Each command is an immutable data
// structure that carries all the necessary information for its processing.
//
// Commands are distinct from Events. A Command is a request to do something,
// whereas an Event is a record of something that has already happened.
type Command interface {
	// CommandID returns the unique identifier for this command instance.
	// This is crucial for idempotency, tracing, and auditability. A command
	// processor can use this ID to ensure a command is not processed more than once.
	CommandID() uuid.UUID

	// CommandType returns a string identifier for the type of command.
	// This is used for routing the command to the correct handler.
	// e.g., "CREATE_ACCOUNT", "SUBMIT_ORDER", "PROCESS_DEPOSIT".
	CommandType() string
}

// Event represents a factual, immutable record of something that has happened
// in the system as a result of a command being successfully processed. Events are
// the source of truth for the system's state.
//
// NOTE: This is a high-level interface. A concrete implementation would be defined
// in a dedicated `pkg/events` package with a more detailed structure, including
// versioning, timestamps, and payload serialization.
type Event interface {
	// EventID returns the unique identifier for the event.
	EventID() uuid.UUID
	// EventType returns the type of the event.
	EventType() string
	// AggregateID returns the ID of the entity (e.g., Account, Order) this event pertains to.
	AggregateID() uuid.UUID
}

// CommandHandler defines the interface for processing a specific type of command.
// Each concrete implementation of CommandHandler encapsulates the business logic
// and validation rules for a single command type. This promotes the Single
// Responsibility Principle and makes the system easier to test, maintain, and reason about.
//
// The Handle method is expected to be idempotent. If the same command is
// processed multiple times, the system state must remain consistent with it
// having been processed only once. Implementations typically achieve this by
// checking for the existence of results from the command's ID before executing.
type CommandHandler interface {
	// Handle processes the given command.
	// It takes a context for cancellation and deadlines, and the command to be executed.
	// It returns a slice of events that were generated as a result of processing
	// the command, or an error if the command could not be processed due to
	// validation failures, invariant violations, or other business logic constraints.
	//
	// The returned events represent the state changes that occurred. If an error is
	// returned, it is guaranteed that no state has changed and no events were generated.
	// This fail-closed semantic is critical for system stability.
	Handle(ctx context.Context, cmd Command) ([]Event, error)
}

// CommandHandlerMap is a registry for mapping command types to their handlers.
// This allows a central dispatcher (e.g., a CommandBus) to route commands
// dynamically without being coupled to specific handler implementations.
type CommandHandlerMap map[string]CommandHandler

// NewCommandHandlerMap creates a new, empty command handler map.
func NewCommandHandlerMap() CommandHandlerMap {
	return make(CommandHandlerMap)
}

// Register associates a command type with a specific handler.
// It will panic if a handler is already registered for the given command type.
// This strict, fail-fast behavior ensures that command routing is unambiguous
// and is configured correctly at application startup, preventing runtime errors.
func (m CommandHandlerMap) Register(commandType string, handler CommandHandler) {
	if _, exists := m[commandType]; exists {
		panic(fmt.Sprintf("command handler already registered for type: %s", commandType))
	}
	if handler == nil {
		panic(fmt.Sprintf("cannot register a nil handler for type: %s", commandType))
	}
	m[commandType] = handler
}

// Get retrieves the handler for a given command type.
// It returns the handler and a boolean indicating if a handler was found. This
// allows the caller to safely handle cases where a command is dispatched for
// which no handler has been registered.
func (m CommandHandlerMap) Get(commandType string) (CommandHandler, bool) {
	handler, found := m[commandType]
	return handler, found
}
### END_OF_FILE_COMPLETED ###
```