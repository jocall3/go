// Package transaction provides service-level logic for handling financial transactions and insights.
package transaction

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/user/financier/internal/domain"
	"github.com/user/financier/internal/platform/errors"
)

// AIClient defines the interface for an AI service client.
// This abstraction allows for swapping different AI providers (e.g., OpenAI, Gemini).
// The actual implementation would be in a separate package (e.g., internal/client/openai).
type AIClient interface {
	GenerateText(ctx context.Context, prompt string) (string, error)
}

// DataPoint represents a single point in a time-series data set, typically for charting.
type DataPoint struct {
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}

// SpendingTrends represents spending over a period of time.
type SpendingTrends struct {
	Period     string      `json:"period"`
	StartDate  time.Time   `json:"start_date"`
	EndDate    time.Time   `json:"end_date"`
	TotalSpend float64     `json:"total_spend"`
	Trends     []DataPoint `json:"trends"`
}

// CategorySpending represents the total spending for a specific category.
type CategorySpending struct {
	Category    string  `json:"category"`
	TotalAmount float64 `json:"total_amount"`
	Percentage  float64 `json:"percentage"`
}

// CategoryBreakdown represents the distribution of spending across categories.
type CategoryBreakdown struct {
	StartDate   time.Time          `json:"start_date"`
	EndDate     time.Time          `json:"end_date"`
	TotalSpend  float64            `json:"total_spend"`
	Breakdown   []CategorySpending `json:"breakdown"`
}

// AIInsight holds the AI-generated analysis of a user's spending habits.
type AIInsight struct {
	GeneratedAt time.Time `json:"generated_at"`
	InsightText string    `json:"insight_text"`
	Suggestions []string  `json:"suggestions"`
}

// SpendingForecast represents a prediction of future spending.
type SpendingForecast struct {
	ForecastPeriodDays int         `json:"forecast_period_days"`
	ProjectedSpend     float64     `json:"projected_spend"`
	ConfidenceLevel    float64     `json:"confidence_level"` // e.g., 0.85 for 85%
	ForecastedTrend    []DataPoint `json:"forecasted_trend"`
}

// InsightService defines the interface for transaction insight operations.
// It provides methods to analyze transaction data and generate meaningful reports.
type InsightService interface {
	GetSpendingTrends(ctx context.Context, userID string, period string) (*SpendingTrends, error)
	GetCategoryBreakdown(ctx context.Context, userID string, startDate, endDate time.Time) (*CategoryBreakdown, error)
	GenerateAIInsights(ctx context.Context, userID string) (*AIInsight, error)
	GetSpendingForecast(ctx context.Context, userID string) (*SpendingForecast, error)
}

// insightService is the concrete implementation of the InsightService.
type insightService struct {
	repo     domain.TransactionRepository
	aiClient AIClient
}

// NewInsightService creates a new instance of the insight service.
// It requires a transaction repository and an AI client as dependencies.
func NewInsightService(repo domain.TransactionRepository, aiClient AIClient) InsightService {
	return &insightService{
		repo:     repo,
		aiClient: aiClient,
	}
}

// GetSpendingTrends generates a report of spending trends over a given period.
// It aggregates spending by day to show a time-series view of expenditures.
func (s *insightService) GetSpendingTrends(ctx context.Context, userID string, period string) (*SpendingTrends, error) {
	startDate, endDate, err := parsePeriod(period)
	if err != nil {
		return nil, errors.NewInvalidInput("period", err.Error())
	}

	transactions, err := s.repo.FindByUserIDAndDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch transactions for spending trends")
	}

	// Aggregate spending by day
	dailySpending := make(map[time.Time]float64)
	totalSpend := 0.0
	for _, t := range transactions {
		// Only consider debit transactions (spending)
		if t.Amount < 0 {
			day := t.Date.Truncate(24 * time.Hour)
			dailySpending[day] += -t.Amount // Use positive value for spending
			totalSpend += -t.Amount
		}
	}

	// Convert map to a slice of DataPoints
	trends := make([]DataPoint, 0, len(dailySpending))
	for day, amount := range dailySpending {
		trends = append(trends, DataPoint{Date: day, Amount: amount})
	}

	// Sort by date for a clean timeline
	sort.Slice(trends, func(i, j int) bool {
		return trends[i].Date.Before(trends[j].Date)
	})

	return &SpendingTrends{
		Period:     period,
		StartDate:  startDate,
		EndDate:    endDate,
		TotalSpend: totalSpend,
		Trends:     trends,
	}, nil
}

// GetCategoryBreakdown provides a breakdown of spending by category for a given date range.
func (s *insightService) GetCategoryBreakdown(ctx context.Context, userID string, startDate, endDate time.Time) (*CategoryBreakdown, error) {
	if startDate.IsZero() || endDate.IsZero() || startDate.After(endDate) {
		return nil, errors.NewInvalidInput("date range", "start and end dates must be valid and in correct order")
	}

	transactions, err := s.repo.FindByUserIDAndDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch transactions for category breakdown")
	}

	categorySpending := make(map[string]float64)
	totalSpend := 0.0
	for _, t := range transactions {
		if t.Amount < 0 {
			amount := -t.Amount
			category := t.Category
			if category == "" {
				category = "Uncategorized"
			}
			categorySpending[category] += amount
			totalSpend += amount
		}
	}

	if totalSpend == 0 {
		return &CategoryBreakdown{
			StartDate:  startDate,
			EndDate:    endDate,
			TotalSpend: 0,
			Breakdown:  []CategorySpending{},
		}, nil
	}

	breakdown := make([]CategorySpending, 0, len(categorySpending))
	for category, amount := range categorySpending {
		breakdown = append(breakdown, CategorySpending{
			Category:    category,
			TotalAmount: amount,
			Percentage:  (amount / totalSpend) * 100,
		})
	}

	// Sort by amount descending to show highest spending categories first
	sort.Slice(breakdown, func(i, j int) bool {
		return breakdown[i].TotalAmount > breakdown[j].TotalAmount
	})

	return &CategoryBreakdown{
		StartDate:  startDate,
		EndDate:    endDate,
		TotalSpend: totalSpend,
		Breakdown:  breakdown,
	}, nil
}

// GenerateAIInsights uses an AI model to provide personalized financial insights based on recent transactions.
func (s *insightService) GenerateAIInsights(ctx context.Context, userID string) (*AIInsight, error) {
	// Fetch last 90 days of transactions for a meaningful analysis
	startDate := time.Now().AddDate(0, -3, 0)
	endDate := time.Now()

	transactions, err := s.repo.FindByUserIDAndDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch transactions for AI insights")
	}

	if len(transactions) < 10 { // Require a minimum number of transactions for useful insights
		return nil, errors.NewNotFound("transactions", "not enough transaction data to generate insights")
	}

	// Format transaction data for the AI prompt
	var promptBuilder strings.Builder
	promptBuilder.WriteString("You are a friendly and insightful financial advisor AI. Analyze the following list of transactions for a user and provide a concise summary of their spending habits, identify 1-2 potentially unusual or noteworthy transactions, and offer 3 actionable, personalized saving suggestions. Format your response exactly as follows, with each section on a new line:\n")
	promptBuilder.WriteString("[SUMMARY] Your summary here.\n")
	promptBuilder.WriteString("[SUGGESTIONS] Suggestion 1. | Suggestion 2. | Suggestion 3.\n")
	promptBuilder.WriteString("\nHere are the transactions (negative amount is spending):\n")

	for _, t := range transactions {
		promptBuilder.WriteString(fmt.Sprintf("- Date: %s, Amount: %.2f, Description: %s, Category: %s\n", t.Date.Format("2006-01-02"), t.Amount, t.Description, t.Category))
	}

	// Call the AI client
	aiResponse, err := s.aiClient.GenerateText(ctx, promptBuilder.String())
	if err != nil {
		return nil, errors.Wrap(err, "AI client failed to generate insights")
	}

	// Parse the structured response
	summary, suggestions := parseAIResponse(aiResponse)
	if summary == "" {
		// Fallback if parsing fails
		summary = aiResponse
	}

	return &AIInsight{
		GeneratedAt: time.Now(),
		InsightText: summary,
		Suggestions: suggestions,
	}, nil
}

// GetSpendingForecast projects future spending based on historical data using a simple linear projection.
func (s *insightService) GetSpendingForecast(ctx context.Context, userID string) (*SpendingForecast, error) {
	const historyDays = 90
	const forecastDays = 30

	// Fetch historical data
	startDate := time.Now().AddDate(0, 0, -historyDays)
	endDate := time.Now()

	transactions, err := s.repo.FindByUserIDAndDateRange(ctx, userID, startDate, endDate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch transactions for forecast")
	}

	if len(transactions) < 5 {
		return nil, errors.NewNotFound("transactions", "not enough historical data to create a forecast")
	}

	// Calculate average daily spend
	totalSpend := 0.0
	for _, t := range transactions {
		if t.Amount < 0 {
			totalSpend += -t.Amount
		}
	}
	avgDailySpend := totalSpend / float64(historyDays)

	// Generate future data points
	forecastedTrend := make([]DataPoint, forecastDays)
	today := time.Now().Truncate(24 * time.Hour)
	for i := 0; i < forecastDays; i++ {
		forecastedTrend[i] = DataPoint{
			Date:   today.AddDate(0, 0, i+1),
			Amount: avgDailySpend,
		}
	}

	return &SpendingForecast{
		ForecastPeriodDays: forecastDays,
		ProjectedSpend:     avgDailySpend * float64(forecastDays),
		ConfidenceLevel:    0.75, // Static confidence for this simple model
		ForecastedTrend:    forecastedTrend,
	}, nil
}

// parsePeriod is a helper function to convert a string like "7d" or "30d" into a time range.
func parsePeriod(period string) (time.Time, time.Time, error) {
	now := time.Now()
	endDate := now
	var startDate time.Time

	switch strings.ToLower(period) {
	case "7d":
		startDate = now.AddDate(0, 0, -7)
	case "30d":
		startDate = now.AddDate(0, -1, 0)
	case "90d":
		startDate = now.AddDate(0, -3, 0)
	case "1y":
		startDate = now.AddDate(-1, 0, 0)
	default:
		return time.Time{}, time.Time{}, fmt.Errorf("invalid period '%s', must be one of: 7d, 30d, 90d, 1y", period)
	}

	return startDate, endDate, nil
}

// parseAIResponse is a helper to extract structured data from the AI's text response.
func parseAIResponse(response string) (summary string, suggestions []string) {
	lines := strings.Split(response, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "[SUMMARY]") {
			summary = strings.TrimSpace(strings.TrimPrefix(line, "[SUMMARY]"))
		} else if strings.HasPrefix(line, "[SUGGESTIONS]") {
			suggestionStr := strings.TrimSpace(strings.TrimPrefix(line, "[SUGGESTIONS]"))
			rawSuggestions := strings.Split(suggestionStr, "|")
			for _, s := range rawSuggestions {
				trimmed := strings.TrimSpace(s)
				if trimmed != "" {
					suggestions = append(suggestions, trimmed)
				}
			}
		}
	}
	return
}