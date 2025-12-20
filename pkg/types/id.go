```go
// Copyright (c) 2023-2024 The Bridge
//
// This file is part of the Bridge project.
// See the LICENSE file at the root of the project for licensing information.
//
//
//  pkg/types/id.go
//
// This file defines strongly-typed identifiers used throughout the system.
// By creating distinct types for different kinds of IDs (e.g., AccountID,
// TransactionID), we leverage the Go compiler to prevent accidental misuse,
// such as passing a TransactionID where an AccountID is expected.
//
// All identifiers are based on UUIDs to ensure global uniqueness without
// requiring a central coordinating service.

package types

import (
	"fmt"
	"github.com/google/uuid"
)

// idValidationError is a helper to create a consistent error message for invalid IDs.
func idValidationError(typeName string, idValue string, err error) error {
	return fmt.Errorf("invalid %s: value '%s' is not a valid UUID: %w", typeName, idValue, err)
}

// --- RequestID ---

// RequestID is a unique identifier for an API request, used for idempotency and tracing.
type RequestID string

// NewRequestID generates a new, unique RequestID.
func NewRequestID() RequestID {
	return RequestID(uuid.NewString())
}

// ParseRequestID converts a string into a RequestID, validating its format.
func ParseRequestID(s string) (RequestID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("RequestID", s, err)
	}
	return RequestID(s), nil
}

// String returns the string representation of the RequestID.
func (id RequestID) String() string {
	return string(id)
}

// Validate checks if the RequestID has a valid format.
func (id RequestID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("RequestID", id.String(), err)
	}
	return nil
}

// --- UserID ---

// UserID is a unique identifier for a user or principal in the system.
type UserID string

// NewUserID generates a new, unique UserID.
func NewUserID() UserID {
	return UserID(uuid.NewString())
}

// ParseUserID converts a string into a UserID, validating its format.
func ParseUserID(s string) (UserID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("UserID", s, err)
	}
	return UserID(s), nil
}

// String returns the string representation of the UserID.
func (id UserID) String() string {
	return string(id)
}

// Validate checks if the UserID has a valid format.
func (id UserID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("UserID", id.String(), err)
	}
	return nil
}

// --- AccountID ---

// AccountID is a unique identifier for a financial account.
// An account holds balances in one or more instruments.
type AccountID string

// NewAccountID generates a new, unique AccountID.
func NewAccountID() AccountID {
	return AccountID(uuid.NewString())
}

// ParseAccountID converts a string into an AccountID, validating its format.
func ParseAccountID(s string) (AccountID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("AccountID", s, err)
	}
	return AccountID(s), nil
}

// String returns the string representation of the AccountID.
func (id AccountID) String() string {
	return string(id)
}

// Validate checks if the AccountID has a valid format.
func (id AccountID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("AccountID", id.String(), err)
	}
	return nil
}

// --- LedgerID ---

// LedgerID is a unique identifier for a ledger.
// A ledger records all entries for a specific class of transactions or instrument.
type LedgerID string

// NewLedgerID generates a new, unique LedgerID.
func NewLedgerID() LedgerID {
	return LedgerID(uuid.NewString())
}

// ParseLedgerID converts a string into a LedgerID, validating its format.
func ParseLedgerID(s string) (LedgerID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("LedgerID", s, err)
	}
	return LedgerID(s), nil
}

// String returns the string representation of the LedgerID.
func (id LedgerID) String() string {
	return string(id)
}

// Validate checks if the LedgerID has a valid format.
func (id LedgerID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("LedgerID", id.String(), err)
	}
	return nil
}

// --- InstrumentID ---

// InstrumentID is a unique identifier for a financial instrument (e.g., currency, stock).
// While often represented by a ticker (e.g., "USD", "BTC"), a UUID ensures global uniqueness
// across different asset classes and avoids naming collisions.
type InstrumentID string

// NewInstrumentID generates a new, unique InstrumentID.
func NewInstrumentID() InstrumentID {
	return InstrumentID(uuid.NewString())
}

// ParseInstrumentID converts a string into an InstrumentID, validating its format.
func ParseInstrumentID(s string) (InstrumentID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("InstrumentID", s, err)
	}
	return InstrumentID(s), nil
}

// String returns the string representation of the InstrumentID.
func (id InstrumentID) String() string {
	return string(id)
}

// Validate checks if the InstrumentID has a valid format.
func (id InstrumentID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("InstrumentID", id.String(), err)
	}
	return nil
}

// --- TransactionID ---

// TransactionID is a unique identifier for a set of balanced ledger entries.
// A transaction represents a single, atomic financial event.
type TransactionID string

// NewTransactionID generates a new, unique TransactionID.
func NewTransactionID() TransactionID {
	return TransactionID(uuid.NewString())
}

// ParseTransactionID converts a string into a TransactionID, validating its format.
func ParseTransactionID(s string) (TransactionID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("TransactionID", s, err)
	}
	return TransactionID(s), nil
}

// String returns the string representation of the TransactionID.
func (id TransactionID) String() string {
	return string(id)
}

// Validate checks if the TransactionID has a valid format.
func (id TransactionID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("TransactionID", id.String(), err)
	}
	return nil
}

// --- EntryID ---

// EntryID is a unique identifier for a single ledger entry (a debit or a credit).
type EntryID string

// NewEntryID generates a new, unique EntryID.
func NewEntryID() EntryID {
	return EntryID(uuid.NewString())
}

// ParseEntryID converts a string into an EntryID, validating its format.
func ParseEntryID(s string) (EntryID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("EntryID", s, err)
	}
	return EntryID(s), nil
}

// String returns the string representation of the EntryID.
func (id EntryID) String() string {
	return string(id)
}

// Validate checks if the EntryID has a valid format.
func (id EntryID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("EntryID", id.String(), err)
	}
	return nil
}

// --- OrderID ---

// OrderID is a unique identifier for a trading order.
type OrderID string

// NewOrderID generates a new, unique OrderID.
func NewOrderID() OrderID {
	return OrderID(uuid.NewString())
}

// ParseOrderID converts a string into an OrderID, validating its format.
func ParseOrderID(s string) (OrderID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("OrderID", s, err)
	}
	return OrderID(s), nil
}

// String returns the string representation of the OrderID.
func (id OrderID) String() string {
	return string(id)
}

// Validate checks if the OrderID has a valid format.
func (id OrderID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("OrderID", id.String(), err)
	}
	return nil
}

// --- ExecutionID ---

// ExecutionID is a unique identifier for a trade execution, which is a fill
// or partial fill of an order.
type ExecutionID string

// NewExecutionID generates a new, unique ExecutionID.
func NewExecutionID() ExecutionID {
	return ExecutionID(uuid.NewString())
}

// ParseExecutionID converts a string into an ExecutionID, validating its format.
func ParseExecutionID(s string) (ExecutionID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("ExecutionID", s, err)
	}
	return ExecutionID(s), nil
}

// String returns the string representation of the ExecutionID.
func (id ExecutionID) String() string {
	return string(id)
}

// Validate checks if the ExecutionID has a valid format.
func (id ExecutionID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("ExecutionID", id.String(), err)
	}
	return nil
}

// --- PositionID ---

// PositionID is a unique identifier for a user's net position in a given instrument.
type PositionID string

// NewPositionID generates a new, unique PositionID.
func NewPositionID() PositionID {
	return PositionID(uuid.NewString())
}

// ParsePositionID converts a string into a PositionID, validating its format.
func ParsePositionID(s string) (PositionID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return "", idValidationError("PositionID", s, err)
	}
	return PositionID(s), nil
}

// String returns the string representation of the PositionID.
func (id PositionID) String() string {
	return string(id)
}

// Validate checks if the PositionID has a valid format.
func (id PositionID) Validate() error {
	_, err := uuid.Parse(id.String())
	if err != nil {
		return idValidationError("PositionID", id.String(), err)
	}
	return nil
}
### END_OF_FILE_COMPLETED ###
```