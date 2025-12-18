package ai

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

// Common errors for the AI Ad Service.
var (
	ErrInvalidPrompt = errors.New("invalid prompt: prompt cannot be empty")
	ErrJobNotFound   = errors.New("job not found")
	ErrGenFailed     = errors.New("generation failed")
)

// AdGenerationRequest encapsulates the parameters required to generate an AI video ad.
type AdGenerationRequest struct {
	UserID      string
	ProjectID   string
	Prompt      string
	Style       string // e.g., "cinematic", "minimalist", "cyberpunk"
	Duration    int    // Duration in seconds
	AspectRatio string // e.g., "16:9", "9:16"
}

// AdJobStatus represents the state of an asynchronous generation job.
type AdJobStatus string

const (
	StatusPending    AdJobStatus = "PENDING"
	StatusProcessing AdJobStatus = "PROCESSING"
	StatusCompleted  AdJobStatus = "COMPLETED"
	StatusFailed     AdJobStatus = "FAILED"
)

// AdJob represents the database record for a generation task.
type AdJob struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	ProjectID string      `json:"project_id"`
	Status    AdJobStatus `json:"status"`
	VideoURL  string      `json:"video_url,omitempty"`
	ErrorMsg  string      `json:"error_msg,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// AdRepository defines the data access layer requirements for this service.
// This follows the Consumer-Driven Interface pattern.
type AdRepository interface {
	CreateJob(ctx context.Context, job *AdJob) error
	GetJob(ctx context.Context, id string) (*AdJob, error)
	UpdateJobStatus(ctx context.Context, id string, status AdJobStatus, videoURL string, errorMsg string) error
}

// VideoGeneratorClient defines the interface for the external AI video generation provider.
type VideoGeneratorClient interface {
	// GenerateVideo initiates the video generation and returns a URL or an identifier.
	// If the provider is synchronous, it returns the final URL.
	GenerateVideo(ctx context.Context, prompt, style string, duration int, ratio string) (string, error)
}

// Service defines the public API for the AI Ad Service.
type Service interface {
	// GenerateAd initiates an asynchronous ad generation process.
	// It returns a Job ID that can be used to track status.
	GenerateAd(ctx context.Context, req AdGenerationRequest) (string, error)

	// GetOperationStatus retrieves the current status of a generation job.
	GetOperationStatus(ctx context.Context, jobID string) (*AdJob, error)
}

// adService implements the Service interface.
type adService struct {
	repo      AdRepository
	aiClient  VideoGeneratorClient
	logger    *slog.Logger
	jobTimeout time.Duration
}

// NewAdService creates a new instance of the AI Ad Service with dependencies.
func NewAdService(
	repo AdRepository,
	aiClient VideoGeneratorClient,
	logger *slog.Logger,
) Service {
	return &adService{
		repo:       repo,
		aiClient:   aiClient,
		logger:     logger,
		jobTimeout: 15 * time.Minute, // Default timeout for generation jobs
	}
}

// GenerateAd validates the request, creates a tracking job, and triggers the AI generation.
func (s *adService) GenerateAd(ctx context.Context, req AdGenerationRequest) (string, error) {
	if req.Prompt == "" {
		return "", ErrInvalidPrompt
	}

	// Generate a unique Job ID
	jobID, err := generateRandomID(16)
	if err != nil {
		return "", fmt.Errorf("failed to generate job id: %w", err)
	}

	job := &AdJob{
		ID:        jobID,
		UserID:    req.UserID,
		ProjectID: req.ProjectID,
		Status:    StatusPending,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	// Persist initial job state
	if err := s.repo.CreateJob(ctx, job); err != nil {
		s.logger.Error("failed to create ad job", "error", err, "user_id", req.UserID)
		return "", fmt.Errorf("database error: %w", err)
	}

	s.logger.Info("ad generation job created", "job_id", jobID, "user_id", req.UserID)

	// Trigger asynchronous processing
	// In a distributed system, this would likely push to a message queue (e.g., Kafka, RabbitMQ).
	// For this implementation, we use a goroutine to handle the long-running task.
	go s.processAdGeneration(jobID, req)

	return jobID, nil
}

// GetOperationStatus retrieves the job details from the repository.
func (s *adService) GetOperationStatus(ctx context.Context, jobID string) (*AdJob, error) {
	job, err := s.repo.GetJob(ctx, jobID)
	if err != nil {
		s.logger.Error("failed to fetch job status", "job_id", jobID, "error", err)
		return nil, fmt.Errorf("failed to retrieve job: %w", err)
	}
	if job == nil {
		return nil, ErrJobNotFound
	}
	return job, nil
}

// processAdGeneration handles the interaction with the AI client and updates job status.
// It runs in a background goroutine.
func (s *adService) processAdGeneration(jobID string, req AdGenerationRequest) {
	// Create a background context with a timeout to prevent hanging goroutines
	ctx, cancel := context.WithTimeout(context.Background(), s.jobTimeout)
	defer cancel()

	// Update status to PROCESSING
	if err := s.repo.UpdateJobStatus(ctx, jobID, StatusProcessing, "", ""); err != nil {
		s.logger.Error("failed to update job to processing", "job_id", jobID, "error", err)
		return
	}

	start := time.Now()
	s.logger.Debug("calling ai provider", "job_id", jobID)

	// Call the AI Provider
	videoURL, err := s.aiClient.GenerateVideo(ctx, req.Prompt, req.Style, req.Duration, req.AspectRatio)
	
	duration := time.Since(start)

	if err != nil {
		s.logger.Error("ai generation failed", "job_id", jobID, "error", err, "duration", duration)
		
		// Update status to FAILED
		errMsg := err.Error()
		if updateErr := s.repo.UpdateJobStatus(ctx, jobID, StatusFailed, "", errMsg); updateErr != nil {
			s.logger.Error("failed to update job status to failed", "job_id", jobID, "error", updateErr)
		}
		return
	}

	// Update status to COMPLETED
	if err := s.repo.UpdateJobStatus(ctx, jobID, StatusCompleted, videoURL, ""); err != nil {
		s.logger.Error("failed to update job status to completed", "job_id", jobID, "error", err)
		return
	}

	s.logger.Info("ad generation completed successfully", "job_id", jobID, "duration", duration)
}

// generateRandomID generates a cryptographically secure random hex string.
func generateRandomID(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}