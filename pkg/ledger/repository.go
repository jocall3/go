```go
// Copyright 2023 The Bridge Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ledger

import (
	"context"
	"errors"
	"math/big"
)

var (
	// ErrTransactionNotFound is returned when a requested transaction does not exist.
	ErrTransactionNotFound = errors.New("transaction not found")
	// ErrAccountNotFound is returned when a requested account does not exist.
	ErrAccountNotFound = errors.New("account not found")
	// ErrAccountExists is returned when attempting to create an account that already exists.
	ErrAccountExists = errors.New("account already exists")
	// ErrVersionMismatch is returned by Tx.UpdateAccountBalance when the provided
	// oldVersion does not match the current version of the account in the database,
	// indicating a concurrent modification.
	ErrVersionMismatch = errors.New("account version mismatch")
	// ErrInsufficientBalance is a sentinel error that can be used by implementations
	// to indicate a business rule violation, though balance checks are typically
	// performed in the service layer before calling the repository.
	ErrInsufficientBalance = errors.New("insufficient balance")
)

// Repository defines the interface for persisting and retrieving ledger state.
// It abstracts the underlying storage mechanism (e.g., SQL database, key-value store),
// allowing the core ledger logic to remain agnostic to the persistence layer.
type Repository interface {
	// BeginTx starts a new transaction. All write operations must be performed
	// within a transaction to ensure atomicity.
	BeginTx(ctx context.Context) (Tx, error)

	// FindTransactionByID retrieves a single transaction by its unique ID.
	// This is a read-only operation.
	// Returns ErrTransactionNotFound if the transaction does not exist.
	FindTransactionByID(ctx context.Context, id TransactionID) (*Transaction, error)

	// FindAccountByID retrieves a single account by its unique ID.
	// This is a read-only operation.
	// Returns ErrAccountNotFound if the account does not exist.
	FindAccountByID(ctx context.Context, id AccountID) (*Account, error)

	// FindEntriesByAccountID retrieves a paginated list of entries for a given account.
	// This is essential for generating account statements and auditing.
	FindEntriesByAccountID(ctx context.Context, id AccountID, page PageRequest) ([]Entry, PageInfo, error)

	// GetAccountBalance retrieves the current balance and version for a specific account.
	// This is a specialized, read-only operation optimized for frequent balance checks
	// without loading the full account object.
	// Returns ErrAccountNotFound if the account does not exist.
	GetAccountBalance(ctx context.Context, id AccountID) (*AccountBalance, error)
}

// Tx defines the interface for operations that must be performed within an
// atomic database transaction. A transaction must be explicitly committed or
// rolled back.
type Tx interface {
	// Commit finalizes the transaction, making all changes permanent.
	Commit(ctx context.Context) error

	// Rollback discards all changes made in the transaction.
	Rollback(ctx context.Context) error

	// CreateTransaction inserts a new transaction and its associated entries.
	// This is the fundamental write operation that records financial movements.
	CreateTransaction(ctx context.Context, tx *Transaction) error

	// UpdateAccountBalance updates the balance and increments the version of an account.
	// This method must enforce optimistic concurrency control using the version number.
	// An implementation should return ErrVersionMismatch if oldVersion does not match
	// the stored version, forcing the caller to retry the business logic.
	UpdateAccountBalance(ctx context.Context, id AccountID, oldVersion uint64, newBalance *big.Int) error

	// CreateAccount creates a new account.
	// Returns ErrAccountExists if an account with the same ID already exists.
	CreateAccount(ctx context.Context, acc *Account) error

	// LockAccounts acquires pessimistic locks on the specified account IDs for the
	// duration of the transaction. This is crucial to prevent race conditions
	// and ensure serializable isolation for operations involving multiple accounts.
	// To prevent deadlocks, implementations and callers must ensure a consistent
	// locking order (e.g., lexicographically by AccountID).
	LockAccounts(ctx context.Context, accountIDs ...AccountID) error
}

// AccountBalance is a lightweight struct for retrieving just the balance and version
// of an account, useful for high-frequency read operations.
type AccountBalance struct {
	AccountID AccountID
	Balance   *big.Int
	Version   uint64
}

// PageRequest defines parameters for pagination, enabling efficient traversal
// of large datasets like account histories.
type PageRequest struct {
	// Cursor is an opaque string that marks the starting point for the next page.
	// For the first page, this should be an empty string.
	Cursor string
	// Limit is the maximum number of items to return.
	Limit int
}

// PageInfo contains metadata for the returned page of results, including the
// necessary information to fetch the next page.
type PageInfo struct {
	// HasNextPage is true if there are more items to fetch.
	HasNextPage bool
	// EndCursor is the cursor to use for the next PageRequest to get the next page.
	EndCursor string
}
### END_OF_FILE_COMPLETED ###
```