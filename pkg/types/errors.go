```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"fmt"
	"strings"
)

// --- ErrInvariantViolation ---

// ErrInvariantViolation indicates a critical, unrecoverable state inconsistency.
// This error signals a bug or data corruption and should lead to a system halt
// or safe mode. It represents a breach of a fundamental rule that must always hold true.
// For example, if the sum of debits and credits in a transaction does not equal zero.
type ErrInvariantViolation struct {
	// Invariant is a human-readable description of the rule that was broken.
	Invariant string
	// Details provides specific context about the violation, such as expected vs. actual values.
	Details map[string]any
	// cause is the underlying error, if any, that led to the violation.
	cause error
}

// NewErrInvariantViolation creates a new ErrInvariantViolation.
func NewErrInvariantViolation(invariant string, details map[string]any, cause error) *ErrInvariantViolation {
	return &ErrInvariantViolation{
		Invariant: invariant,
		Details:   details,
		cause:     cause,
	}
}

// Error implements the standard error interface.
func (e *ErrInvariantViolation) Error() string {
	var details []string
	if e.Details != nil {
		for k, v := range e.Details {
			details = append(details, fmt.Sprintf("%s: %v", k, v))
		}
	}
	msg := fmt.Sprintf("invariant violation: %s", e.Invariant)
	if len(details) > 0 {
		msg = fmt.Sprintf("%s [%s]", msg, strings.Join(details, ", "))
	}
	if e.cause != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.cause)
	}
	return msg
}

// Unwrap provides compatibility with errors.Is and errors.As, allowing it to be
// part of an error chain.
func (e *ErrInvariantViolation) Unwrap() error {
	return e.cause
}

// --- ErrConcurrencyConflict ---

// ErrConcurrencyConflict indicates that an operation failed due to a stale state,
// typically in an optimistic locking scenario. This is a transient error, and the
// operation can usually be retried.
type ErrConcurrencyConflict struct {
	// ResourceType is the type of the resource that had a conflict (e.g., "Account", "Order").
	ResourceType string
	// ResourceID is the unique identifier of the resource.
	ResourceID string
	// ExpectedVersion is the version of the resource the operation expected to find.
	ExpectedVersion uint64
	// ActualVersion is the version of the resource that was actually found.
	ActualVersion uint64
	// cause is the underlying error, if any.
	cause error
}

// NewErrConcurrencyConflict creates a new ErrConcurrencyConflict.
func NewErrConcurrencyConflict(resourceType, resourceID string, expected, actual uint64, cause error) *ErrConcurrencyConflict {
	return &ErrConcurrencyConflict{
		ResourceType:    resourceType,
		ResourceID:      resourceID,
		ExpectedVersion: expected,
		ActualVersion:   actual,
		cause:           cause,
	}
}

// Error implements the standard error interface.
func (e *ErrConcurrencyConflict) Error() string {
	msg := fmt.Sprintf(
		"concurrency conflict on %s %s: expected version %d, but found %d",
		e.ResourceType,
		e.ResourceID,
		e.ExpectedVersion,
		e.ActualVersion,
	)
	if e.cause != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.cause)
	}
	return msg
}

// Unwrap provides compatibility with errors.Is and errors.As.
func (e *ErrConcurrencyConflict) Unwrap() error {
	return e.cause
}

// --- ErrInputValidation ---

// ErrInputValidation indicates that the input provided for an operation is invalid.
// This is typically a client-side error and should result in a clear error
// message to the caller without causing a system panic.
type ErrInputValidation struct {
	// Field is the name of the field that failed validation. Can be a path for nested fields.
	Field string
	// Reason is the explanation for why the validation failed.
	Reason string
	// cause is the underlying error, if any.
	cause error
}

// NewErrInputValidation creates a new ErrInputValidation.
func NewErrInputValidation(field, reason string, cause error) *ErrInputValidation {
	return &ErrInputValidation{
		Field:  field,
		Reason: reason,
		cause:  cause,
	}
}

// Error implements the standard error interface.
func (e *ErrInputValidation) Error() string {
	msg := fmt.Sprintf("input validation failed for field '%s': %s", e.Field, e.Reason)
	if e.cause != nil {
		msg = fmt.Sprintf("%s: %v", msg, e.cause)
	}
	return msg
}

// Unwrap provides compatibility with errors.Is and errors.As.
func (e *ErrInputValidation) Unwrap() error {
	return e.cause
}

// --- ErrNotFound ---

// ErrNotFound indicates that a requested resource could not be found.
type ErrNotFound struct {
	// ResourceType is the type of the resource that was not found (e.g., "Account", "Transaction").
	ResourceType string
	// ResourceID is the unique identifier of the resource.
	ResourceID string
}

// NewErrNotFound creates a new ErrNotFound.
func NewErrNotFound(resourceType, resourceID string) *ErrNotFound {
	return &ErrNotFound{
		ResourceType: resourceType,
		ResourceID:   resourceID,
	}
}

// Error implements the standard error interface.
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s with ID '%s' not found", e.ResourceType, e.ResourceID)
}

// --- ErrInsufficientFunds ---

// ErrInsufficientFunds indicates that an operation could not be completed
// because an account lacks the required balance.
type ErrInsufficientFunds struct {
	// AccountID is the identifier of the account with insufficient funds.
	AccountID string
	// Requested is the amount that was requested for the operation.
	// Using string to support arbitrary precision decimal libraries.
	Requested string
	// Available is the amount that was available in the account.
	// Using string to support arbitrary precision decimal libraries.
	Available string
}

// NewErrInsufficientFunds creates a new ErrInsufficientFunds.
func NewErrInsufficientFunds(accountID, requested, available string) *ErrInsufficientFunds {
	return &ErrInsufficientFunds{
		AccountID: accountID,
		Requested: requested,
		Available: available,
	}
}

// Error implements the standard error interface.
func (e *ErrInsufficientFunds) Error() string {
	return fmt.Sprintf(
		"insufficient funds for account %s: requested %s, but only %s available",
		e.AccountID,
		e.Requested,
		e.Available,
	)
}

```