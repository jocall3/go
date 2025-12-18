package identity

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

// Common errors for the KYC service.
var (
	ErrKYCAlreadySubmitted = errors.New("KYC verification is already in progress or approved")
	ErrKYCNotFound         = errors.New("KYC record not found")
	ErrInvalidInput        = errors.New("invalid KYC submission data")
)

// KYCStatus represents the state of a KYC verification.
type KYCStatus string

const (
	KYCStatusDraft    KYCStatus = "DRAFT"
	KYCStatusPending  KYCStatus = "PENDING"
	KYCStatusReview   KYCStatus = "MANUAL_REVIEW"
	KYCStatusApproved KYCStatus = "APPROVED"
	KYCStatusRejected KYCStatus = "REJECTED"
)

// KYCSubmission represents the domain entity for a KYC record.
type KYCSubmission struct {
	ID             string
	UserID         string
	Status         KYCStatus
	FirstName      string
	LastName       string
	DateOfBirth    time.Time
	DocumentType   string
	DocumentNumber string
	DocumentURL    string
	SubmittedAt    time.Time
	UpdatedAt      time.Time
	ReviewNotes    string
}

// SubmitKYCRequest holds the payload for a new KYC submission.
type SubmitKYCRequest struct {
	UserID         string
	FirstName      string
	LastName       string
	DateOfBirth    time.Time
	DocumentType   string
	DocumentNumber string
	DocumentURL    string
}

// Validate checks if the request data is valid.
func (r *SubmitKYCRequest) Validate() error {
	if r.UserID == "" {
		return fmt.Errorf("%w: user ID is required", ErrInvalidInput)
	}
	if r.FirstName == "" || r.LastName == "" {
		return fmt.Errorf("%w: full name is required", ErrInvalidInput)
	}
	if r.DocumentNumber == "" || r.DocumentURL == "" {
		return fmt.Errorf("%w: document details are required", ErrInvalidInput)
	}
	if r.DateOfBirth.After(time.Now().AddDate(-18, 0, 0)) {
		return fmt.Errorf("%w: user must be at least 18 years old", ErrInvalidInput)
	}
	return nil
}

// KYCRepository defines the data access layer requirements for KYC operations.
type KYCRepository interface {
	Save(ctx context.Context, submission *KYCSubmission) error
	GetByUserID(ctx context.Context, userID string) (*KYCSubmission, error)
	UpdateStatus(ctx context.Context, id string, status KYCStatus, notes string) error
}

// KYCService defines the business logic for KYC operations.
type KYCService interface {
	SubmitKYC(ctx context.Context, req SubmitKYCRequest) (*KYCSubmission, error)
	GetKYCStatus(ctx context.Context, userID string) (*KYCSubmission, error)
	ProcessDecision(ctx context.Context, submissionID string, approved bool, notes string) error
}

// kycService is the concrete implementation of KYCService.
type kycService struct {
	repo   KYCRepository
	logger *slog.Logger
}

// NewKYCService creates a new instance of the KYC service.
func NewKYCService(repo KYCRepository, logger *slog.Logger) KYCService {
	return &kycService{
		repo:   repo,
		logger: logger,
	}
}

// SubmitKYC handles the submission of user documents for verification.
func (s *kycService) SubmitKYC(ctx context.Context, req SubmitKYCRequest) (*KYCSubmission, error) {
	logger := s.logger.With("method", "SubmitKYC", "user_id", req.UserID)

	if err := req.Validate(); err != nil {
		logger.Warn("invalid kyc request", "error", err)
		return nil, err
	}

	// Check for existing submission
	existing, err := s.repo.GetByUserID(ctx, req.UserID)
	if err != nil && !errors.Is(err, ErrKYCNotFound) {
		logger.Error("failed to check existing kyc", "error", err)
		return nil, fmt.Errorf("failed to check existing records: %w", err)
	}

	if existing != nil {
		// If already approved or pending, deny new submission
		if existing.Status == KYCStatusApproved || existing.Status == KYCStatusPending || existing.Status == KYCStatusReview {
			logger.Info("duplicate kyc submission attempt", "status", existing.Status)
			return nil, ErrKYCAlreadySubmitted
		}
		// If rejected, we might allow re-submission, but for this implementation, we create a new record.
	}

	submission := &KYCSubmission{
		ID:             uuid.New().String(),
		UserID:         req.UserID,
		Status:         KYCStatusPending, // Default to pending for async verification
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DateOfBirth:    req.DateOfBirth,
		DocumentType:   req.DocumentType,
		DocumentNumber: req.DocumentNumber,
		DocumentURL:    req.DocumentURL,
		SubmittedAt:    time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	if err := s.repo.Save(ctx, submission); err != nil {
		logger.Error("failed to save kyc submission", "error", err)
		return nil, fmt.Errorf("failed to save submission: %w", err)
	}

	logger.Info("kyc submitted successfully", "submission_id", submission.ID)

	// In a real-world scenario, you might trigger an async job here (e.g., via a message queue)
	// to send the data to a 3rd party KYC provider (SumSub, Onfido, etc.).

	return submission, nil
}

// GetKYCStatus retrieves the current KYC status for a user.
func (s *kycService) GetKYCStatus(ctx context.Context, userID string) (*KYCSubmission, error) {
	logger := s.logger.With("method", "GetKYCStatus", "user_id", userID)

	submission, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, ErrKYCNotFound) {
			return nil, ErrKYCNotFound
		}
		logger.Error("failed to retrieve kyc record", "error", err)
		return nil, fmt.Errorf("failed to retrieve kyc status: %w", err)
	}

	return submission, nil
}

// ProcessDecision allows an admin or system webhook to update the status of a KYC submission.
func (s *kycService) ProcessDecision(ctx context.Context, submissionID string, approved bool, notes string) error {
	logger := s.logger.With("method", "ProcessDecision", "submission_id", submissionID)

	status := KYCStatusRejected
	if approved {
		status = KYCStatusApproved
	}

	logger.Info("processing kyc decision", "new_status", status)

	if err := s.repo.UpdateStatus(ctx, submissionID, status, notes); err != nil {
		logger.Error("failed to update kyc status", "error", err)
		return fmt.Errorf("failed to update status: %w", err)
	}

	// Here you would typically emit an event (e.g., IdentityVerifiedEvent)
	// so that other services (Accounts, Wallets) can unlock features.

	return nil
}