package corporate

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

// -----------------------------------------------------------------------------
// Domain Models & Enums
// -----------------------------------------------------------------------------

// CardStatus represents the current state of a corporate card.
type CardStatus string

const (
	CardStatusActive   CardStatus = "ACTIVE"
	CardStatusFrozen   CardStatus = "FROZEN"
	CardStatusCanceled CardStatus = "CANCELED"
	CardStatusPending  CardStatus = "PENDING"
)

// CardType distinguishes between virtual and physical cards.
type CardType string

const (
	CardTypeVirtual  CardType = "VIRTUAL"
	CardTypePhysical CardType = "PHYSICAL"
)

// Currency represents the ISO currency code.
type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
)

// CardControls defines spending limits and restrictions.
type CardControls struct {
	DailyLimit     int64    `json:"daily_limit"`     // In smallest currency unit (e.g., cents)
	MonthlyLimit   int64    `json:"monthly_limit"`   // In smallest currency unit
	AllowedMCCs    []string `json:"allowed_mccs"`    // Merchant Category Codes
	BlockedMCCs    []string `json:"blocked_mccs"`    // Merchant Category Codes
	AllowForeign   bool     `json:"allow_foreign"`   // Allow non-domestic transactions
	SingleTxLimit  int64    `json:"single_tx_limit"` // Max amount per transaction
}

// Card represents a corporate credit/debit card.
type Card struct {
	ID             string       `json:"id"`
	AccountID      string       `json:"account_id"`
	HolderName     string       `json:"holder_name"`
	Last4          string       `json:"last_4"`
	ExpiryMonth    int          `json:"expiry_month"`
	ExpiryYear     int          `json:"expiry_year"`
	CVV            string       `json:"cvv,omitempty"` // Usually not stored, but returned on creation
	PAN            string       `json:"pan,omitempty"` // Primary Account Number (sensitive)
	Type           CardType     `json:"type"`
	Status         CardStatus   `json:"status"`
	Currency       Currency     `json:"currency"`
	Controls       CardControls `json:"controls"`
	ExternalCardID string       `json:"external_card_id"` // ID from the issuer (e.g., Stripe/Marqeta)
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}

// -----------------------------------------------------------------------------
// Service Interfaces
// -----------------------------------------------------------------------------

// CardRepository defines the persistence layer requirements for cards.
type CardRepository interface {
	Create(ctx context.Context, card *Card) error
	GetByID(ctx context.Context, id string) (*Card, error)
	ListByAccount(ctx context.Context, accountID string) ([]*Card, error)
	UpdateStatus(ctx context.Context, id string, status CardStatus) error
	UpdateControls(ctx context.Context, id string, controls CardControls) error
}

// CardIssuerProvider defines the interface for the external card issuing platform.
type CardIssuerProvider interface {
	IssueCard(ctx context.Context, req IssueCardRequest) (*IssuedCardDetails, error)
	FreezeCard(ctx context.Context, externalID string) error
	UnfreezeCard(ctx context.Context, externalID string) error
	CancelCard(ctx context.Context, externalID string) error
	UpdateControls(ctx context.Context, externalID string, controls CardControls) error
}

// -----------------------------------------------------------------------------
// DTOs
// -----------------------------------------------------------------------------

// IssueCardRequest contains data required to issue a new card.
type IssueCardRequest struct {
	AccountID  string       `json:"account_id"`
	HolderName string       `json:"holder_name"`
	Type       CardType     `json:"type"`
	Currency   Currency     `json:"currency"`
	Controls   CardControls `json:"controls"`
}

// IssuedCardDetails contains sensitive data returned by the issuer.
type IssuedCardDetails struct {
	ExternalID  string
	PAN         string
	CVV         string
	ExpiryMonth int
	ExpiryYear  int
	Last4       string
}

// -----------------------------------------------------------------------------
// Service Implementation
// -----------------------------------------------------------------------------

// CardService handles business logic for corporate cards.
type CardService struct {
	repo     CardRepository
	provider CardIssuerProvider
	logger   *slog.Logger
}

// NewCardService creates a new instance of CardService.
func NewCardService(repo CardRepository, provider CardIssuerProvider, logger *slog.Logger) *CardService {
	return &CardService{
		repo:     repo,
		provider: provider,
		logger:   logger,
	}
}

// IssueVirtualCard creates a new virtual card via the provider and stores it.
func (s *CardService) IssueVirtualCard(ctx context.Context, req IssueCardRequest) (*Card, error) {
	log := s.logger.With("method", "IssueVirtualCard", "account_id", req.AccountID)
	log.Info("initiating virtual card issuance")

	if req.Type != CardTypeVirtual {
		return nil, errors.New("only virtual cards can be issued via this endpoint")
	}

	// 1. Call external provider to issue the card
	issuedDetails, err := s.provider.IssueCard(ctx, req)
	if err != nil {
		log.Error("failed to issue card with provider", "error", err)
		return nil, fmt.Errorf("provider failed to issue card: %w", err)
	}

	// 2. Construct domain entity
	now := time.Now().UTC()
	card := &Card{
		ID:             uuid.New().String(),
		AccountID:      req.AccountID,
		HolderName:     req.HolderName,
		Last4:          issuedDetails.Last4,
		ExpiryMonth:    issuedDetails.ExpiryMonth,
		ExpiryYear:     issuedDetails.ExpiryYear,
		PAN:            issuedDetails.PAN, // Note: In production, ensure PAN is encrypted or tokenized
		CVV:            issuedDetails.CVV, // Note: Usually not stored persistently
		Type:           CardTypeVirtual,
		Status:         CardStatusActive,
		Currency:       req.Currency,
		Controls:       req.Controls,
		ExternalCardID: issuedDetails.ExternalID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	// 3. Persist to database
	// Note: We might want to mask PAN/CVV before saving depending on PCI-DSS compliance requirements.
	// For this implementation, we assume the repository handles encryption or we store it as is (dev mode).
	if err := s.repo.Create(ctx, card); err != nil {
		log.Error("failed to persist card", "error", err)
		// Attempt to rollback provider creation if DB fails
		_ = s.provider.CancelCard(ctx, issuedDetails.ExternalID)
		return nil, fmt.Errorf("failed to save card record: %w", err)
	}

	log.Info("virtual card issued successfully", "card_id", card.ID)
	return card, nil
}

// GetCard retrieves a card by its ID.
func (s *CardService) GetCard(ctx context.Context, cardID string) (*Card, error) {
	card, err := s.repo.GetByID(ctx, cardID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve card: %w", err)
	}
	if card == nil {
		return nil, errors.New("card not found")
	}
	return card, nil
}

// ListCards retrieves all cards for a specific corporate account.
func (s *CardService) ListCards(ctx context.Context, accountID string) ([]*Card, error) {
	cards, err := s.repo.ListByAccount(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to list cards: %w", err)
	}
	return cards, nil
}

// FreezeCard temporarily disables a card.
func (s *CardService) FreezeCard(ctx context.Context, cardID string) error {
	log := s.logger.With("method", "FreezeCard", "card_id", cardID)

	card, err := s.repo.GetByID(ctx, cardID)
	if err != nil {
		return fmt.Errorf("failed to fetch card: %w", err)
	}
	if card.Status == CardStatusFrozen {
		return nil // Already frozen
	}

	// 1. Update external provider
	if err := s.provider.FreezeCard(ctx, card.ExternalCardID); err != nil {
		log.Error("failed to freeze card at provider", "error", err)
		return fmt.Errorf("provider failed to freeze card: %w", err)
	}

	// 2. Update local state
	if err := s.repo.UpdateStatus(ctx, cardID, CardStatusFrozen); err != nil {
		log.Error("failed to update card status in db", "error", err)
		return fmt.Errorf("failed to update card status: %w", err)
	}

	log.Info("card frozen successfully")
	return nil
}

// UnfreezeCard reactivates a frozen card.
func (s *CardService) UnfreezeCard(ctx context.Context, cardID string) error {
	log := s.logger.With("method", "UnfreezeCard", "card_id", cardID)

	card, err := s.repo.GetByID(ctx, cardID)
	if err != nil {
		return fmt.Errorf("failed to fetch card: %w", err)
	}
	if card.Status != CardStatusFrozen {
		return errors.New("card is not frozen")
	}

	// 1. Update external provider
	if err := s.provider.UnfreezeCard(ctx, card.ExternalCardID); err != nil {
		log.Error("failed to unfreeze card at provider", "error", err)
		return fmt.Errorf("provider failed to unfreeze card: %w", err)
	}

	// 2. Update local state
	if err := s.repo.UpdateStatus(ctx, cardID, CardStatusActive); err != nil {
		log.Error("failed to update card status in db", "error", err)
		return fmt.Errorf("failed to update card status: %w", err)
	}

	log.Info("card unfrozen successfully")
	return nil
}

// UpdateCardControls modifies spending limits and restrictions.
func (s *CardService) UpdateCardControls(ctx context.Context, cardID string, controls CardControls) error {
	log := s.logger.With("method", "UpdateCardControls", "card_id", cardID)

	card, err := s.repo.GetByID(ctx, cardID)
	if err != nil {
		return fmt.Errorf("failed to fetch card: %w", err)
	}

	// 1. Update external provider
	if err := s.provider.UpdateControls(ctx, card.ExternalCardID, controls); err != nil {
		log.Error("failed to update controls at provider", "error", err)
		return fmt.Errorf("provider failed to update controls: %w", err)
	}

	// 2. Update local state
	if err := s.repo.UpdateControls(ctx, cardID, controls); err != nil {
		log.Error("failed to update controls in db", "error", err)
		return fmt.Errorf("failed to update card controls: %w", err)
	}

	log.Info("card controls updated successfully")
	return nil
}

// CancelCard permanently terminates a card.
func (s *CardService) CancelCard(ctx context.Context, cardID string) error {
	log := s.logger.With("method", "CancelCard", "card_id", cardID)

	card, err := s.repo.GetByID(ctx, cardID)
	if err != nil {
		return fmt.Errorf("failed to fetch card: %w", err)
	}
	if card.Status == CardStatusCanceled {
		return nil
	}

	// 1. Update external provider
	if err := s.provider.CancelCard(ctx, card.ExternalCardID); err != nil {
		log.Error("failed to cancel card at provider", "error", err)
		return fmt.Errorf("provider failed to cancel card: %w", err)
	}

	// 2. Update local state
	if err := s.repo.UpdateStatus(ctx, cardID, CardStatusCanceled); err != nil {
		log.Error("failed to update card status in db", "error", err)
		return fmt.Errorf("failed to update card status: %w", err)
	}

	log.Info("card canceled successfully")
	return nil
}