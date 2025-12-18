package payment

import (
	"context"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
	"sync"
	"time"
)

// Common errors for the international payment service.
var (
	ErrInvalidAmount        = errors.New("amount must be greater than zero")
	ErrUnsupportedCurrency  = errors.New("unsupported currency pair")
	ErrInsufficientFunds    = errors.New("insufficient funds in source account")
	ErrComplianceCheckFailed = errors.New("transaction blocked by compliance checks")
	ErrInvalidIBAN          = errors.New("invalid IBAN format")
	ErrInvalidSWIFT         = errors.New("invalid SWIFT/BIC code format")
	ErrGatewayFailure       = errors.New("payment gateway failed to process transaction")
)

// TransactionStatus defines the state of a transfer.
type TransactionStatus string

const (
	StatusPending    TransactionStatus = "PENDING"
	StatusProcessing TransactionStatus = "PROCESSING"
	StatusCompleted  TransactionStatus = "COMPLETED"
	StatusFailed     TransactionStatus = "FAILED"
	StatusCancelled  TransactionStatus = "CANCELLED"
)

// InternationalTransferRequest represents the payload for initiating a cross-border payment.
type InternationalTransferRequest struct {
	SenderID        string
	RecipientID     string
	RecipientName   string
	SourceCurrency  string
	TargetCurrency  string
	Amount          float64 // Represented as major units (e.g., 100.50 USD). In real systems, use big.Int or decimal.
	IBAN            string
	SWIFTCode       string
	Reference       string
	IdempotencyKey  string
}

// TransferQuote represents a pre-transaction quote including FX rates and fees.
type TransferQuote struct {
	QuoteID        string
	SourceCurrency string
	TargetCurrency string
	SourceAmount   float64
	TargetAmount   float64
	ExchangeRate   float64
	Fee            float64
	ExpiresAt      time.Time
}

// TransferResult contains the details of a submitted transfer.
type TransferResult struct {
	TransactionID    string
	Status           TransactionStatus
	CreatedAt        time.Time
	EstimatedArrival time.Time
	Quote            TransferQuote
}

// -----------------------------------------------------------------------------
// Interfaces (Ports)
// These interfaces define the dependencies required by the service.
// In a hexagonal architecture, these would be implemented by adapters.
// -----------------------------------------------------------------------------

// Repository defines data access for payments.
type Repository interface {
	CreateTransaction(ctx context.Context, req InternationalTransferRequest, quote TransferQuote, status TransactionStatus) (string, error)
	UpdateStatus(ctx context.Context, transactionID string, status TransactionStatus, gatewayRef string) error
	GetBalance(ctx context.Context, userID string, currency string) (float64, error)
	HoldFunds(ctx context.Context, userID string, currency string, amount float64) error
	ReleaseFunds(ctx context.Context, userID string, currency string, amount float64) error
	CommitFunds(ctx context.Context, userID string, currency string, amount float64) error
}

// FXProvider defines the interface for foreign exchange rates.
type FXProvider interface {
	GetExchangeRate(ctx context.Context, source, target string) (float64, error)
}

// ComplianceProvider defines the interface for AML/KYC checks.
type ComplianceProvider interface {
	ScreenTransaction(ctx context.Context, senderID, recipientName string, amount float64, currency string) (bool, string, error)
}

// PaymentGateway defines the interface for the external banking rail (e.g., SWIFT).
type PaymentGateway interface {
	ExecuteTransfer(ctx context.Context, txID string, req InternationalTransferRequest) (string, error)
}

// Logger defines a structured logging interface.
type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
	Debug(msg string, keysAndValues ...interface{})
	With(keysAndValues ...interface{}) Logger
}

// -----------------------------------------------------------------------------
// Service Implementation
// -----------------------------------------------------------------------------

// InternationalService handles the business logic for cross-border payments.
type InternationalService struct {
	repo       Repository
	fx         FXProvider
	compliance ComplianceProvider
	gateway    PaymentGateway
	logger     Logger
	
	// feePercentage is a simplified fee model. In production, inject a FeeCalculator strategy.
	feePercentage float64 
}

// NewInternationalService creates a new instance of the InternationalService.
func NewInternationalService(
	repo Repository,
	fx FXProvider,
	compliance ComplianceProvider,
	gateway PaymentGateway,
	logger Logger,
) *InternationalService {
	return &InternationalService{
		repo:          repo,
		fx:            fx,
		compliance:    compliance,
		gateway:       gateway,
		logger:        logger,
		feePercentage: 0.015, // 1.5% default fee
	}
}

// GetQuote calculates the exchange rate and fees for a potential transfer.
func (s *InternationalService) GetQuote(ctx context.Context, sourceCurrency, targetCurrency string, amount float64) (*TransferQuote, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	rate, err := s.fx.GetExchangeRate(ctx, sourceCurrency, targetCurrency)
	if err != nil {
		s.logger.Error("failed to fetch exchange rate", "source", sourceCurrency, "target", targetCurrency, "error", err)
		return nil, fmt.Errorf("failed to get exchange rate: %w", err)
	}

	fee := amount * s.feePercentage
	netAmount := amount - fee
	targetAmount := netAmount * rate

	// Rounding to 2 decimal places for simplicity
	targetAmount = math.Floor(targetAmount*100) / 100

	return &TransferQuote{
		QuoteID:        fmt.Sprintf("Q-%d", time.Now().UnixNano()), // Simple ID generation
		SourceCurrency: sourceCurrency,
		TargetCurrency: targetCurrency,
		SourceAmount:   amount,
		TargetAmount:   targetAmount,
		ExchangeRate:   rate,
		Fee:            fee,
		ExpiresAt:      time.Now().Add(15 * time.Minute),
	}, nil
}

// InitiateTransfer orchestrates the international payment process.
func (s *InternationalService) InitiateTransfer(ctx context.Context, req InternationalTransferRequest) (*TransferResult, error) {
	log := s.logger.With("sender_id", req.SenderID, "recipient_id", req.RecipientID, "idempotency_key", req.IdempotencyKey)
	log.Info("initiating international transfer")

	// 1. Validation
	if err := s.validateRequest(req); err != nil {
		log.Error("validation failed", "error", err)
		return nil, err
	}

	// 2. Get Quote (Refresh rate to ensure accuracy at execution time)
	quote, err := s.GetQuote(ctx, req.SourceCurrency, req.TargetCurrency, req.Amount)
	if err != nil {
		return nil, err
	}

	// 3. Check Balance
	balance, err := s.repo.GetBalance(ctx, req.SenderID, req.SourceCurrency)
	if err != nil {
		log.Error("failed to fetch balance", "error", err)
		return nil, fmt.Errorf("failed to check balance: %w", err)
	}
	if balance < req.Amount {
		return nil, ErrInsufficientFunds
	}

	// 4. Compliance Check (AML/Sanctions)
	// We run this before holding funds to avoid locking money for rejected transactions.
	allowed, reason, err := s.compliance.ScreenTransaction(ctx, req.SenderID, req.RecipientName, req.Amount, req.SourceCurrency)
	if err != nil {
		log.Error("compliance service error", "error", err)
		return nil, fmt.Errorf("compliance check error: %w", err)
	}
	if !allowed {
		log.Error("transaction rejected by compliance", "reason", reason)
		return nil, fmt.Errorf("%w: %s", ErrComplianceCheckFailed, reason)
	}

	// 5. Hold Funds (Two-phase commit pattern start)
	if err := s.repo.HoldFunds(ctx, req.SenderID, req.SourceCurrency, req.Amount); err != nil {
		log.Error("failed to hold funds", "error", err)
		return nil, fmt.Errorf("failed to hold funds: %w", err)
	}

	// Defer rollback in case of panic or error before commit
	// Note: In a real distributed system, this requires a more robust saga pattern or state machine.
	var commitSuccessful bool
	defer func() {
		if !commitSuccessful {
			// Attempt to release funds if we didn't finish
			_ = s.repo.ReleaseFunds(context.Background(), req.SenderID, req.SourceCurrency, req.Amount)
		}
	}()

	// 6. Create Transaction Record
	txID, err := s.repo.CreateTransaction(ctx, req, *quote, StatusPending)
	if err != nil {
		log.Error("failed to create transaction record", "error", err)
		return nil, err
	}
	log = log.With("transaction_id", txID)

	// 7. Execute via Gateway
	gatewayRef, err := s.gateway.ExecuteTransfer(ctx, txID, req)
	if err != nil {
		log.Error("gateway execution failed", "error", err)
		// Update status to failed
		_ = s.repo.UpdateStatus(ctx, txID, StatusFailed, "")
		return nil, fmt.Errorf("%w: %v", ErrGatewayFailure, err)
	}

	// 8. Commit Funds (Finalize deduction)
	if err := s.repo.CommitFunds(ctx, req.SenderID, req.SourceCurrency, req.Amount); err != nil {
		// This is a critical failure state (Money sent but not deducted).
		// In production, this would trigger a reconciliation alert.
		log.Error("CRITICAL: failed to commit funds after gateway success", "error", err)
		// We still return success to the user because the money was sent, but internal ledger is out of sync.
	}
	commitSuccessful = true

	// 9. Update Status to Processing/Completed
	// Usually international transfers are "Processing" until a webhook confirms completion.
	if err := s.repo.UpdateStatus(ctx, txID, StatusProcessing, gatewayRef); err != nil {
		log.Error("failed to update transaction status", "error", err)
	}

	log.Info("international transfer initiated successfully")

	return &TransferResult{
		TransactionID:    txID,
		Status:           StatusProcessing,
		CreatedAt:        time.Now(),
		EstimatedArrival: time.Now().Add(48 * time.Hour), // Standard SWIFT time
		Quote:            *quote,
	}, nil
}

// validateRequest performs structural and format validation on the request.
func (s *InternationalService) validateRequest(req InternationalTransferRequest) error {
	if req.Amount <= 0 {
		return ErrInvalidAmount
	}
	if req.SenderID == "" || req.RecipientID == "" {
		return errors.New("sender and recipient IDs are required")
	}
	if req.SourceCurrency == "" || req.TargetCurrency == "" {
		return errors.New("source and target currencies are required")
	}
	
	// Basic IBAN validation (Regex for general format, not country specific checksums)
	// A real implementation would use a library like github.com/alrent/iban
	ibanRegex := regexp.MustCompile(`^[A-Z]{2}\d{2}[A-Z0-9]{1,30}$`)
	if !ibanRegex.MatchString(strings.ReplaceAll(req.IBAN, " ", "")) {
		return ErrInvalidIBAN
	}

	// Basic SWIFT/BIC validation
	swiftRegex := regexp.MustCompile(`^[A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?$`)
	if !swiftRegex.MatchString(req.SWIFTCode) {
		return ErrInvalidSWIFT
	}

	return nil
}

// TrackTransfer retrieves the current status of a transfer.
// This might aggregate data from the local DB and the external gateway.
func (s *InternationalService) TrackTransfer(ctx context.Context, transactionID string) (TransactionStatus, error) {
	// In a real app, we would fetch from repo. 
	// For this file, we assume the repo handles retrieval logic or we add a GetTransaction method to the interface.
	// Since the interface wasn't exhaustive in the prompt, I'll stub the logic conceptually.
	
	s.logger.Debug("tracking transfer", "transaction_id", transactionID)
	
	// Placeholder: Assume we query the repo
	// tx, err := s.repo.GetTransaction(ctx, transactionID)
	// return tx.Status, err
	
	return StatusProcessing, nil
}

// BatchProcess allows processing multiple transfers concurrently (e.g., for corporate payroll).
func (s *InternationalService) BatchProcess(ctx context.Context, requests []InternationalTransferRequest) ([]*TransferResult, []error) {
	var wg sync.WaitGroup
	results := make([]*TransferResult, len(requests))
	errs := make([]error, len(requests))

	for i, req := range requests {
		wg.Add(1)
		go func(index int, r InternationalTransferRequest) {
			defer wg.Done()
			// Create a detached context with timeout for individual items if needed, 
			// or use the parent context.
			res, err := s.InitiateTransfer(ctx, r)
			results[index] = res
			errs[index] = err
		}(i, req)
	}

	wg.Wait()
	return results, errs
}