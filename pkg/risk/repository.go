```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package risk

import (
	"context"
	"errors"
	"time"
)

// --- Placeholder Types ---
// In a real application, these types would be defined in their own files
// (e.g., policy.go, model.go) with full details and business logic.

// PolicyID uniquely identifies a risk policy.
type PolicyID string

// ModelID uniquely identifies a risk model configuration.
type ModelID string

// Version is used for optimistic locking to ensure data consistency.
type Version uint64

// Policy defines a set of risk rules and limits that govern trading activity.
// It is versioned to provide a full audit trail of changes.
type Policy struct {
	ID          PolicyID
	Name        string
	Description string
	// Limits is a collection of specific risk limits under this policy.
	Limits []Limit
	// Version is used for optimistic locking to prevent stale writes.
	// The persistence layer is responsible for incrementing this on update.
	Version Version
	// CreatedAt is the timestamp when the policy was first created.
	CreatedAt time.Time
	// UpdatedAt is the timestamp of the last update.
	UpdatedAt time.Time
	// EffectiveFrom is the timestamp from which the policy is active.
	EffectiveFrom time.Time
	// IsArchived indicates if the policy is no longer in use.
	IsArchived bool
}

// Limit defines a specific quantitative risk constraint within a policy.
type Limit struct {
	ID          string
	Type        string // e.g., "MaxPositionValue", "MaxDailyLoss", "CounterpartyExposure"
	Value       float64
	Currency    string
	Timeframe   time.Duration // e.g., 24h for a daily limit
	Dimensions  map[string]string // e.g., {"instrument": "BTC-USD", "account": "X123"}
}

// ModelConfiguration holds the parameters and settings for a risk model.
// This allows for tuning and swapping models without code changes.
type ModelConfiguration struct {
	ID         ModelID
	Name       string
	Type       string // e.g., "VaR", "StressTest", "ScenarioAnalysis"
	Parameters map[string]string
	// Version is used for optimistic locking.
	Version Version
	// CreatedAt is the timestamp when the configuration was first created.
	CreatedAt time.Time
	// UpdatedAt is the timestamp of the last update.
	UpdatedAt time.Time
	// IsActive indicates if the model configuration is currently in use.
	IsActive bool
}

// Exposure represents a point-in-time snapshot of a calculated risk value.
// These are persisted for auditing, reporting, and historical analysis.
type Exposure struct {
	// Key uniquely identifies the subject of the exposure (e.g., "Portfolio:A", "Counterparty:B").
	Key string
	// Value is the calculated risk exposure amount.
	Value float64
	// Timestamp is when the exposure was calculated.
	Timestamp time.Time
	// ModelID is the ID of the model used for the calculation.
	ModelID ModelID
	// PolicyID is the ID of the policy in effect during the calculation.
	PolicyID PolicyID
}

// --- End Placeholder Types ---

// Common repository errors.
var (
	// ErrNotFound indicates that a requested entity was not found.
	ErrNotFound = errors.New("entity not found")

	// ErrConflict indicates a write conflict, such as an optimistic locking failure
	// or a unique constraint violation.
	ErrConflict = errors.New("entity conflict")
)

// PolicyRepository defines the interface for persisting risk policies.
// Policies are versioned to ensure auditability and prevent race conditions.
type PolicyRepository interface {
	// SavePolicy creates a new policy or updates an existing one.
	// If the policy is new (e.g., Version is 0), it's created.
	// If the policy exists, the provided version must match the stored version
	// to prevent stale writes (optimistic locking).
	// On successful update, the policy's version is incremented by the persistence layer.
	// It should return ErrConflict on a version mismatch.
	SavePolicy(ctx context.Context, policy *Policy) error

	// FindPolicyByID retrieves the latest version of a policy by its ID.
	// It should return ErrNotFound if the policy does not exist.
	FindPolicyByID(ctx context.Context, id PolicyID) (*Policy, error)

	// FindActivePolicies retrieves all policies that are currently active at a given time.
	FindActivePolicies(ctx context.Context, at time.Time) ([]*Policy, error)
}

// ModelRepository defines the interface for persisting risk model configurations.
// Model configurations are also versioned to track changes over time.
type ModelRepository interface {
	// SaveModelConfiguration creates or updates a risk model's configuration.
	// It employs optimistic locking based on the configuration's version.
	// It should return ErrConflict on a version mismatch.
	SaveModelConfiguration(ctx context.Context, config *ModelConfiguration) error

	// FindModelConfigurationByID retrieves the latest version of a model configuration.
	// It should return ErrNotFound if the configuration does not exist.
	FindModelConfigurationByID(ctx context.Context, id ModelID) (*ModelConfiguration, error)

	// FindActiveModelConfigurations retrieves all model configurations that are currently active.
	FindActiveModelConfigurations(ctx context.Context) ([]*ModelConfiguration, error)
}

// ExposureRepository defines the interface for persisting risk exposure snapshots.
// Exposures are typically high-volume data, so the interface focuses on writing
// snapshots and retrieving historical data for analysis and audit.
type ExposureRepository interface {
	// SaveExposureSnapshot records a point-in-time snapshot of a calculated risk exposure.
	// This is an append-only operation for auditability.
	SaveExposureSnapshot(ctx context.Context, snapshot *Exposure) error

	// FindLatestExposureSnapshot retrieves the most recent exposure snapshot for a given key.
	// The key could represent an instrument, a counterparty, or a portfolio.
	// It should return ErrNotFound if no snapshot exists for the key.
	FindLatestExposureSnapshot(ctx context.Context, key string) (*Exposure, error)

	// FindExposureHistory retrieves a series of exposure snapshots over a specified time range.
	// This is used for historical analysis, reporting, and debugging.
	FindExposureHistory(ctx context.Context, key string, start, end time.Time) ([]*Exposure, error)
}

// Repository is a composite interface that combines all risk-related persistence operations.
// This allows for implementations that can manage transactions across different
// risk entities (e.g., updating a policy and a model configuration atomically).
type Repository interface {
	PolicyRepository
	ModelRepository
	ExposureRepository

	// RunInTransaction executes the given function within a single atomic transaction.
	// If the function returns an error, the transaction is rolled back. Otherwise, it's committed.
	// The provided repository instance within the function should be used for all
	// database operations to ensure they are part of the same transaction.
	RunInTransaction(ctx context.Context, fn func(txRepo Repository) error) error
}

```