package payment

import (
	"errors"
	"fmt"
	"time"
)

// Common domain errors
var (
	ErrInvalidAmount          = errors.New("amount must be positive")
	ErrInvalidCurrency        = errors.New("invalid currency code")
	ErrQuoteExpired           = errors.New("fx quote has expired")
	ErrInvalidStateTransition = errors.New("invalid state transition")
	ErrSameCurrencyTransfer   = errors.New("source and target currency cannot be the same for international transfer")
)

// TransferStatus represents the lifecycle state of a payment.
type TransferStatus string

const (
	TransferStatusDraft        TransferStatus = "DRAFT"
	TransferStatusQuoted       TransferStatus = "QUOTED"
	TransferStatusPendingFunds TransferStatus = "PENDING_FUNDS"
	TransferStatusProcessing   TransferStatus = "PROCESSING"
	TransferStatusCompleted    TransferStatus = "COMPLETED"
	TransferStatusFailed       TransferStatus = "FAILED"
	TransferStatusCancelled    TransferStatus = "CANCELLED"
)

// Currency represents a 3-letter ISO 4217 currency code.
type Currency string

// Money represents a monetary value in minor units (e.g., cents) to avoid floating-point errors.
type Money struct {
	Amount   int64    `json:"amount"`   // Amount in minor units
	Currency Currency `json:"currency"` // ISO 4217 code
}

// NewMoney creates a new Money value object.
func NewMoney(amount int64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, ErrInvalidAmount
	}
	if len(currency) != 3 {
		return Money{}, ErrInvalidCurrency
	}
	return Money{
		Amount:   amount,
		Currency: Currency(currency),
	}, nil
}

// FXRate represents an exchange rate between two currencies.
type FXRate struct {
	BaseCurrency   Currency  `json:"base_currency"`
	TargetCurrency Currency  `json:"target_currency"`
	Rate           float64   `json:"rate"` // Multiplier: Target = Base * Rate
	ExpiresAt      time.Time `json:"expires_at"`
	QuoteID        string    `json:"quote_id"`
}

// IsExpired checks if the FX rate is no longer valid.
func (r FXRate) IsExpired() bool {
	return time.Now().After(r.ExpiresAt)
}

// Convert calculates the target amount based on the source amount and rate.
// It returns the converted Money object.
func (r FXRate) Convert(source Money) (Money, error) {
	if r.IsExpired() {
		return Money{}, ErrQuoteExpired
	}
	if source.Currency != r.BaseCurrency {
		return Money{}, fmt.Errorf("source currency %s does not match rate base currency %s", source.Currency, r.BaseCurrency)
	}

	// Calculate converted amount (simple implementation, real world might need specific rounding modes)
	convertedAmount := int64(float64(source.Amount) * r.Rate)

	return Money{
		Amount:   convertedAmount,
		Currency: r.TargetCurrency,
	}, nil
}

// Transfer represents the aggregate root for an international payment.
type Transfer struct {
	ID             string            `json:"id"`
	CustomerID     string            `json:"customer_id"`
	BeneficiaryID  string            `json:"beneficiary_id"`
	SourceAmount   Money             `json:"source_amount"`
	TargetAmount   Money             `json:"target_amount"`
	Rate           FXRate            `json:"applied_rate"`
	Status         TransferStatus    `json:"status"`
	Reference      string            `json:"reference"`
	Metadata       map[string]string `json:"metadata,omitempty"`
	FailureReason  string            `json:"failure_reason,omitempty"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	CompletedAt    *time.Time        `json:"completed_at,omitempty"`
}

// NewTransfer initializes a new transfer in Draft state.
func NewTransfer(id, customerID, beneficiaryID string, source Money, targetCurrency string) (*Transfer, error) {
	if source.Currency == Currency(targetCurrency) {
		return nil, ErrSameCurrencyTransfer
	}

	now := time.Now().UTC()
	return &Transfer{
		ID:            id,
		CustomerID:    customerID,
		BeneficiaryID: beneficiaryID,
		SourceAmount:  source,
		TargetAmount:  Money{Amount: 0, Currency: Currency(targetCurrency)}, // Calculated later via quote
		Status:        TransferStatusDraft,
		CreatedAt:     now,
		UpdatedAt:     now,
		Metadata:      make(map[string]string),
	}, nil
}

// ApplyQuote applies an FX rate to the transfer and transitions it to Quoted status.
func (t *Transfer) ApplyQuote(rate FXRate) error {
	if t.Status != TransferStatusDraft {
		return fmt.Errorf("%w: cannot apply quote to transfer in status %s", ErrInvalidStateTransition, t.Status)
	}
	if rate.BaseCurrency != t.SourceAmount.Currency {
		return fmt.Errorf("quote base currency mismatch")
	}
	if rate.TargetCurrency != t.TargetAmount.Currency {
		return fmt.Errorf("quote target currency mismatch")
	}

	targetMoney, err := rate.Convert(t.SourceAmount)
	if err != nil {
		return err
	}

	t.Rate = rate
	t.TargetAmount = targetMoney
	t.Status = TransferStatusQuoted
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// Confirm locks the transfer and moves it to Pending Funds.
func (t *Transfer) Confirm() error {
	if t.Status != TransferStatusQuoted {
		return fmt.Errorf("%w: only quoted transfers can be confirmed", ErrInvalidStateTransition)
	}
	if t.Rate.IsExpired() {
		return ErrQuoteExpired
	}

	t.Status = TransferStatusPendingFunds
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// MarkProcessing moves the transfer to Processing state (funds received).
func (t *Transfer) MarkProcessing() error {
	if t.Status != TransferStatusPendingFunds {
		return fmt.Errorf("%w: transfer must be pending funds to start processing", ErrInvalidStateTransition)
	}

	t.Status = TransferStatusProcessing
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// Complete finalizes the transfer.
func (t *Transfer) Complete() error {
	if t.Status != TransferStatusProcessing {
		return fmt.Errorf("%w: only processing transfers can be completed", ErrInvalidStateTransition)
	}

	now := time.Now().UTC()
	t.Status = TransferStatusCompleted
	t.UpdatedAt = now
	t.CompletedAt = &now
	return nil
}

// Fail marks the transfer as failed with a reason.
func (t *Transfer) Fail(reason string) error {
	// Can fail from almost any active state
	if t.IsTerminal() {
		return fmt.Errorf("%w: cannot fail a terminal transfer", ErrInvalidStateTransition)
	}

	t.Status = TransferStatusFailed
	t.FailureReason = reason
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// Cancel allows the user to cancel the transfer if it hasn't been processed yet.
func (t *Transfer) Cancel(reason string) error {
	if t.Status == TransferStatusProcessing || t.IsTerminal() {
		return fmt.Errorf("%w: cannot cancel transfer in status %s", ErrInvalidStateTransition, t.Status)
	}

	t.Status = TransferStatusCancelled
	t.FailureReason = reason
	t.UpdatedAt = time.Now().UTC()
	return nil
}

// IsTerminal checks if the transfer is in a final state.
func (t *Transfer) IsTerminal() bool {
	return t.Status == TransferStatusCompleted ||
		t.Status == TransferStatusFailed ||
		t.Status == TransferStatusCancelled
}