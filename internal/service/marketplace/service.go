// Package marketplace provides the business logic for the marketplace feature,
// including listing products, simulating redemptions, and processing redemptions.
package marketplace

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	// Assuming these domain and repository packages exist within the project structure.
	// These would define the core data structures and data access layer contracts.
	"github.com/your-org/your-app/internal/domain"
)

// --- Custom Errors ---

var (
	// ErrProductNotFound is returned when a requested product does not exist or is not active.
	ErrProductNotFound = errors.New("product not found")
	// ErrUserNotFound is returned when a user does not exist.
	ErrUserNotFound = errors.New("user not found")
	// ErrInsufficientFunds is returned when a user tries to redeem a product they cannot afford.
	ErrInsufficientFunds = errors.New("insufficient funds for redemption")
	// ErrRedemptionFailed is a generic error for when the redemption process fails for an unexpected reason.
	ErrRedemptionFailed = errors.new("redemption failed")
)

// --- Repository Interfaces (Dependencies) ---
// Note: In a real application, these interfaces would be defined in their respective
// repository packages (e.g., internal/repository/product/repository.go).

// ProductRepository defines the data access methods for products.
type ProductRepository interface {
	// GetByID retrieves a single product by its unique identifier.
	GetByID(ctx context.Context, id string) (*domain.Product, error)
	// ListAll retrieves all available products in the marketplace.
	ListAll(ctx context.Context) ([]domain.Product, error)
}

// UserRepository defines the data access methods for users.
type UserRepository interface {
	// GetByID retrieves a single user by their unique identifier.
	GetByID(ctx context.Context, id string) (*domain.User, error)
	// UpdateBalance updates the point balance for a specific user.
	// This operation should be atomic and ideally handle concurrency.
	UpdateBalance(ctx context.Context, userID string, newBalance int) error
}

// RedemptionRepository defines the data access methods for redemption records.
type RedemptionRepository interface {
	// Create records a new redemption event in the data store.
	Create(ctx context.Context, redemption *domain.Redemption) error
}

// TransactionManager defines an interface for managing database transactions.
// This ensures that multi-step operations like redemption are atomic.
type TransactionManager interface {
	// WithTransaction executes the given function within a database transaction.
	// It commits the transaction if the function returns no error, and rolls back otherwise.
	WithTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
}

// --- Service Definition ---

// Service provides the business logic for the marketplace.
// It orchestrates operations between users, products, and redemptions,
// ensuring business rules are enforced.
type Service interface {
	// ListProducts retrieves a list of all products available for redemption.
	ListProducts(ctx context.Context) ([]domain.Product, error)

	// SimulateRedemption calculates the impact of a redemption without actually performing it.
	// It returns a simulation result, indicating if the user can afford the product and
	// what their new balance would be.
	SimulateRedemption(ctx context.Context, userID, productID string) (*domain.RedemptionSimulation, error)

	// RedeemProduct allows a user to redeem a product using their points.
	// This is a transactional operation that deducts points from the user's balance
	// and records the redemption. It returns a detailed result upon success.
	RedeemProduct(ctx context.Context, userID, productID string) (*domain.RedemptionResult, error)
}

// service is the concrete implementation of the Service interface.
type service struct {
	productRepo    ProductRepository
	userRepo       UserRepository
	redemptionRepo RedemptionRepository
	txManager      TransactionManager
}

// Config holds the dependencies for creating a new marketplace service.
// Using a config struct makes initialization cleaner and more extensible.
type Config struct {
	ProductRepo    ProductRepository
	UserRepo       UserRepository
	RedemptionRepo RedemptionRepository
	TxManager      TransactionManager
}

// NewService creates and returns a new marketplace service instance.
// It validates that all necessary dependencies are provided.
func NewService(cfg Config) (Service, error) {
	if cfg.ProductRepo == nil {
		return nil, errors.New("marketplace service: product repository is required")
	}
	if cfg.UserRepo == nil {
		return nil, errors.New("marketplace service: user repository is required")
	}
	if cfg.RedemptionRepo == nil {
		return nil, errors.New("marketplace service: redemption repository is required")
	}
	if cfg.TxManager == nil {
		return nil, errors.New("marketplace service: transaction manager is required")
	}

	return &service{
		productRepo:    cfg.ProductRepo,
		userRepo:       cfg.UserRepo,
		redemptionRepo: cfg.RedemptionRepo,
		txManager:      cfg.TxManager,
	}, nil
}

// ListProducts retrieves a list of all products available for redemption.
func (s *service) ListProducts(ctx context.Context) ([]domain.Product, error) {
	products, err := s.productRepo.ListAll(ctx)
	if err != nil {
		// In a real app, we would log the internal error here.
		// e.g., s.logger.ErrorContext(ctx, "failed to list products", "error", err)
		return nil, fmt.Errorf("failed to retrieve products: %w", err)
	}
	return products, nil
}

// SimulateRedemption calculates the impact of a redemption without actually performing it.
func (s *service) SimulateRedemption(ctx context.Context, userID, productID string) (*domain.RedemptionSimulation, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		// Here we check for a specific repository error if possible, otherwise wrap.
		// For simplicity, we wrap it in our domain-specific error.
		return nil, fmt.Errorf("could not find user %s: %w", userID, ErrUserNotFound)
	}

	product, err := s.productRepo.GetByID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("could not find product %s: %w", productID, ErrProductNotFound)
	}

	canAfford := user.PointsBalance >= product.Cost
	newBalance := user.PointsBalance
	if canAfford {
		newBalance -= product.Cost
	}

	simulation := &domain.RedemptionSimulation{
		CanAfford:        canAfford,
		CurrentBalance:   user.PointsBalance,
		ProductCost:      product.Cost,
		ResultingBalance: newBalance,
	}

	return simulation, nil
}

// RedeemProduct allows a user to redeem a product using their points.
func (s *service) RedeemProduct(ctx context.Context, userID, productID string) (*domain.RedemptionResult, error) {
	var redemptionResult *domain.RedemptionResult

	// Use the transaction manager to ensure the entire operation is atomic.
	// If any step inside this function fails, all database changes will be rolled back.
	err := s.txManager.WithTransaction(ctx, func(txCtx context.Context) error {
		// Fetch user and product within the transaction to get the latest state and lock rows if necessary.
		user, err := s.userRepo.GetByID(txCtx, userID)
		if err != nil {
			return fmt.Errorf("failed to find user %s: %w", userID, ErrUserNotFound)
		}

		product, err := s.productRepo.GetByID(txCtx, productID)
		if err != nil {
			return fmt.Errorf("failed to find product %s: %w", productID, ErrProductNotFound)
		}

		// --- Business Rule Enforcement ---
		if user.PointsBalance < product.Cost {
			return ErrInsufficientFunds
		}

		// --- State Changes ---
		newBalance := user.PointsBalance - product.Cost
		if err := s.userRepo.UpdateBalance(txCtx, userID, newBalance); err != nil {
			// This could be a concurrency issue (e.g., optimistic lock failure) or a database error.
			return fmt.Errorf("failed to update user balance: %w", err)
		}

		// Create a record of the redemption for auditing and history.
		redemption := &domain.Redemption{
			ID:        uuid.NewString(),
			UserID:    userID,
			ProductID: productID,
			Cost:      product.Cost,
			Timestamp: time.Now().UTC(),
		}

		if err := s.redemptionRepo.Create(txCtx, redemption); err != nil {
			return fmt.Errorf("failed to record redemption: %w", err)
		}

		// Prepare the successful result to be returned outside the transaction scope.
		redemptionResult = &domain.RedemptionResult{
			ConfirmationID: redemption.ID,
			ProductID:      product.ID,
			ProductName:    product.Name,
			Cost:           product.Cost,
			NewBalance:     newBalance,
			Timestamp:      redemption.Timestamp,
		}

		return nil // Returning nil commits the transaction.
	})

	if err != nil {
		// The transaction was rolled back. The returned error will be one of the custom
		// errors (e.g., ErrInsufficientFunds) or a wrapped infrastructure error.
		return nil, err
	}

	// The transaction was successful.
	return redemptionResult, nil
}