package lending

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

// Common errors returned by the lending service.
var (
	ErrInvalidInput       = errors.New("invalid input parameters")
	ErrApplicationNotFound = errors.New("loan application not found")
	ErrOfferExpired       = errors.New("loan offer has expired")
	ErrCreditScoreTooLow  = errors.New("credit score does not meet minimum requirements")
	ErrInternal           = errors.New("internal system error")
)

// Status represents the state of a loan application.
type Status string

const (
	StatusPending    Status = "PENDING"
	StatusApproved   Status = "APPROVED"
	StatusRejected   Status = "REJECTED"
	StatusDisbursed  Status = "DISBURSED"
	StatusRepaid     Status = "REPAID"
)

// LoanApplication represents a user's request for funds.
type LoanApplication struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	AmountCents    int64     `json:"amount_cents"`
	Currency       string    `json:"currency"`
	DurationMonths int       `json:"duration_months"`
	Purpose        string    `json:"purpose"`
	Status         Status    `json:"status"`
	InterestRate   float64   `json:"interest_rate"` // Annual percentage
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Offer represents a pre-approved loan offer presented to the user.
type Offer struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	MaxAmountCents int64     `json:"max_amount_cents"`
	InterestRate   float64   `json:"interest_rate"`
	ValidUntil     time.Time `json:"valid_until"`
	Terms          string    `json:"terms"`
}

// Repository defines the data persistence layer requirements.
type Repository interface {
	SaveApplication(ctx context.Context, app *LoanApplication) error
	GetApplicationByID(ctx context.Context, id string) (*LoanApplication, error)
	UpdateApplicationStatus(ctx context.Context, id string, status Status) error
	GetActiveOffers(ctx context.Context, userID string) ([]Offer, error)
	SaveOffer(ctx context.Context, offer *Offer) error
}

// CreditScorer defines the interface for external credit scoring systems.
type CreditScorer interface {
	GetScore(ctx context.Context, userID string) (int, error)
}

// NotificationService defines the interface for sending alerts to users.
type NotificationService interface {
	Notify(ctx context.Context, userID string, message string) error
}

// Service defines the public API for the lending domain.
type Service interface {
	ApplyForLoan(ctx context.Context, userID string, amountCents int64, durationMonths int, purpose string) (*LoanApplication, error)
	GetPreApprovedOffers(ctx context.Context, userID string) ([]Offer, error)
	AcceptOffer(ctx context.Context, userID string, offerID string, amountCents int64) (*LoanApplication, error)
}

// service implements the Service interface.
type service struct {
	repo     Repository
	scorer   CreditScorer
	notifier NotificationService
	// config could go here (min score, max amount, etc.)
}

// NewService creates a new instance of the lending service.
func NewService(repo Repository, scorer CreditScorer, notifier NotificationService) Service {
	return &service{
		repo:     repo,
		scorer:   scorer,
		notifier: notifier,
	}
}

// ApplyForLoan processes a new loan application.
func (s *service) ApplyForLoan(ctx context.Context, userID string, amountCents int64, durationMonths int, purpose string) (*LoanApplication, error) {
	if userID == "" || amountCents <= 0 || durationMonths <= 0 {
		return nil, ErrInvalidInput
	}

	// 1. Check Credit Score
	score, err := s.scorer.GetScore(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch credit score: %w", err)
	}

	// 2. Decision Logic
	status := StatusPending
	rate := 0.0

	// Simple rule engine for demonstration
	const minScore = 600
	if score < minScore {
		status = StatusRejected
	} else {
		status = StatusApproved
		// Calculate rate based on score (higher score = lower rate)
		// Base rate 5% + risk premium
		riskFactor := float64(850-score) / 100.0
		rate = 5.0 + riskFactor
	}

	// 3. Create Application Model
	app := &LoanApplication{
		ID:             generateID(), // Assumes a helper or UUID lib
		UserID:         userID,
		AmountCents:    amountCents,
		Currency:       "USD",
		DurationMonths: durationMonths,
		Purpose:        purpose,
		Status:         status,
		InterestRate:   math.Round(rate*100) / 100, // Round to 2 decimal places
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	// 4. Persist
	if err := s.repo.SaveApplication(ctx, app); err != nil {
		return nil, fmt.Errorf("failed to save application: %w", err)
	}

	// 5. Notify asynchronously (fire and forget or handled by queue in real prod)
	go func() {
		msg := fmt.Sprintf("Your loan application for %s has been %s.", formatCurrency(amountCents), status)
		_ = s.notifier.Notify(context.Background(), userID, msg)
	}()

	if status == StatusRejected {
		return app, ErrCreditScoreTooLow
	}

	return app, nil
}

// GetPreApprovedOffers retrieves valid offers for a user based on their history and score.
func (s *service) GetPreApprovedOffers(ctx context.Context, userID string) ([]Offer, error) {
	if userID == "" {
		return nil, ErrInvalidInput
	}

	// Check if we have cached/persisted offers
	offers, err := s.repo.GetActiveOffers(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch offers: %w", err)
	}

	// If no active offers, try to generate one dynamically
	if len(offers) == 0 {
		score, err := s.scorer.GetScore(ctx, userID)
		if err != nil {
			// If we can't get a score, we can't offer anything
			return []Offer{}, nil
		}

		if score >= 700 {
			// Generate a premium offer
			newOffer := Offer{
				ID:             generateID(),
				UserID:         userID,
				MaxAmountCents: 5000000, // $50,000.00
				InterestRate:   4.5,
				ValidUntil:     time.Now().Add(72 * time.Hour),
				Terms:          "Premium pre-approved offer for high credit score.",
			}
			if err := s.repo.SaveOffer(ctx, &newOffer); err == nil {
				offers = append(offers, newOffer)
			}
		}
	}

	return offers, nil
}

// AcceptOffer converts a pre-approved offer into an active loan application.
func (s *service) AcceptOffer(ctx context.Context, userID string, offerID string, amountCents int64) (*LoanApplication, error) {
	if userID == "" || offerID == "" || amountCents <= 0 {
		return nil, ErrInvalidInput
	}

	// 1. Retrieve Offers
	offers, err := s.repo.GetActiveOffers(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve offers: %w", err)
	}

	var selectedOffer *Offer
	for _, o := range offers {
		if o.ID == offerID {
			selectedOffer = &o
			break
		}
	}

	if selectedOffer == nil {
		return nil, errors.New("offer not found or does not belong to user")
	}

	// 2. Validate Offer
	if time.Now().After(selectedOffer.ValidUntil) {
		return nil, ErrOfferExpired
	}
	if amountCents > selectedOffer.MaxAmountCents {
		return nil, fmt.Errorf("requested amount exceeds offer limit of %s", formatCurrency(selectedOffer.MaxAmountCents))
	}

	// 3. Create Application (Immediately Approved)
	app := &LoanApplication{
		ID:             generateID(),
		UserID:         userID,
		AmountCents:    amountCents,
		Currency:       "USD",
		DurationMonths: 24, // Default for offers, or passed in
		Purpose:        "Offer Acceptance: " + selectedOffer.Terms,
		Status:         StatusApproved,
		InterestRate:   selectedOffer.InterestRate,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	if err := s.repo.SaveApplication(ctx, app); err != nil {
		return nil, fmt.Errorf("failed to create loan from offer: %w", err)
	}

	// 4. Notify
	_ = s.notifier.Notify(ctx, userID, "You have successfully accepted the loan offer. Funds are being processed.")

	return app, nil
}

// Helpers

// generateID is a placeholder for a UUID generator.
// In a real application, use "github.com/google/uuid".
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func formatCurrency(cents int64) string {
	return fmt.Sprintf("$%.2f", float64(cents)/100.0)
}