package corporate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Common domain errors
var (
	ErrInvalidAmount       = errors.New("amount must be non-negative")
	ErrInvalidCurrency     = errors.New("invalid currency code")
	ErrCardExpired         = errors.New("corporate card has expired")
	ErrCardLimitExceeded   = errors.New("transaction exceeds card limit")
	ErrCardInactive        = errors.New("corporate card is not active")
	ErrInvalidRiskScore    = errors.New("risk score must be between 0 and 100")
	ErrTreasuryLocked      = errors.New("treasury account is locked for rebalancing")
	ErrComplianceReviewReq = errors.New("compliance review required before action")
)

// Money represents a monetary value in the smallest unit (e.g., cents).
type Money struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

// NewMoney creates a new Money instance.
func NewMoney(amount int64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, ErrInvalidAmount
	}
	if len(currency) != 3 {
		return Money{}, ErrInvalidCurrency
	}
	return Money{Amount: amount, Currency: strings.ToUpper(currency)}, nil
}

// -----------------------------------------------------------------------------
// Corporate Cards Domain
// -----------------------------------------------------------------------------

type CardStatus string
type CardType string

const (
	CardStatusActive    CardStatus = "ACTIVE"
	CardStatusFrozen    CardStatus = "FROZEN"
	CardStatusCancelled CardStatus = "CANCELLED"
	CardStatusPending   CardStatus = "PENDING"

	CardTypePhysical CardType = "PHYSICAL"
	CardTypeVirtual  CardType = "VIRTUAL"
)

// CorporateCard represents a credit or debit card issued to a corporate entity.
type CorporateCard struct {
	ID             string     `json:"id"`
	CorporateID    string     `json:"corporate_id"`
	HolderName     string     `json:"holder_name"`
	LastFourDigits string     `json:"last_four_digits"`
	ExpiryMonth    int        `json:"expiry_month"`
	ExpiryYear     int        `json:"expiry_year"`
	Status         CardStatus `json:"status"`
	Type           CardType   `json:"type"`
	SpendingLimit  Money      `json:"spending_limit"`
	CurrentSpend   Money      `json:"current_spend"`
	IssuedAt       time.Time  `json:"issued_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// IsActive checks if the card is usable.
func (c *CorporateCard) IsActive() bool {
	return c.Status == CardStatusActive
}

// IsExpired checks if the card is past its expiry date.
func (c *CorporateCard) IsExpired() bool {
	now := time.Now()
	// Expiry is typically the end of the expiry month
	expiryDate := time.Date(c.ExpiryYear, time.Month(c.ExpiryMonth)+1, 0, 23, 59, 59, 0, time.UTC)
	return now.After(expiryDate)
}

// CanAuthorize checks if a transaction amount can be authorized.
func (c *CorporateCard) CanAuthorize(amount int64, currency string) error {
	if !c.IsActive() {
		return ErrCardInactive
	}
	if c.IsExpired() {
		return ErrCardExpired
	}
	if c.SpendingLimit.Currency != currency {
		return fmt.Errorf("currency mismatch: card is %s, tx is %s", c.SpendingLimit.Currency, currency)
	}
	if c.CurrentSpend.Amount+amount > c.SpendingLimit.Amount {
		return ErrCardLimitExceeded
	}
	return nil
}

// -----------------------------------------------------------------------------
// Anomaly & Fraud Detection Domain
// -----------------------------------------------------------------------------

type AnomalySeverity string
type AnomalyStatus string

const (
	SeverityLow      AnomalySeverity = "LOW"
	SeverityMedium   AnomalySeverity = "MEDIUM"
	SeverityHigh     AnomalySeverity = "HIGH"
	SeverityCritical AnomalySeverity = "CRITICAL"

	AnomalyStatusDetected      AnomalyStatus = "DETECTED"
	AnomalyStatusInvestigating AnomalyStatus = "INVESTIGATING"
	AnomalyStatusConfirmed     AnomalyStatus = "CONFIRMED_FRAUD"
	AnomalyStatusFalsePositive AnomalyStatus = "FALSE_POSITIVE"
	AnomalyStatusResolved      AnomalyStatus = "RESOLVED"
)

// TransactionAnomaly represents a detected irregularity in transaction patterns.
type TransactionAnomaly struct {
	ID            string          `json:"id"`
	TransactionID string          `json:"transaction_id"`
	CorporateID   string          `json:"corporate_id"`
	Severity      AnomalySeverity `json:"severity"`
	Score         float64         `json:"score"` // 0.0 to 1.0
	Description   string          `json:"description"`
	DetectedAt    time.Time       `json:"detected_at"`
	Status        AnomalyStatus   `json:"status"`
	ResolvedAt    *time.Time      `json:"resolved_at,omitempty"`
	ResolvedBy    string          `json:"resolved_by,omitempty"`
}

// FlagAsFraud marks the anomaly as confirmed fraud.
func (a *TransactionAnomaly) FlagAsFraud(resolverID string) {
	now := time.Now()
	a.Status = AnomalyStatusConfirmed
	a.ResolvedAt = &now
	a.ResolvedBy = resolverID
}

// Dismiss marks the anomaly as a false positive.
func (a *TransactionAnomaly) Dismiss(resolverID string) {
	now := time.Now()
	a.Status = AnomalyStatusFalsePositive
	a.ResolvedAt = &now
	a.ResolvedBy = resolverID
}

// -----------------------------------------------------------------------------
// Compliance Domain (KYC/KYB/AML)
// -----------------------------------------------------------------------------

type ComplianceCheckType string
type ComplianceStatus string

const (
	CheckTypeKYC ComplianceCheckType = "KYC" // Know Your Customer
	CheckTypeKYB ComplianceCheckType = "KYB" // Know Your Business
	CheckTypeAML ComplianceCheckType = "AML" // Anti-Money Laundering

	ComplianceStatusPending  ComplianceStatus = "PENDING"
	ComplianceStatusApproved ComplianceStatus = "APPROVED"
	ComplianceStatusRejected ComplianceStatus = "REJECTED"
	ComplianceStatusReview   ComplianceStatus = "MANUAL_REVIEW"
)

// ComplianceProfile represents the regulatory standing of a corporate entity.
type ComplianceProfile struct {
	ID             string           `json:"id"`
	CorporateID    string           `json:"corporate_id"`
	Status         ComplianceStatus `json:"status"`
	RiskScore      int              `json:"risk_score"` // 0 (Low) - 100 (High)
	LastCheckedAt  time.Time        `json:"last_checked_at"`
	NextReviewDate time.Time        `json:"next_review_date"`
	Documents      []Document       `json:"documents"`
	AuditLog       []AuditEntry     `json:"audit_log"`
}

type Document struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	URL      string    `json:"url"`
	Uploaded time.Time `json:"uploaded_at"`
	Verified bool      `json:"verified"`
}

type AuditEntry struct {
	Action    string    `json:"action"`
	ActorID   string    `json:"actor_id"`
	Timestamp time.Time `json:"timestamp"`
	Note      string    `json:"note"`
}

// UpdateRiskScore updates the risk score and adjusts status if necessary.
func (cp *ComplianceProfile) UpdateRiskScore(score int) error {
	if score < 0 || score > 100 {
		return ErrInvalidRiskScore
	}
	cp.RiskScore = score
	cp.LastCheckedAt = time.Now()

	// Auto-flag for review if risk is high
	if score > 75 && cp.Status == ComplianceStatusApproved {
		cp.Status = ComplianceStatusReview
	}
	return nil
}

// IsCompliant returns true if the entity is allowed to transact.
func (cp *ComplianceProfile) IsCompliant() bool {
	return cp.Status == ComplianceStatusApproved
}

// -----------------------------------------------------------------------------
// Treasury Domain
// -----------------------------------------------------------------------------

type TreasuryAccountType string

const (
	TreasuryOperating TreasuryAccountType = "OPERATING"
	TreasuryReserve   TreasuryAccountType = "RESERVE"
	TreasuryYield     TreasuryAccountType = "YIELD_BEARING"
)

// TreasuryPosition represents a corporate treasury account or position.
type TreasuryPosition struct {
	ID             string              `json:"id"`
	CorporateID    string              `json:"corporate_id"`
	AccountType    TreasuryAccountType `json:"account_type"`
	Balance        Money               `json:"balance"`
	LockedBalance  Money               `json:"locked_balance"`
	InterestRate   float64             `json:"interest_rate_bps"` // Basis points
	LastRebalanced time.Time           `json:"last_rebalanced"`
	IsLocked       bool                `json:"is_locked"`
}

// AllocateFunds moves funds from available balance to locked balance (e.g., for pending investments).
func (t *TreasuryPosition) AllocateFunds(amount int64) error {
	if t.IsLocked {
		return ErrTreasuryLocked
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if t.Balance.Amount < amount {
		return errors.New("insufficient treasury funds")
	}

	t.Balance.Amount -= amount
	t.LockedBalance.Amount += amount
	return nil
}

// ReleaseFunds moves funds from locked back to available.
func (t *TreasuryPosition) ReleaseFunds(amount int64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if t.LockedBalance.Amount < amount {
		return errors.New("insufficient locked funds")
	}

	t.LockedBalance.Amount -= amount
	t.Balance.Amount += amount
	return nil
}

// CalculateAccruedInterest estimates interest earned since last rebalance.
// This is a simplified calculation for domain modeling purposes.
func (t *TreasuryPosition) CalculateAccruedInterest(now time.Time) Money {
	if t.InterestRate <= 0 {
		return Money{Amount: 0, Currency: t.Balance.Currency}
	}

	daysElapsed := now.Sub(t.LastRebalanced).Hours() / 24
	if daysElapsed <= 0 {
		return Money{Amount: 0, Currency: t.Balance.Currency}
	}

	// Formula: Principal * (Rate in BPS / 10000) * (Days / 365)
	principal := float64(t.Balance.Amount + t.LockedBalance.Amount)
	rateDecimal := t.InterestRate / 10000.0
	interest := principal * rateDecimal * (daysElapsed / 365.0)

	return Money{
		Amount:   int64(interest),
		Currency: t.Balance.Currency,
	}
}

// -----------------------------------------------------------------------------
// Value Objects & Helpers
// -----------------------------------------------------------------------------

// MaskCardNumber returns a masked version of a PAN (e.g., ************1234).
func MaskCardNumber(pan string) string {
	clean := regexp.MustCompile(`\D`).ReplaceAllString(pan, "")
	if len(clean) < 4 {
		return clean
	}
	return strings.Repeat("*", len(clean)-4) + clean[len(clean)-4:]
}