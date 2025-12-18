package budget

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a requested budget cannot be found.
	ErrNotFound = errors.New("budget not found")
	// ErrInvalidInput is returned when the input data is invalid.
	ErrInvalidInput = errors.New("invalid input parameters")
	// ErrUnauthorized is returned when the user is not authorized to perform the action.
	ErrUnauthorized = errors.New("unauthorized access to budget")
)

// Budget represents the core domain entity for a user's budget.
type Budget struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateBudgetRequest contains the payload for creating a new budget.
type CreateBudgetRequest struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	Currency  string    `json:"currency"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// UpdateBudgetRequest contains the payload for updating an existing budget.
// Pointers are used to distinguish between zero-values and unset fields.
type UpdateBudgetRequest struct {
	Name      *string    `json:"name,omitempty"`
	Amount    *float64   `json:"amount,omitempty"`
	Currency  *string    `json:"currency,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// Repository defines the interface for budget persistence operations.
// This allows for dependency injection of the storage layer.
type Repository interface {
	Create(ctx context.Context, budget *Budget) error
	GetByID(ctx context.Context, id uuid.UUID) (*Budget, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Budget, error)
	Update(ctx context.Context, budget *Budget) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// Service defines the interface for the budget business logic.
type Service interface {
	CreateBudget(ctx context.Context, req CreateBudgetRequest) (*Budget, error)
	GetBudget(ctx context.Context, id uuid.UUID) (*Budget, error)
	ListUserBudgets(ctx context.Context, userID uuid.UUID) ([]*Budget, error)
	UpdateBudget(ctx context.Context, id uuid.UUID, req UpdateBudgetRequest) (*Budget, error)
	DeleteBudget(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo   Repository
	logger *slog.Logger
}

// NewService initializes a new Budget Service with the given repository and logger.
func NewService(repo Repository, logger *slog.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// CreateBudget validates the request and creates a new budget record.
func (s *service) CreateBudget(ctx context.Context, req CreateBudgetRequest) (*Budget, error) {
	if err := s.validateCreateRequest(req); err != nil {
		s.logger.Warn("invalid create budget request", "error", err)
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	now := time.Now().UTC()
	budget := &Budget{
		ID:        uuid.New(),
		UserID:    req.UserID,
		Name:      req.Name,
		Amount:    req.Amount,
		Currency:  req.Currency,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, budget); err != nil {
		s.logger.Error("failed to create budget in repository", "error", err)
		return nil, err
	}

	s.logger.Info("budget created successfully", "budget_id", budget.ID, "user_id", budget.UserID)
	return budget, nil
}

// GetBudget retrieves a specific budget by its ID.
func (s *service) GetBudget(ctx context.Context, id uuid.UUID) (*Budget, error) {
	budget, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}
		s.logger.Error("failed to get budget", "budget_id", id, "error", err)
		return nil, err
	}

	return budget, nil
}

// ListUserBudgets retrieves all budgets associated with a specific user.
func (s *service) ListUserBudgets(ctx context.Context, userID uuid.UUID) ([]*Budget, error) {
	budgets, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		s.logger.Error("failed to list user budgets", "user_id", userID, "error", err)
		return nil, err
	}

	return budgets, nil
}

// UpdateBudget modifies an existing budget. It performs partial updates based on provided fields.
func (s *service) UpdateBudget(ctx context.Context, id uuid.UUID, req UpdateBudgetRequest) (*Budget, error) {
	// Retrieve existing budget first
	budget, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrNotFound
		}
		s.logger.Error("failed to retrieve budget for update", "budget_id", id, "error", err)
		return nil, err
	}

	// Apply updates
	hasUpdates := false
	if req.Name != nil {
		if len(*req.Name) < 3 {
			return nil, fmt.Errorf("%w: name must be at least 3 characters", ErrInvalidInput)
		}
		budget.Name = *req.Name
		hasUpdates = true
	}
	if req.Amount != nil {
		if *req.Amount < 0 {
			return nil, fmt.Errorf("%w: amount cannot be negative", ErrInvalidInput)
		}
		budget.Amount = *req.Amount
		hasUpdates = true
	}
	if req.Currency != nil {
		if len(*req.Currency) != 3 {
			return nil, fmt.Errorf("%w: currency must be a 3-letter code", ErrInvalidInput)
		}
		budget.Currency = *req.Currency
		hasUpdates = true
	}
	if req.StartDate != nil {
		budget.StartDate = *req.StartDate
		hasUpdates = true
	}
	if req.EndDate != nil {
		budget.EndDate = *req.EndDate
		hasUpdates = true
	}

	// Validate logical consistency if dates changed
	if budget.EndDate.Before(budget.StartDate) {
		return nil, fmt.Errorf("%w: end date cannot be before start date", ErrInvalidInput)
	}

	if !hasUpdates {
		return budget, nil
	}

	budget.UpdatedAt = time.Now().UTC()

	if err := s.repo.Update(ctx, budget); err != nil {
		s.logger.Error("failed to update budget", "budget_id", id, "error", err)
		return nil, err
	}

	s.logger.Info("budget updated successfully", "budget_id", id)
	return budget, nil
}

// DeleteBudget removes a budget from the system.
func (s *service) DeleteBudget(ctx context.Context, id uuid.UUID) error {
	// Check existence first (optional, depends on repo implementation, but good for explicit errors)
	_, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrNotFound
		}
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Error("failed to delete budget", "budget_id", id, "error", err)
		return err
	}

	s.logger.Info("budget deleted successfully", "budget_id", id)
	return nil
}

// validateCreateRequest performs basic validation on the creation payload.
func (s *service) validateCreateRequest(req CreateBudgetRequest) error {
	if req.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if len(req.Name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}
	if req.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	if len(req.Currency) != 3 {
		return errors.New("currency must be a valid 3-letter ISO code")
	}
	if req.EndDate.Before(req.StartDate) {
		return errors.New("end date cannot be before start date")
	}
	return nil
}