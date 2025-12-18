package notification

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"gothunder/internal/model"
	"gothunder/internal/repository"
	"gothunder/internal/service/user"
)

// Error definitions for the notification service layer.
var (
	ErrNotFound         = errors.New("notification not found")
	ErrPermissionDenied = errors.New("permission denied to access this notification")
	ErrInvalidArgument  = errors.New("invalid argument provided")
)

// Service defines the interface for notification-related business logic.
// It abstracts the underlying data storage and provides a clean API for notification operations.
type Service interface {
	// List retrieves a paginated list of notifications for a specific user.
	List(ctx context.Context, userID uuid.UUID, page, pageSize int) ([]model.Notification, int, error)

	// MarkAsRead marks a single notification as read. It ensures the notification belongs to the user.
	MarkAsRead(ctx context.Context, userID, notificationID uuid.UUID) error

	// MarkAllAsRead marks all unread notifications for a user as read.
	MarkAllAsRead(ctx context.Context, userID uuid.UUID) error

	// GetSettings retrieves the notification settings for a user. If no settings exist, it creates and returns default settings.
	GetSettings(ctx context.Context, userID uuid.UUID) (*model.NotificationSettings, error)

	// UpdateSettings updates the notification settings for a user based on the provided request.
	UpdateSettings(ctx context.Context, userID uuid.UUID, req *model.UpdateNotificationSettingsRequest) (*model.NotificationSettings, error)

	// Create creates a new notification. This method is typically called by other services within the application
	// (e.g., when a user gets a new follower or a comment on their post).
	Create(ctx context.Context, notification *model.Notification) error
}

// service is the concrete implementation of the Service interface.
type service struct {
	notificationRepo repository.NotificationRepository
	userRepo         repository.UserRepository
	logger           *slog.Logger
}

// Config holds the configuration and dependencies for the notification service.
type Config struct {
	NotificationRepo repository.NotificationRepository
	UserRepo         repository.UserRepository
	Logger           *slog.Logger
}

// NewService creates and returns a new notification service instance.
// It validates that all necessary dependencies are provided.
func NewService(cfg Config) (Service, error) {
	if cfg.NotificationRepo == nil {
		return nil, errors.New("notification repository is required")
	}
	if cfg.UserRepo == nil {
		return nil, errors.New("user repository is required")
	}
	if cfg.Logger == nil {
		return nil, errors.New("logger is required")
	}

	return &service{
		notificationRepo: cfg.NotificationRepo,
		userRepo:         cfg.UserRepo,
		logger:           cfg.Logger.With("service", "notification"),
	}, nil
}

// List retrieves a paginated list of notifications for a specific user.
func (s *service) List(ctx context.Context, userID uuid.UUID, page, pageSize int) ([]model.Notification, int, error) {
	const defaultPage = 1
	const defaultPageSize = 20
	const maxPageSize = 100

	if page < 1 {
		page = defaultPage
	}
	if pageSize < 1 || pageSize > maxPageSize {
		pageSize = defaultPageSize
	}

	offset := (page - 1) * pageSize
	log := s.logger.With("method", "List", "userID", userID.String(), "page", page, "pageSize", pageSize)

	// Concurrently fetch notifications and total count for better performance.
	var notifications []model.Notification
	var total int
	var errNotifications, errTotal error
	errGroup, gCtx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		notifications, errNotifications = s.notificationRepo.ListByUserID(gCtx, userID, pageSize, offset)
		return errNotifications
	})

	errGroup.Go(func() error {
		total, errTotal = s.notificationRepo.CountByUserID(gCtx, userID)
		return errTotal
	})

	if err := errGroup.Wait(); err != nil {
		log.Error("Failed to retrieve notifications data", "error", err)
		return nil, 0, fmt.Errorf("failed to retrieve notifications: %w", err)
	}

	return notifications, total, nil
}

// MarkAsRead marks a single notification as read.
func (s *service) MarkAsRead(ctx context.Context, userID, notificationID uuid.UUID) error {
	log := s.logger.With("method", "MarkAsRead", "userID", userID.String(), "notificationID", notificationID.String())

	notification, err := s.notificationRepo.FindByID(ctx, notificationID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Warn("Notification not found")
			return ErrNotFound
		}
		log.Error("Failed to find notification by ID", "error", err)
		return fmt.Errorf("failed to retrieve notification: %w", err)
	}

	if notification.UserID != userID {
		log.Warn("User attempted to access another user's notification")
		return ErrPermissionDenied
	}

	// If already read, do nothing to avoid unnecessary database writes.
	if notification.ReadAt != nil {
		return nil
	}

	now := time.Now().UTC()
	notification.ReadAt = &now

	if err := s.notificationRepo.Update(ctx, notification); err != nil {
		log.Error("Failed to update notification status", "error", err)
		return fmt.Errorf("failed to mark notification as read: %w", err)
	}

	log.Info("Notification marked as read")
	return nil
}

// MarkAllAsRead marks all unread notifications for a user as read.
func (s *service) MarkAllAsRead(ctx context.Context, userID uuid.UUID) error {
	log := s.logger.With("method", "MarkAllAsRead", "userID", userID.String())

	// We can directly call the repository to perform a bulk update,
	// which is more efficient than fetching all notifications first.
	if err := s.notificationRepo.MarkAllAsReadByUserID(ctx, userID); err != nil {
		log.Error("Failed to mark all notifications as read in repository", "error", err)
		return fmt.Errorf("failed to mark all notifications as read: %w", err)
	}

	log.Info("All unread notifications marked as read for user")
	return nil
}

// GetSettings retrieves the notification settings for a user.
func (s *service) GetSettings(ctx context.Context, userID uuid.UUID) (*model.NotificationSettings, error) {
	log := s.logger.With("method", "GetSettings", "userID", userID.String())

	settings, err := s.notificationRepo.GetSettingsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			// If no settings exist, create and return default settings.
			// This ensures a user always has a settings record, simplifying client-side logic.
			log.Info("No settings found for user, creating default settings")
			defaultSettings := model.NewDefaultNotificationSettings(userID)
			if err := s.notificationRepo.CreateSettings(ctx, defaultSettings); err != nil {
				log.Error("Failed to create default settings for user", "error", err)
				return nil, fmt.Errorf("failed to create default settings: %w", err)
			}
			return defaultSettings, nil
		}
		log.Error("Failed to get notification settings from repository", "error", err)
		return nil, fmt.Errorf("failed to retrieve settings: %w", err)
	}

	return settings, nil
}

// UpdateSettings updates the notification settings for a user.
func (s *service) UpdateSettings(ctx context.Context, userID uuid.UUID, req *model.UpdateNotificationSettingsRequest) (*model.NotificationSettings, error) {
	log := s.logger.With("method", "UpdateSettings", "userID", userID.String())

	// GetSettings will create settings if they don't exist, ensuring we always have a record to update.
	currentSettings, err := s.GetSettings(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve current settings: %w", err)
	}

	// Apply updates from the request. Pointers are used to differentiate between a value being false and not being provided.
	if req.EmailEnabled != nil {
		currentSettings.EmailEnabled = *req.EmailEnabled
	}
	if req.PushEnabled != nil {
		currentSettings.PushEnabled = *req.PushEnabled
	}
	if req.NewFollower != nil {
		currentSettings.NewFollower = *req.NewFollower
	}
	if req.NewComment != nil {
		currentSettings.NewComment = *req.NewComment
	}
	if req.NewLike != nil {
		currentSettings.NewLike = *req.NewLike
	}

	currentSettings.UpdatedAt = time.Now().UTC()

	if err := s.notificationRepo.UpdateSettings(ctx, currentSettings); err != nil {
		log.Error("Failed to update notification settings in repository", "error", err)
		return nil, fmt.Errorf("failed to update settings: %w", err)
	}

	log.Info("Notification settings updated")
	return currentSettings, nil
}

// Create creates a new notification.
func (s *service) Create(ctx context.Context, notification *model.Notification) error {
	if err := s.validateNotification(notification); err != nil {
		return err
	}

	log := s.logger.With("method", "Create", "userID", notification.UserID.String(), "type", notification.Type)

	// Check if user exists before creating notification to maintain data integrity.
	if _, err := s.userRepo.FindByID(ctx, notification.UserID); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Warn("Attempted to create notification for non-existent user")
			return user.ErrNotFound
		}
		log.Error("Failed to check user existence", "error", err)
		return fmt.Errorf("failed to verify user: %w", err)
	}

	// Set default values for new notification
	notification.ID = uuid.New()
	notification.CreatedAt = time.Now().UTC()
	notification.ReadAt = nil

	if err := s.notificationRepo.Create(ctx, notification); err != nil {
		log.Error("Failed to create notification in repository", "error", err)
		return fmt.Errorf("failed to create notification: %w", err)
	}

	log.Info("Notification created successfully")

	// In a real-world application, this is where you would trigger the actual delivery.
	// This is best handled asynchronously, e.g., by publishing an event to a message queue (Kafka, RabbitMQ, etc.).
	// Example: s.eventPublisher.Publish(ctx, "notification.created", notification)

	return nil
}

// validateNotification performs basic validation on a notification object.
func (s *service) validateNotification(n *model.Notification) error {
	if n.UserID == uuid.Nil {
		return fmt.Errorf("%w: user ID is required", ErrInvalidArgument)
	}
	if n.Message == "" {
		return fmt.Errorf("%w: message is required", ErrInvalidArgument)
	}
	if n.Type == "" {
		return fmt.Errorf("%w: type is required", ErrInvalidArgument)
	}
	return nil
}