package developer

import (
	"time"

	"github.com/google/uuid"
)

// Webhook represents a developer webhook subscription.
// It defines an endpoint where events are sent to a developer's application.
type Webhook struct {
	ID          uuid.UUID `json:"id"`           // Unique identifier for the webhook.
	DeveloperID uuid.UUID `json:"developer_id"` // The ID of the developer who owns this webhook.
	URL         string    `json:"url"`          // The URL endpoint to send events to.
	Events      []string  `json:"events"`       // List of event types this webhook subscribes to (e.g., "user.created", "order.updated").
	Secret      string    `json:"secret"`       // Secret used to sign webhook payloads for verification by the receiver.
	Status      string    `json:"status"`       // Current status of the webhook (e.g., "active", "disabled", "errored").
	CreatedAt   time.Time `json:"created_at"`   // Timestamp when the webhook was created.
	UpdatedAt   time.Time `json:"updated_at"`   // Timestamp when the webhook was last updated.
}

// APIKey represents a developer API key.
// It grants access to specific API resources and actions.
type APIKey struct {
	ID          uuid.UUID  `json:"id"`                     // Unique identifier for the API key.
	DeveloperID uuid.UUID  `json:"developer_id"`           // The ID of the developer who owns this API key.
	Name        string     `json:"name"`                   // A user-friendly name for the API key (e.g., "My Production Key").
	Key         string     `json:"key"`                    // The actual API key string (should be stored hashed/encrypted in DB, not plaintext).
	Permissions []string   `json:"permissions"`            // List of permissions associated with this key (e.g., "read:users", "write:products").
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`   // Optional expiration date for the key. Nil if it does not expire.
	Revoked     bool       `json:"revoked"`                // True if the key has been revoked and is no longer valid.
	CreatedAt   time.Time  `json:"created_at"`             // Timestamp when the API key was created.
	UpdatedAt   time.Time  `json:"updated_at"`             // Timestamp when the API key was last updated.
	LastUsedAt  *time.Time `json:"last_used_at,omitempty"` // Timestamp of the last time this key was used. Nil if never used.
}

// NewWebhook creates a new Webhook instance with a generated ID, default status, and timestamps.
func NewWebhook(developerID uuid.UUID, url string, events []string, secret string) *Webhook {
	now := time.Now().UTC()
	return &Webhook{
		ID:          uuid.New(),
		DeveloperID: developerID,
		URL:         url,
		Events:      events,
		Secret:      secret,
		Status:      "active", // Default status for a new webhook
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// NewAPIKey creates a new APIKey instance with a generated ID, default revoked status, and timestamps.
// The 'key' parameter here is the raw key string, which should typically be hashed before persistence
// and never exposed directly in responses.
func NewAPIKey(developerID uuid.UUID, name string, key string, permissions []string, expiresAt *time.Time) *APIKey {
	now := time.Now().UTC()
	return &APIKey{
		ID:          uuid.New(),
		DeveloperID: developerID,
		Name:        name,
		Key:         key, // This should be the hashed key when stored in a database
		Permissions: permissions,
		ExpiresAt:   expiresAt,
		Revoked:     false, // New keys are not revoked by default
		CreatedAt:   now,
		UpdatedAt:   now,
		LastUsedAt:  nil,   // A new key has not been used yet
	}
}