// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/apiquery"
	"github.com/stainless-sdks/1231-go/internal/param"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// InvestmentPortfolioService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestmentPortfolioService] method instead.
type InvestmentPortfolioService struct {
	Options []option.RequestOption
}

// NewInvestmentPortfolioService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestmentPortfolioService(opts ...option.RequestOption) (r *InvestmentPortfolioService) {
	r = &InvestmentPortfolioService{}
	r.Options = opts
	return
}

// Creates a new investment portfolio, with options for initial asset allocation.
func (r *InvestmentPortfolioService) New(ctx context.Context, body InvestmentPortfolioNewParams, opts ...option.RequestOption) (res *InvestmentPortfolio, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/portfolios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information for a specific investment portfolio, including
// holdings, performance, and AI insights.
func (r *InvestmentPortfolioService) Get(ctx context.Context, portfolioID string, opts ...option.RequestOption) (res *InvestmentPortfolio, err error) {
	opts = slices.Concat(r.Options, opts)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates high-level details of an investment portfolio, such as name or risk
// tolerance.
func (r *InvestmentPortfolioService) Update(ctx context.Context, portfolioID string, body InvestmentPortfolioUpdateParams, opts ...option.RequestOption) (res *InvestmentPortfolio, err error) {
	opts = slices.Concat(r.Options, opts)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a summary of all investment portfolios linked to the user's account.
func (r *InvestmentPortfolioService) List(ctx context.Context, query InvestmentPortfolioListParams, opts ...option.RequestOption) (res *InvestmentPortfolioListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/portfolios"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Triggers an AI-driven rebalancing process for a specific investment portfolio
// based on a target risk tolerance or strategy.
func (r *InvestmentPortfolioService) Rebalance(ctx context.Context, portfolioID string, body InvestmentPortfolioRebalanceParams, opts ...option.RequestOption) (res *InvestmentPortfolioRebalanceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if portfolioID == "" {
		err = errors.New("missing required portfolioId parameter")
		return
	}
	path := fmt.Sprintf("investments/portfolios/%s/rebalance", portfolioID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type InvestmentPortfolio struct {
	ID                    string                           `json:"id"`
	AIPerformanceInsights []AIInsight                      `json:"aiPerformanceInsights"`
	Currency              string                           `json:"currency"`
	Holdings              []InvestmentPortfolioHolding     `json:"holdings"`
	LastUpdated           time.Time                        `json:"lastUpdated" format:"date-time"`
	Name                  string                           `json:"name"`
	RiskTolerance         InvestmentPortfolioRiskTolerance `json:"riskTolerance"`
	TodayGainLoss         float64                          `json:"todayGainLoss"`
	TotalValue            float64                          `json:"totalValue"`
	Type                  string                           `json:"type"`
	UnrealizedGainLoss    float64                          `json:"unrealizedGainLoss"`
	JSON                  investmentPortfolioJSON          `json:"-"`
}

// investmentPortfolioJSON contains the JSON metadata for the struct
// [InvestmentPortfolio]
type investmentPortfolioJSON struct {
	ID                    apijson.Field
	AIPerformanceInsights apijson.Field
	Currency              apijson.Field
	Holdings              apijson.Field
	LastUpdated           apijson.Field
	Name                  apijson.Field
	RiskTolerance         apijson.Field
	TodayGainLoss         apijson.Field
	TotalValue            apijson.Field
	Type                  apijson.Field
	UnrealizedGainLoss    apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvestmentPortfolio) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioHolding struct {
	AverageCost           float64                        `json:"averageCost"`
	CurrentPrice          float64                        `json:"currentPrice"`
	EsgScore              float64                        `json:"esgScore"`
	MarketValue           float64                        `json:"marketValue"`
	Name                  string                         `json:"name"`
	PercentageOfPortfolio float64                        `json:"percentageOfPortfolio"`
	Quantity              float64                        `json:"quantity"`
	Symbol                string                         `json:"symbol"`
	JSON                  investmentPortfolioHoldingJSON `json:"-"`
}

// investmentPortfolioHoldingJSON contains the JSON metadata for the struct
// [InvestmentPortfolioHolding]
type investmentPortfolioHoldingJSON struct {
	AverageCost           apijson.Field
	CurrentPrice          apijson.Field
	EsgScore              apijson.Field
	MarketValue           apijson.Field
	Name                  apijson.Field
	PercentageOfPortfolio apijson.Field
	Quantity              apijson.Field
	Symbol                apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvestmentPortfolioHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioHoldingJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioRiskTolerance string

const (
	InvestmentPortfolioRiskToleranceLow        InvestmentPortfolioRiskTolerance = "low"
	InvestmentPortfolioRiskToleranceMedium     InvestmentPortfolioRiskTolerance = "medium"
	InvestmentPortfolioRiskToleranceHigh       InvestmentPortfolioRiskTolerance = "high"
	InvestmentPortfolioRiskToleranceAggressive InvestmentPortfolioRiskTolerance = "aggressive"
)

func (r InvestmentPortfolioRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRiskToleranceLow, InvestmentPortfolioRiskToleranceMedium, InvestmentPortfolioRiskToleranceHigh, InvestmentPortfolioRiskToleranceAggressive:
		return true
	}
	return false
}

type InvestmentPortfolioListResponse struct {
	Data []InvestmentPortfolio               `json:"data"`
	JSON investmentPortfolioListResponseJSON `json:"-"`
	PaginatedList
}

// investmentPortfolioListResponseJSON contains the JSON metadata for the struct
// [InvestmentPortfolioListResponse]
type investmentPortfolioListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestmentPortfolioListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioRebalanceResponse struct {
	ConfirmationExpiresAt time.Time                                  `json:"confirmationExpiresAt" format:"date-time"`
	ConfirmationRequired  bool                                       `json:"confirmationRequired"`
	EstimatedImpact       string                                     `json:"estimatedImpact"`
	PortfolioID           string                                     `json:"portfolioId"`
	RebalanceID           string                                     `json:"rebalanceId"`
	Status                InvestmentPortfolioRebalanceResponseStatus `json:"status"`
	StatusMessage         string                                     `json:"statusMessage"`
	JSON                  investmentPortfolioRebalanceResponseJSON   `json:"-"`
}

// investmentPortfolioRebalanceResponseJSON contains the JSON metadata for the
// struct [InvestmentPortfolioRebalanceResponse]
type investmentPortfolioRebalanceResponseJSON struct {
	ConfirmationExpiresAt apijson.Field
	ConfirmationRequired  apijson.Field
	EstimatedImpact       apijson.Field
	PortfolioID           apijson.Field
	RebalanceID           apijson.Field
	Status                apijson.Field
	StatusMessage         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *InvestmentPortfolioRebalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentPortfolioRebalanceResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentPortfolioRebalanceResponseStatus string

const (
	InvestmentPortfolioRebalanceResponseStatusAnalyzing           InvestmentPortfolioRebalanceResponseStatus = "analyzing"
	InvestmentPortfolioRebalanceResponseStatusPendingConfirmation InvestmentPortfolioRebalanceResponseStatus = "pending_confirmation"
	InvestmentPortfolioRebalanceResponseStatusInProgress          InvestmentPortfolioRebalanceResponseStatus = "in_progress"
	InvestmentPortfolioRebalanceResponseStatusCompleted           InvestmentPortfolioRebalanceResponseStatus = "completed"
	InvestmentPortfolioRebalanceResponseStatusFailed              InvestmentPortfolioRebalanceResponseStatus = "failed"
)

func (r InvestmentPortfolioRebalanceResponseStatus) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRebalanceResponseStatusAnalyzing, InvestmentPortfolioRebalanceResponseStatusPendingConfirmation, InvestmentPortfolioRebalanceResponseStatusInProgress, InvestmentPortfolioRebalanceResponseStatusCompleted, InvestmentPortfolioRebalanceResponseStatusFailed:
		return true
	}
	return false
}

type InvestmentPortfolioNewParams struct {
	Currency          param.Field[string]                                    `json:"currency,required"`
	Name              param.Field[string]                                    `json:"name,required"`
	RiskTolerance     param.Field[InvestmentPortfolioNewParamsRiskTolerance] `json:"riskTolerance,required"`
	Type              param.Field[string]                                    `json:"type,required"`
	AIAutoAllocate    param.Field[bool]                                      `json:"aiAutoAllocate"`
	InitialInvestment param.Field[float64]                                   `json:"initialInvestment"`
	LinkedAccountID   param.Field[string]                                    `json:"linkedAccountId"`
}

func (r InvestmentPortfolioNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestmentPortfolioNewParamsRiskTolerance string

const (
	InvestmentPortfolioNewParamsRiskToleranceLow        InvestmentPortfolioNewParamsRiskTolerance = "low"
	InvestmentPortfolioNewParamsRiskToleranceMedium     InvestmentPortfolioNewParamsRiskTolerance = "medium"
	InvestmentPortfolioNewParamsRiskToleranceHigh       InvestmentPortfolioNewParamsRiskTolerance = "high"
	InvestmentPortfolioNewParamsRiskToleranceAggressive InvestmentPortfolioNewParamsRiskTolerance = "aggressive"
)

func (r InvestmentPortfolioNewParamsRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioNewParamsRiskToleranceLow, InvestmentPortfolioNewParamsRiskToleranceMedium, InvestmentPortfolioNewParamsRiskToleranceHigh, InvestmentPortfolioNewParamsRiskToleranceAggressive:
		return true
	}
	return false
}

type InvestmentPortfolioUpdateParams struct {
	AIRebalancingFrequency param.Field[InvestmentPortfolioUpdateParamsAIRebalancingFrequency] `json:"aiRebalancingFrequency"`
	Name                   param.Field[string]                                                `json:"name"`
	RiskTolerance          param.Field[InvestmentPortfolioUpdateParamsRiskTolerance]          `json:"riskTolerance"`
}

func (r InvestmentPortfolioUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestmentPortfolioUpdateParamsAIRebalancingFrequency string

const (
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencyMonthly   InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "monthly"
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencyQuarterly InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "quarterly"
	InvestmentPortfolioUpdateParamsAIRebalancingFrequencyYearly    InvestmentPortfolioUpdateParamsAIRebalancingFrequency = "yearly"
)

func (r InvestmentPortfolioUpdateParamsAIRebalancingFrequency) IsKnown() bool {
	switch r {
	case InvestmentPortfolioUpdateParamsAIRebalancingFrequencyMonthly, InvestmentPortfolioUpdateParamsAIRebalancingFrequencyQuarterly, InvestmentPortfolioUpdateParamsAIRebalancingFrequencyYearly:
		return true
	}
	return false
}

type InvestmentPortfolioUpdateParamsRiskTolerance string

const (
	InvestmentPortfolioUpdateParamsRiskToleranceLow        InvestmentPortfolioUpdateParamsRiskTolerance = "low"
	InvestmentPortfolioUpdateParamsRiskToleranceMedium     InvestmentPortfolioUpdateParamsRiskTolerance = "medium"
	InvestmentPortfolioUpdateParamsRiskToleranceHigh       InvestmentPortfolioUpdateParamsRiskTolerance = "high"
	InvestmentPortfolioUpdateParamsRiskToleranceAggressive InvestmentPortfolioUpdateParamsRiskTolerance = "aggressive"
)

func (r InvestmentPortfolioUpdateParamsRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioUpdateParamsRiskToleranceLow, InvestmentPortfolioUpdateParamsRiskToleranceMedium, InvestmentPortfolioUpdateParamsRiskToleranceHigh, InvestmentPortfolioUpdateParamsRiskToleranceAggressive:
		return true
	}
	return false
}

type InvestmentPortfolioListParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [InvestmentPortfolioListParams]'s query parameters as
// `url.Values`.
func (r InvestmentPortfolioListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type InvestmentPortfolioRebalanceParams struct {
	TargetRiskTolerance  param.Field[InvestmentPortfolioRebalanceParamsTargetRiskTolerance] `json:"targetRiskTolerance,required"`
	ConfirmationRequired param.Field[bool]                                                  `json:"confirmationRequired"`
	// If true, returns the proposed changes without executing them.
	DryRun param.Field[bool] `json:"dryRun"`
}

func (r InvestmentPortfolioRebalanceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestmentPortfolioRebalanceParamsTargetRiskTolerance string

const (
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceLow        InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "low"
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceMedium     InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "medium"
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceHigh       InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "high"
	InvestmentPortfolioRebalanceParamsTargetRiskToleranceAggressive InvestmentPortfolioRebalanceParamsTargetRiskTolerance = "aggressive"
)

func (r InvestmentPortfolioRebalanceParamsTargetRiskTolerance) IsKnown() bool {
	switch r {
	case InvestmentPortfolioRebalanceParamsTargetRiskToleranceLow, InvestmentPortfolioRebalanceParamsTargetRiskToleranceMedium, InvestmentPortfolioRebalanceParamsTargetRiskToleranceHigh, InvestmentPortfolioRebalanceParamsTargetRiskToleranceAggressive:
		return true
	}
	return false
}
