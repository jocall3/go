package user

import (
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// -----------------------------------------------------------------------------
// Domain Errors
// -----------------------------------------------------------------------------

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInvalidToken       = errors.New("invalid or expired token")
	ErrInternalServer     = errors.New("internal server error")
)

// -----------------------------------------------------------------------------
// Data Transfer Objects (DTOs)
// -----------------------------------------------------------------------------

type RegisterRequest struct {
	FullName string
	Email    string
	Password string
}

type RegisterResponse struct {
	ID        string
	Email     string
	CreatedAt time.Time
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64 // Seconds
	TokenType    string
}

type ResetPasswordRequest struct {
	Token       string
	NewPassword string
}

// -----------------------------------------------------------------------------
// Domain Models (Assumed to be in this package or imported)
// -----------------------------------------------------------------------------

// User represents the user entity.
type User struct {
	ID           string
	FullName     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// -----------------------------------------------------------------------------
// Interfaces (Dependencies)
// -----------------------------------------------------------------------------

// UserRepository defines the data access layer for users.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	UpdatePassword(ctx context.Context, userID, newHash string) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

// TokenService defines operations for JWT handling.
type TokenService interface {
	GenerateTokenPair(userID string, email string) (accessToken, refreshToken string, exp int64, err error)
	ValidateResetToken(token string) (string, error) // Returns email associated with token
	GenerateResetToken(email string) (string, error)
}

// EmailService defines operations for sending notifications.
type EmailService interface {
	SendWelcomeEmail(ctx context.Context, toEmail, name string) error
	SendPasswordResetEmail(ctx context.Context, toEmail, token string) error
}

// IDGenerator defines a strategy for generating unique identifiers.
type IDGenerator interface {
	NewID() string
}

// -----------------------------------------------------------------------------
// Service Implementation
// -----------------------------------------------------------------------------

// AuthService handles business logic for user authentication.
type AuthService struct {
	userRepo     UserRepository
	tokenService TokenService
	emailService EmailService
	idGen        IDGenerator
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(
	userRepo UserRepository,
	tokenService TokenService,
	emailService EmailService,
	idGen IDGenerator,
) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		tokenService: tokenService,
		emailService: emailService,
		idGen:        idGen,
	}
}

// Register creates a new user account, hashes the password, and sends a welcome email.
func (s *AuthService) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {
	// 1. Validate input (basic validation)
	email := strings.ToLower(strings.TrimSpace(req.Email))
	if email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// 2. Check if user already exists
	exists, err := s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrEmailAlreadyExists
	}

	// 3. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	// 4. Create User entity
	now := time.Now().UTC()
	user := &User{
		ID:           s.idGen.NewID(),
		FullName:     strings.TrimSpace(req.FullName),
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// 5. Persist user
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	// 6. Send welcome email (async to not block response)
	go func() {
		// Create a detached context for async operation
		bgCtx := context.Background()
		_ = s.emailService.SendWelcomeEmail(bgCtx, user.Email, user.FullName)
	}()

	return &RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

// Login authenticates a user and returns JWT tokens.
func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	email := strings.ToLower(strings.TrimSpace(req.Email))

	// 1. Find user by email
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// 2. Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// 3. Generate Tokens
	accessToken, refreshToken, exp, err := s.tokenService.GenerateTokenPair(user.ID, user.Email)
	if err != nil {
		return nil, ErrInternalServer
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    exp,
		TokenType:    "Bearer",
	}, nil
}

// ForgotPassword initiates the password reset flow by sending an email with a reset token.
func (s *AuthService) ForgotPassword(ctx context.Context, email string) error {
	email = strings.ToLower(strings.TrimSpace(email))

	// 1. Check if user exists
	// We generally don't want to reveal if an email exists, but for logic we need to know.
	// A common security practice is to return nil even if user not found, 
	// but internally we stop processing.
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil // Return nil to prevent email enumeration
		}
		return err
	}

	// 2. Generate reset token
	token, err := s.tokenService.GenerateResetToken(user.Email)
	if err != nil {
		return err
	}

	// 3. Send email
	// Using a goroutine or synchronous depending on reliability requirements. 
	// For critical paths like auth, synchronous is often safer to ensure delivery handoff,
	// but here we'll do it synchronously to report errors.
	if err := s.emailService.SendPasswordResetEmail(ctx, user.Email, token); err != nil {
		return errors.New("failed to send reset email")
	}

	return nil
}

// ResetPassword validates the token and updates the user's password.
func (s *AuthService) ResetPassword(ctx context.Context, req ResetPasswordRequest) error {
	// 1. Validate token
	email, err := s.tokenService.ValidateResetToken(req.Token)
	if err != nil {
		return ErrInvalidToken
	}

	// 2. Find user
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return ErrUserNotFound
	}

	// 3. Hash new password
	newHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to process new password")
	}

	// 4. Update password in repo
	if err := s.userRepo.UpdatePassword(ctx, user.ID, string(newHash)); err != nil {
		return err
	}

	return nil
}