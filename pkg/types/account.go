```go
// Package types defines the core domain models and data structures used throughout the financial system.
// These types represent fundamental concepts like accounts, transactions, and ledgers,
// and are designed to be explicit, deterministic, and safe for concurrent use.
package types

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AccountStatus represents the lifecycle state of an account.
// The status dictates which operations are permissible on the account, forming a
// core part of the system's state machine and safety invariants.
type AccountStatus string

const (
	// AccountStatusPending indicates the account has been created but is not yet active.
	// This state is typically used for accounts undergoing review or setup processes.
	// No financial transactions are allowed in this state.
	AccountStatusPending AccountStatus = "PENDING"

	// AccountStatusOpen indicates the account is active and can participate in transactions.
	// This is the normal operational state for an account.
	AccountStatusOpen AccountStatus = "OPEN"

	// AccountStatusFrozen indicates the account is temporarily suspended.
	// Existing funds are held, but no new debits or credits are permitted.
	// This state is typically used for compliance, risk mitigation, or administrative actions.
	AccountStatusFrozen AccountStatus = "FROZEN"

	// AccountStatusClosed indicates the account has been permanently closed.
	// The balance must be zero, and no further activity is possible. This is a terminal state.
	AccountStatusClosed AccountStatus = "CLOSED"
)

// IsValid checks if the AccountStatus is one of the predefined valid statuses.
// This is useful for input validation and ensuring state integrity.
func (s AccountStatus) IsValid() bool {
	switch s {
	case AccountStatusPending, AccountStatusOpen, AccountStatusFrozen, AccountStatusClosed:
		return true
	default:
		return false
	}
}

// Account represents a single, uniquely identifiable store of value for a specific
// currency. It is the fundamental building block of the ledger system.
//
// Invariants:
// 1. The Balance can never be negative. This is enforced at the transaction execution layer.
// 2. An Account's Currency is immutable once created.
// 3. A closed account must have a zero balance. This is enforced by the account closing process.
// 4. State transitions must follow a defined lifecycle (e.g., cannot move from CLOSED to OPEN).
type Account struct {
	// ID is the unique, immutable identifier for the account (UUID v4).
	ID uuid.UUID `json:"id" db:"id"`

	// OwnerID identifies the entity (e.g., customer, internal desk) that owns this account.
	OwnerID uuid.UUID `json:"owner_id" db:"owner_id"`

	// Currency is the code for the asset held in this account (e.g., "USD", "BTC").
	// This is immutable after account creation. It should conform to a known set of
	// currencies managed by the system.
	Currency string `json:"currency" db:"currency"`

	// Status represents the current state of the account, governing permissible actions.
	Status AccountStatus `json:"status" db:"status"`

	// Balance is the amount of the currency held in the account.
	// It uses a high-precision decimal type to prevent floating-point errors, which is
	// non-negotiable for financial systems.
	Balance decimal.Decimal `json:"balance" db:"balance"`

	// Version is an integer used for optimistic concurrency control.
	// It is incremented on every state change, ensuring that updates are not
	// based on stale data. Any attempt to update an account with an incorrect
	// version number will be rejected.
	Version int64 `json:"version" db:"version"`

	// CreatedAt is the timestamp when the account was created in the system.
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// UpdatedAt is the timestamp of the last modification to the account.
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	// Metadata is a flexible field for storing additional, non-critical information
	// as a JSON object. This can be used for application-specific data (e.g., labels,
	// external system references) without altering the core account schema.
	Metadata json.RawMessage `json:"metadata" db:"metadata"`
}

// NewAccount creates a new account instance in a valid initial state.
// By default, a new account is PENDING with a zero balance.
func NewAccount(ownerID uuid.UUID, currency string) *Account {
	now := time.Now().UTC()
	return &Account{
		ID:        uuid.New(),
		OwnerID:   ownerID,
		Currency:  currency,
		Status:    AccountStatusPending,
		Balance:   decimal.Zero,
		Version:   1,
		CreatedAt: now,
		UpdatedAt: now,
		Metadata:  json.RawMessage("{}"),
	}
}

// CanTransact returns true if the account is in a state that allows for
// debits or credits. This is a critical check before initiating any transaction.
// This enforces a core system invariant: only open accounts can transact.
func (a *Account) CanTransact() bool {
	return a.Status == AccountStatusOpen
}

// IsClosed returns true if the account is in a terminal, permanently closed state.
func (a *Account) IsClosed() bool {
	return a.Status == AccountStatusClosed
}
### END_OF_FILE_COMPLETED ###
```