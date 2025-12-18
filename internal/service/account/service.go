package account

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	// Assuming domain models are in a separate package.
	// "github.com/your-repo/internal/domain/account"
)

// For demonstration, we define domain models here. In a real project,
// these would be in a package like `internal/domain/account`.
type (
	// Account represents a user's primary account within the system.
	Account struct {
		ID             uuid.UUID
		UserID         uuid.UUID
		AccountNumber  string
		Balance        float64
		Currency       string
		OverdraftLimit float64
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}

	// LinkedAccount represents an external financial account linked to a primary account.
	LinkedAccount struct {
		ID                uuid.UUID
		AccountID         uuid.UUID // FK to our internal Account
		Provider          string    // e.g., "Plaid", "Stripe"
		ProviderAccountID string    // The ID of the account at the provider
		MaskedNumber      string    // e.g., "****1234"
		Status            string    // e.g., "active", "revoked"
		CreatedAt         time.Time
	}
)

// Custom errors for the account service layer.
var (
	ErrAccountNotFound       = errors.New("account not found")
	ErrUserHasNoAccounts     = errors.New("user has no accounts to link to")
	ErrInvalidOverdraftLimit = errors.New("invalid overdraft limit provided")
	ErrLinkingFailed         = errors.New("failed to link external account")
	ErrPermissionDenied      = errors.New("permission denied to perform this action")
)

// Repository defines the persistence layer interface for accounts.
// This allows us to swap the database implementation without changing the service logic.
type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Account, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]*Account, error)
	Update(ctx context.Context, account *Account) error
	CreateLinkedAccount(ctx context.Context, linkedAccount *LinkedAccount) error
	FindLinkedAccountsByAccountID(ctx context.Context, accountID uuid.UUID) ([]*LinkedAccount, error)
}

// ExternalLinker defines the interface for a third-party account linking service (e.g., Plaid).
type ExternalLinker interface {
	ExchangePublicToken(ctx context.Context, publicToken string) (accessToken string, providerAccountID string, err error)
	GetAccountInfo(ctx context.Context, accessToken string) (maskedNumber string, err error)
}

// Service provides the business logic for account management.
type Service struct {
	repo   Repository
	linker ExternalLinker
}

// NewService creates a new account service instance.
func NewService(repo Repository, linker ExternalLinker) (*Service, error) {
	if repo == nil {
		return nil, errors.New("repository cannot be nil")
	}
	if linker == nil {
		return nil, errors.New("external linker cannot be nil")
	}
	return &Service{
		repo:   repo,
		linker: linker,
	}, nil
}

// --- DTOs (Data Transfer Objects) ---
// DTOs are used to decouple the service layer from the transport layer (e.g., HTTP handlers).

type AccountDetailsDTO struct {
	ID             uuid.UUID           `json:"id"`
	AccountNumber  string              `json:"account_number"`
	Balance        float64             `json:"balance"`
	Currency       string              `json:"currency"`
	OverdraftLimit float64             `json:"overdraft_limit"`
	CreatedAt      time.Time           `json:"created_at"`
	LinkedAccounts []LinkedAccountDTO  `json:"linked_accounts"`
}

type LinkedAccountDTO struct {
	ID           uuid.UUID `json:"id"`
	Provider     string    `json:"provider"`
	MaskedNumber string    `json:"masked_number"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type OverdraftSettingsDTO struct {
	AccountID      uuid.UUID `json:"account_id"`
	OverdraftLimit float64   `json:"overdraft_limit"`
	Currency       string    `json:"currency"`
}

type AccountSummaryDTO struct {
	ID            uuid.UUID `json:"id"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	Currency      string    `json:"currency"`
}

// --- Service Methods ---

// GetAccountDetails retrieves comprehensive details for a specific account, including linked accounts.
// It assumes authorization has been checked by a preceding layer (e.g., middleware).
func (s *Service) GetAccountDetails(ctx context.Context, accountID uuid.UUID) (*AccountDetailsDTO, error) {
	account, err := s.repo.FindByID(ctx, accountID)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) { // Assuming repository returns a specific error
			return nil, ErrAccountNotFound
		}
		return nil, fmt.Errorf("failed to find account by id: %w", err)
	}

	linkedAccounts, err := s.repo.FindLinkedAccountsByAccountID(ctx, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to find linked accounts: %w", err)
	}

	return mapToAccountDetailsDTO(account, linkedAccounts), nil
}

// ListUserAccounts retrieves a summary of all accounts belonging to a user.
func (s *Service) ListUserAccounts(ctx context.Context, userID uuid.UUID) ([]*AccountSummaryDTO, error) {
	accounts, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find accounts for user %s: %w", userID, err)
	}

	if len(accounts) == 0 {
		return []*AccountSummaryDTO{}, nil // Return empty slice, not nil
	}

	dtos := make([]*AccountSummaryDTO, len(accounts))
	for i, acc := range accounts {
		dtos[i] = &AccountSummaryDTO{
			ID:            acc.ID,
			AccountNumber: acc.AccountNumber,
			Balance:       acc.Balance,
			Currency:      acc.Currency,
		}
	}

	return dtos, nil
}

// LinkExternalAccount handles the flow of linking a third-party bank account.
func (s *Service) LinkExternalAccount(ctx context.Context, userID uuid.UUID, provider, publicToken string) (*LinkedAccountDTO, error) {
	// Step 1: Find the user's primary internal account to link to.
	// For simplicity, we'll link to the first account found for the user.
	// A more complex app might require specifying which internal account to link.
	userAccounts, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve user accounts for linking: %w", err)
	}
	if len(userAccounts) == 0 {
		return nil, ErrUserHasNoAccounts
	}
	primaryAccount := userAccounts[0]

	// Step 2: Exchange the public token for an access token and provider account ID.
	accessToken, providerAccountID, err := s.linker.ExchangePublicToken(ctx, publicToken)
	if err != nil {
		return nil, fmt.Errorf("%w: could not exchange public token: %v", ErrLinkingFailed, err)
	}

	// Step 3: Use the access token to get account details (like masked number).
	maskedNumber, err := s.linker.GetAccountInfo(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("%w: could not get account info: %v", ErrLinkingFailed, err)
	}

	// Step 4: Create and persist the new linked account record.
	newLinkedAccount := &LinkedAccount{
		ID:                uuid.New(),
		AccountID:         primaryAccount.ID,
		Provider:          provider,
		ProviderAccountID: providerAccountID,
		MaskedNumber:      maskedNumber,
		Status:            "active",
		CreatedAt:         time.Now().UTC(),
	}

	if err := s.repo.CreateLinkedAccount(ctx, newLinkedAccount); err != nil {
		return nil, fmt.Errorf("could not save linked account: %w", err)
	}

	return mapToLinkedAccountDTO(newLinkedAccount), nil
}

// SetOverdraftLimit updates the overdraft protection limit for an account.
// It assumes the caller's identity (e.g., userID from context) has been verified
// to have ownership of the accountID.
func (s *Service) SetOverdraftLimit(ctx context.Context, accountID uuid.UUID, newLimit float64) error {
	if newLimit < 0 {
		return ErrInvalidOverdraftLimit
	}

	// In a real app, you might have a global or user-specific max limit.
	const maxOverdraftLimit = 5000.00
	if newLimit > maxOverdraftLimit {
		return fmt.Errorf("%w: limit cannot exceed %.2f", ErrInvalidOverdraftLimit, maxOverdraftLimit)
	}

	account, err := s.repo.FindByID(ctx, accountID)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) {
			return ErrAccountNotFound
		}
		return fmt.Errorf("failed to find account for update: %w", err)
	}

	// Here you would add an authorization check, e.g.:
	// userID, ok := ctx.Value("userID").(uuid.UUID)
	// if !ok || userID != account.UserID {
	// 	 return ErrPermissionDenied
	// }

	account.OverdraftLimit = newLimit
	account.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, account); err != nil {
		return fmt.Errorf("failed to update overdraft limit: %w", err)
	}

	return nil
}

// GetOverdraftSettings retrieves the current overdraft settings for an account.
func (s *Service) GetOverdraftSettings(ctx context.Context, accountID uuid.UUID) (*OverdraftSettingsDTO, error) {
	account, err := s.repo.FindByID(ctx, accountID)
	if err != nil {
		if errors.Is(err, ErrAccountNotFound) {
			return nil, ErrAccountNotFound
		}
		return nil, fmt.Errorf("failed to find account for overdraft settings: %w", err)
	}

	return &OverdraftSettingsDTO{
		AccountID:      account.ID,
		OverdraftLimit: account.OverdraftLimit,
		Currency:       account.Currency,
	}, nil
}

// --- Mappers ---

func mapToAccountDetailsDTO(acc *Account, linked []*LinkedAccount) *AccountDetailsDTO {
	linkedDTOs := make([]LinkedAccountDTO, len(linked))
	for i, la := range linked {
		linkedDTOs[i] = *mapToLinkedAccountDTO(la)
	}

	return &AccountDetailsDTO{
		ID:             acc.ID,
		AccountNumber:  acc.AccountNumber,
		Balance:        acc.Balance,
		Currency:       acc.Currency,
		OverdraftLimit: acc.OverdraftLimit,
		CreatedAt:      acc.CreatedAt,
		LinkedAccounts: linkedDTOs,
	}
}

func mapToLinkedAccountDTO(la *LinkedAccount) *LinkedAccountDTO {
	return &LinkedAccountDTO{
		ID:           la.ID,
		Provider:     la.Provider,
		MaskedNumber: la.MaskedNumber,
		Status:       la.Status,
		CreatedAt:    la.CreatedAt,
	}
}