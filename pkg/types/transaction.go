```go
// Copyright (c) 2023-2024 The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TransactionStatus represents the state of a transaction in its lifecycle.
type TransactionStatus string

const (
	// StatusPending indicates that the transaction has been created but not yet processed.
	// It is not yet reflected in any account balances.
	StatusPending TransactionStatus = "PENDING"

	// StatusPosted indicates that the transaction has been successfully processed and its
	// entries have been applied to the respective accounts. This is a terminal state.
	StatusPosted TransactionStatus = "POSTED"

	// StatusFailed indicates that the transaction could not be processed due to an error
	// (e.g., insufficient funds, invalid account). This is a terminal state.
	StatusFailed TransactionStatus = "FAILED"

	// StatusCancelled indicates that the transaction was intentionally voided before posting.
	// This is a terminal state.
	StatusCancelled TransactionStatus = "CANCELLED"
)

// IsTerminal returns true if the status is a final state from which no further transitions are expected.
func (s TransactionStatus) IsTerminal() bool {
	switch s {
	case StatusPosted, StatusFailed, StatusCancelled:
		return true
	default:
		return false
	}
}

// EntryDirection specifies whether an entry is a debit or a credit.
type EntryDirection string

const (
	// Debit represents a debit entry, which increases asset/expense accounts and decreases liability/equity/revenue accounts.
	Debit EntryDirection = "DEBIT"
	// Credit represents a credit entry, which decreases asset/expense accounts and increases liability/equity/revenue accounts.
	Credit EntryDirection = "CREDIT"
)

// Entry represents a single leg of a financial transaction.
// It records a debit or a credit to a specific account.
type Entry struct {
	// ID is the unique identifier for the entry.
	ID uuid.UUID `json:"id"`

	// TransactionID is the identifier of the transaction this entry belongs to.
	TransactionID uuid.UUID `json:"transaction_id"`

	// AccountID is the identifier of the account this entry affects.
	AccountID string `json:"account_id"`

	// Amount is the value of the entry. It is always a positive value.
	// The direction (debit or credit) determines its effect on the account balance.
	Amount decimal.Decimal `json:"amount"`

	// Direction specifies whether this entry is a debit or a credit.
	Direction EntryDirection `json:"direction"`

	// Metadata holds arbitrary key-value pairs of data associated with the entry.
	Metadata map[string]string `json:"metadata"`

	// CreatedAt is the timestamp when the entry was created.
	CreatedAt time.Time `json:"created_at"`
}

// Validate checks the integrity of the Entry's fields.
func (e *Entry) Validate() error {
	if e.ID == uuid.Nil {
		return fmt.Errorf("entry id is required")
	}
	if e.TransactionID == uuid.Nil {
		return fmt.Errorf("entry transaction_id is required")
	}
	if e.AccountID == "" {
		return fmt.Errorf("entry account_id is required")
	}
	if e.Amount.IsNegative() || e.Amount.IsZero() {
		return fmt.Errorf("entry amount must be a positive value, got %s", e.Amount.String())
	}
	switch e.Direction {
	case Debit, Credit:
		// valid
	default:
		return fmt.Errorf("entry direction is invalid: %s", e.Direction)
	}
	return nil
}

// SignedAmount returns the amount with its sign determined by the direction.
// For the purpose of balancing a transaction, we use the convention that
// credits are positive and debits are negative. The sum of all signed amounts
// in a valid transaction must be zero.
func (e *Entry) SignedAmount() decimal.Decimal {
	if e.Direction == Debit {
		return e.Amount.Neg()
	}
	return e.Amount
}

// Transaction represents a complete, balanced financial event.
// It is composed of a set of entries that must sum to zero, enforcing the
// principle of double-entry bookkeeping.
type Transaction struct {
	// ID is the unique identifier for the transaction.
	ID uuid.UUID `json:"id"`

	// IdempotencyKey is a client-provided key to ensure that the same transaction
	// is not processed multiple times.
	IdempotencyKey string `json:"idempotency_key"`

	// Entries is the collection of debits and credits that make up this transaction.
	Entries []Entry `json:"entries"`

	// Status is the current state of the transaction.
	Status TransactionStatus `json:"status"`

	// Metadata holds arbitrary key-value pairs of data associated with the transaction.
	Metadata map[string]string `json:"metadata"`

	// CreatedAt is the timestamp when the transaction was created.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt is the timestamp when the transaction was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// NewTransaction creates a new transaction with a given idempotency key and entries.
// It initializes the transaction with a new UUID, pending status, and timestamps.
// The provided entries should already be populated with their respective data,
// but their TransactionID and ID will be set by this function if not already present.
func NewTransaction(idempotencyKey string, entries []Entry, metadata map[string]string) (*Transaction, error) {
	if idempotencyKey == "" {
		return nil, fmt.Errorf("idempotency key is required")
	}

	now := time.Now().UTC()
	txID := uuid.New()

	// Assign transaction ID to all entries and generate entry IDs if missing.
	for i := range entries {
		entries[i].TransactionID = txID
		if entries[i].ID == uuid.Nil {
			entries[i].ID = uuid.New()
		}
		if entries[i].CreatedAt.IsZero() {
			entries[i].CreatedAt = now
		}
	}

	tx := &Transaction{
		ID:             txID,
		IdempotencyKey: idempotencyKey,
		Entries:        entries,
		Status:         StatusPending,
		Metadata:       metadata,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := tx.Validate(); err != nil {
		return nil, fmt.Errorf("new transaction failed validation: %w", err)
	}

	return tx, nil
}

// Validate checks the structural and logical integrity of the Transaction.
// It enforces the core invariants of a financial transaction.
func (t *Transaction) Validate() error {
	if t.ID == uuid.Nil {
		return fmt.Errorf("transaction id is required")
	}
	if t.IdempotencyKey == "" {
		return fmt.Errorf("transaction idempotency_key is required")
	}

	// Invariant: A transaction must have at least two entries for double-entry bookkeeping.
	if len(t.Entries) < 2 {
		return fmt.Errorf("transaction must have at least two entries")
	}

	// Invariant: The sum of all entries must be zero (debits must equal credits).
	balance := decimal.Zero
	for i, entry := range t.Entries {
		if err := entry.Validate(); err != nil {
			return fmt.Errorf("invalid entry at index %d: %w", i, err)
		}
		// Ensure entry belongs to this transaction.
		if entry.TransactionID != t.ID {
			return fmt.Errorf("entry at index %d has mismatched transaction_id", i)
		}
		balance = balance.Add(entry.SignedAmount())
	}

	if !balance.IsZero() {
		return fmt.Errorf("transaction entries do not balance: sum is %s", balance.String())
	}

	return nil
}

// IsBalanced checks if the transaction's entries sum to zero.
// This is a convenience method to check the primary invariant.
func (t *Transaction) IsBalanced() bool {
	balance := decimal.Zero
	for _, entry := range t.Entries {
		balance = balance.Add(entry.SignedAmount())
	}
	return balance.IsZero()
}

```