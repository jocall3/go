```go
// Copyright (c) 2024. The Bridge Project. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"time"

	"github.com/google/uuid"
)

// This file defines the concrete event structs for the Account aggregate.
// These events represent the immutable facts, or the "journal entries", of an account's lifecycle.
// Each event is a self-contained record of something that has happened.

// AggregateType constant for the Account aggregate.
const (
	AccountAggregate AggregateType = "Account"
)

// EventType constants for account-related events.
const (
	AccountCreatedType  EventType = "AccountCreated"
	AccountCreditedType EventType = "AccountCredited"
	AccountDebitedType  EventType = "AccountDebited"
	AccountFrozenType   EventType = "AccountFrozen"
	AccountUnfrozenType EventType = "AccountUnfrozen"
	AccountClosedType   EventType = "AccountClosed"
)

// --- AccountCreated Event ---

// AccountCreated is the event recorded when a new financial account is successfully opened.
// It represents the initial state of an account in the system.
type AccountCreated struct {
	BaseEvent
	OwnerID     uuid.UUID `json:"owner_id"`
	Currency    string    `json:"currency"`     // e.g., "USD", "BTC". Should be a validated type in a real system.
	AccountType string    `json:"account_type"` // e.g., "customer", "internal", "settlement"
}

// NewAccountCreated creates a new AccountCreated event.
func NewAccountCreated(aggregateID, ownerID uuid.UUID, currency, accountType string) *AccountCreated {
	return &AccountCreated{
		BaseEvent: BaseEvent{
			EventID:       uuid.New(),
			EventType:     AccountCreatedType,
			AggregateID:   aggregateID,
			AggregateType: AccountAggregate,
			Version:       1, // The first event for this aggregate
			Timestamp:     time.Now().UTC(),
		},
		OwnerID:     ownerID,
		Currency:    currency,
		AccountType: accountType,
	}
}

// --- AccountCredited Event ---

// AccountCredited is the event recorded when funds are added to an account.
// This is an immutable record of a credit operation.
type AccountCredited struct {
	BaseEvent
	Amount        int64     `json:"amount"`         // Amount in the smallest unit of the currency (e.g., cents)
	TransactionID uuid.UUID `json:"transaction_id"` // The transaction that caused this credit
	Reason        string    `json:"reason"`         // e.g., "deposit", "transfer_in"
	NewBalance    int64     `json:"new_balance"`    // The balance of the account *after* this credit
}

// NewAccountCredited creates a new AccountCredited event.
func NewAccountCredited(aggregateID, transactionID uuid.UUID, version int, amount, newBalance int64, reason string) *AccountCredited {
	return &AccountCredited{
		BaseEvent: BaseEvent{
			EventID:       uuid.New(),
			EventType:     AccountCreditedType,
			AggregateID:   aggregateID,
			AggregateType: AccountAggregate,
			Version:       version,
			Timestamp:     time.Now().UTC(),
		},
		Amount:        amount,
		TransactionID: transactionID,
		Reason:        reason,
		NewBalance:    newBalance,
	}
}

// --- AccountDebited Event ---

// AccountDebited is the event recorded when funds are removed from an account.
// This is an immutable record of a debit operation.
type AccountDebited struct {
	BaseEvent
	Amount        int64     `json:"amount"`         // Amount in the smallest unit of the currency (e.g., cents)
	TransactionID uuid.UUID `json:"transaction_id"` // The transaction that caused this debit
	Reason        string    `json:"reason"`         // e.g., "withdrawal", "transfer_out", "fee"
	NewBalance    int64     `json:"new_balance"`    // The balance of the account *after* this debit
}

// NewAccountDebited creates a new AccountDebited event.
func NewAccountDebited(aggregateID, transactionID uuid.UUID, version int, amount, newBalance int64, reason string) *AccountDebited {
	return &AccountDebited{
		BaseEvent: BaseEvent{
			EventID:       uuid.New(),
			EventType:     AccountDebitedType,
			AggregateID:   aggregateID,
			AggregateType: AccountAggregate,
			Version:       version,
			Timestamp:     time.Now().UTC(),
		},
		Amount:        amount,
		TransactionID: transactionID,
		Reason:        reason,
		NewBalance:    newBalance,
	}
}

// --- AccountFrozen Event ---

// AccountFrozen is the event recorded when an account's ability to transact is suspended.
// This is a critical event for risk and compliance management.
type AccountFrozen struct {
	BaseEvent
	Reason   string    `json:"reason"`    // A structured reason for the freeze, e.g., "compliance_review", "fraud_suspicion"
	FrozenBy uuid.UUID `json:"frozen_by"` // ID of the user or system component that initiated the freeze
}

// NewAccountFrozen creates a new AccountFrozen event.
func NewAccountFrozen(aggregateID, frozenBy uuid.UUID, version int, reason string) *AccountFrozen {
	return &AccountFrozen{
		BaseEvent: BaseEvent{
			EventID:       uuid.New(),
			EventType:     AccountFrozenType,
			AggregateID:   aggregateID,
			AggregateType: AccountAggregate,
			Version:       version,
			Timestamp:     time.Now().UTC(),
		},
		Reason:   reason,
		FrozenBy: frozenBy,
	}
}

// --- AccountUnfrozen Event ---

// AccountUnfrozen is the event recorded when a previously frozen account is made active again.
type AccountUnfrozen struct {
	BaseEvent
	Reason     string    `json:"reason"`      // A structured reason for the unfreeze, e.g., "compliance_review_cleared"
	UnfrozenBy uuid.UUID `json:"unfrozen_by"` // ID of the user or system component that initiated the unfreeze
}

// NewAccountUnfrozen creates a new AccountUnfrozen event.
func NewAccountUnfrozen(aggregateID, unfrozenBy uuid.UUID, version int, reason string) *AccountUnfrozen {
	return &AccountUnfrozen{
		BaseEvent: BaseEvent{
			EventID:       uuid.New(),
			EventType:     AccountUnfrozenType,
			AggregateID:   aggregateID,
			AggregateType: AccountAggregate,
			Version:       version,
			Timestamp:     time.Now().UTC(),
		},
		Reason:     reason,
		UnfrozenBy: unfrozenBy,
	}
}

// --- AccountClosed Event ---

// AccountClosed is the event recorded when an account is permanently closed.
// This is a terminal state for an account aggregate.
type AccountClosed struct {
	BaseEvent
	Reason       string    `json:"reason"`        // A structured reason for the closure, e.g., "customer_request", "dormancy"
	ClosedBy     uuid.UUID `json:"closed_by"`     // ID of the user or system component that initiated the closure
	FinalBalance int64     `json:"final_balance"` // The balance at the time of closure (should typically be zero)
}

// NewAccountClosed creates a new AccountClosed event.
func NewAccountClosed(aggregateID, closedBy uuid.UUID, version int, finalBalance int64, reason string) *AccountClosed {
	return &AccountClosed{
		BaseEvent: BaseEvent{
			EventID:       uuid.New(),
			EventType:     AccountClosedType,
			AggregateID:   aggregateID,
			AggregateType: AccountAggregate,
			Version:       version,
			Timestamp:     time.Now().UTC(),
		},
		Reason:       reason,
		ClosedBy:     closedBy,
		FinalBalance: finalBalance,
	}
}

```