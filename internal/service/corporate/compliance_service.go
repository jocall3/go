package corporate

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

// Common errors for the compliance service.
var (
	ErrInvalidScreeningRequest = errors.New("invalid screening request: missing required fields")
	ErrSanctionCheckFailed     = errors.New("failed to perform sanction screening")
	ErrAuditLogFailed          = errors.New("failed to record audit log")
)

// RiskLevel represents the calculated risk associated with an entity.
type RiskLevel string

const (
	RiskLevelLow      RiskLevel = "LOW"
	RiskLevelMedium   RiskLevel = "MEDIUM"
	RiskLevelHigh     RiskLevel = "HIGH"
	RiskLevelCritical RiskLevel = "CRITICAL"
)

// EntityType defines the type of entity being screened.
type EntityType string

const (
	EntityTypeIndividual EntityType = "INDIVIDUAL"
	EntityTypeCompany    EntityType = "COMPANY"
)

// ScreeningRequest contains the data required to screen an entity.
type ScreeningRequest struct {
	ReferenceID string            `json:"reference_id"`
	Name        string            `json:"name"`
	Type        EntityType        `json:"type"`
	CountryCode string            `json:"country_code"` // ISO 3166-1 alpha-2
	TaxID       string            `json:"tax_id,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// SanctionMatch represents a potential match found in a sanction list.
type SanctionMatch struct {
	ListSource   string  `json:"list_source"`
	EntityName   string  `json:"entity_name"`
	MatchScore   float64 `json:"match_score"` // 0.0 to 1.0
	Reason       string  `json:"reason"`
	SanctionDate string  `json:"sanction_date"`
}

// ScreeningResult is the outcome of a compliance check.
type ScreeningResult struct {
	ScreeningID  string          `json:"screening_id"`
	Timestamp    time.Time       `json:"timestamp"`
	RiskLevel    RiskLevel       `json:"risk_level"`
	IsSanctioned bool            `json:"is_sanctioned"`
	Matches      []SanctionMatch `json:"matches"`
	Flags        []string        `json:"flags"`
}

// AuditEvent represents an immutable record of a compliance action.
type AuditEvent struct {
	ID        string                 `json:"id"`
	Timestamp time.Time              `json:"timestamp"`
	ActorID   string                 `json:"actor_id"`
	Action    string                 `json:"action"`
	Resource  string                 `json:"resource"`
	Payload   map[string]interface{} `json:"payload"`
	Hash      string                 `json:"hash"` // HMAC signature for integrity
}

// SanctionRepository defines the data access layer for sanction lists.
type SanctionRepository interface {
	// SearchSanctionLists searches for entities matching the name and country.
	SearchSanctionLists(ctx context.Context, name string, countryCode string) ([]SanctionMatch, error)
	// IsHighRiskCountry checks if the country is on a watchlist (e.g., FATF).
	IsHighRiskCountry(ctx context.Context, countryCode string) (bool, error)
}

// AuditRepository defines the storage for compliance audit logs.
type AuditRepository interface {
	SaveEvent(ctx context.Context, event AuditEvent) error
	GetEventsByResource(ctx context.Context, resourceID string) ([]AuditEvent, error)
}

// ComplianceService defines the business logic for compliance operations.
type ComplianceService interface {
	// ScreenEntity performs a full compliance check on an entity.
	ScreenEntity(ctx context.Context, req ScreeningRequest) (*ScreeningResult, error)
	// LogAction records a compliance-related action for audit purposes.
	LogAction(ctx context.Context, actorID, action, resource string, details map[string]interface{}) error
	// GetAuditTrail retrieves the history of actions for a specific resource.
	GetAuditTrail(ctx context.Context, resourceID string) ([]AuditEvent, error)
}

// complianceService implements the ComplianceService interface.
type complianceService struct {
	sanctionRepo SanctionRepository
	auditRepo    AuditRepository
	logger       *slog.Logger
	hmacSecret   []byte
}

// NewComplianceService creates a new instance of the compliance service.
func NewComplianceService(
	sr SanctionRepository,
	ar AuditRepository,
	logger *slog.Logger,
	secretKey string,
) ComplianceService {
	return &complianceService{
		sanctionRepo: sr,
		auditRepo:    ar,
		logger:       logger,
		hmacSecret:   []byte(secretKey),
	}
}

// ScreenEntity orchestrates the sanction screening process.
func (s *complianceService) ScreenEntity(ctx context.Context, req ScreeningRequest) (*ScreeningResult, error) {
	start := time.Now()
	log := s.logger.With("method", "ScreenEntity", "ref_id", req.ReferenceID)

	if req.Name == "" || req.CountryCode == "" {
		log.Warn("Invalid screening request received")
		return nil, ErrInvalidScreeningRequest
	}

	// 1. Check Country Risk
	isHighRiskCountry, err := s.sanctionRepo.IsHighRiskCountry(ctx, req.CountryCode)
	if err != nil {
		log.Error("Failed to check country risk", "error", err)
		return nil, fmt.Errorf("%w: country check", ErrSanctionCheckFailed)
	}

	// 2. Search Sanction Lists
	matches, err := s.sanctionRepo.SearchSanctionLists(ctx, req.Name, req.CountryCode)
	if err != nil {
		log.Error("Failed to search sanction lists", "error", err)
		return nil, fmt.Errorf("%w: list search", ErrSanctionCheckFailed)
	}

	// 3. Analyze Results
	result := &ScreeningResult{
		ScreeningID: generateID(), // Helper function assumed
		Timestamp:   time.Now(),
		Matches:     matches,
		Flags:       make([]string, 0),
	}

	if isHighRiskCountry {
		result.Flags = append(result.Flags, "HIGH_RISK_JURISDICTION")
	}

	// Determine Risk Level
	result.RiskLevel = calculateRiskLevel(matches, isHighRiskCountry)
	result.IsSanctioned = result.RiskLevel == RiskLevelCritical || result.RiskLevel == RiskLevelHigh

	// 4. Audit the Screening
	auditPayload := map[string]interface{}{
		"request":    req,
		"result_id":  result.ScreeningID,
		"risk_level": result.RiskLevel,
		"match_count": len(matches),
	}
	
	// We log the screening action internally as 'SYSTEM'
	if err := s.LogAction(ctx, "SYSTEM", "SCREEN_ENTITY", req.ReferenceID, auditPayload); err != nil {
		// We don't fail the screening if audit fails, but we log the error critically
		log.Error("Failed to audit screening result", "error", err)
	}

	log.Info("Screening completed", 
		"risk_level", result.RiskLevel, 
		"matches", len(matches), 
		"duration", time.Since(start))

	return result, nil
}

// LogAction creates a tamper-evident audit log entry.
func (s *complianceService) LogAction(ctx context.Context, actorID, action, resource string, details map[string]interface{}) error {
	event := AuditEvent{
		ID:        generateID(),
		Timestamp: time.Now().UTC(),
		ActorID:   actorID,
		Action:    action,
		Resource:  resource,
		Payload:   details,
	}

	// Generate integrity hash
	hash, err := s.generateEventHash(event)
	if err != nil {
		return fmt.Errorf("failed to generate audit hash: %w", err)
	}
	event.Hash = hash

	if err := s.auditRepo.SaveEvent(ctx, event); err != nil {
		s.logger.Error("Failed to persist audit event", "error", err, "action", action)
		return ErrAuditLogFailed
	}

	return nil
}

// GetAuditTrail retrieves logs and verifies their integrity (optional implementation detail).
func (s *complianceService) GetAuditTrail(ctx context.Context, resourceID string) ([]AuditEvent, error) {
	events, err := s.auditRepo.GetEventsByResource(ctx, resourceID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve audit trail: %w", err)
	}

	// Verify integrity of fetched events
	for i, event := range events {
		expectedHash, err := s.generateEventHash(event)
		if err != nil {
			s.logger.Warn("Could not verify hash for event", "event_id", event.ID)
			continue
		}
		if event.Hash != expectedHash {
			s.logger.Error("Audit log integrity failure detected", "event_id", event.ID)
			// In a strict system, we might return an error or mark the event as compromised.
			// For now, we flag it in the logs.
			events[i].Payload["_integrity_check"] = "FAILED"
		}
	}

	return events, nil
}

// Internal Helpers

func calculateRiskLevel(matches []SanctionMatch, highRiskCountry bool) RiskLevel {
	if len(matches) > 0 {
		// Check for exact matches or high scores
		for _, m := range matches {
			if m.MatchScore >= 0.95 {
				return RiskLevelCritical
			}
			if m.MatchScore >= 0.80 {
				return RiskLevelHigh
			}
		}
		return RiskLevelMedium
	}

	if highRiskCountry {
		return RiskLevelHigh
	}

	return RiskLevelLow
}

func (s *complianceService) generateEventHash(event AuditEvent) (string, error) {
	// Create a canonical string representation of the critical fields
	payloadBytes, err := json.Marshal(event.Payload)
	if err != nil {
		return "", err
	}

	data := fmt.Sprintf("%s|%s|%s|%s|%s|%s",
		event.ID,
		event.Timestamp.Format(time.RFC3339Nano),
		event.ActorID,
		event.Action,
		event.Resource,
		string(payloadBytes),
	)

	h := hmac.New(sha256.New, s.hmacSecret)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil)), nil
}

// generateID generates a pseudo-unique identifier.
// In a real production app, use "github.com/google/uuid".
// Using a simple implementation here to ensure zero external dependencies if required by the prompt constraints.
func generateID() string {
	now := time.Now().UnixNano()
	return fmt.Sprintf("evt-%x", now)
}