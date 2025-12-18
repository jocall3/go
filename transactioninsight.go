// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// TransactionInsightService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionInsightService] method instead.
type TransactionInsightService struct {
	Options []option.RequestOption
}

// NewTransactionInsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionInsightService(opts ...option.RequestOption) (r *TransactionInsightService) {
	r = &TransactionInsightService{}
	r.Options = opts
	return
}

// Retrieves AI-generated insights into user spending trends over time, identifying
// patterns and anomalies.
func (r *TransactionInsightService) GetSpendingTrends(ctx context.Context, opts ...option.RequestOption) (res *TransactionInsightGetSpendingTrendsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/insights/spending-trends"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AIInsight struct {
	ID                       string            `json:"id"`
	ActionableRecommendation string            `json:"actionableRecommendation"`
	Category                 AIInsightCategory `json:"category"`
	Description              string            `json:"description"`
	Severity                 AIInsightSeverity `json:"severity"`
	Timestamp                time.Time         `json:"timestamp" format:"date-time"`
	Title                    string            `json:"title"`
	JSON                     aiInsightJSON     `json:"-"`
}

// aiInsightJSON contains the JSON metadata for the struct [AIInsight]
type aiInsightJSON struct {
	ID                       apijson.Field
	ActionableRecommendation apijson.Field
	Category                 apijson.Field
	Description              apijson.Field
	Severity                 apijson.Field
	Timestamp                apijson.Field
	Title                    apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AIInsight) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInsightJSON) RawJSON() string {
	return r.raw
}

type AIInsightCategory string

const (
	AIInsightCategorySpending       AIInsightCategory = "spending"
	AIInsightCategorySaving         AIInsightCategory = "saving"
	AIInsightCategoryInvesting      AIInsightCategory = "investing"
	AIInsightCategoryBudget         AIInsightCategory = "budget"
	AIInsightCategorySecurity       AIInsightCategory = "security"
	AIInsightCategoryCompliance     AIInsightCategory = "compliance"
	AIInsightCategoryTreasury       AIInsightCategory = "treasury"
	AIInsightCategorySustainability AIInsightCategory = "sustainability"
)

func (r AIInsightCategory) IsKnown() bool {
	switch r {
	case AIInsightCategorySpending, AIInsightCategorySaving, AIInsightCategoryInvesting, AIInsightCategoryBudget, AIInsightCategorySecurity, AIInsightCategoryCompliance, AIInsightCategoryTreasury, AIInsightCategorySustainability:
		return true
	}
	return false
}

type AIInsightSeverity string

const (
	AIInsightSeverityLow    AIInsightSeverity = "low"
	AIInsightSeverityMedium AIInsightSeverity = "medium"
	AIInsightSeverityHigh   AIInsightSeverity = "high"
)

func (r AIInsightSeverity) IsKnown() bool {
	switch r {
	case AIInsightSeverityLow, AIInsightSeverityMedium, AIInsightSeverityHigh:
		return true
	}
	return false
}

type TransactionInsightGetSpendingTrendsResponse struct {
	AIInsights            []AIInsight                                                        `json:"aiInsights"`
	ForecastNextMonth     float64                                                            `json:"forecastNextMonth"`
	OverallTrend          TransactionInsightGetSpendingTrendsResponseOverallTrend            `json:"overallTrend"`
	PercentageChange      float64                                                            `json:"percentageChange"`
	Period                string                                                             `json:"period"`
	TopCategoriesByChange []TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange `json:"topCategoriesByChange"`
	JSON                  transactionInsightGetSpendingTrendsResponseJSON                    `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseJSON contains the JSON metadata for
// the struct [TransactionInsightGetSpendingTrendsResponse]
type transactionInsightGetSpendingTrendsResponseJSON struct {
	AIInsights            apijson.Field
	ForecastNextMonth     apijson.Field
	OverallTrend          apijson.Field
	PercentageChange      apijson.Field
	Period                apijson.Field
	TopCategoriesByChange apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *TransactionInsightGetSpendingTrendsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionInsightGetSpendingTrendsResponseOverallTrend string

const (
	TransactionInsightGetSpendingTrendsResponseOverallTrendIncreasing TransactionInsightGetSpendingTrendsResponseOverallTrend = "increasing"
	TransactionInsightGetSpendingTrendsResponseOverallTrendDecreasing TransactionInsightGetSpendingTrendsResponseOverallTrend = "decreasing"
	TransactionInsightGetSpendingTrendsResponseOverallTrendStable     TransactionInsightGetSpendingTrendsResponseOverallTrend = "stable"
)

func (r TransactionInsightGetSpendingTrendsResponseOverallTrend) IsKnown() bool {
	switch r {
	case TransactionInsightGetSpendingTrendsResponseOverallTrendIncreasing, TransactionInsightGetSpendingTrendsResponseOverallTrendDecreasing, TransactionInsightGetSpendingTrendsResponseOverallTrendStable:
		return true
	}
	return false
}

type TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange struct {
	AbsoluteChange   float64                                                              `json:"absoluteChange"`
	Category         string                                                               `json:"category"`
	PercentageChange float64                                                              `json:"percentageChange"`
	JSON             transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON `json:"-"`
}

// transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON contains
// the JSON metadata for the struct
// [TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange]
type transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON struct {
	AbsoluteChange   apijson.Field
	Category         apijson.Field
	PercentageChange apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransactionInsightGetSpendingTrendsResponseTopCategoriesByChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionInsightGetSpendingTrendsResponseTopCategoriesByChangeJSON) RawJSON() string {
	return r.raw
}
