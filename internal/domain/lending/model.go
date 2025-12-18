package lending

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrInvalidAmount indicates the monetary amount is invalid (e.g., negative or zero).
	ErrInvalidAmount = errors.New("amount must be greater than zero")
	// ErrInvalidDuration indicates the loan duration is invalid.
	ErrInvalidDuration = errors.New("duration must be between 1 and 120 months")
	// ErrApplicationNotPending indicates an operation was attempted on a non-pending application.
	ErrApplicationNotPending = errors.New("loan application is not in a pending state")
	// ErrOfferExpired indicates an attempt to accept an expired offer.
	ErrOfferExpired = errors.New("loan offer has expired")
	// ErrOfferNotAvailable indicates the offer is no longer available for acceptance.
	ErrOfferNotAvailable = errors.New("loan offer is not in a generated state")
)

// ApplicationStatus represents the lifecycle state of a LoanApplication.
type ApplicationStatus string

const (
	ApplicationStatusPending     ApplicationStatus = "PENDING"
	ApplicationStatusUnderReview ApplicationStatus = "UNDER_REVIEW"
	ApplicationStatusApproved    ApplicationStatus = "APPROVED"
	ApplicationStatusRejected    ApplicationStatus = "REJECTED"
	ApplicationStatusCancelled   ApplicationStatus = "CANCELLED"
)

// OfferStatus represents the lifecycle state of a LoanOffer.
type OfferStatus string

const (
	OfferStatusGenerated OfferStatus = "GENERATED"
	OfferStatusAccepted  OfferStatus = "ACCEPTED"
	OfferStatusRejected  OfferStatus = "REJECTED"
	OfferStatusExpired   OfferStatus = "EXPIRED"
)

// Money represents a monetary value safely using integer math for minor units.
type Money struct {
	Amount   int64  `json:"amount"`   // Amount in minor units (e.g., cents)
	Currency string `json:"currency"` // ISO 4217 Currency Code (e.g., "USD")
}

// LoanApplication is the aggregate root representing a user's request for a loan.
type LoanApplication struct {
	ID              uuid.UUID         `json:"id"`
	BorrowerID      uuid.UUID         `json:"borrower_id"`
	RequestedAmount Money             `json:"requested_amount"`
	DurationMonths  int               `json:"duration_months"`
	Purpose         string            `json:"purpose"`
	Status          ApplicationStatus `json:"status"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}

// NewLoanApplication creates a new LoanApplication with initial validation.
func NewLoanApplication(borrowerID uuid.UUID, amount int64, currency string, durationMonths int, purpose string) (*LoanApplication, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if durationMonths < 1 || durationMonths > 120 {
		return nil, ErrInvalidDuration
	}
	if currency == "" {
		return nil, errors.New("currency code is required")
	}

	now := time.Now().UTC()
	return &LoanApplication{
		ID:         uuid.New(),
		BorrowerID: borrowerID,
		RequestedAmount: Money{
			Amount:   amount,
			Currency: currency,
		},
		DurationMonths: durationMonths,
		Purpose:        purpose,
		Status:         ApplicationStatusPending,
		Metadata:       make(map[string]string),
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

// Review transitions the application to the UnderReview state.
func (la *LoanApplication) Review() error {
	if la.Status != ApplicationStatusPending {
		return ErrApplicationNotPending
	}
	la.Status = ApplicationStatusUnderReview
	la.UpdatedAt = time.Now().UTC()
	return nil
}

// Reject marks the application as rejected.
func (la *LoanApplication) Reject() {
	la.Status = ApplicationStatusRejected
	la.UpdatedAt = time.Now().UTC()
}

// Approve marks the application as approved.
func (la *LoanApplication) Approve() {
	la.Status = ApplicationStatusApproved
	la.UpdatedAt = time.Now().UTC()
}

// LoanOffer represents a formal offer extended to the borrower based on an application.
type LoanOffer struct {
	ID              uuid.UUID   `json:"id"`
	ApplicationID   uuid.UUID   `json:"application_id"`
	Principal       Money       `json:"principal"`
	InterestRateBps int64       `json:"interest_rate_bps"` // Basis Points (e.g., 500 = 5.00%)
	TermMonths      int         `json:"term_months"`
	MonthlyPayment  Money       `json:"monthly_payment"`
	TotalRepayment  Money       `json:"total_repayment"`
	Status          OfferStatus `json:"status"`
	ExpiresAt       time.Time   `json:"expires_at"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

// NewLoanOffer creates a new offer for a specific application.
func NewLoanOffer(appID uuid.UUID, principal Money, rateBps int64, term int, monthlyPayment Money, totalRepayment Money, validityDuration time.Duration) *LoanOffer {
	now := time.Now().UTC()
	return &LoanOffer{
		ID:              uuid.New(),
		ApplicationID:   appID,
		Principal:       principal,
		InterestRateBps: rateBps,
		TermMonths:      term,
		MonthlyPayment:  monthlyPayment,
		TotalRepayment:  totalRepayment,
		Status:          OfferStatusGenerated,
		ExpiresAt:       now.Add(validityDuration),
		CreatedAt:       now,
		UpdatedAt:       now,
	}
}

// Accept marks the offer as accepted by the borrower.
func (lo *LoanOffer) Accept() error {
	if lo.Status != OfferStatusGenerated {
		return ErrOfferNotAvailable
	}
	if time.Now().UTC().After(lo.ExpiresAt) {
		lo.Status = OfferStatusExpired
		lo.UpdatedAt = time.Now().UTC()
		return ErrOfferExpired
	}

	lo.Status = OfferStatusAccepted
	lo.UpdatedAt = time.Now().UTC()
	return nil
}

// Reject marks the offer as rejected by the borrower.
func (lo *LoanOffer) Reject() error {
	if lo.Status != OfferStatusGenerated {
		return ErrOfferNotAvailable
	}
	lo.Status = OfferStatusRejected
	lo.UpdatedAt = time.Now().UTC()
	return nil
}

// IsExpired checks if the offer is past its expiration time.
func (lo *LoanOffer) IsExpired() bool {
	return time.Now().UTC().After(lo.ExpiresAt)
}

// Expire forces the offer status to expired.
func (lo *LoanOffer) Expire() {
	if lo.Status == OfferStatusGenerated {
		lo.Status = OfferStatusExpired
		lo.UpdatedAt = time.Now().UTC()
	}
}