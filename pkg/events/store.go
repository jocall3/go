```go
// Copyright (c) 2024. The Bridge Project Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"context"
	"errors"
)

// Sequence represents the monotonic, unique, and ordered identifier for an event in the store.
// It is the primary key for the event log, ensuring a total ordering of all state changes
// across the entire system. This is fundamental for deterministic replay and auditability.
type Sequence uint64

// Store is the core interface for an append-only event log.
// It provides the fundamental contract for persisting and retrieving ordered events,
// which is the foundation for replayable, inspectable, and auditable systems.
// This abstraction allows the underlying persistence mechanism to be swapped without
// affecting the business logic that relies on the event stream.
//
// Implementations could range from in-memory stores for testing to distributed
// databases like PostgreSQL, TiDB, or specialized event stores for production.
type Store interface {
	// Append adds a new event to the store. It must guarantee atomicity and
	// assign a new, unique, and monotonically increasing Sequence number.
	//
	// This method is the single point of entry for all state changes in the system.
	// Implementations must handle concurrency control. A common pattern is to use
	// optimistic locking based on an expected aggregate version to prevent
	// conflicting writes, returning ErrConflict if the state has changed since read.
	//
	// The returned Sequence is the globally unique identifier assigned to the
	// persisted event, confirming its position in the total order.
	Append(ctx context.Context, event Event) (Sequence, error)

	// Load retrieves a sequence of events from the global log, starting from the
	// event immediately after the given 'after' sequence number.
	//
	// 'after' specifies the exclusive starting point. To load from the beginning
	// of the log, a caller should use a sequence number of 0.
	//
	// 'limit' constrains the maximum number of events returned in a single call.
	// This is crucial for managing memory and network load during system-wide
	// replays, projections, or audits. Implementations are encouraged to enforce
	// a sane maximum even if the caller provides a large or zero value.
	Load(ctx context.Context, after Sequence, limit uint64) ([]Event, error)

	// LoadByAggregate retrieves all events for a specific aggregate instance,
	// starting after a given sequence number within that aggregate's stream.
	// This is the primary query pattern for reconstituting the state of an
	// aggregate root from its history.
	//
	// 'aggregateID' is the unique identifier of the aggregate.
	// 'after' specifies the exclusive starting point (version) within the aggregate's
	// own event stream. This is useful for loading events since the last snapshot.
	// A value of 0 will load the aggregate's entire history.
	LoadByAggregate(ctx context.Context, aggregateID string, after Sequence) ([]Event, error)
}

// Pre-defined errors for store operations.
// These allow consumers of the Store interface to handle specific failure modes
// in a standardized way, regardless of the underlying implementation. This promotes
// predictable, fail-closed behavior in the layers above.
var (
	// ErrConflict indicates that an event could not be appended due to a
	// concurrency conflict, such as an optimistic lock failure. This typically
	// means the state of an aggregate has changed since it was last read.
	// The calling operation should typically be retried after reloading the
	// aggregate's state and re-evaluating the command.
	ErrConflict = errors.New("events: concurrency conflict")

	// ErrNotFound indicates that a requested entity (like an aggregate or a
	// specific event sequence) does not exist in the store.
	ErrNotFound = errors.New("events: not found")
)


```