package identity

import (
	"errors"
	"strings"
	"time"
)

// KYCStatus represents the lifecycle state of an Identity Verification request.
type KYCStatus string

const (
	// KYCStatusDraft indicates the submission is being prepared but not yet finalized.
	KYCStatusDraft KYCStatus = "DRAFT"
	// KYCStatusPending indicates the submission has been received and is awaiting processing.
	KYCStatusPending KYCStatus = "PENDING"
	// KYCStatusInReview indicates the submission is currently being analyzed by an agent or automated system.
	KYCStatusInReview KYCStatus = "IN_REVIEW"
	// KYCStatusApproved indicates the identity has been successfully verified.
	KYCStatusApproved KYCStatus = "APPROVED"
	// KYCStatusRejected indicates the verification failed permanently for this submission.
	KYCStatusRejected KYCStatus = "REJECTED"
	// KYCStatusRequiresAction indicates the user must provide additional information or re-upload documents.
	KYCStatusRequiresAction KYCStatus = "REQUIRES_ACTION"
)

// IsTerminal returns true if the status represents a completed state.
func (s KYCStatus) IsTerminal() bool {
	return s == KYCStatusApproved || s == KYCStatusRejected
}

// DocumentType defines the category of government-issued ID provided.
type DocumentType string

const (
	DocumentTypePassport        DocumentType = "PASSPORT"
	DocumentTypeDriversLicense  DocumentType = "DRIVERS_LICENSE"
	DocumentTypeNationalID      DocumentType = "NATIONAL_ID"
	DocumentTypeResidencePermit DocumentType = "RESIDENCE_PERMIT"
)

// KYCRecord represents the aggregate root for a user's Know Your Customer (KYC) verification process.
// It tracks the history, current status, and outcome of the verification.
type KYCRecord struct {
	// ID is the unique identifier for this KYC attempt.
	ID string `json:"id"`

	// UserID links this record to a specific user in the system.
	UserID string `json:"user_id"`

	// Status denotes the current progress of the verification.
	Status KYCStatus `json:"status"`

	// CurrentTier represents the verification level achieved (e.g., 1 for basic, 2 for advanced).
	CurrentTier int `json:"current_tier"`

	// Submission contains the data provided by the user.
	Submission KYCSubmission `json:"submission"`

	// VerificationResult contains details from the verification provider/admin.
	VerificationResult *VerificationResult `json:"verification_result,omitempty"`

	// Audit metadata.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// KYCSubmission encapsulates the input data required to perform identity verification.
type KYCSubmission struct {
	DocumentType   DocumentType `json:"document_type"`
	DocumentNumber string       `json:"document_number"`
	FirstName      string       `json:"first_name"`
	LastName       string       `json:"last_name"`
	
	// DateOfBirth is expected in ISO 8601 format (YYYY-MM-DD).
	DateOfBirth string `json:"date_of_birth"`
	
	// Nationality is the ISO 3166-1 alpha-2 country code.
	Nationality string `json:"nationality"`
	
	Address Address `json:"address"`

	// Secure URLs or references to the stored document images.
	DocumentFrontImageID string `json:"document_front_image_id"`
	DocumentBackImageID  string `json:"document_back_image_id,omitempty"`
	SelfieImageID        string `json:"selfie_image_id"`
	
	// ClientIP captures the IP address from where the submission originated (fraud prevention).
	ClientIP string `json:"client_ip,omitempty"`
}

// Address represents a physical location for residence verification.
type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2,omitempty"`
	City       string `json:"city"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"` // ISO 3166-1 alpha-2
}

// VerificationResult holds the outcome details of a review.
type VerificationResult struct {
	VerifiedAt      time.Time `json:"verified_at"`
	ReviewedBy      string    `json:"reviewed_by,omitempty"` // Admin ID or System
	RejectionReason string    `json:"rejection_reason,omitempty"`
	RiskScore       float64   `json:"risk_score,omitempty"` // 0.0 to 1.0
	Notes           string    `json:"notes,omitempty"`
}

// NewKYCRecord creates a new KYC record in pending state.
func NewKYCRecord(id, userID string, submission KYCSubmission) (*KYCRecord, error) {
	if err := submission.Validate(); err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	return &KYCRecord{
		ID:          id,
		UserID:      userID,
		Status:      KYCStatusPending,
		Submission:  submission,
		CreatedAt:   now,
		UpdatedAt:   now,
		CurrentTier: 0,
	}, nil
}

// Validate performs domain-level validation on the submission data.
func (s *KYCSubmission) Validate() error {
	if s.DocumentType == "" {
		return errors.New("document type is required")
	}
	if strings.TrimSpace(s.DocumentNumber) == "" {
		return errors.New("document number is required")
	}
	if strings.TrimSpace(s.FirstName) == "" || strings.TrimSpace(s.LastName) == "" {
		return errors.New("full name is required")
	}
	if s.DateOfBirth == "" {
		return errors.New("date of birth is required")
	}
	if len(s.Nationality) != 2 {
		return errors.New("nationality must be a valid 2-letter ISO country code")
	}
	if s.DocumentFrontImageID == "" {
		return errors.New("front document image is required")
	}
	if s.SelfieImageID == "" {
		return errors.New("selfie image is required")
	}
	
	// Basic Address Validation
	if strings.TrimSpace(s.Address.Line1) == "" {
		return errors.New("address line 1 is required")
	}
	if strings.TrimSpace(s.Address.City) == "" {
		return errors.New("city is required")
	}
	if strings.TrimSpace(s.Address.Country) == "" {
		return errors.New("address country is required")
	}

	return nil
}

// Approve transitions the record to Approved status.
func (r *KYCRecord) Approve(reviewerID string, tier int) error {
	if r.Status == KYCStatusApproved {
		return errors.New("record is already approved")
	}
	
	now := time.Now().UTC()
	r.Status = KYCStatusApproved
	r.CurrentTier = tier
	r.UpdatedAt = now
	r.VerificationResult = &VerificationResult{
		VerifiedAt: now,
		ReviewedBy: reviewerID,
	}
	return nil
}

// Reject transitions the record to Rejected status with a reason.
func (r *KYCRecord) Reject(reviewerID, reason string) error {
	if r.Status.IsTerminal() {
		return errors.New("cannot reject a record in a terminal state")
	}

	now := time.Now().UTC()
	r.Status = KYCStatusRejected
	r.UpdatedAt = now
	r.VerificationResult = &VerificationResult{
		VerifiedAt:      now,
		ReviewedBy:      reviewerID,
		RejectionReason: reason,
	}
	return nil
}

// RequestMoreInfo transitions the record to RequiresAction.
func (r *KYCRecord) RequestMoreInfo(reviewerID, notes string) error {
	if r.Status.IsTerminal() {
		return errors.New("cannot request info for a record in a terminal state")
	}

	now := time.Now().UTC()
	r.Status = KYCStatusRequiresAction
	r.UpdatedAt = now
	
	// We preserve existing verification result history if needed, 
	// but here we update the latest context.
	if r.VerificationResult == nil {
		r.VerificationResult = &VerificationResult{}
	}
	r.VerificationResult.ReviewedBy = reviewerID
	r.VerificationResult.Notes = notes
	
	return nil
}