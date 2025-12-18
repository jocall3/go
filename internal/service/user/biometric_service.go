package user

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

// Common errors for biometric operations.
var (
	ErrBiometricNotFound      = errors.New("biometric credential not found")
	ErrInvalidBiometricParams = errors.New("invalid biometric parameters")
	ErrBiometricVerification  = errors.New("biometric verification failed")
	ErrBiometricOwnership     = errors.New("biometric credential does not belong to user")
)

// BiometricType defines the category of biometric data.
type BiometricType string

const (
	BiometricTypeFingerprint BiometricType = "FINGERPRINT"
	BiometricTypeFaceID      BiometricType = "FACE_ID"
	BiometricTypeIris        BiometricType = "IRIS"
	BiometricTypeFIDO2       BiometricType = "FIDO2"
)

// BiometricCredential represents the stored metadata and public key for a user's biometric authenticator.
type BiometricCredential struct {
	ID          string        `json:"id"`
	UserID      string        `json:"user_id"`
	Type        BiometricType `json:"type"`
	Label       string        `json:"label"`       // User-friendly name (e.g., "John's iPhone")
	PublicKey   []byte        `json:"public_key"`  // Raw public key bytes
	Attestation []byte        `json:"attestation"` // Attestation object if available
	DeviceID    string        `json:"device_id"`
	Counter     uint32        `json:"counter"` // Sign count for replay protection
	CreatedAt   time.Time     `json:"created_at"`
	LastUsedAt  time.Time     `json:"last_used_at"`
}

// BiometricRepository defines the data access layer requirements for biometric credentials.
type BiometricRepository interface {
	Create(ctx context.Context, cred *BiometricCredential) error
	GetByID(ctx context.Context, id string) (*BiometricCredential, error)
	GetByUserID(ctx context.Context, userID string) ([]*BiometricCredential, error)
	Delete(ctx context.Context, id string) error
	UpdateCounterAndUsage(ctx context.Context, id string, newCounter uint32, lastUsed time.Time) error
}

// SignatureVerifier defines the strategy for verifying cryptographic signatures.
// This allows injecting specific logic for WebAuthn, TPM, or raw key verification.
type SignatureVerifier interface {
	// Verify checks if the signature is valid for the given data and public key.
	Verify(ctx context.Context, pubKey []byte, data []byte, signature []byte) (bool, error)
}

// BiometricService manages the lifecycle of user biometric credentials.
type BiometricService struct {
	repo     BiometricRepository
	verifier SignatureVerifier
	logger   *slog.Logger
}

// NewBiometricService creates a new instance of BiometricService.
func NewBiometricService(repo BiometricRepository, verifier SignatureVerifier, logger *slog.Logger) *BiometricService {
	return &BiometricService{
		repo:     repo,
		verifier: verifier,
		logger:   logger,
	}
}

// EnrollParams contains the necessary data to register a new biometric credential.
type EnrollParams struct {
	UserID      string
	Type        BiometricType
	Label       string
	PublicKey   []byte
	Attestation []byte
	DeviceID    string
}

// Enroll registers a new biometric credential for a user.
func (s *BiometricService) Enroll(ctx context.Context, params EnrollParams) (*BiometricCredential, error) {
	if params.UserID == "" || len(params.PublicKey) == 0 {
		return nil, ErrInvalidBiometricParams
	}

	// Generate a secure unique identifier for the credential
	credID, err := generateSecureID()
	if err != nil {
		s.logger.Error("failed to generate credential ID", "error", err)
		return nil, fmt.Errorf("generation failed: %w", err)
	}

	cred := &BiometricCredential{
		ID:          credID,
		UserID:      params.UserID,
		Type:        params.Type,
		Label:       params.Label,
		PublicKey:   params.PublicKey,
		Attestation: params.Attestation,
		DeviceID:    params.DeviceID,
		Counter:     0,
		CreatedAt:   time.Now().UTC(),
		LastUsedAt:  time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, cred); err != nil {
		s.logger.Error("failed to persist biometric credential", "user_id", params.UserID, "error", err)
		return nil, fmt.Errorf("persistence failed: %w", err)
	}

	s.logger.Info("biometric credential enrolled", "id", cred.ID, "user_id", cred.UserID, "type", cred.Type)
	return cred, nil
}

// VerifyParams contains the data required to verify a biometric assertion.
type VerifyParams struct {
	CredentialID string
	Challenge    []byte // The original challenge sent to the client
	Signature    []byte // The signature returned by the client
	ClientData   []byte // Additional client data (e.g., collected client data in WebAuthn)
}

// VerifyAuthenticate checks if the provided signature matches the stored credential.
// It updates the usage timestamp and counter upon success.
func (s *BiometricService) VerifyAuthenticate(ctx context.Context, params VerifyParams) (bool, error) {
	cred, err := s.repo.GetByID(ctx, params.CredentialID)
	if err != nil {
		if errors.Is(err, ErrBiometricNotFound) {
			return false, ErrBiometricNotFound
		}
		s.logger.Error("failed to retrieve credential", "id", params.CredentialID, "error", err)
		return false, err
	}

	// Construct the data that was signed.
	// Note: In a real WebAuthn implementation, this involves parsing the authenticator data
	// and hashing the client data JSON. Here we assume the verifier handles the protocol specifics
	// or that 'ClientData' combined with 'Challenge' represents the signed payload.
	signedPayload := append(params.ClientData, params.Challenge...)

	valid, err := s.verifier.Verify(ctx, cred.PublicKey, signedPayload, params.Signature)
	if err != nil {
		s.logger.Warn("verification process error", "id", cred.ID, "error", err)
		return false, ErrBiometricVerification
	}

	if !valid {
		s.logger.Warn("invalid signature detected", "id", cred.ID, "user_id", cred.UserID)
		return false, nil
	}

	// Update usage stats
	// Note: Real implementations should also validate and increment the counter to prevent replay attacks.
	// We assume the counter is extracted from the signature payload in a full implementation.
	newCounter := cred.Counter + 1
	if err := s.repo.UpdateCounterAndUsage(ctx, cred.ID, newCounter, time.Now().UTC()); err != nil {
		s.logger.Error("failed to update credential usage", "id", cred.ID, "error", err)
		// We do not fail the authentication if the update fails, but we log it.
	}

	s.logger.Info("biometric authentication successful", "user_id", cred.UserID, "cred_id", cred.ID)
	return true, nil
}

// Deregister removes a specific biometric credential.
// It enforces ownership checks to ensure a user can only delete their own credentials.
func (s *BiometricService) Deregister(ctx context.Context, userID string, credentialID string) error {
	cred, err := s.repo.GetByID(ctx, credentialID)
	if err != nil {
		return err
	}

	if cred.UserID != userID {
		s.logger.Warn("unauthorized deregister attempt", "user_id", userID, "cred_owner", cred.UserID, "cred_id", credentialID)
		return ErrBiometricOwnership
	}

	if err := s.repo.Delete(ctx, credentialID); err != nil {
		s.logger.Error("failed to delete credential", "id", credentialID, "error", err)
		return err
	}

	s.logger.Info("biometric credential deregistered", "user_id", userID, "cred_id", credentialID)
	return nil
}

// ListCredentials retrieves all enrolled biometrics for a specific user.
func (s *BiometricService) ListCredentials(ctx context.Context, userID string) ([]*BiometricCredential, error) {
	creds, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		s.logger.Error("failed to list credentials", "user_id", userID, "error", err)
		return nil, err
	}
	return creds, nil
}

// generateSecureID creates a random hex string for use as an ID.
func generateSecureID() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}