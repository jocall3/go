```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"context"
	"fmt"

	"github.com/shopspring/decimal"
)

// Invariant represents a condition that must always hold true for the system to be in a valid state.
// Invariants are enforced at runtime to prevent state corruption and ensure financial safety.
// They are a core component of the system's fail-closed semantics. If an invariant check fails,
// the operation is halted immediately.
type Invariant interface {
	// Name returns a human-readable, unique name for the invariant.
	// This is used for logging, monitoring, and error reporting.
	Name() string

	// Check evaluates the invariant. It returns nil if the invariant holds true,
	// and a descriptive error if it is violated. The context can be used for
	// cancellation or to pass request-scoped data if needed.
	Check(ctx context.Context) error
}

// Asset represents a financial asset, such as USD or BTC.
type Asset string

// BalanceHolder is an interface for any entity that holds a balance of one or more assets.
// This allows invariants to operate on different types of accounts (e.g., internal, user, omnibus)
// without being coupled to their concrete implementations.
type BalanceHolder interface {
	// ID returns the unique identifier for the balance holder.
	ID() string
	// Balance returns the current balance for a specific asset.
	Balance(asset Asset) decimal.Decimal
}

// -----------------------------------------------------------------------------
// Concrete Invariant Implementations
// -----------------------------------------------------------------------------

// PositiveBalanceInvariant ensures that an account's balance for a given asset is not negative.
// This is a fundamental safety check in any ledger system.
type PositiveBalanceInvariant struct {
	account BalanceHolder
	asset   Asset
}

// NewPositiveBalanceInvariant creates a new invariant to check for a non-negative balance.
func NewPositiveBalanceInvariant(account BalanceHolder, asset Asset) *PositiveBalanceInvariant {
	return &PositiveBalanceInvariant{
		account: account,
		asset:   asset,
	}
}

// Name returns the name of the invariant.
func (i *PositiveBalanceInvariant) Name() string {
	return "PositiveBalance"
}

// Check verifies that the account's balance is greater than or equal to zero.
func (i *PositiveBalanceInvariant) Check(_ context.Context) error {
	balance := i.account.Balance(i.asset)
	if balance.IsNegative() {
		return fmt.Errorf(
			"invariant violated: %s: negative balance for account %s, asset %s: %s",
			i.Name(),
			i.account.ID(),
			i.asset,
			balance.String(),
		)
	}
	return nil
}

// SufficientBalanceInvariant ensures that an account has enough funds for a debit operation.
type SufficientBalanceInvariant struct {
	account     BalanceHolder
	asset       Asset
	debitAmount decimal.Decimal
}

// NewSufficientBalanceInvariant creates a new invariant to check for sufficient funds.
func NewSufficientBalanceInvariant(account BalanceHolder, asset Asset, debitAmount decimal.Decimal) *SufficientBalanceInvariant {
	return &SufficientBalanceInvariant{
		account:     account,
		asset:       asset,
		debitAmount: debitAmount,
	}
}

// Name returns the name of the invariant.
func (i *SufficientBalanceInvariant) Name() string {
	return "SufficientBalance"
}

// Check verifies that the account's balance is greater than or equal to the amount being debited.
func (i *SufficientBalanceInvariant) Check(_ context.Context) error {
	balance := i.account.Balance(i.asset)
	if balance.LessThan(i.debitAmount) {
		return fmt.Errorf(
			"invariant violated: %s: insufficient balance for account %s, asset %s. Required: %s, Available: %s",
			i.Name(),
			i.account.ID(),
			i.asset,
			i.debitAmount.String(),
			balance.String(),
		)
	}
	return nil
}

// NonZeroAmountInvariant ensures that a transactional amount is strictly positive.
// This prevents zero-value or negative-value transfers which are typically meaningless or malicious.
type NonZeroAmountInvariant struct {
	amount    decimal.Decimal
	operation string // e.g., "transfer", "withdrawal", "payment"
}

// NewNonZeroAmountInvariant creates a new invariant to check for a positive transaction amount.
func NewNonZeroAmountInvariant(amount decimal.Decimal, operation string) *NonZeroAmountInvariant {
	return &NonZeroAmountInvariant{
		amount:    amount,
		operation: operation,
	}
}

// Name returns the name of the invariant.
func (i *NonZeroAmountInvariant) Name() string {
	return "NonZeroAmount"
}

// Check verifies that the amount is strictly greater than zero.
func (i *NonZeroAmountInvariant) Check(_ context.Context) error {
	if !i.amount.IsPositive() {
		return fmt.Errorf(
			"invariant violated: %s: non-positive amount for operation '%s': %s",
			i.Name(),
			i.operation,
			i.amount.String(),
		)
	}
	return nil
}

// -----------------------------------------------------------------------------
// Composite Invariant
// -----------------------------------------------------------------------------

// CompositeInvariant groups multiple invariants to be checked as a single unit.
// It follows fail-fast semantics: the first violated invariant halts the check.
type CompositeInvariant struct {
	name       string
	invariants []Invariant
}

// NewCompositeInvariant creates a new composite invariant.
func NewCompositeInvariant(name string, invariants ...Invariant) *CompositeInvariant {
	return &CompositeInvariant{
		name:       name,
		invariants: invariants,
	}
}

// Name returns the name of the composite invariant.
func (c *CompositeInvariant) Name() string {
	return c.name
}

// Check evaluates each contained invariant in order. It returns immediately
// upon the first failure.
func (c *CompositeInvariant) Check(ctx context.Context) error {
	for _, inv := range c.invariants {
		if err := inv.Check(ctx); err != nil {
			// Wrap the error to provide context from the composite check
			return fmt.Errorf("in composite invariant '%s': %w", c.name, err)
		}
	}
	return nil
}

// CheckInvariants is a helper function to evaluate a slice of invariants.
// It provides a convenient way to run multiple checks without manually creating
// a CompositeInvariant. It stops and returns on the first error encountered.
func CheckInvariants(ctx context.Context, invariants ...Invariant) error {
	for _, inv := range invariants {
		if err := inv.Check(ctx); err != nil {
			return err
		}
	}
	return nil
}

```