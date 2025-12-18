package ai

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// =============================================================================
// Common Types & Enums
// =============================================================================

// ProcessingStatus represents the lifecycle state of an AI task.
type ProcessingStatus string

const (
	StatusPending    ProcessingStatus = "pending"
	StatusProcessing ProcessingStatus = "processing"
	StatusCompleted  ProcessingStatus = "completed"
	StatusFailed     ProcessingStatus = "failed"
)

// ModelProvider defines which underlying AI model is used (e.g., GPT-4, Claude, Llama).
type ModelProvider string

const (
	ProviderOpenAI    ModelProvider = "openai"
	ProviderAnthropic ModelProvider = "anthropic"
	ProviderLocal     ModelProvider = "local_llm"
)

// =============================================================================
// Feature: Advisor Chat
// Purpose: Strategic business advice from specific AI personas.
// =============================================================================

// AdvisorPersona defines the role the AI assumes during a chat session.
type AdvisorPersona string

const (
	PersonaGeneralStrategist AdvisorPersona = "general_strategist"
	PersonaLegalCounsel      AdvisorPersona = "legal_counsel"
	PersonaMarketingGuru     AdvisorPersona = "marketing_guru"
	PersonaTechLead          AdvisorPersona = "tech_lead"
	PersonaVentureCapitalist AdvisorPersona = "vc_investor"
)

// MessageRole indicates who sent the message.
type MessageRole string

const (
	RoleSystem    MessageRole = "system"
	RoleUser      MessageRole = "user"
	RoleAssistant MessageRole = "assistant"
)

// AdvisorSession represents a continuous conversation thread with a specific advisor persona.
type AdvisorSession struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	UserID    uuid.UUID      `json:"user_id" db:"user_id"`
	ProjectID *uuid.UUID     `json:"project_id,omitempty" db:"project_id"` // Optional context
	Persona   AdvisorPersona `json:"persona" db:"persona"`
	Topic     string         `json:"topic" db:"topic"`
	Summary   string         `json:"summary,omitempty" db:"summary"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}

// AdvisorMessage represents a single exchange within a session.
type AdvisorMessage struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	SessionID uuid.UUID   `json:"session_id" db:"session_id"`
	Role      MessageRole `json:"role" db:"role"`
	Content   string      `json:"content" db:"content"`
	// TokenUsage tracks cost/complexity
	TokenUsage int       `json:"token_usage" db:"token_usage"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// NewAdvisorSession creates a new session instance.
func NewAdvisorSession(userID uuid.UUID, persona AdvisorPersona, topic string) *AdvisorSession {
	return &AdvisorSession{
		ID:        uuid.New(),
		UserID:    userID,
		Persona:   persona,
		Topic:     topic,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

// =============================================================================
// Feature: Oracle Simulations
// Purpose: Predictive modeling and scenario analysis for business decisions.
// =============================================================================

// SimulationType defines the category of the simulation.
type SimulationType string

const (
	SimMarketEntry   SimulationType = "market_entry"
	SimPricingModel  SimulationType = "pricing_model"
	SimCrisisMgmt    SimulationType = "crisis_management"
	SimCompetitorWar SimulationType = "competitor_wargame"
)

// OracleSimulation represents a request to run a predictive scenario.
type OracleSimulation struct {
	ID             uuid.UUID        `json:"id" db:"id"`
	ProjectID      uuid.UUID        `json:"project_id" db:"project_id"`
	Type           SimulationType   `json:"type" db:"type"`
	Name           string           `json:"name" db:"name"`
	Parameters     json.RawMessage  `json:"parameters" db:"parameters"` // Specific inputs for the sim
	Status         ProcessingStatus `json:"status" db:"status"`
	ResultSummary  string           `json:"result_summary,omitempty" db:"result_summary"`
	Confidence     float64          `json:"confidence_score" db:"confidence_score"` // 0.0 to 1.0
	ExecutionTime  int64            `json:"execution_time_ms" db:"execution_time_ms"`
	CreatedAt      time.Time        `json:"created_at" db:"created_at"`
	CompletedAt    *time.Time       `json:"completed_at,omitempty" db:"completed_at"`
}

// SimulationOutcome details the specific results of a simulation run.
type SimulationOutcome struct {
	ID           uuid.UUID       `json:"id" db:"id"`
	SimulationID uuid.UUID       `json:"simulation_id" db:"simulation_id"`
	Scenario     string          `json:"scenario" db:"scenario"` // e.g., "Best Case", "Worst Case"
	Probability  float64         `json:"probability" db:"probability"`
	Impact       string          `json:"impact" db:"impact"` // High, Medium, Low
	Metrics      json.RawMessage `json:"metrics" db:"metrics"` // Key-value pairs of projected numbers
	Narrative    string          `json:"narrative" db:"narrative"`
}

// =============================================================================
// Feature: Incubator Pitches
// Purpose: Refining pitch decks and business models.
// =============================================================================

// PitchDeck represents the user's business proposal.
type PitchDeck struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Industry    string    `json:"industry" db:"industry"`
	Problem     string    `json:"problem" db:"problem"`
	Solution    string    `json:"solution" db:"solution"`
	TargetMrkt  string    `json:"target_market" db:"target_market"`
	BusinessMod string    `json:"business_model" db:"business_model"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// PitchAnalysis contains the AI's critique and scoring of a pitch.
type PitchAnalysis struct {
	ID          uuid.UUID `json:"id" db:"id"`
	PitchDeckID uuid.UUID `json:"pitch_deck_id" db:"pitch_deck_id"`
	OverallScore int      `json:"overall_score" db:"overall_score"` // 0-100
	Clarity      int      `json:"clarity_score" db:"clarity_score"`
	Viability    int      `json:"viability_score" db:"viability_score"`
	Innovation   int      `json:"innovation_score" db:"innovation_score"`
	Strengths    []string `json:"strengths" db:"strengths"`     // Stored as JSON array in DB
	Weaknesses   []string `json:"weaknesses" db:"weaknesses"`   // Stored as JSON array in DB
	Suggestions  []string `json:"suggestions" db:"suggestions"` // Stored as JSON array in DB
	InvestorPOV  string   `json:"investor_pov" db:"investor_pov"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// =============================================================================
// Feature: Ad Generation
// Purpose: Creating marketing copy and visual prompts.
// =============================================================================

// AdPlatform defines where the ad is intended to run.
type AdPlatform string

const (
	PlatformGoogleSearch AdPlatform = "google_search"
	PlatformFacebook     AdPlatform = "facebook_feed"
	PlatformInstagram    AdPlatform = "instagram_story"
	PlatformLinkedIn     AdPlatform = "linkedin_sponsored"
	PlatformTwitter      AdPlatform = "twitter_post"
	PlatformEmail        AdPlatform = "email_newsletter"
)

// AdCampaign represents a container for generated creatives.
type AdCampaign struct {
	ID             uuid.UUID        `json:"id" db:"id"`
	ProjectID      uuid.UUID        `json:"project_id" db:"project_id"`
	Name           string           `json:"name" db:"name"`
	Objective      string           `json:"objective" db:"objective"` // e.g., "Brand Awareness", "Conversion"
	TargetAudience string           `json:"target_audience" db:"target_audience"`
	Tone           string           `json:"tone" db:"tone"`
	Status         ProcessingStatus `json:"status" db:"status"`
	CreatedAt      time.Time        `json:"created_at" db:"created_at"`
}

// AdCreative represents a single generated ad variation.
type AdCreative struct {
	ID           uuid.UUID       `json:"id" db:"id"`
	CampaignID   uuid.UUID       `json:"campaign_id" db:"campaign_id"`
	Platform     AdPlatform      `json:"platform" db:"platform"`
	Headline     string          `json:"headline" db:"headline"`
	BodyCopy     string          `json:"body_copy" db:"body_copy"`
	CallToAction string          `json:"call_to_action" db:"call_to_action"`
	ImagePrompt  string          `json:"image_prompt,omitempty" db:"image_prompt"` // For DALL-E/Midjourney
	Keywords     []string        `json:"keywords,omitempty" db:"keywords"`
	Metadata     json.RawMessage `json:"metadata,omitempty" db:"metadata"` // Platform specific fields
	CreatedAt    time.Time       `json:"created_at" db:"created_at"`
}

// =============================================================================
// Domain Errors
// =============================================================================

var (
	ErrSessionNotFound    = errors.New("advisor session not found")
	ErrSimulationFailed   = errors.New("simulation execution failed")
	ErrInvalidPitchData   = errors.New("pitch deck data is incomplete")
	ErrUnsupportedPlatform = errors.New("unsupported ad platform")
)

// =============================================================================
// Repository Interface
// =============================================================================

// Repository defines the persistence contract for AI domain entities.
type Repository interface {
	// Advisor
	SaveSession(session *AdvisorSession) error
	GetSession(id uuid.UUID) (*AdvisorSession, error)
	SaveMessage(msg *AdvisorMessage) error
	GetHistory(sessionID uuid.UUID) ([]AdvisorMessage, error)

	// Oracle
	CreateSimulation(sim *OracleSimulation) error
	UpdateSimulationStatus(id uuid.UUID, status ProcessingStatus, result string) error
	SaveSimulationOutcomes(outcomes []SimulationOutcome) error

	// Incubator
	SavePitch(pitch *PitchDeck) error
	SavePitchAnalysis(analysis *PitchAnalysis) error
	GetPitchAnalysis(pitchID uuid.UUID) (*PitchAnalysis, error)

	// Ad Gen
	CreateCampaign(campaign *AdCampaign) error
	SaveCreatives(creatives []AdCreative) error
	GetCampaignCreatives(campaignID uuid.UUID) ([]AdCreative, error)
}

// =============================================================================
// Service Interface
// =============================================================================

// Service defines the business logic contract for AI operations.
type Service interface {
	// Advisor
	StartChat(userID uuid.UUID, persona AdvisorPersona, topic string) (*AdvisorSession, error)
	SendMessage(sessionID uuid.UUID, content string) (*AdvisorMessage, error)

	// Oracle
	RunSimulation(projectID uuid.UUID, simType SimulationType, params map[string]interface{}) (*OracleSimulation, error)

	// Incubator
	AnalyzePitch(pitch *PitchDeck) (*PitchAnalysis, error)

	// Ad Gen
	GenerateAds(projectID uuid.UUID, req CreateAdRequest) ([]AdCreative, error)
}

// CreateAdRequest is a DTO for generating ads.
type CreateAdRequest struct {
	Name           string       `json:"name"`
	Platforms      []AdPlatform `json:"platforms"`
	ProductDesc    string       `json:"product_description"`
	TargetAudience string       `json:"target_audience"`
	Tone           string       `json:"tone"`
}