package budget

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Domain errors for Budget operations.
var (
	ErrInvalidID        = errors.New("invalid budget ID")
	ErrInvalidUserID    = errors.New("invalid user ID")
	ErrEmptyName        = errors.New("budget name cannot be empty")
	ErrInvalidAmount    = errors.New("budget amount must be greater than zero")
	ErrInvalidCurrency  = errors.New("invalid currency code")
	ErrInvalidPeriod    = errors.New("invalid budget period")
	ErrInvalidDateRange = errors.New("start date must be before end date")
)

// Period defines the recurrence or time span type of a budget.
type Period string

const (
	PeriodWeekly  Period = "WEEKLY"
	PeriodMonthly Period = "MONTHLY"
	PeriodYearly  Period = "YEARLY"
	PeriodCustom  Period = "CUSTOM"
)

// IsValid checks if the period is a known type.
func (p Period) IsValid() bool {
	switch p {
	case PeriodWeekly, PeriodMonthly, PeriodYearly, PeriodCustom:
		return true
	}
	return false
}

// Budget represents a financial constraint or plan set by a user for a specific duration.
type Budget struct {
	ID          uuid.UUID   `json:"id"`
	UserID      uuid.UUID   `json:"user_id"`
	Name        string      `json:"name"`
	Amount      float64     `json:"amount"` // Amount in major currency units
	Currency    string      `json:"currency"`
	Period      Period      `json:"period"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	CategoryIDs []uuid.UUID `json:"category_ids,omitempty"` // Specific categories this budget applies to; empty means general
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// NewBudget creates a new Budget entity with validation and default timestamps.
func NewBudget(userID uuid.UUID, name string, amount float64, currency string, period Period, start, end time.Time, categories []uuid.UUID) (*Budget, error) {
	if userID == uuid.Nil {
		return nil, ErrInvalidUserID
	}
	if name == "" {
		return nil, ErrEmptyName
	}
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}
	if len(currency) != 3 {
		return nil, ErrInvalidCurrency
	}
	if !period.IsValid() {
		return nil, ErrInvalidPeriod
	}
	if start.After(end) {
		return nil, ErrInvalidDateRange
	}

	now := time.Now().UTC()

	return &Budget{
		ID:          uuid.New(),
		UserID:      userID,
		Name:        name,
		Amount:      amount,
		Currency:    currency,
		Period:      period,
		StartDate:   start,
		EndDate:     end,
		CategoryIDs: categories,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

// Update modifies the budget details and updates the timestamp.
func (b *Budget) Update(name string, amount float64, currency string, period Period, start, end time.Time, categories []uuid.UUID) error {
	if name == "" {
		return ErrEmptyName
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if len(currency) != 3 {
		return ErrInvalidCurrency
	}
	if !period.IsValid() {
		return ErrInvalidPeriod
	}
	if start.After(end) {
		return ErrInvalidDateRange
	}

	b.Name = name
	b.Amount = amount
	b.Currency = currency
	b.Period = period
	b.StartDate = start
	b.EndDate = end
	b.CategoryIDs = categories
	b.UpdatedAt = time.Now().UTC()

	return nil
}

// IsActive checks if the budget is currently active based on the provided time.
func (b *Budget) IsActive(t time.Time) bool {
	// Normalize to UTC for comparison if needed, assuming t is passed correctly
	return (t.Equal(b.StartDate) || t.After(b.StartDate)) && (t.Equal(b.EndDate) || t.Before(b.EndDate))
}

// Remaining calculates the remaining budget given the spent amount.
func (b *Budget) Remaining(spent float64) float64 {
	return b.Amount - spent
}

// PercentUsed calculates the percentage of the budget that has been utilized.
func (b *Budget) PercentUsed(spent float64) float64 {
	if b.Amount == 0 {
		return 0
	}
	return (spent / b.Amount) * 100
}