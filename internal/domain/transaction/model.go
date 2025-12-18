package transaction

import (
	"errors"
	"strings"
	"time"
)

// Common errors for transaction domain validation.
var (
	ErrInvalidAmount      = errors.New("transaction amount must be positive")
	ErrInvalidCurrency    = errors.New("invalid currency code")
	ErrMissingAccount     = errors.New("account identifier is required")
	ErrInvalidInterval    = errors.New("invalid recurring interval")
	ErrDisputeReasonEmpty = errors.New("dispute reason cannot be empty")
)

// Status represents the lifecycle state of a transaction.
type Status string

const (
	StatusPending    Status = "PENDING"
	StatusProcessing Status = "PROCESSING"
	StatusCompleted  Status = "COMPLETED"
	StatusFailed     Status = "FAILED"
	StatusCancelled  Status = "CANCELLED"
	StatusRefunded   Status = "REFUNDED"
)

// Type defines the nature of the financial movement.
type Type string

const (
	TypeDeposit    Type = "DEPOSIT"
	TypeWithdrawal Type = "WITHDRAWAL"
	TypeTransfer   Type = "TRANSFER"
	TypePayment    Type = "PAYMENT"
	TypeFee        Type = "FEE"
	TypeRefund     Type = "REFUND"
)

// Transaction represents a core financial event within the system.
// It uses int64 for Amount to represent minor units (e.g., cents) to ensure precision.
type Transaction struct {
	ID              string            `json:"id"`
	UserID          string            `json:"user_id"`
	SourceAccountID string            `json:"source_account_id"`
	TargetAccountID string            `json:"target_account_id,omitempty"`
	Amount          int64             `json:"amount"` // Amount in minor units
	Currency        string            `json:"currency"`
	Type            Type              `json:"type"`
	Status          Status            `json:"status"`
	Reference       string            `json:"reference,omitempty"`
	Description     string            `json:"description,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	FailureReason   string            `json:"failure_reason,omitempty"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	CompletedAt     *time.Time        `json:"completed_at,omitempty"`
}

// NewTransaction creates a new transaction in the Pending state.
func NewTransaction(id, userID, sourceID string, amount int64, currency string, txnType Type) (*Transaction, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if len(currency) != 3 {
		return nil, ErrInvalidCurrency
	}
	if sourceID == "" {
		return nil, ErrMissingAccount
	}

	now := time.Now().UTC()
	return &Transaction{
		ID:              id,
		UserID:          userID,
		SourceAccountID: sourceID,
		Amount:          amount,
		Currency:        strings.ToUpper(currency),
		Type:            txnType,
		Status:          StatusPending,
		CreatedAt:       now,
		UpdatedAt:       now,
		Metadata:        make(map[string]string),
	}, nil
}

// Complete marks the transaction as successful.
func (t *Transaction) Complete() {
	now := time.Now().UTC()
	t.Status = StatusCompleted
	t.CompletedAt = &now
	t.UpdatedAt = now
}

// Fail marks the transaction as failed with a reason.
func (t *Transaction) Fail(reason string) {
	t.Status = StatusFailed
	t.FailureReason = reason
	t.UpdatedAt = time.Now().UTC()
}

// RecurringStatus represents the state of a recurring schedule.
type RecurringStatus string

const (
	RecurringActive    RecurringStatus = "ACTIVE"
	RecurringPaused    RecurringStatus = "PAUSED"
	RecurringCancelled RecurringStatus = "CANCELLED"
	RecurringFinished  RecurringStatus = "FINISHED"
)

// Interval defines the frequency of recurring transactions.
type Interval string

const (
	IntervalDaily   Interval = "DAILY"
	IntervalWeekly  Interval = "WEEKLY"
	IntervalMonthly Interval = "MONTHLY"
	IntervalYearly  Interval = "YEARLY"
)

// RecurringTransaction defines a schedule for automated transaction creation.
type RecurringTransaction struct {
	ID              string          `json:"id"`
	UserID          string          `json:"user_id"`
	SourceAccountID string          `json:"source_account_id"`
	TargetAccountID string          `json:"target_account_id"`
	Amount          int64           `json:"amount"`
	Currency        string          `json:"currency"`
	Interval        Interval        `json:"interval"`
	Status          RecurringStatus `json:"status"`
	NextRunAt       time.Time       `json:"next_run_at"`
	LastRunAt       *time.Time      `json:"last_run_at,omitempty"`
	TotalRuns       int             `json:"total_runs"`
	MaxRuns         *int            `json:"max_runs,omitempty"` // nil means infinite
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// CalculateNextRun computes the next execution time based on the interval.
func (rt *RecurringTransaction) CalculateNextRun(from time.Time) time.Time {
	switch rt.Interval {
	case IntervalDaily:
		return from.AddDate(0, 0, 1)
	case IntervalWeekly:
		return from.AddDate(0, 0, 7)
	case IntervalMonthly:
		return from.AddDate(0, 1, 0)
	case IntervalYearly:
		return from.AddDate(1, 0, 0)
	default:
		return from
	}
}

// ShouldRun checks if the recurring transaction is due for execution.
func (rt *RecurringTransaction) ShouldRun() bool {
	if rt.Status != RecurringActive {
		return false
	}
	if rt.MaxRuns != nil && rt.TotalRuns >= *rt.MaxRuns {
		return false
	}
	return time.Now().UTC().After(rt.NextRunAt)
}

// DisputeStatus represents the lifecycle of a dispute.
type DisputeStatus string

const (
	DisputeOpen     DisputeStatus = "OPEN"
	DisputeReview   DisputeStatus = "UNDER_REVIEW"
	DisputeWon      DisputeStatus = "WON"
	DisputeLost     DisputeStatus = "LOST"
	DisputeCanceled DisputeStatus = "CANCELED"
)

// Dispute represents a customer challenge to a specific transaction.
type Dispute struct {
	ID            string        `json:"id"`
	TransactionID string        `json:"transaction_id"`
	UserID        string        `json:"user_id"`
	Reason        string        `json:"reason"`
	Description   string        `json:"description,omitempty"`
	Status        DisputeStatus `json:"status"`
	Amount        int64         `json:"amount"` // Amount being disputed
	Evidence      []string      `json:"evidence,omitempty"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	ResolvedAt    *time.Time    `json:"resolved_at,omitempty"`
	Resolution    string        `json:"resolution,omitempty"`
}

// NewDispute creates a new dispute for a transaction.
func NewDispute(id, txnID, userID, reason string, amount int64) (*Dispute, error) {
	if reason == "" {
		return nil, ErrDisputeReasonEmpty
	}
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	now := time.Now().UTC()
	return &Dispute{
		ID:            id,
		TransactionID: txnID,
		UserID:        userID,
		Reason:        reason,
		Status:        DisputeOpen,
		Amount:        amount,
		CreatedAt:     now,
		UpdatedAt:     now,
		Evidence:      make([]string, 0),
	}, nil
}

// Resolve finalizes the dispute.
func (d *Dispute) Resolve(won bool, resolutionNote string) {
	now := time.Now().UTC()
	if won {
		d.Status = DisputeWon
	} else {
		d.Status = DisputeLost
	}
	d.Resolution = resolutionNote
	d.ResolvedAt = &now
	d.UpdatedAt = now
}