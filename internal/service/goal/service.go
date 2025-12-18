package goal

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"
)

// Common errors for the goal service
var (
	ErrGoalNotFound      = errors.New("financial goal not found")
	ErrUnauthorized      = errors.New("unauthorized access to goal")
	ErrInvalidInput      = errors.New("invalid input data")
	ErrDeadlineInPast    = errors.New("deadline cannot be in the past")
	ErrTargetAmountZero  = errors.New("target amount must be greater than zero")
	ErrContributionError = errors.New("failed to process contribution")
)

// GoalStatus represents the state of a financial goal
type GoalStatus string

const (
	StatusActive    GoalStatus = "ACTIVE"
	StatusCompleted GoalStatus = "COMPLETED"
	StatusFailed    GoalStatus = "FAILED"
	StatusPaused    GoalStatus = "PAUSED"
)

// Goal represents the domain model for a financial goal
type Goal struct {
	ID            string     `json:"id"`
	UserID        string     `json:"user_id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	TargetAmount  float64    `json:"target_amount"`
	CurrentAmount float64    `json:"current_amount"`
	Currency      string     `json:"currency"`
	StartDate     time.Time  `json:"start_date"`
	Deadline      time.Time  `json:"deadline"`
	Status        GoalStatus `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// IsAchieved checks if the goal target has been met
func (g *Goal) IsAchieved() bool {
	return g.CurrentAmount >= g.TargetAmount
}

// ProgressPercentage calculates the completion percentage
func (g *Goal) ProgressPercentage() float64 {
	if g.TargetAmount <= 0 {
		return 0
	}
	percent := (g.CurrentAmount / g.TargetAmount) * 100
	return math.Min(percent, 100.0)
}

// Repository defines the interface for goal persistence
type Repository interface {
	Create(ctx context.Context, goal *Goal) error
	GetByID(ctx context.Context, id string) (*Goal, error)
	Update(ctx context.Context, goal *Goal) error
	Delete(ctx context.Context, id string) error
	ListByUserID(ctx context.Context, userID string, status *GoalStatus) ([]*Goal, error)
}

// Service defines the business logic interface for goals
type Service interface {
	CreateGoal(ctx context.Context, userID string, input CreateGoalInput) (*Goal, error)
	GetGoal(ctx context.Context, userID, goalID string) (*Goal, error)
	UpdateGoal(ctx context.Context, userID, goalID string, input UpdateGoalInput) (*Goal, error)
	AddContribution(ctx context.Context, userID, goalID string, amount float64) (*Goal, error)
	DeleteGoal(ctx context.Context, userID, goalID string) error
	ListUserGoals(ctx context.Context, userID string, statusFilter *GoalStatus) ([]*Goal, error)
	GetGoalProgress(ctx context.Context, userID, goalID string) (*GoalProgressOutput, error)
}

// CreateGoalInput DTO for creating a goal
type CreateGoalInput struct {
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	TargetAmount float64   `json:"target_amount"`
	Currency     string    `json:"currency"`
	Deadline     time.Time `json:"deadline"`
}

// UpdateGoalInput DTO for updating a goal
type UpdateGoalInput struct {
	Name         *string    `json:"name"`
	Description  *string    `json:"description"`
	TargetAmount *float64   `json:"target_amount"`
	Deadline     *time.Time `json:"deadline"`
	Status       *GoalStatus `json:"status"`
}

// GoalProgressOutput DTO for reporting progress
type GoalProgressOutput struct {
	GoalID           string  `json:"goal_id"`
	CurrentAmount    float64 `json:"current_amount"`
	TargetAmount     float64 `json:"target_amount"`
	Percentage       float64 `json:"percentage"`
	RemainingAmount  float64 `json:"remaining_amount"`
	DaysRemaining    int     `json:"days_remaining"`
	Status           GoalStatus `json:"status"`
}

type service struct {
	repo Repository
}

// NewService creates a new instance of the Goal Service
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateGoal(ctx context.Context, userID string, input CreateGoalInput) (*Goal, error) {
	if input.TargetAmount <= 0 {
		return nil, ErrTargetAmountZero
	}
	if input.Deadline.Before(time.Now()) {
		return nil, ErrDeadlineInPast
	}
	if input.Name == "" {
		return nil, fmt.Errorf("%w: name is required", ErrInvalidInput)
	}
	if input.Currency == "" {
		input.Currency = "USD" // Default currency
	}

	now := time.Now().UTC()
	goal := &Goal{
		UserID:        userID,
		Name:          input.Name,
		Description:   input.Description,
		TargetAmount:  input.TargetAmount,
		CurrentAmount: 0,
		Currency:      input.Currency,
		StartDate:     now,
		Deadline:      input.Deadline.UTC(),
		Status:        StatusActive,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	// Repository is responsible for generating ID if not handled by DB
	if err := s.repo.Create(ctx, goal); err != nil {
		return nil, fmt.Errorf("failed to create goal: %w", err)
	}

	return goal, nil
}

func (s *service) GetGoal(ctx context.Context, userID, goalID string) (*Goal, error) {
	goal, err := s.repo.GetByID(ctx, goalID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve goal: %w", err)
	}
	if goal == nil {
		return nil, ErrGoalNotFound
	}

	if goal.UserID != userID {
		return nil, ErrUnauthorized
	}

	return goal, nil
}

func (s *service) UpdateGoal(ctx context.Context, userID, goalID string, input UpdateGoalInput) (*Goal, error) {
	goal, err := s.GetGoal(ctx, userID, goalID)
	if err != nil {
		return nil, err
	}

	isUpdated := false

	if input.Name != nil && *input.Name != "" {
		goal.Name = *input.Name
		isUpdated = true
	}
	if input.Description != nil {
		goal.Description = *input.Description
		isUpdated = true
	}
	if input.TargetAmount != nil {
		if *input.TargetAmount <= 0 {
			return nil, ErrTargetAmountZero
		}
		goal.TargetAmount = *input.TargetAmount
		isUpdated = true
		
		// Re-evaluate status based on new target
		if goal.IsAchieved() && goal.Status == StatusActive {
			goal.Status = StatusCompleted
		} else if !goal.IsAchieved() && goal.Status == StatusCompleted {
			goal.Status = StatusActive
		}
	}
	if input.Deadline != nil {
		if input.Deadline.Before(time.Now()) {
			return nil, ErrDeadlineInPast
		}
		goal.Deadline = input.Deadline.UTC()
		isUpdated = true
	}
	if input.Status != nil {
		goal.Status = *input.Status
		isUpdated = true
	}

	if isUpdated {
		goal.UpdatedAt = time.Now().UTC()
		if err := s.repo.Update(ctx, goal); err != nil {
			return nil, fmt.Errorf("failed to update goal: %w", err)
		}
	}

	return goal, nil
}

func (s *service) AddContribution(ctx context.Context, userID, goalID string, amount float64) (*Goal, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("%w: contribution amount must be positive", ErrInvalidInput)
	}

	goal, err := s.GetGoal(ctx, userID, goalID)
	if err != nil {
		return nil, err
	}

	if goal.Status == StatusCompleted || goal.Status == StatusFailed {
		return nil, fmt.Errorf("%w: cannot contribute to a %s goal", ErrInvalidInput, goal.Status)
	}

	goal.CurrentAmount += amount
	goal.UpdatedAt = time.Now().UTC()

	// Check for completion
	if goal.IsAchieved() {
		goal.Status = StatusCompleted
	}

	if err := s.repo.Update(ctx, goal); err != nil {
		return nil, fmt.Errorf("failed to save contribution: %w", err)
	}

	return goal, nil
}

func (s *service) DeleteGoal(ctx context.Context, userID, goalID string) error {
	goal, err := s.GetGoal(ctx, userID, goalID)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, goal.ID); err != nil {
		return fmt.Errorf("failed to delete goal: %w", err)
	}

	return nil
}

func (s *service) ListUserGoals(ctx context.Context, userID string, statusFilter *GoalStatus) ([]*Goal, error) {
	goals, err := s.repo.ListByUserID(ctx, userID, statusFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to list goals: %w", err)
	}
	return goals, nil
}

func (s *service) GetGoalProgress(ctx context.Context, userID, goalID string) (*GoalProgressOutput, error) {
	goal, err := s.GetGoal(ctx, userID, goalID)
	if err != nil {
		return nil, err
	}

	remaining := math.Max(0, goal.TargetAmount-goal.CurrentAmount)
	daysRemaining := int(time.Until(goal.Deadline).Hours() / 24)
	if daysRemaining < 0 {
		daysRemaining = 0
	}

	return &GoalProgressOutput{
		GoalID:          goal.ID,
		CurrentAmount:   goal.CurrentAmount,
		TargetAmount:    goal.TargetAmount,
		Percentage:      goal.ProgressPercentage(),
		RemainingAmount: remaining,
		DaysRemaining:   daysRemaining,
		Status:          goal.Status,
	}, nil
}