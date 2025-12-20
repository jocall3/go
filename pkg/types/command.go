```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// CommandID is a unique, content-addressable identifier for a Command.
// It ensures idempotency and provides a stable reference for auditing.
// Typically, this is a cryptographic hash of the command's serialized data.
type CommandID string

// CommandType is a string identifier that categorizes a command.
// It is used for routing commands to their appropriate handlers within the
// execution engine.
type CommandType string

// Command represents a verifiable, auditable intent to change the system's state.
//
// Commands are the fundamental unit of change within the system. They are distinct
// from Events, which represent the outcome of a state change. A Command is a
// request ("please do this"), while an Event is a fact ("this happened").
//
// This distinction is critical for building a system that is replayable, auditable,
// and resilient. The execution engine processes Commands to produce Events.
//
// Invariants for all Commands:
// 1. Deterministic: Given the same command and the same initial state, the
//    resulting state change and events must be identical.
// 2. Idempotent: Applying the same command multiple times should have the same
//    effect as applying it once. This is typically achieved via content-addressable
//    CommandIDs.
// 3. Serializable: Commands must be fully serializable to a stable format
//    (e.g., JSON, Protobuf) for storage, transmission, and auditing.
// 4. Self-Contained: A command should contain all the information necessary for
//    its own validation and execution, without relying on external context that
//    may change over time (e.g., session state).
type Command interface {
	// CommandID returns a unique, deterministic identifier for the command.
	// This is typically a content-based hash of the command's data, which
	// is crucial for ensuring idempotency. The system can use this ID to
	// detect and discard duplicate submissions.
	CommandID() CommandID

	// Type returns a string identifier for the command's type. This is used
	// by the execution engine to dispatch the command to the correct handler.
	// e.g., "CREATE_ACCOUNT", "POST_TRANSFER".
	Type() CommandType

	// Validate performs stateless validation checks on the command's data.
	// This ensures the command is well-formed and internally consistent before
	// it is considered for execution. This check should not depend on the
	// current state of the system (e.g., checking account balances).
	// State-dependent validation occurs during command handling.
	Validate() error

	// Target returns the primary identifier of the entity this command acts upon.
	// For example, for a "DEBIT_ACCOUNT" command, this would be the account ID.
	// This is essential for routing, sharding, and ensuring that commands
	// related to the same entity are processed sequentially.
	// An empty string can be returned if the command is not entity-specific
	// (e.g., a system-level configuration change).
	Target() string
}


```