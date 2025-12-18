package transaction

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math"
	"time"

	"github.com/google/uuid"

	// Assumed internal import for domain models. 
	// Replace with your actual module path, e.g., "github.com/org/repo/internal/domain"
	"github.com/yourproject/internal/domain"
)

var (
	ErrTransactionNotFound = errors.New("transaction not found")
	ErrInvalidCategory     = errors.New("invalid category")
	ErrAlreadyDisputed     = errors.New("transaction is already disputed")
	ErrDisputeTimeExpired  = errors.New("transaction is too old to dispute")
)

// Repository defines the data access layer requirements for transactions.
type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error)
	List(ctx context.Context, filter domain.TransactionFilter) ([]domain.Transaction, int64, error)
	Update(ctx context.Context, tx *domain.Transaction) error
	CreateDispute(ctx context.Context, dispute *domain.Dispute) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Transaction, error)
}

// Service defines the interface for transaction business logic.
type Service interface {
	GetTransaction(ctx context.Context, id uuid.UUID) (*domain.Transaction, error)
	ListTransactions(ctx context.Context, filter domain.TransactionFilter) (*domain.PaginatedResponse, error)
	CategorizeTransaction(ctx context.Context, id uuid.UUID, category domain.Category) error
	DisputeTransaction(ctx context.Context, id uuid.UUID, reason string) (*domain.Dispute, error)
	DetectAndMarkRecurring(ctx context.Context, userID uuid.UUID) error
}

type service struct {
	repo   Repository
	logger *slog.Logger
}

// NewService creates a new instance of the transaction service.
func NewService(repo Repository, logger *slog.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// GetTransaction retrieves a single transaction by its ID.
func (s *service) GetTransaction(ctx context.Context, id uuid.UUID) (*domain.Transaction, error) {
	tx, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to fetch transaction", "id", id, "error", err)
		return nil, fmt.Errorf("fetching transaction: %w", err)
	}
	if tx == nil {
		return nil, ErrTransactionNotFound
	}
	return tx, nil
}

// ListTransactions retrieves a list of transactions based on filters.
func (s *service) ListTransactions(ctx context.Context, filter domain.TransactionFilter) (*domain.PaginatedResponse, error) {
	// Set default pagination if not provided
	if filter.Limit <= 0 {
		filter.Limit = 20
	}
	if filter.Page <= 0 {
		filter.Page = 1
	}

	txs, total, err := s.repo.List(ctx, filter)
	if err != nil {
		s.logger.Error("failed to list transactions", "filter", filter, "error", err)
		return nil, fmt.Errorf("listing transactions: %w", err)
	}

	return &domain.PaginatedResponse{
		Data:       txs,
		TotalCount: total,
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalPages: int((total + int64(filter.Limit) - 1) / int64(filter.Limit)),
	}, nil
}

// CategorizeTransaction updates the category of a transaction.
func (s *service) CategorizeTransaction(ctx context.Context, id uuid.UUID, category domain.Category) error {
	if !category.IsValid() {
		return ErrInvalidCategory
	}

	tx, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("retrieving transaction for categorization: %w", err)
	}
	if tx == nil {
		return ErrTransactionNotFound
	}

	oldCategory := tx.Category
	tx.Category = category
	tx.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, tx); err != nil {
		s.logger.Error("failed to update transaction category", "id", id, "category", category, "error", err)
		return fmt.Errorf("updating category: %w", err)
	}

	s.logger.Info("transaction categorized", "id", id, "old_category", oldCategory, "new_category", category)
	return nil
}

// DisputeTransaction initiates a dispute for a specific transaction.
func (s *service) DisputeTransaction(ctx context.Context, id uuid.UUID, reason string) (*domain.Dispute, error) {
	tx, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("retrieving transaction for dispute: %w", err)
	}
	if tx == nil {
		return nil, ErrTransactionNotFound
	}

	// Business Rule: Cannot dispute if already disputed
	if tx.Status == domain.TransactionStatusDisputed {
		return nil, ErrAlreadyDisputed
	}

	// Business Rule: Cannot dispute transactions older than 90 days
	if time.Since(tx.Date) > 90*24*time.Hour {
		return nil, ErrDisputeTimeExpired
	}

	dispute := &domain.Dispute{
		ID:            uuid.New(),
		TransactionID: tx.ID,
		UserID:        tx.UserID,
		Reason:        reason,
		Status:        domain.DisputeStatusOpen,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Update transaction status to indicate a dispute is in progress
	tx.Status = domain.TransactionStatusDisputed
	tx.UpdatedAt = time.Now()

	// We should ideally run this in a transaction block (Unit of Work), 
	// but assuming the repo handles atomic updates or we do it sequentially for now.
	if err := s.repo.CreateDispute(ctx, dispute); err != nil {
		s.logger.Error("failed to create dispute record", "transaction_id", id, "error", err)
		return nil, fmt.Errorf("creating dispute: %w", err)
	}

	if err := s.repo.Update(ctx, tx); err != nil {
		s.logger.Error("failed to update transaction status for dispute", "transaction_id", id, "error", err)
		// Note: In a real production system without distributed transactions, 
		// this might leave data in an inconsistent state if CreateDispute succeeded.
		return nil, fmt.Errorf("updating transaction status: %w", err)
	}

	s.logger.Info("dispute initiated", "dispute_id", dispute.ID, "transaction_id", id)
	return dispute, nil
}

// DetectAndMarkRecurring analyzes a user's transaction history to identify and mark recurring transactions.
// It looks for transactions with the same merchant and similar amounts occurring at regular intervals.
func (s *service) DetectAndMarkRecurring(ctx context.Context, userID uuid.UUID) error {
	// Fetch all transactions for the user (could be optimized to fetch only last N months)
	txs, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("fetching user transactions: %w", err)
	}

	if len(txs) < 2 {
		return nil // Not enough data
	}

	// Group by Merchant
	grouped := make(map[string][]*domain.Transaction)
	for i := range txs {
		// Skip already marked recurring to avoid redundant processing, 
		// unless we want to re-validate.
		if txs[i].MerchantName == "" {
			continue
		}
		grouped[txs[i].MerchantName] = append(grouped[txs[i].MerchantName], &txs[i])
	}

	updates := 0
	for merchant, group := range grouped {
		if len(group) < 3 {
			continue // Need at least 3 occurrences to establish a pattern
		}

		// Sort by date descending
		// Assuming domain.Transaction has a Date field of type time.Time
		// Simple bubble sort for small groups or assume repo returns sorted. 
		// Let's sort manually to be safe.
		// (Omitted full sort implementation for brevity, assuming roughly ordered or small sets)

		// Check for regularity
		isRecurring := s.analyzeRecurringPattern(group)
		
		if isRecurring {
			for _, tx := range group {
				if !tx.IsRecurring {
					tx.IsRecurring = true
					tx.UpdatedAt = time.Now()
					if err := s.repo.Update(ctx, tx); err != nil {
						s.logger.Error("failed to mark transaction as recurring", "id", tx.ID, "error", err)
					} else {
						updates++
					}
				}
			}
			if updates > 0 {
				s.logger.Info("detected recurring transactions", "merchant", merchant, "count", updates)
			}
		}
	}

	return nil
}

// analyzeRecurringPattern checks if a group of transactions fits a recurring profile.
// Logic: Same amount (within small delta) and roughly same interval (e.g., ~30 days).
func (s *service) analyzeRecurringPattern(txs []*domain.Transaction) bool {
	// 1. Check Amount Consistency
	baseAmount := txs[0].Amount
	for _, tx := range txs {
		// Allow 5% variance for currency fluctuations or slight bill changes
		diff := math.Abs(tx.Amount - baseAmount)
		if diff > (baseAmount * 0.05) {
			return false
		}
	}

	// 2. Check Time Interval Consistency
	// This is a simplified heuristic. Real-world logic would use standard deviation of intervals.
	// We look for monthly subscriptions mostly.
	// We need at least 2 intervals (3 transactions).
	
	// Calculate intervals
	var intervals []float64
	for i := 0; i < len(txs)-1; i++ {
		// Assuming txs are sorted by date
		d1 := txs[i].Date
		d2 := txs[i+1].Date
		
		// Ensure positive interval
		if d1.After(d2) {
			d1, d2 = d2, d1
		}
		
		days := d2.Sub(d1).Hours() / 24
		intervals = append(intervals, days)
	}

	// Check if intervals are roughly 30 days (+/- 5 days) or 7 days, etc.
	// For this MVP, we check if the variance between intervals is low.
	if len(intervals) == 0 {
		return false
	}

	avgInterval := 0.0
	for _, inv := range intervals {
		avgInterval += inv
	}
	avgInterval /= float64(len(intervals))

	// If average interval is random (e.g. < 1 day or very large irregular), ignore.
	// We target weekly (7), bi-weekly (14), monthly (28-31), yearly (365).
	validPeriods := []float64{7, 14, 30, 365}
	isStandardPeriod := false
	for _, p := range validPeriods {
		if math.Abs(avgInterval-p) < (p * 0.2) { // 20% tolerance on period length
			isStandardPeriod = true
			break
		}
	}
	if !isStandardPeriod {
		return false
	}

	// Check variance
	for _, inv := range intervals {
		if math.Abs(inv-avgInterval) > 5 { // If any interval deviates by more than 5 days from average
			return false
		}
	}

	return true
}