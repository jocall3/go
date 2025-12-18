package account

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Domain-specific errors related to account entities.
var (
	ErrInvalidAccountType     = errors.New("invalid account type")
	ErrInvalidAccountStatus   = errors.New("invalid account status")
	ErrInvalidTransactionType = errors.New("invalid transaction type")
	ErrInvalidCurrency        = errors.New("invalid currency code, must be 3-letter ISO 4217")
)

// AccountType defines the type of a bank account.
type AccountType string

const (
	// Checking is a standard transactional account.
	Checking AccountType = "checking"
	// Savings is an account for saving money, often with interest.
	Savings AccountType = "savings"
)

// IsValid checks if the AccountType is a valid, defined type.
func (at AccountType) IsValid() bool {
	switch at {
	case Checking, Savings:
		return true
	default:
		return false
	}
}

// AccountStatus defines the current status of an account.
type AccountStatus string

const (
	// Active status means the account is open and can be used.
	Active AccountStatus = "active"
	// Frozen status means the account is temporarily suspended.
	Frozen AccountStatus = "frozen"
	// Closed status means the account has been permanently closed.
	Closed AccountStatus = "closed"
)

// IsValid checks if the AccountStatus is a valid, defined status.
func (as AccountStatus) IsValid() bool {
	switch as {
	case Active, Frozen, Closed:
		return true
	default:
		return false
	}
}

// Account represents the core domain entity for a user's bank account.
// It holds the balance, type, status, and other essential information.
type Account struct {
	ID                uuid.UUID       `json:"id" db:"id"`
	UserID            uuid.UUID       `json:"user_id" db:"user_id"`
	AccountNumber     string          `json:"account_number" db:"account_number"`
	AccountType       AccountType     `json:"account_type" db:"account_type"`
	Balance           decimal.Decimal `json:"balance" db:"balance"`
	Currency          string          `json:"currency" db:"currency"` // ISO 4217 currency code (e.g., "USD")
	Status            AccountStatus   `json:"status" db:"status"`
	OverdraftSettings OverdraftSettings `json:"overdraft_settings"` // Often loaded separately or joined
	CreatedAt         time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" db:"updated_at"`
}

// NewAccount creates a new Account instance with default values.
// It's the designated factory function for creating valid accounts.
func NewAccount(userID uuid.UUID, accountNumber, currency string, accountType AccountType) (*Account, error) {
	if !accountType.IsValid() {
		return nil, ErrInvalidAccountType
	}
	if len(currency) != 3 {
		return nil, ErrInvalidCurrency
	}

	now := time.Now().UTC()
	return &Account{
		ID:            uuid.New(),
		UserID:        userID,
		AccountNumber: accountNumber,
		AccountType:   accountType,
		Balance:       decimal.Zero,
		Currency:      currency,
		Status:        Active,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// OverdraftSettings configures the overdraft protection for an account.
type OverdraftSettings struct {
	ID              uuid.UUID       `json:"id" db:"id"`
	AccountID       uuid.UUID       `json:"account_id" db:"account_id"`
	Enabled         bool            `json:"enabled" db:"enabled"`
	Limit           decimal.Decimal `json:"limit" db:"limit"`
	Fee             decimal.Decimal `json:"fee" db:"fee"`
	LinkedAccountID *uuid.UUID      `json:"linked_account_id,omitempty" db:"linked_account_id"` // Optional linked account for funding
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}

// LinkedAccount represents an external bank account linked to a primary account
// for transfers or other operations.
type LinkedAccount struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	AccountID           uuid.UUID `json:"account_id" db:"account_id"` // The primary account it's linked to
	UserID              uuid.UUID `json:"user_id" db:"user_id"`
	BankName            string    `json:"bank_name" db:"bank_name"`
	AccountNumberMasked string    `json:"account_number_masked" db:"account_number_masked"` // e.g., "******1234"
	RoutingNumber       string    `json:"routing_number,omitempty" db:"routing_number"`
	Nickname            string    `json:"nickname" db:"nickname"`
	Verified            bool      `json:"verified" db:"verified"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}

// Statement represents a periodic summary of financial transactions
// that have occurred over a given period. This is typically a read-only, generated entity.
type Statement struct {
	ID             uuid.UUID       `json:"id"`
	AccountID      uuid.UUID       `json:"account_id"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        time.Time       `json:"end_date"`
	OpeningBalance decimal.Decimal `json:"opening_balance"`
	ClosingBalance decimal.Decimal `json:"closing_balance"`
	Transactions   []Transaction   `json:"transactions"`
	GeneratedAt    time.Time       `json:"generated_at"`
}

// TransactionType defines the nature of a financial transaction.
type TransactionType string

const (
	Deposit     TransactionType = "deposit"
	Withdrawal  TransactionType = "withdrawal"
	TransferIn  TransactionType = "transfer_in"
	TransferOut TransactionType = "transfer_out"
	Fee         TransactionType = "fee"
	Interest    TransactionType = "interest"
	Payment     TransactionType = "payment"
)

// IsValid checks if the TransactionType is a valid, defined type.
func (tt TransactionType) IsValid() bool {
	switch tt {
	case Deposit, Withdrawal, TransferIn, TransferOut, Fee, Interest, Payment:
		return true
	default:
		return false
	}
}

// Transaction represents a single financial event, such as a deposit,
// withdrawal, or transfer.
type Transaction struct {
	ID             uuid.UUID       `json:"id" db:"id"`
	AccountID      uuid.UUID       `json:"account_id" db:"account_id"`
	Type           TransactionType `json:"type" db:"type"`
	Amount         decimal.Decimal `json:"amount" db:"amount"` // Positive for credits (deposits), negative for debits (withdrawals)
	Currency       string          `json:"currency" db:"currency"`
	Description    string          `json:"description" db:"description"`
	Timestamp      time.Time       `json:"timestamp" db:"timestamp"`
	RunningBalance decimal.Decimal `json:"running_balance" db:"running_balance"`
	ReferenceID    *string         `json:"reference_id,omitempty" db:"reference_id"` // e.g., for linking transfers
	// Metadata can store additional, unstructured data.
	// Note: Requires custom database handling (e.g., as a JSON/JSONB column).
	Metadata map[string]any `json:"metadata,omitempty" db:"metadata"`
}

// IsCredit returns true if the transaction amount is positive.
func (t *Transaction) IsCredit() bool {
	return t.Amount.IsPositive()
}

// IsDebit returns true if the transaction amount is negative.
func (t *Transaction) IsDebit() bool {
	return t.Amount.IsNegative()
}