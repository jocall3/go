```go
// Copyright (c) 2023-2024 The Corredor Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ledger

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// --- Custom Types ---
// In a larger system, these would likely live in a shared types package.

type AccountID string
type TransactionID string
type Currency string

// --- Enums ---

// EntryDirection specifies whether an entry is a debit or a credit.
// Debits decrease an asset or liability account, while credits increase them.
type EntryDirection int

const (
	// Debit represents a debit entry, which decreases the balance of an account.
	// For example, withdrawing money from a checking account.
	Debit EntryDirection = -1
	// Credit represents a credit entry, which increases the balance of an account.
	// For example, depositing money into a checking account.
	Credit EntryDirection = 1
)

func (d EntryDirection) String() string {
	switch d {
	case Debit:
		return "DEBIT"
	case Credit:
		return "CREDIT"
	default:
		return "UNKNOWN"
	}
}

// TransactionStatus represents the lifecycle state of a transaction.
type TransactionStatus string

const (
	StatusPending   TransactionStatus = "PENDING"
	StatusCommitted TransactionStatus = "COMMITTED"
	StatusFailed    TransactionStatus = "FAILED"
)

// --- Core Data Structures ---

// Account represents a single account in the ledger.
// It holds the balance and metadata for a specific entity and currency.
// The principle of immutability is encouraged; state changes should only
// occur via transactions processed by the Ledger.
type Account struct {
	ID        AccountID
	Currency  Currency
	Balance   decimal.Decimal
	Version   int64 // For optimistic concurrency control in a persistent store.
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Entry represents a single debit or credit operation on an account.
// A transaction is composed of one or more entries.
type Entry struct {
	AccountID AccountID
	Amount    decimal.Decimal // Always a positive value.
	Direction EntryDirection
}

// Transaction represents a set of balanced, atomic entries.
// The fundamental invariant of a transaction is that the sum of all debits
// must equal the sum of all credits.
type Transaction struct {
	ID          TransactionID
	Entries     []Entry
	Timestamp   time.Time
	Description string
	Status      TransactionStatus
}

// Ledger is the heart of the double-entry bookkeeping system.
// It holds the canonical state of all accounts and provides the sole mechanism
// for mutating that state by applying balanced, deterministic transactions.
// It is designed to be thread-safe.
type Ledger struct {
	mu           sync.RWMutex
	accounts     map[AccountID]*Account
	transactions map[TransactionID]TransactionStatus // For idempotency and audit.
}

// --- Errors ---

var (
	ErrAccountNotFound      = fmt.Errorf("account not found")
	ErrAccountExists        = fmt.Errorf("account already exists")
	ErrInsufficientFunds    = fmt.Errorf("insufficient funds")
	ErrTransactionUnbalanced = fmt.Errorf("transaction is unbalanced (debits do not equal credits)")
	ErrTransactionInvalid   = fmt.Errorf("transaction is invalid")
	ErrTransactionExists    = fmt.Errorf("transaction with the same ID already exists")
	ErrMismatchedCurrencies = fmt.Errorf("transaction entries involve mismatched currencies")
)

// --- Ledger Implementation ---

// NewLedger creates and initializes a new Ledger instance.
func NewLedger() *Ledger {
	return &Ledger{
		accounts:     make(map[AccountID]*Account),
		transactions: make(map[TransactionID]TransactionStatus),
	}
}

// CreateAccount adds a new account to the ledger with a zero balance.
// It fails if an account with the same ID already exists.
func (l *Ledger) CreateAccount(id AccountID, currency Currency) (*Account, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, exists := l.accounts[id]; exists {
		return nil, fmt.Errorf("%w: %s", ErrAccountExists, id)
	}

	now := time.Now().UTC()
	account := &Account{
		ID:        id,
		Currency:  currency,
		Balance:   decimal.Zero,
		Version:   1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	l.accounts[id] = account
	// Return a copy to prevent external modification.
	accCopy := *account
	return &accCopy, nil
}

// GetAccount retrieves an account by its ID.
// It returns a copy of the account to prevent direct mutation of the ledger's state.
// Returns ErrAccountNotFound if the account does not exist.
func (l *Ledger) GetAccount(id AccountID) (*Account, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	account, exists := l.accounts[id]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrAccountNotFound, id)
	}

	// Return a copy to ensure the ledger's internal state is not mutated by callers.
	accCopy := *account
	return &accCopy, nil
}

// ApplyTransaction validates and applies a transaction to the ledger.
// This is the sole entry point for state mutation and is designed to be atomic and deterministic.
// The process follows a strict validate-then-mutate pattern:
// 1. Idempotency Check: Prevents duplicate processing.
// 2. Validation Phase: All conditions (account existence, currency consistency, balance) are checked without altering state.
// 3. Funds Check: Ensures no account balance will become negative.
// 4. Mutation Phase: If all checks pass, the state is atomically updated.
// If any check fails, the system halts the operation, leaving the state untouched and guaranteeing capital safety.
func (l *Ledger) ApplyTransaction(tx Transaction) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// 1. Idempotency Check
	if status, exists := l.transactions[tx.ID]; exists {
		if status == StatusCommitted {
			return nil // Successfully processed before, return success.
		}
		return fmt.Errorf("%w: %s (status: %s)", ErrTransactionExists, tx.ID, status)
	}

	// 2. Validation Phase (read-only checks)
	if err := l.validateTransaction(tx); err != nil {
		l.transactions[tx.ID] = StatusFailed
		return err
	}

	// 3. Funds Check (read-only)
	deltas, err := l.calculateDeltas(tx)
	if err != nil {
		l.transactions[tx.ID] = StatusFailed
		return err // Should not happen if validateTransaction passed, but defensive.
	}
	if err := l.checkSufficientFunds(deltas); err != nil {
		l.transactions[tx.ID] = StatusFailed
		return err
	}

	// 4. Mutation Phase (write operations)
	// All checks passed. It is now safe to mutate the state.
	now := time.Now().UTC()
	for accID, delta := range deltas {
		account := l.accounts[accID] // We know the account exists from validation.
		account.Balance = account.Balance.Add(delta)
		account.Version++
		account.UpdatedAt = now
	}

	l.transactions[tx.ID] = StatusCommitted
	return nil
}

// validateTransaction performs a series of read-only checks on a transaction.
func (l *Ledger) validateTransaction(tx Transaction) error {
	if err := tx.validateStructure(); err != nil {
		return err
	}

	var transactionCurrency Currency
	sum := decimal.Zero

	for i, entry := range tx.Entries {
		account, exists := l.accounts[entry.AccountID]
		if !exists {
			return fmt.Errorf("%w: %s", ErrAccountNotFound, entry.AccountID)
		}

		if i == 0 {
			transactionCurrency = account.Currency
		} else if account.Currency != transactionCurrency {
			return fmt.Errorf("%w: expected %s, found %s for account %s",
				ErrMismatchedCurrencies, transactionCurrency, account.Currency, account.ID)
		}

		if entry.Direction == Credit {
			sum = sum.Add(entry.Amount)
		} else {
			sum = sum.Sub(entry.Amount)
		}
	}

	if !sum.IsZero() {
		return ErrTransactionUnbalanced
	}

	return nil
}

// calculateDeltas computes the net change for each account in a transaction.
func (l *Ledger) calculateDeltas(tx Transaction) (map[AccountID]decimal.Decimal, error) {
	deltas := make(map[AccountID]decimal.Decimal)
	for _, entry := range tx.Entries {
		change := entry.Amount
		if entry.Direction == Debit {
			change = change.Neg()
		}
		deltas[entry.AccountID] = deltas[entry.AccountID].Add(change)
	}
	return deltas, nil
}

// checkSufficientFunds ensures that applying the deltas will not result in any negative balances.
func (l *Ledger) checkSufficientFunds(deltas map[AccountID]decimal.Decimal) error {
	for accID, delta := range deltas {
		// This check is safe because validateTransaction already confirmed the account exists.
		account := l.accounts[accID]
		if account.Balance.Add(delta).IsNegative() {
			return fmt.Errorf("%w for account %s", ErrInsufficientFunds, accID)
		}
	}
	return nil
}

// --- Transaction Helper Methods ---

// NewTransaction creates a new transaction with a unique ID and a pending status.
func NewTransaction(description string) Transaction {
	return Transaction{
		ID:          TransactionID(uuid.NewString()),
		Entries:     make([]Entry, 0),
		Timestamp:   time.Now().UTC(),
		Description: description,
		Status:      StatusPending,
	}
}

// AddEntry adds a debit or credit entry to the transaction.
// The amount must always be positive; the direction determines the effect.
func (t *Transaction) AddEntry(accountID AccountID, amount decimal.Decimal, direction EntryDirection) {
	t.Entries = append(t.Entries, Entry{
		AccountID: accountID,
		Amount:    amount,
		Direction: direction,
	})
}

// validateStructure performs basic structural checks on a transaction.
func (t *Transaction) validateStructure() error {
	if len(t.Entries) < 2 {
		return fmt.Errorf("%w: must have at least two entries", ErrTransactionInvalid)
	}

	for _, entry := range t.Entries {
		if entry.Amount.IsNegative() || entry.Amount.IsZero() {
			return fmt.Errorf("%w: entry amount must be positive", ErrTransactionInvalid)
		}
		if entry.Direction != Debit && entry.Direction != Credit {
			return fmt.Errorf("%w: entry has an invalid direction", ErrTransactionInvalid)
		}
	}
	return nil
}
### END_OF_FILE_COMPLETED ###
```