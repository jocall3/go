```go
// Copyright 2024 The Bridge Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ledger

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// AccountType represents the fundamental classification of an account in double-entry bookkeeping.
// The five standard types form the basis of the balance sheet and income statement.
// This strict enumeration enforces core accounting principles at the type level.
type AccountType string

// Defines the five core account types, ensuring that all ledger accounts conform
// to standard accounting practices.
const (
	AssetAccount     AccountType = "ASSET"
	LiabilityAccount AccountType = "LIABILITY"
	EquityAccount    AccountType = "EQUITY"
	RevenueAccount   AccountType = "REVENUE"
	ExpenseAccount   AccountType = "EXPENSE"
)

// IsValid checks if the AccountType is one of the predefined valid types.
// This function acts as a runtime invariant check, preventing the creation or
// processing of accounts with non-standard types.
func (at AccountType) IsValid() bool {
	switch at {
	case AssetAccount, LiabilityAccount, EquityAccount, RevenueAccount, ExpenseAccount:
		return true
	default:
		return false
	}
}

// NormalBalance indicates whether an account's balance is typically increased by a debit or a credit.
// This is a fundamental property derived from the account's type and is critical for
// maintaining the integrity of the accounting equation (Assets = Liabilities + Equity).
type NormalBalance string

const (
	DebitBalance  NormalBalance = "DEBIT"
	CreditBalance NormalBalance = "CREDIT"
)

// GetNormalBalance returns the normal balance for a given account type.
// This function enforces a core accounting rule deterministically, removing ambiguity
// in how transactions should affect an account's balance.
//   - Assets and Expenses have a normal debit balance.
//   - Liabilities, Equity, and Revenue have a normal credit balance.
func (at AccountType) GetNormalBalance() (NormalBalance, error) {
	switch at {
	case AssetAccount, ExpenseAccount:
		return DebitBalance, nil
	case LiabilityAccount, EquityAccount, RevenueAccount:
		return CreditBalance, nil
	default:
		// This path should be unreachable if AccountType.IsValid() is used correctly.
		return "", fmt.Errorf("unknown account type: %s", at)
	}
}

// AccountStatus represents the lifecycle state of an account.
// Using an explicit status prevents ambiguity and ensures that operations
// are only performed on accounts in an appropriate state (e.g., no new transactions
// on an archived account).
type AccountStatus string

const (
	// StatusPending indicates an account has been created but is not yet active for transactions.
	// This allows for review or setup processes before an account goes live.
	StatusPending AccountStatus = "PENDING"
	// StatusActive indicates an account is open and can have transactions posted to it.
	StatusActive AccountStatus = "ACTIVE"
	// StatusArchived indicates an account is closed and cannot have new transactions posted.
	// Its balance and history are preserved for auditing.
	StatusArchived AccountStatus = "ARCHIVED"
)

// IsValid checks if the AccountStatus is one of the predefined valid statuses.
func (as AccountStatus) IsValid() bool {
	switch as {
	case StatusPending, StatusActive, StatusArchived:
		return true
	default:
		return false
	}
}

// Account represents a specific record within the ledger for tracking financial value.
// It is the fundamental unit for recording debits and credits.
// The struct is designed to be immutable in its core properties (ID, Type), with state changes
// managed through versioning to ensure auditability and deterministic behavior.
type Account struct {
	// ID is the unique, immutable identifier for the account (UUID v4).
	ID uuid.UUID

	// Name is a human-readable identifier for the account (e.g., "Cash", "Accounts Payable").
	// This should be unique within a given chart of accounts.
	Name string

	// Type is the fundamental classification of the account. It is immutable.
	Type AccountType

	// NormalBalance is derived from the Type and indicates how the account balance increases.
	// It is immutable and set upon creation.
	NormalBalance NormalBalance

	// Status indicates the current lifecycle state of the account.
	Status AccountStatus

	// Version is used for optimistic concurrency control. Each modification to an account
	// must increment its version, preventing stale writes and ensuring a linear,
	// auditable history of changes.
	Version int64

	// CreatedAt is the timestamp when the account was first created, in UTC.
	CreatedAt time.Time

	// UpdatedAt is the timestamp of the last modification to the account, in UTC.
	UpdatedAt time.Time
}

// NewAccount creates and validates a new Account instance.
// This factory function is the sole entry point for creating accounts, ensuring that
// all instances are created in a consistent and valid state. It enforces system
// invariants from the moment of instantiation, embodying the "fail-closed" principle.
func NewAccount(id uuid.UUID, name string, accountType AccountType) (*Account, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("account ID cannot be nil")
	}
	if name == "" {
		return nil, fmt.Errorf("account name cannot be empty")
	}
	if !accountType.IsValid() {
		return nil, fmt.Errorf("invalid account type: %s", accountType)
	}

	normalBalance, err := accountType.GetNormalBalance()
	if err != nil {
		// This should be unreachable if accountType.IsValid() passes,
		// but serves as a safeguard for internal consistency.
		return nil, fmt.Errorf("could not determine normal balance for type %s: %w", accountType, err)
	}

	now := time.Now().UTC()

	return &Account{
		ID:            id,
		Name:          name,
		Type:          accountType,
		NormalBalance: normalBalance,
		Status:        StatusPending, // Accounts always start in a pending state for safety.
		Version:       1,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

```