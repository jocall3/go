package user

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Common domain errors
var (
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrPasswordTooWeak = errors.New("password is too weak")
	ErrUserNotFound    = errors.New("user not found")
)

// Role represents the authorization level of a user.
type Role string

const (
	RoleUser       Role = "user"
	RoleAdmin      Role = "admin"
	RoleSuperAdmin Role = "super_admin"
	RoleModerator  Role = "moderator"
)

// Status represents the current state of the user account.
type Status string

const (
	StatusPending   Status = "pending"
	StatusActive    Status = "active"
	StatusSuspended Status = "suspended"
	StatusBanned    Status = "banned"
	StatusArchived  Status = "archived"
)

// User is the aggregate root for the user domain.
type User struct {
	ID           uuid.UUID   `json:"id" db:"id"`
	Email        string      `json:"email" db:"email"`
	Phone        *string     `json:"phone,omitempty" db:"phone"`
	PasswordHash string      `json:"-" db:"password_hash"` // Never serialize password hash
	FirstName    string      `json:"first_name" db:"first_name"`
	LastName     string      `json:"last_name" db:"last_name"`
	Role         Role        `json:"role" db:"role"`
	Status       Status      `json:"status" db:"status"`
	AvatarURL    *string     `json:"avatar_url,omitempty" db:"avatar_url"`
	Preferences  Preferences `json:"preferences" db:"preferences"`
	Addresses    []Address   `json:"addresses,omitempty" db:"-"` // Loaded via relation
	Devices      []Device    `json:"devices,omitempty" db:"-"`   // Loaded via relation
	Biometrics   []Biometric `json:"biometrics,omitempty" db:"-"`
	LastLoginAt  *time.Time  `json:"last_login_at,omitempty" db:"last_login_at"`
	CreatedAt    time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time  `json:"deleted_at,omitempty" db:"deleted_at"`
}

// NewUser creates a new user instance with default values.
func NewUser(email, firstName, lastName string) (*User, error) {
	if email == "" || !strings.Contains(email, "@") {
		return nil, ErrInvalidEmail
	}

	now := time.Now().UTC()
	return &User{
		ID:        uuid.New(),
		Email:     strings.ToLower(strings.TrimSpace(email)),
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
		Role:      RoleUser,
		Status:    StatusPending,
		Preferences: Preferences{
			Theme:             "system",
			Language:          "en",
			Notifications:     true,
			MarketingEmails:   false,
			TwoFactorEnabled:  false,
			Timezone:          "UTC",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// FullName returns the concatenated first and last name.
func (u *User) FullName() string {
	return strings.TrimSpace(u.FirstName + " " + u.LastName)
}

// IsActive checks if the user status allows for login.
func (u *User) IsActive() bool {
	return u.Status == StatusActive && u.DeletedAt == nil
}

// Activate sets the user status to active.
func (u *User) Activate() {
	u.Status = StatusActive
	u.UpdatedAt = time.Now().UTC()
}

// Address represents a physical location associated with a user.
type Address struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Label     string    `json:"label" db:"label"` // e.g., "Home", "Work"
	Line1     string    `json:"line1" db:"line1"`
	Line2     string    `json:"line2,omitempty" db:"line2"`
	City      string    `json:"city" db:"city"`
	State     string    `json:"state" db:"state"`
	ZipCode   string    `json:"zip_code" db:"zip_code"`
	Country   string    `json:"country" db:"country"`
	IsDefault bool      `json:"is_default" db:"is_default"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Preferences stores user-specific configuration settings.
// This is often stored as a JSONB column in the database.
type Preferences struct {
	Theme             string `json:"theme"`             // "light", "dark", "system"
	Language          string `json:"language"`          // ISO code
	Timezone          string `json:"timezone"`          // IANA Timezone string
	Notifications     bool   `json:"notifications"`     // Global toggle
	MarketingEmails   bool   `json:"marketing_emails"`  // Newsletter opt-in
	TwoFactorEnabled  bool   `json:"two_factor_enabled"`
	DataSharingOptIn  bool   `json:"data_sharing_opt_in"`
}

// Device represents a trusted device or active session.
type Device struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	UserID       uuid.UUID  `json:"user_id" db:"user_id"`
	DeviceToken  string     `json:"-" db:"device_token"` // Push notification token, kept private
	Name         string     `json:"name" db:"name"`       // e.g., "iPhone 13"
	Type         string     `json:"type" db:"type"`       // "ios", "android", "web"
	UserAgent    string     `json:"user_agent" db:"user_agent"`
	IPAddress    string     `json:"ip_address" db:"ip_address"`
	LastActiveAt time.Time  `json:"last_active_at" db:"last_active_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty" db:"expires_at"`
}

// Biometric represents FIDO2/WebAuthn credentials or other biometric data references.
type Biometric struct {
	ID              uuid.UUID `json:"id" db:"id"`
	UserID          uuid.UUID `json:"user_id" db:"user_id"`
	CredentialID    []byte    `json:"credential_id" db:"credential_id"` // WebAuthn Credential ID
	PublicKey       []byte    `json:"-" db:"public_key"`                // Public Key for verification
	AttestationType string    `json:"attestation_type" db:"attestation_type"`
	AAGUID          uuid.UUID `json:"aaguid" db:"aaguid"`
	SignCount       uint32    `json:"sign_count" db:"sign_count"`
	DeviceName      string    `json:"device_name" db:"device_name"` // Friendly name for the authenticator
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	LastUsedAt      time.Time `json:"last_used_at" db:"last_used_at"`
}

// UpdateLastLogin updates the user's last login timestamp.
func (u *User) UpdateLastLogin() {
	now := time.Now().UTC()
	u.LastLoginAt = &now
}

// HasRole checks if the user possesses a specific role.
func (u *User) HasRole(role Role) bool {
	return u.Role == role
}

// IsAdmin checks if the user has administrative privileges.
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin || u.Role == RoleSuperAdmin
}