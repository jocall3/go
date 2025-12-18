package goal

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

// Common domain errors for Financial Goals.
var (
	ErrInvalidID            = errors.New("id cannot be empty")
	ErrInvalidUserID        = errors.New("user id cannot be empty")
	ErrInvalidName          = errors.New("goal name cannot be empty")
	ErrInvalidTargetAmount  = errors.New("target amount must be greater than zero")
	ErrInvalidCurrency      = errors.New("currency code must be valid (e.g., USD, EUR)")
	ErrTargetDateInPast     = errors.New("target date must be in the future")
	ErrNegativeContribution = errors.New("contribution amount cannot be negative")
	ErrInsufficientFunds    = errors.New("insufficient funds in goal to withdraw")
	ErrGoalAlreadyCompleted = errors.New("goal is already marked as completed")
	ErrGoalArchived         = errors.New("cannot modify an archived goal")
)

// Status represents the lifecycle state of a Goal.
type Status string

const (
	StatusActive    Status = "ACTIVE"
	StatusPaused    Status = "PAUSED"
	StatusCompleted Status = "COMPLETED"
	StatusArchived  Status = "ARCHIVED"
)

// Priority indicates the importance of the financial goal.
type Priority int

const (
	PriorityLow    Priority = 1
	PriorityMedium Priority = 2
	PriorityHigh   Priority = 3
)

// Goal represents a financial target a user wants to achieve.
// It encapsulates the state and business logic for tracking progress.
type Goal struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description,omitempty"`
	TargetAmount  int64     `json:"target_amount"`  // Stored in minor units (e.g., cents)
	CurrentAmount int64     `json:"current_amount"` // Stored in minor units (e.g., cents)
	Currency      string    `json:"currency"`       // ISO 4217 Currency Code
	TargetDate    time.Time `json:"target_date"`
	Status        Status    `json:"status"`
	Priority      Priority  `json:"priority"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// New creates a new Goal entity with validation.
// amount is expected in minor units (e.g., 1000 = $10.00).
func New(id, userID, name, description, currency string, targetAmount int64, targetDate time.Time, priority Priority) (*Goal, error) {
	if id == "" {
		return nil, ErrInvalidID
	}
	if userID == "" {
		return nil, ErrInvalidUserID
	}
	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidName
	}
	if targetAmount <= 0 {
		return nil, ErrInvalidTargetAmount
	}
	if len(currency) != 3 {
		return nil, ErrInvalidCurrency
	}
	if targetDate.Before(time.Now()) {
		return nil, ErrTargetDateInPast
	}

	// Default priority if invalid
	if priority < PriorityLow || priority > PriorityHigh {
		priority = PriorityMedium
	}

	now := time.Now().UTC()

	return &Goal{
		ID:            id,
		UserID:        userID,
		Name:          strings.TrimSpace(name),
		Description:   strings.TrimSpace(description),
		TargetAmount:  targetAmount,
		CurrentAmount: 0,
		Currency:      strings.ToUpper(currency),
		TargetDate:    targetDate.UTC(),
		Status:        StatusActive,
		Priority:      priority,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// Deposit adds funds to the goal.
func (g *Goal) Deposit(amount int64) error {
	if err := g.canModify(); err != nil {
		return err
	}
	if amount < 0 {
		return ErrNegativeContribution
	}

	g.CurrentAmount += amount
	g.touch()

	// Auto-complete if target reached?
	// Business decision: Let's keep it active until explicitly completed,
	// but we could check here.
	return nil
}

// Withdraw removes funds from the goal.
func (g *Goal) Withdraw(amount int64) error {
	if err := g.canModify(); err != nil {
		return err
	}
	if amount < 0 {
		return ErrNegativeContribution
	}
	if g.CurrentAmount < amount {
		return ErrInsufficientFunds
	}

	g.CurrentAmount -= amount
	g.touch()
	return nil
}

// MarkCompleted changes the status to Completed.
func (g *Goal) MarkCompleted() error {
	if g.Status == StatusArchived {
		return ErrGoalArchived
	}
	g.Status = StatusCompleted
	g.touch()
	return nil
}

// Archive moves the goal to an archived state, making it read-only.
func (g *Goal) Archive() {
	g.Status = StatusArchived
	g.touch()
}

// ProgressPercentage calculates the completion percentage (0-100).
func (g *Goal) ProgressPercentage() float64 {
	if g.TargetAmount == 0 {
		return 0
	}
	percentage := (float64(g.CurrentAmount) / float64(g.TargetAmount)) * 100
	return math.Min(percentage, 100.0)
}

// RemainingAmount returns how much is left to reach the target.
func (g *Goal) RemainingAmount() int64 {
	remaining := g.TargetAmount - g.CurrentAmount
	if remaining < 0 {
		return 0
	}
	return remaining
}

// IsOnTrack checks if the goal is on track based on time elapsed vs money saved.
// This is a simple linear projection.
func (g *Goal) IsOnTrack() bool {
	if g.Status == StatusCompleted {
		return true
	}
	
	totalDuration := g.TargetDate.Sub(g.CreatedAt)
	elapsedDuration := time.Since(g.CreatedAt)

	// Avoid division by zero
	if totalDuration <= 0 {
		return g.CurrentAmount >= g.TargetAmount
	}

	timeProgress := float64(elapsedDuration) / float64(totalDuration)
	moneyProgress := float64(g.CurrentAmount) / float64(g.TargetAmount)

	// If we have saved more percentage-wise than time has passed, we are on track.
	return moneyProgress >= timeProgress
}

// UpdateDetails allows changing mutable fields.
func (g *Goal) UpdateDetails(name, description string, targetAmount int64, targetDate time.Time, priority Priority) error {
	if err := g.canModify(); err != nil {
		return err
	}
	if strings.TrimSpace(name) == "" {
		return ErrInvalidName
	}
	if targetAmount <= 0 {
		return ErrInvalidTargetAmount
	}
	// Allow extending date, but check if new date is in past relative to now
	if targetDate.Before(time.Now()) {
		return ErrTargetDateInPast
	}

	g.Name = strings.TrimSpace(name)
	g.Description = strings.TrimSpace(description)
	g.TargetAmount = targetAmount
	g.TargetDate = targetDate.UTC()
	g.Priority = priority
	g.touch()

	return nil
}

// canModify checks if the entity is in a state that allows modification.
func (g *Goal) canModify() error {
	if g.Status == StatusArchived {
		return ErrGoalArchived
	}
	if g.Status == StatusCompleted {
		return ErrGoalAlreadyCompleted
	}
	return nil
}

// touch updates the UpdatedAt timestamp.
func (g *Goal) touch() {
	g.UpdatedAt = time.Now().UTC()
}

// String implements the Stringer interface for logging.
func (g *Goal) String() string {
	return fmt.Sprintf("Goal<%s: %s (%d/%d %s)>", g.ID, g.Name, g.CurrentAmount, g.TargetAmount, g.Currency)
}