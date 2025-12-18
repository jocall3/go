package developer

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/my-org/my-project/internal/repository"
	"github.com/my-org/my-project/internal/repository/model"
	"github.com/my-org/my-project/pkg/logger"
	"github.com/my-org/my-project/pkg/security"
)

// Service defines the interface for developer-related operations,
// specifically managing webhook subscriptions and API keys.
type Service interface {
	// API Key Management
	// GenerateAPIKey creates a new API key for a developer.
	// It returns the created APIKey model and the raw, unhashed key.
	// The raw key should only be shown once to the developer.
	GenerateAPIKey(ctx context.Context, developerID string, name string) (*model.APIKey, string, error)
	// GetAPIKeys retrieves all active API keys for a specific developer.
	GetAPIKeys(ctx context.Context, developerID string) ([]model.APIKey, error)
	// RevokeAPIKey deactivates an API key, making it unusable for authentication.
	RevokeAPIKey(ctx context.Context, developerID, keyID string) error
	// ValidateAPIKey validates a raw API key against stored hashed keys.
	// It returns the APIKey model if valid and active, otherwise an error.
	// This method is typically used by authentication middleware.
	ValidateAPIKey(ctx context.Context, rawKey string) (*model.APIKey, error)

	// Webhook Subscription Management
	// CreateWebhookSubscription creates a new webhook endpoint for a developer.
	CreateWebhookSubscription(ctx context.Context, developerID string, url string, events []string, secret string) (*model.WebhookSubscription, error)
	// GetWebhookSubscriptions retrieves all webhook subscriptions for a specific developer.
	GetWebhookSubscriptions(ctx context.Context, developerID string) ([]model.WebhookSubscription, error)
	// UpdateWebhookSubscription modifies an existing webhook subscription.
	// Pointers are used for optional fields that may or may not be updated.
	UpdateWebhookSubscription(ctx context.Context, developerID, subscriptionID string, url *string, events *[]string, secret *string, isActive *bool) (*model.WebhookSubscription, error)
	// DeleteWebhookSubscription removes a webhook subscription.
	DeleteWebhookSubscription(ctx context.Context, developerID, subscriptionID string) error
}

// service implements the Service interface.
type service struct {
	repo         repository.DeveloperRepository
	logger       *logger.Logger
	apiKeyHasher security.APIKeyHasher
}

// NewService creates a new developer service with the given repository, logger, and API key hasher.
func NewService(repo repository.DeveloperRepository, log *logger.Logger, hasher security.APIKeyHasher) Service {
	return &service{
		repo:         repo,
		logger:       log,
		apiKeyHasher: hasher,
	}
}

// generateRandomAPIKey creates a cryptographically secure random string for an API key.
func generateRandomAPIKey() (string, error) {
	b := make([]byte, 32) // 256 bits of randomness
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes for API key: %w", err)
	}
	return hex.EncodeToString(b), nil
}

// GenerateAPIKey generates a new API key for a developer.
// It returns the created APIKey model and the raw key (which should be shown only once).
func (s *service) GenerateAPIKey(ctx context.Context, developerID string, name string) (*model.APIKey, string, error) {
	s.logger.Debugf(ctx, "Generating API key for developer %s with name '%s'", developerID, name)

	rawKey, err := generateRandomAPIKey()
	if err != nil {
		s.logger.Errorf(ctx, "Error generating random API key: %v", err)
		return nil, "", fmt.Errorf("failed to generate API key: %w", err)
	}

	hashedKey, err := s.apiKeyHasher.HashAPIKey(rawKey)
	if err != nil {
		s.logger.Errorf(ctx, "Error hashing API key: %v", err)
		return nil, "", fmt.Errorf("failed to hash API key: %w", err)
	}

	apiKey := &model.APIKey{
		DeveloperID: developerID,
		Name:        name,
		HashedKey:   hashedKey,
		CreatedAt:   time.Now().UTC(),
		IsActive:    true,
	}

	createdKey, err := s.repo.CreateAPIKey(ctx, apiKey)
	if err != nil {
		s.logger.Errorf(ctx, "Error creating API key in repository for developer %s: %v", developerID, err)
		return nil, "", fmt.Errorf("failed to store API key: %w", err)
	}

	s.logger.Infof(ctx, "API key %s generated for developer %s", createdKey.ID, developerID)
	return createdKey, rawKey, nil
}

// GetAPIKeys retrieves all active API keys for a developer.
func (s *service) GetAPIKeys(ctx context.Context, developerID string) ([]model.APIKey, error) {
	s.logger.Debugf(ctx, "Retrieving API keys for developer %s", developerID)
	keys, err := s.repo.GetAPIKeysByDeveloperID(ctx, developerID)
	if err != nil {
		s.logger.Errorf(ctx, "Error retrieving API keys for developer %s: %v", developerID, err)
		return nil, fmt.Errorf("failed to retrieve API keys: %w", err)
	}
	return keys, nil
}

// RevokeAPIKey deactivates an API key, making it unusable.
func (s *service) RevokeAPIKey(ctx context.Context, developerID, keyID string) error {
	s.logger.Debugf(ctx, "Revoking API key %s for developer %s", keyID, developerID)

	apiKey, err := s.repo.GetAPIKeyByID(ctx, keyID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return fmt.Errorf("API key not found: %w", err)
		}
		s.logger.Errorf(ctx, "Error getting API key %s for developer %s: %v", keyID, developerID, err)
		return fmt.Errorf("failed to retrieve API key: %w", err)
	}

	if apiKey.DeveloperID != developerID {
		s.logger.Warnf(ctx, "Attempt to revoke API key %s by unauthorized developer %s (owner: %s)", keyID, developerID, apiKey.DeveloperID)
		return fmt.Errorf("unauthorized: API key does not belong to developer")
	}

	if !apiKey.IsActive {
		s.logger.Debugf(ctx, "API key %s for developer %s is already inactive", keyID, developerID)
		return nil // Idempotent: already revoked
	}

	err = s.repo.UpdateAPIKeyStatus(ctx, keyID, false)
	if err != nil {
		s.logger.Errorf(ctx, "Error revoking API key %s for developer %s: %v", keyID, developerID, err)
		return fmt.Errorf("failed to revoke API key: %w", err)
	}

	s.logger.Infof(ctx, "API key %s revoked for developer %s", keyID, developerID)
	return nil
}

// ValidateAPIKey validates a raw API key against stored hashed keys.
// It returns the APIKey model if valid and active, otherwise an error.
// This method is typically used by authentication middleware.
func (s *service) ValidateAPIKey(ctx context.Context, rawKey string) (*model.APIKey, error) {
	// Log only a prefix of the key for security
	logKeyPrefix := "N/A"
	if len(rawKey) > 8 {
		logKeyPrefix = rawKey[:8] + "..."
	} else if len(rawKey) > 0 {
		logKeyPrefix = rawKey
	}
	s.logger.Debugf(ctx, "Validating API key (prefix): %s", logKeyPrefix)

	// The repository is responsible for finding the API key by comparing the raw key
	// against stored hashed keys using the provided comparer function.
	// A robust implementation would involve extracting an ID from the rawKey
	// to perform a direct lookup, rather than iterating through all keys.
	apiKey, err := s.repo.FindActiveAPIKeyByHashedKey(ctx, rawKey, s.apiKeyHasher.CompareAPIKey)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			s.logger.Debugf(ctx, "API key validation failed: key not found or inactive.")
			return nil, errors.New("invalid or inactive API key")
		}
		s.logger.Errorf(ctx, "Error finding API key for validation: %v", err)
		return nil, fmt.Errorf("internal error during API key validation: %w", err)
	}

	// Update LastUsedAt asynchronously to avoid blocking the request path.
	go func() {
		updateCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.repo.UpdateAPIKeyLastUsed(updateCtx, apiKey.ID, time.Now().UTC()); err != nil {
			s.logger.Errorf(updateCtx, "Failed to update LastUsedAt for API key %s: %v", apiKey.ID, err)
		}
	}()

	s.logger.Debugf(ctx, "API key %s validated successfully for developer %s", apiKey.ID, apiKey.DeveloperID)
	return apiKey, nil
}

// CreateWebhookSubscription creates a new webhook subscription for a developer.
func (s *service) CreateWebhookSubscription(ctx context.Context, developerID string, url string, events []string, secret string) (*model.WebhookSubscription, error) {
	s.logger.Debugf(ctx, "Creating webhook subscription for developer %s to URL %s with events %v", developerID, url, events)

	if url == "" {
		return nil, errors.New("webhook URL cannot be empty")
	}
	if len(events) == 0 {
		return nil, errors.New("at least one event must be specified for webhook subscription")
	}
	// TODO: Add more robust URL validation (e.g., valid scheme, host, HTTPS enforcement in production).
	// TODO: Validate events against a predefined list of supported event types.

	var hashedSecret string
	if secret != "" {
		var err error
		hashedSecret, err = s.apiKeyHasher.HashAPIKey(secret) // Reusing APIKeyHasher for webhook secret hashing
		if err != nil {
			s.logger.Errorf(ctx, "Error hashing webhook secret: %v", err)
			return nil, fmt.Errorf("failed to hash webhook secret: %w", err)
		}
	}

	subscription := &model.WebhookSubscription{
		DeveloperID: developerID,
		URL:         url,
		Events:      events,
		Secret:      hashedSecret, // Store hashed secret
		IsActive:    true,
		CreatedAt:   time.Now().UTC(),
	}

	createdSub, err := s.repo.CreateWebhookSubscription(ctx, subscription)
	if err != nil {
		s.logger.Errorf(ctx, "Error creating webhook subscription in repository for developer %s: %v", developerID, err)
		return nil, fmt.Errorf("failed to store webhook subscription: %w", err)
	}

	s.logger.Infof(ctx, "Webhook subscription %s created for developer %s to URL %s", createdSub.ID, developerID, url)
	return createdSub, nil
}

// GetWebhookSubscriptions retrieves all webhook subscriptions for a developer.
func (s *service) GetWebhookSubscriptions(ctx context.Context, developerID string) ([]model.WebhookSubscription, error) {
	s.logger.Debugf(ctx, "Retrieving webhook subscriptions for developer %s", developerID)
	subs, err := s.repo.GetWebhookSubscriptionsByDeveloperID(ctx, developerID)
	if err != nil {
		s.logger.Errorf(ctx, "Error retrieving webhook subscriptions for developer %s: %v", developerID, err)
		return nil, fmt.Errorf("failed to retrieve webhook subscriptions: %w", err)
	}
	return subs, nil
}

// UpdateWebhookSubscription modifies an existing webhook subscription.
// Fields are updated only if their corresponding pointer is not nil.
func (s *service) UpdateWebhookSubscription(ctx context.Context, developerID, subscriptionID string, url *string, events *[]string, secret *string, isActive *bool) (*model.WebhookSubscription, error) {
	s.logger.Debugf(ctx, "Updating webhook subscription %s for developer %s", subscriptionID, developerID)

	existingSub, err := s.repo.GetWebhookSubscriptionByID(ctx, subscriptionID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, fmt.Errorf("webhook subscription not found: %w", err)
		}
		s.logger.Errorf(ctx, "Error getting webhook subscription %s for developer %s: %v", subscriptionID, developerID, err)
		return nil, fmt.Errorf("failed to retrieve webhook subscription: %w", err)
	}

	if existingSub.DeveloperID != developerID {
		s.logger.Warnf(ctx, "Attempt to update webhook subscription %s by unauthorized developer %s (owner: %s)", subscriptionID, developerID, existingSub.DeveloperID)
		return nil, fmt.Errorf("unauthorized: webhook subscription does not belong to developer")
	}

	// Apply updates if provided
	if url != nil {
		if *url == "" {
			return nil, errors.New("webhook URL cannot be empty")
		}
		existingSub.URL = *url
		// TODO: Re-validate URL format
	}
	if events != nil {
		if len(*events) == 0 {
			return nil, errors.New("at least one event must be specified for webhook subscription")
		}
		existingSub.Events = *events
		// TODO: Re-validate events against supported types
	}
	if secret != nil {
		if *secret == "" {
			existingSub.Secret = "" // Clear secret if empty string is passed
		} else {
			hashedSecret, err := s.apiKeyHasher.HashAPIKey(*secret)
			if err != nil {
				s.logger.Errorf(ctx, "Error hashing new webhook secret for subscription %s: %v", subscriptionID, err)
				return nil, fmt.Errorf("failed to hash new webhook secret: %w", err)
			}
			existingSub.Secret = hashedSecret
		}
	}
	if isActive != nil {
		existingSub.IsActive = *isActive
	}

	existingSub.UpdatedAt = func() *time.Time { t := time.Now().UTC(); return &t }()

	updatedSub, err := s.repo.UpdateWebhookSubscription(ctx, existingSub)
	if err != nil {
		s.logger.Errorf(ctx, "Error updating webhook subscription %s for developer %s: %v", subscriptionID, developerID, err)
		return nil, fmt.Errorf("failed to update webhook subscription: %w", err)
	}

	s.logger.Infof(ctx, "Webhook subscription %s updated for developer %s", subscriptionID, developerID)
	return updatedSub, nil
}

// DeleteWebhookSubscription removes a webhook subscription.
func (s *service) DeleteWebhookSubscription(ctx context.Context, developerID, subscriptionID string) error {
	s.logger.Debugf(ctx, "Deleting webhook subscription %s for developer %s", subscriptionID, developerID)

	existingSub, err := s.repo.GetWebhookSubscriptionByID(ctx, subscriptionID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return fmt.Errorf("webhook subscription not found: %w", err)
		}
		s.logger.Errorf(ctx, "Error getting webhook subscription %s for developer %s: %v", subscriptionID, developerID, err)
		return fmt.Errorf("failed to retrieve webhook subscription: %w", err)
	}

	if existingSub.DeveloperID != developerID {
		s.logger.Warnf(ctx, "Attempt to delete webhook subscription %s by unauthorized developer %s (owner: %s)", subscriptionID, developerID, existingSub.DeveloperID)
		return fmt.Errorf("unauthorized: webhook subscription does not belong to developer")
	}

	err = s.repo.DeleteWebhookSubscription(ctx, subscriptionID)
	if err != nil {
		s.logger.Errorf(ctx, "Error deleting webhook subscription %s for developer %s: %v", subscriptionID, developerID, err)
		return fmt.Errorf("failed to delete webhook subscription: %w", err)
	}

	s.logger.Infof(ctx, "Webhook subscription %s deleted for developer %s", subscriptionID, developerID)
	return nil
}