package corporate

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"log/slog"
)

// Common errors for the Treasury Service.
var (
	ErrInvalidAccountID = errors.New("invalid account ID provided")
	ErrInvalidTimeRange = errors.New("invalid time range for forecasting")
	ErrDataSourceFailed = errors.New("failed to retrieve data from underlying source")
)

// Currency represents a standard ISO 4217 currency code.
type Currency string

// Money represents a monetary value with high precision.
type Money struct {
	Amount   *big.Float
	Currency Currency
}

// LiquidityPosition represents the current cash position of an account or entity.
type LiquidityPosition struct {
	AccountID   string    `json:"account_id"`
	EntityID    string    `json:"entity_id"`
	TotalAssets Money     `json:"total_assets"`
	Liabilities Money     `json:"liabilities"`
	NetPosition Money     `json:"net_position"`
	LastUpdated time.Time `json:"last_updated"`
}

// CashFlowForecastRequest defines parameters for generating a forecast.
type CashFlowForecastRequest struct {
	AccountIDs []string  `json:"account_ids"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Interval   string    `json:"interval"` // e.g., "daily", "weekly", "monthly"
}

// ForecastInterval represents a single time slice in the forecast.
type ForecastInterval struct {
	PeriodStart    time.Time `json:"period_start"`
	PeriodEnd      time.Time `json:"period_end"`
	ProjectedIn    Money     `json:"projected_in"`
	ProjectedOut   Money     `json:"projected_out"`
	NetFlow        Money     `json:"net_flow"`
	ConfidenceRate float64   `json:"confidence_rate"` // 0.0 to 1.0
}

// CashFlowForecastResponse contains the generated forecast data.
type CashFlowForecastResponse struct {
	GeneratedAt time.Time          `json:"generated_at"`
	Currency    Currency           `json:"currency"`
	Intervals   []ForecastInterval `json:"intervals"`
	TotalNet    Money              `json:"total_net"`
}

// TreasuryRepository defines the data access layer requirements for treasury operations.
type TreasuryRepository interface {
	GetAccountBalance(ctx context.Context, accountID string) (Money, error)
	GetPendingLiabilities(ctx context.Context, accountID string) (Money, error)
	GetScheduledTransactions(ctx context.Context, accountIDs []string, start, end time.Time) ([]ScheduledTransaction, error)
	GetHistoricalFlows(ctx context.Context, accountIDs []string, start, end time.Time) ([]HistoricalFlow, error)
}

// ScheduledTransaction represents a known future transaction.
type ScheduledTransaction struct {
	ID        string
	AccountID string
	Amount    *big.Float
	Direction string // "INFLOW" or "OUTFLOW"
	DueDate   time.Time
	Currency  Currency
}

// HistoricalFlow represents past transaction data for trend analysis.
type HistoricalFlow struct {
	Date      time.Time
	Amount    *big.Float
	Direction string
	Currency  Currency
}

// TreasuryService defines the interface for corporate treasury operations.
type TreasuryService interface {
	GetLiquidityPosition(ctx context.Context, accountID string) (*LiquidityPosition, error)
	GetConsolidatedLiquidity(ctx context.Context, entityID string, accountIDs []string) (*LiquidityPosition, error)
	GenerateCashFlowForecast(ctx context.Context, req CashFlowForecastRequest) (*CashFlowForecastResponse, error)
}

// treasuryServiceImpl implements the TreasuryService interface.
type treasuryServiceImpl struct {
	repo   TreasuryRepository
	logger *slog.Logger
}

// NewTreasuryService creates a new instance of the Treasury Service.
func NewTreasuryService(repo TreasuryRepository, logger *slog.Logger) TreasuryService {
	return &treasuryServiceImpl{
		repo:   repo,
		logger: logger,
	}
}

// GetLiquidityPosition retrieves the current liquidity stance for a specific account.
func (s *treasuryServiceImpl) GetLiquidityPosition(ctx context.Context, accountID string) (*LiquidityPosition, error) {
	if accountID == "" {
		return nil, ErrInvalidAccountID
	}

	s.logger.InfoContext(ctx, "fetching liquidity position", "account_id", accountID)

	// Fetch assets and liabilities concurrently
	var (
		assets      Money
		liabilities Money
		errAssets   error
		errLiabs    error
		wg          sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		assets, errAssets = s.repo.GetAccountBalance(ctx, accountID)
	}()

	go func() {
		defer wg.Done()
		liabilities, errLiabs = s.repo.GetPendingLiabilities(ctx, accountID)
	}()

	wg.Wait()

	if errAssets != nil {
		s.logger.ErrorContext(ctx, "failed to fetch assets", "error", errAssets)
		return nil, fmt.Errorf("%w: assets", ErrDataSourceFailed)
	}
	if errLiabs != nil {
		s.logger.ErrorContext(ctx, "failed to fetch liabilities", "error", errLiabs)
		return nil, fmt.Errorf("%w: liabilities", ErrDataSourceFailed)
	}

	// Calculate Net Position
	netVal := new(big.Float).Sub(assets.Amount, liabilities.Amount)

	return &LiquidityPosition{
		AccountID:   accountID,
		TotalAssets: assets,
		Liabilities: liabilities,
		NetPosition: Money{Amount: netVal, Currency: assets.Currency},
		LastUpdated: time.Now().UTC(),
	}, nil
}

// GetConsolidatedLiquidity aggregates liquidity positions across multiple accounts for an entity.
func (s *treasuryServiceImpl) GetConsolidatedLiquidity(ctx context.Context, entityID string, accountIDs []string) (*LiquidityPosition, error) {
	if len(accountIDs) == 0 {
		return nil, errors.New("no accounts provided for consolidation")
	}

	s.logger.InfoContext(ctx, "calculating consolidated liquidity", "entity_id", entityID, "account_count", len(accountIDs))

	totalAssets := new(big.Float).SetFloat64(0)
	totalLiabilities := new(big.Float).SetFloat64(0)
	// Assuming single currency consolidation for simplicity, or base currency conversion logic would go here.
	// For this implementation, we assume all accounts are in USD or normalized by the repo.
	baseCurrency := Currency("USD")

	// Use a worker pool if account list is large, but simple concurrency for now
	type result struct {
		pos *LiquidityPosition
		err error
	}
	results := make(chan result, len(accountIDs))
	var wg sync.WaitGroup

	for _, accID := range accountIDs {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			pos, err := s.GetLiquidityPosition(ctx, id)
			results <- result{pos: pos, err: err}
		}(accID)
	}

	wg.Wait()
	close(results)

	for res := range results {
		if res.err != nil {
			s.logger.WarnContext(ctx, "partial failure in consolidation", "error", res.err)
			continue // Partial success strategy: skip failed accounts but log them
		}
		totalAssets.Add(totalAssets, res.pos.TotalAssets.Amount)
		totalLiabilities.Add(totalLiabilities, res.pos.Liabilities.Amount)
	}

	netVal := new(big.Float).Sub(totalAssets, totalLiabilities)

	return &LiquidityPosition{
		EntityID:    entityID,
		TotalAssets: Money{Amount: totalAssets, Currency: baseCurrency},
		Liabilities: Money{Amount: totalLiabilities, Currency: baseCurrency},
		NetPosition: Money{Amount: netVal, Currency: baseCurrency},
		LastUpdated: time.Now().UTC(),
	}, nil
}

// GenerateCashFlowForecast predicts future cash flows based on scheduled transactions and historical trends.
func (s *treasuryServiceImpl) GenerateCashFlowForecast(ctx context.Context, req CashFlowForecastRequest) (*CashFlowForecastResponse, error) {
	if req.StartDate.After(req.EndDate) {
		return nil, ErrInvalidTimeRange
	}

	s.logger.InfoContext(ctx, "generating cash flow forecast", "start", req.StartDate, "end", req.EndDate)

	// 1. Fetch Scheduled Transactions (Hard commitments)
	scheduled, err := s.repo.GetScheduledTransactions(ctx, req.AccountIDs, req.StartDate, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch scheduled transactions: %w", err)
	}

	// 2. Fetch Historical Data (For trend projection)
	// Look back period equal to forecast horizon for simple seasonality
	horizon := req.EndDate.Sub(req.StartDate)
	historyStart := req.StartDate.Add(-horizon)
	history, err := s.repo.GetHistoricalFlows(ctx, req.AccountIDs, historyStart, req.StartDate)
	if err != nil {
		s.logger.WarnContext(ctx, "failed to fetch historical data, proceeding with scheduled only", "error", err)
	}

	// 3. Build Intervals
	intervals := s.buildForecastIntervals(req.StartDate, req.EndDate, req.Interval)
	
	// 4. Populate Intervals
	s.populateForecast(intervals, scheduled, history)

	// 5. Aggregate Totals
	totalNet := new(big.Float).SetFloat64(0)
	for _, i := range intervals {
		totalNet.Add(totalNet, i.NetFlow.Amount)
	}

	return &CashFlowForecastResponse{
		GeneratedAt: time.Now().UTC(),
		Currency:    "USD", // Defaulting to base currency
		Intervals:   intervals,
		TotalNet:    Money{Amount: totalNet, Currency: "USD"},
	}, nil
}

// Helper to slice time range into intervals
func (s *treasuryServiceImpl) buildForecastIntervals(start, end time.Time, intervalType string) []ForecastInterval {
	var intervals []ForecastInterval
	current := start

	for current.Before(end) {
		next := current
		switch intervalType {
		case "weekly":
			next = current.AddDate(0, 0, 7)
		case "monthly":
			next = current.AddDate(0, 1, 0)
		default: // daily
			next = current.AddDate(0, 0, 1)
		}

		if next.After(end) {
			next = end
		}

		intervals = append(intervals, ForecastInterval{
			PeriodStart:  current,
			PeriodEnd:    next,
			ProjectedIn:  Money{Amount: new(big.Float), Currency: "USD"},
			ProjectedOut: Money{Amount: new(big.Float), Currency: "USD"},
			NetFlow:      Money{Amount: new(big.Float), Currency: "USD"},
			ConfidenceRate: 1.0, // Default confidence
		})
		current = next
	}
	return intervals
}

// Core logic to map transactions to intervals
func (s *treasuryServiceImpl) populateForecast(intervals []ForecastInterval, scheduled []ScheduledTransaction, history []HistoricalFlow) {
	// Map scheduled transactions
	for _, tx := range scheduled {
		for i := range intervals {
			// Check if transaction falls in interval [Start, End)
			if (tx.DueDate.Equal(intervals[i].PeriodStart) || tx.DueDate.After(intervals[i].PeriodStart)) && tx.DueDate.Before(intervals[i].PeriodEnd) {
				if tx.Direction == "INFLOW" {
					intervals[i].ProjectedIn.Amount.Add(intervals[i].ProjectedIn.Amount, tx.Amount)
				} else {
					intervals[i].ProjectedOut.Amount.Add(intervals[i].ProjectedOut.Amount, tx.Amount)
				}
			}
		}
	}

	// Apply simple historical average for uncommitted flows (Naive projection)
	// In a real ML model, this would be much more complex.
	if len(history) > 0 {
		avgDailyInflow, avgDailyOutflow := calculateDailyAverages(history)
		
		for i := range intervals {
			days := intervals[i].PeriodEnd.Sub(intervals[i].PeriodStart).Hours() / 24
			
			projectedIn := new(big.Float).Mul(avgDailyInflow, big.NewFloat(days))
			projectedOut := new(big.Float).Mul(avgDailyOutflow, big.NewFloat(days))

			// Add projection to scheduled (weighted lower confidence)
			intervals[i].ProjectedIn.Amount.Add(intervals[i].ProjectedIn.Amount, projectedIn)
			intervals[i].ProjectedOut.Amount.Add(intervals[i].ProjectedOut.Amount, projectedOut)
			
			// Adjust confidence because we mixed in estimates
			intervals[i].ConfidenceRate = 0.85
		}
	}

	// Calculate Net for each interval
	for i := range intervals {
		intervals[i].NetFlow.Amount.Sub(intervals[i].ProjectedIn.Amount, intervals[i].ProjectedOut.Amount)
	}
}

func calculateDailyAverages(history []HistoricalFlow) (*big.Float, *big.Float) {
	totalIn := new(big.Float)
	totalOut := new(big.Float)
	
	if len(history) == 0 {
		return totalIn, totalOut
	}

	minDate := history[0].Date
	maxDate := history[0].Date

	for _, h := range history {
		if h.Direction == "INFLOW" {
			totalIn.Add(totalIn, h.Amount)
		} else {
			totalOut.Add(totalOut, h.Amount)
		}
		if h.Date.Before(minDate) { minDate = h.Date }
		if h.Date.After(maxDate) { maxDate = h.Date }
	}

	days := maxDate.Sub(minDate).Hours() / 24
	if days < 1 {
		days = 1
	}

	avgIn := new(big.Float).Quo(totalIn, big.NewFloat(days))
	avgOut := new(big.Float).Quo(totalOut, big.NewFloat(days))

	return avgIn, avgOut
}