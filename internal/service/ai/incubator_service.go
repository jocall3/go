package ai

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// -----------------------------------------------------------------------------
// Domain Models
// -----------------------------------------------------------------------------

// PitchStatus represents the lifecycle state of a startup pitch.
type PitchStatus string

const (
	PitchStatusDraft      PitchStatus = "DRAFT"
	PitchStatusSubmitted  PitchStatus = "SUBMITTED"
	PitchStatusProcessing PitchStatus = "PROCESSING"
	PitchStatusReviewed   PitchStatus = "REVIEWED"
	PitchStatusRefining   PitchStatus = "REFINING"
	PitchStatusFunded     PitchStatus = "FUNDED" // Hypothetical end state
	PitchStatusRejected   PitchStatus = "REJECTED"
)

// Pitch represents a startup idea submitted to the incubator.
type Pitch struct {
	ID          string      `json:"id"`
	UserID      string      `json:"user_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Market      string      `json:"market"`
	TechStack   []string    `json:"tech_stack"`
	Status      PitchStatus `json:"status"`
	Score       float64     `json:"score"` // 0.0 to 10.0
	Feedback    []Feedback  `json:"feedback"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Version     int         `json:"version"`
}

// Feedback represents AI or mentor generated insights.
type Feedback struct {
	ID        string    `json:"id"`
	PitchID   string    `json:"pitch_id"`
	Category  string    `json:"category"` // e.g., "Viability", "Competition", "Tech"
	Content   string    `json:"content"`
	Sentiment string    `json:"sentiment"` // "Positive", "Neutral", "Negative"
	CreatedAt time.Time `json:"created_at"`
}

// AnalysisResult is the raw output from the AI provider.
type AnalysisResult struct {
	OverallScore    float64
	KeyStrengths    []string
	Weaknesses      []string
	PivotSuggestion string
	FeedbackItems   []Feedback
}

// -----------------------------------------------------------------------------
// Interfaces
// -----------------------------------------------------------------------------

// PitchRepository defines the data access layer for pitches.
type PitchRepository interface {
	Save(ctx context.Context, pitch *Pitch) error
	GetByID(ctx context.Context, id string) (*Pitch, error)
	Update(ctx context.Context, pitch *Pitch) error
	ListByUserID(ctx context.Context, userID string) ([]*Pitch, error)
	Delete(ctx context.Context, id string) error
}

// AIProvider defines the interface for the Large Language Model service.
type AIProvider interface {
	AnalyzePitch(ctx context.Context, pitch *Pitch) (*AnalysisResult, error)
	GeneratePivotIdeas(ctx context.Context, pitch *Pitch) ([]string, error)
}

// EventBus defines a mechanism to publish events (e.g., for notifications).
type EventBus interface {
	Publish(ctx context.Context, topic string, payload interface{}) error
}

// -----------------------------------------------------------------------------
// Service Implementation
// -----------------------------------------------------------------------------

// IncubatorService manages the lifecycle of startup pitches and AI interactions.
type IncubatorService struct {
	repo     PitchRepository
	ai       AIProvider
	events   EventBus
	logger   *slog.Logger
	timeout  time.Duration
}

// NewIncubatorService creates a new instance of the service.
func NewIncubatorService(
	repo PitchRepository,
	ai AIProvider,
	events EventBus,
	logger *slog.Logger,
) *IncubatorService {
	return &IncubatorService{
		repo:    repo,
		ai:      ai,
		events:  events,
		logger:  logger,
		timeout: 2 * time.Minute, // Default timeout for AI operations
	}
}

// SubmitPitch accepts a new pitch from a user and initiates the analysis process.
func (s *IncubatorService) SubmitPitch(ctx context.Context, userID, title, description, market string, techStack []string) (*Pitch, error) {
	if title == "" || description == "" {
		return nil, errors.New("title and description are required")
	}

	pitch := &Pitch{
		ID:          uuid.New().String(),
		UserID:      userID,
		Title:       title,
		Description: description,
		Market:      market,
		TechStack:   techStack,
		Status:      PitchStatusSubmitted,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Version:     1,
		Feedback:    make([]Feedback, 0),
	}

	if err := s.repo.Save(ctx, pitch); err != nil {
		s.logger.Error("failed to save pitch", "error", err, "user_id", userID)
		return nil, fmt.Errorf("failed to submit pitch: %w", err)
	}

	s.logger.Info("pitch submitted", "pitch_id", pitch.ID, "title", pitch.Title)

	// Trigger async analysis
	go s.processPitchAnalysis(pitch.ID)

	return pitch, nil
}

// GetPitch retrieves a pitch by its ID.
func (s *IncubatorService) GetPitch(ctx context.Context, id string) (*Pitch, error) {
	pitch, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("pitch not found: %w", err)
	}
	return pitch, nil
}

// GetUserPitches retrieves all pitches for a specific user.
func (s *IncubatorService) GetUserPitches(ctx context.Context, userID string) ([]*Pitch, error) {
	return s.repo.ListByUserID(ctx, userID)
}

// RefinePitch allows a user to update their pitch based on feedback, triggering a re-analysis.
func (s *IncubatorService) RefinePitch(ctx context.Context, pitchID, newDescription string, newTechStack []string) (*Pitch, error) {
	pitch, err := s.repo.GetByID(ctx, pitchID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve pitch: %w", err)
	}

	if pitch.Status == PitchStatusProcessing {
		return nil, errors.New("cannot refine pitch while it is being processed")
	}

	pitch.Description = newDescription
	if len(newTechStack) > 0 {
		pitch.TechStack = newTechStack
	}
	pitch.Status = PitchStatusRefining
	pitch.Version++
	pitch.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, pitch); err != nil {
		return nil, fmt.Errorf("failed to update pitch: %w", err)
	}

	s.logger.Info("pitch refined", "pitch_id", pitch.ID, "version", pitch.Version)

	// Trigger re-analysis
	go s.processPitchAnalysis(pitch.ID)

	return pitch, nil
}

// processPitchAnalysis handles the AI interaction asynchronously.
func (s *IncubatorService) processPitchAnalysis(pitchID string) {
	// Create a detached context with timeout for the background job
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	s.logger.Info("starting pitch analysis", "pitch_id", pitchID)

	// 1. Retrieve Pitch
	pitch, err := s.repo.GetByID(ctx, pitchID)
	if err != nil {
		s.logger.Error("analysis failed: could not fetch pitch", "pitch_id", pitchID, "error", err)
		return
	}

	// 2. Update Status to Processing
	pitch.Status = PitchStatusProcessing
	_ = s.repo.Update(ctx, pitch)

	// 3. Call AI Provider
	analysis, err := s.ai.AnalyzePitch(ctx, pitch)
	if err != nil {
		s.logger.Error("analysis failed: ai provider error", "pitch_id", pitchID, "error", err)
		// Revert status or set to error state? For now, keep as is or add an ERROR status.
		// Ideally, we might want a retry mechanism here.
		return
	}

	// 4. Apply Analysis Results
	pitch.Score = analysis.OverallScore
	pitch.Feedback = append(pitch.Feedback, analysis.FeedbackItems...)
	pitch.Status = PitchStatusReviewed
	pitch.UpdatedAt = time.Now()

	// 5. Save Updated Pitch
	if err := s.repo.Update(ctx, pitch); err != nil {
		s.logger.Error("analysis failed: could not save results", "pitch_id", pitchID, "error", err)
		return
	}

	// 6. Publish Event
	if s.events != nil {
		eventPayload := map[string]interface{}{
			"pitch_id": pitch.ID,
			"user_id":  pitch.UserID,
			"score":    pitch.Score,
			"status":   pitch.Status,
		}
		if err := s.events.Publish(ctx, "pitch.reviewed", eventPayload); err != nil {
			s.logger.Warn("failed to publish event", "pitch_id", pitchID, "error", err)
		}
	}

	s.logger.Info("pitch analysis completed", "pitch_id", pitchID, "score", pitch.Score)
}

// RequestPivotIdeas asks the AI for alternative directions based on the current pitch.
func (s *IncubatorService) RequestPivotIdeas(ctx context.Context, pitchID string) ([]string, error) {
	pitch, err := s.repo.GetByID(ctx, pitchID)
	if err != nil {
		return nil, err
	}

	ideas, err := s.ai.GeneratePivotIdeas(ctx, pitch)
	if err != nil {
		return nil, fmt.Errorf("failed to generate pivot ideas: %w", err)
	}

	// Log the interaction
	s.logger.Info("pivot ideas generated", "pitch_id", pitchID, "count", len(ideas))

	// Optionally save these ideas as a special type of feedback
	feedbackContent := fmt.Sprintf("Pivot Suggestions:\n- %s", strings.Join(ideas, "\n- "))
	pivotFeedback := Feedback{
		ID:        uuid.New().String(),
		PitchID:   pitchID,
		Category:  "Pivot Strategy",
		Content:   feedbackContent,
		Sentiment: "Neutral",
		CreatedAt: time.Now(),
	}

	pitch.Feedback = append(pitch.Feedback, pivotFeedback)
	if err := s.repo.Update(ctx, pitch); err != nil {
		s.logger.Warn("failed to save pivot feedback", "pitch_id", pitchID, "error", err)
	}

	return ideas, nil
}

// BatchAnalyze is an administrative function to re-analyze multiple pitches (e.g., after model upgrade).
func (s *IncubatorService) BatchAnalyze(ctx context.Context, pitchIDs []string) error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5) // Semaphore to limit concurrency to 5

	for _, id := range pitchIDs {
		wg.Add(1)
		go func(pid string) {
			defer wg.Done()
			sem <- struct{}{}        // Acquire
			defer func() { <-sem }() // Release

			s.processPitchAnalysis(pid)
		}(id)
	}

	// Wait for all routines to finish (or context cancellation if we implemented that logic inside the loop)
	// Note: processPitchAnalysis runs on its own detached context, so this waits for completion.
	wg.Wait()
	return nil
}

// CalculateMarketViability is a helper to aggregate scores and feedback sentiment.
func (s *IncubatorService) CalculateMarketViability(pitch *Pitch) string {
	if pitch.Score > 8.5 {
		return "High Potential - Unicorn Territory"
	} else if pitch.Score > 6.0 {
		return "Viable - Needs Execution"
	} else if pitch.Score > 4.0 {
		return "Risky - Significant Pivots Required"
	}
	return "Low Viability - Reassess Core Assumptions"
}