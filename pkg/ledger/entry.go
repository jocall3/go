```go
package ledger

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Layer represents an accounting layer within an account's balance.
// This allows for tracking funds in different states (e.g., settled, pending).
type Layer string

const (
	// SettledLayer represents funds that are fully cleared and available.
	SettledLayer Layer = "SETTLED"
	// PendingLayer represents funds that are part of an in-flight transaction.
	PendingLayer Layer = "PENDING"
	// OnHoldLayer represents funds that are reserved for a specific purpose (e.g., collateral).
	OnHoldLayer Layer = "ON_HOLD"
)

// Validate checks if the Layer value is one of the predefined constants.
func (l Layer) Validate() error {
	switch l {
	case SettledLayer, PendingLayer, OnHoldLayer:
		return nil
	default:
		return fmt.Errorf("invalid layer: %s", l)
	}
}

// Direction indicates whether an entry is a debit or a credit.
type Direction string

const (
	// Debit represents a debit entry, which typically increases an asset account or decreases a liability account.
	Debit Direction = "DEBIT"
	// Credit represents a credit entry, which typically decreases an asset account or increases a liability account.
	Credit Direction = "CREDIT"
)

// Validate checks if the Direction value is one of the predefined constants.
func (d Direction) Validate() error {
	switch d {
	case Debit, Credit:
		return nil
	default:
		return fmt.Errorf("invalid direction: %s", d)
	}
}

// TransactionStatus represents the lifecycle state of a transaction.
type TransactionStatus string

const (
	// StatusPending is the initial state of a transaction before it is processed.
	StatusPending TransactionStatus = "PENDING"
	// StatusPosted indicates the transaction has been successfully processed and its entries are applied.
	StatusPosted TransactionStatus = "POSTED"
	// StatusRejected indicates the transaction was rejected due to validation or business rule failure.
	StatusRejected TransactionStatus = "REJECTED"
)

// Entry is the atomic unit of a transaction, representing a single debit or credit.
// It is immutable after creation.
type Entry struct {
	ID            uuid.UUID
	TransactionID uuid.UUID
	AccountID     uuid.UUID
	Direction     Direction
	Amount        decimal.Decimal
	Layer         Layer
	CreatedAt     time.Time
	Description   string
}

// NewEntry creates and validates a new Entry.
func NewEntry(transactionID, accountID uuid.UUID, direction Direction, amount decimal.Decimal, layer Layer, description string) (*Entry, error) {
	if accountID == uuid.Nil {
		return nil, errors.New("entry requires a valid account id")
	}
	if err := direction.Validate(); err != nil {
		return nil, err
	}
	if err := layer.Validate(); err != nil {
		return nil, err
	}
	if amount.IsNegative() || amount.IsZero() {
		return nil, errors.New("entry amount must be positive")
	}

	return &Entry{
		ID:            uuid.New(),
		TransactionID: transactionID,
		AccountID:     accountID,
		Direction:     direction,
		Amount:        amount,
		Layer:         layer,
		CreatedAt:     time.Now().UTC(),
		Description:   description,
	}, nil
}

// Transaction is a collection of balanced debit and credit entries.
// It is the sole unit of state change in the ledger. A transaction is atomic;
// either all its entries are applied, or none are.
type Transaction struct {
	ID             uuid.UUID
	IdempotencyKey string
	Entries        []*Entry
	Status         TransactionStatus
	CreatedAt      time.Time
	PostedAt       *time.Time // Pointer to allow for null
	Description    string
}

// NewTransaction creates a new transaction in a pending state.
func NewTransaction(idempotencyKey, description string) *Transaction {
	return &Transaction{
		ID:             uuid.New(),
		IdempotencyKey: idempotencyKey,
		Entries:        make([]*Entry, 0),
		Status:         StatusPending,
		CreatedAt:      time.Now().UTC(),
		PostedAt:       nil,
		Description:    description,
	}
}

// AddEntry creates a new entry and adds it to the transaction.
func (t *Transaction) AddEntry(accountID uuid.UUID, direction Direction, amount decimal.Decimal, layer Layer, description string) error {
	if t.Status != StatusPending {
		return fmt.Errorf("cannot add entry to transaction with status: %s", t.Status)
	}
	entry, err := NewEntry(t.ID, accountID, direction, amount, layer, description)
	if err != nil {
		return fmt.Errorf("failed to create new entry: %w", err)
	}
	t.Entries = append(t.Entries, entry)
	return nil
}

// Validate checks the internal consistency and invariants of the transaction.
// The primary invariant is that the sum of debits must equal the sum of credits.
func (t *Transaction) Validate() error {
	if len(t.Entries) < 2 {
		return errors.New("transaction must have at least two entries")
	}

	debitSum := decimal.Zero
	creditSum := decimal.Zero

	for i, entry := range t.Entries {
		if entry == nil {
			return fmt.Errorf("entry at index %d is nil", i)
		}
		if entry.TransactionID != t.ID {
			return fmt.Errorf("entry %s has mismatched transaction id %s (expected %s)", entry.ID, entry.TransactionID, t.ID)
		}
		// NewEntry already validates this, but this is a defense-in-depth check.
		if entry.Amount.IsNegative() || entry.Amount.IsZero() {
			return fmt.Errorf("entry %s has a non-positive amount: %s", entry.ID, entry.Amount.String())
		}

		switch entry.Direction {
		case Debit:
			debitSum = debitSum.Add(entry.Amount)
		case Credit:
			creditSum = creditSum.Add(entry.Amount)
		default:
			return fmt.Errorf("entry %s has an invalid direction: %s", entry.ID, entry.Direction)
		}
	}

	if debitSum.IsZero() {
		return errors.New("transaction total value must be greater than zero")
	}

	if !debitSum.Equal(creditSum) {
		return fmt.Errorf("transaction is unbalanced: debits (%s) != credits (%s)", debitSum.String(), creditSum.String())
	}

	return nil
}

// Post marks the transaction as posted, setting its posted timestamp.
// This should only be called after successful validation and persistence.
func (t *Transaction) Post() error {
	if t.Status != StatusPending {
		return fmt.Errorf("cannot post transaction with status: %s", t.Status)
	}
	if err := t.Validate(); err != nil {
		return fmt.Errorf("cannot post invalid transaction: %w", err)
	}
	now := time.Now().UTC()
	t.PostedAt = &now
	t.Status = StatusPosted
	return nil
}

// Reject marks the transaction as rejected.
func (t *Transaction) Reject() error {
	if t.Status != StatusPending {
		return fmt.Errorf("cannot reject transaction with status: %s", t.Status)
	}
	t.Status = StatusRejected
	return nil
}

```