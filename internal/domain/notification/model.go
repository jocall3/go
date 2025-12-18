// Package notification defines the core domain models for notifications and user settings.
// These entities are central to how notifications are created, managed, and delivered.
package notification

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// NotificationType represents the category of a notification.
// Using a custom type ensures type safety and clarity.
type NotificationType string

// Constants for all supported notification types within the system.
const (
	// Social-related notifications
	NewFollower  NotificationType = "new_follower"
	PostLike     NotificationType = "post_like"
	CommentReply NotificationType = "comment_reply"
	Mention      NotificationType = "mention"

	// System-related notifications
	SystemAnnouncement NotificationType = "system_announcement"
	SecurityAlert      NotificationType = "security_alert"
	FeatureUpdate      NotificationType = "feature_update"
)

// IsValid checks if the notification type is a defined and supported constant.
func (nt NotificationType) IsValid() bool {
	switch nt {
	case NewFollower, PostLike, CommentReply, Mention, SystemAnnouncement, SecurityAlert, FeatureUpdate:
		return true
	}
	return false
}

// Notification represents a single notification message sent to a user.
// It is an immutable entity once created, though its state (e.g., 'read') can change.
type Notification struct {
	ID        uuid.UUID              `json:"id"`
	UserID    uuid.UUID              `json:"user_id"`
	Type      NotificationType       `json:"type"`
	Title     string                 `json:"title"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data,omitempty"` // Optional context-specific data (e.g., post_id, actor_id)
	Read      bool                   `json:"read"`
	CreatedAt time.Time              `json:"created_at"`
	ReadAt    *time.Time             `json:"read_at,omitempty"` // Pointer to allow null value
}

// NewNotification is a factory function to create a new, valid Notification instance.
// It ensures all required fields are present and valid before creation.
func NewNotification(userID uuid.UUID, nType NotificationType, title, message string, data map[string]interface{}) (*Notification, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user ID cannot be nil")
	}
	if !nType.IsValid() {
		return nil, fmt.Errorf("invalid notification type: %s", nType)
	}
	if title == "" || message == "" {
		return nil, errors.New("title and message cannot be empty")
	}

	return &Notification{
		ID:        uuid.New(),
		UserID:    userID,
		Type:      nType,
		Title:     title,
		Message:   message,
		Data:      data,
		Read:      false,
		CreatedAt: time.Now().UTC(),
		ReadAt:    nil,
	}, nil
}

// MarkAsRead updates the notification's state to 'read'.
// This is an idempotent operation.
func (n *Notification) MarkAsRead() {
	if !n.Read {
		n.Read = true
		now := time.Now().UTC()
		n.ReadAt = &now
	}
}

// MarkAsUnread updates the notification's state to 'unread'.
func (n *Notification) MarkAsUnread() {
	if n.Read {
		n.Read = false
		n.ReadAt = nil
	}
}

// --- User Notification Settings ---

// DeliveryChannel represents the medium through which a notification is sent.
type DeliveryChannel string

const (
	ChannelEmail DeliveryChannel = "email"
	ChannelPush  DeliveryChannel = "push"
	ChannelInApp DeliveryChannel = "in_app"
	ChannelSMS   DeliveryChannel = "sms"
)

// ChannelPreferences defines a user's delivery preferences for a single notification type.
type ChannelPreferences struct {
	Email bool `json:"email"`
	Push  bool `json:"push"`
	InApp bool `json:"in_app"`
	SMS   bool `json:"sms"`
}

// Settings is an aggregate root that represents a user's complete notification preferences.
// This entity controls which notifications are generated and delivered for a specific user.
type Settings struct {
	UserID      uuid.UUID                               `json:"user_id"`
	Preferences map[NotificationType]ChannelPreferences `json:"preferences"`
	UpdatedAt   time.Time                               `json:"updated_at"`
}

// NewSettings is a factory function to create a settings object for a new user.
// It populates the settings with sensible, system-wide defaults.
func NewSettings(userID uuid.UUID) (*Settings, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user ID cannot be nil for settings")
	}
	return &Settings{
		UserID:      userID,
		Preferences: defaultPreferences(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}

// defaultPreferences provides a baseline set of notification settings for new users.
// This can be configured to match the application's desired default behavior.
func defaultPreferences() map[NotificationType]ChannelPreferences {
	return map[NotificationType]ChannelPreferences{
		NewFollower:        {Email: true, Push: true, InApp: true, SMS: false},
		PostLike:           {Email: false, Push: true, InApp: true, SMS: false},
		CommentReply:       {Email: true, Push: true, InApp: true, SMS: false},
		Mention:            {Email: true, Push: true, InApp: true, SMS: false},
		SystemAnnouncement: {Email: true, Push: true, InApp: true, SMS: true},
		SecurityAlert:      {Email: true, Push: true, InApp: true, SMS: true},
		FeatureUpdate:      {Email: true, Push: false, InApp: true, SMS: false},
	}
}

// CanReceive checks if a user's settings permit a specific notification type
// to be delivered via a specific channel.
func (s *Settings) CanReceive(nType NotificationType, channel DeliveryChannel) bool {
	prefs, ok := s.Preferences[nType]
	if !ok {
		// A safe default: if a preference is not explicitly set, deny delivery.
		return false
	}

	switch channel {
	case ChannelEmail:
		return prefs.Email
	case ChannelPush:
		return prefs.Push
	case ChannelInApp:
		return prefs.InApp
	case ChannelSMS:
		return prefs.SMS
	default:
		return false
	}
}

// UpdatePreference changes a user's preference for a specific notification type and channel.
// It ensures the UpdatedAt timestamp is modified.
func (s *Settings) UpdatePreference(nType NotificationType, channel DeliveryChannel, enabled bool) error {
	if !nType.IsValid() {
		return fmt.Errorf("invalid notification type: %s", nType)
	}

	prefs, ok := s.Preferences[nType]
	if !ok {
		// If the type doesn't exist in the map (e.g., a new type was added), initialize it.
		prefs = ChannelPreferences{}
	}

	switch channel {
	case ChannelEmail:
		prefs.Email = enabled
	case ChannelPush:
		prefs.Push = enabled
	case ChannelInApp:
		prefs.InApp = enabled
	case ChannelSMS:
		prefs.SMS = enabled
	default:
		return fmt.Errorf("unknown delivery channel: %s", channel)
	}

	s.Preferences[nType] = prefs
	s.UpdatedAt = time.Now().UTC()
	return nil
}