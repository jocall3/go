// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// SustainabilityInvestmentService contains methods and other services that help
// with interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityInvestmentService] method instead.
type SustainabilityInvestmentService struct {
	Options []option.RequestOption
}

// NewSustainabilityInvestmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSustainabilityInvestmentService(opts ...option.RequestOption) (r *SustainabilityInvestmentService) {
	r = &SustainabilityInvestmentService{}
	r.Options = opts
	return
}

// Provides an AI-driven analysis of the Environmental, Social, and Governance
// (ESG) impact of the user's entire investment portfolio, benchmarking against
// industry standards and suggesting more sustainable alternatives.
func (r *SustainabilityInvestmentService) AnalyzeImpact(ctx context.Context, opts ...option.RequestOption) (res *SustainabilityInvestmentAnalyzeImpactResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/investments/impact"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SustainabilityInvestmentAnalyzeImpactResponse struct {
	AIRecommendations     []AIInsight                                                        `json:"aiRecommendations"`
	BenchmarkEsgScore     float64                                                            `json:"benchmarkESGScore"`
	BreakdownByEsgFactors SustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactors `json:"breakdownByESGFactors"`
	LowestEsgHoldings     []SustainabilityInvestmentAnalyzeImpactResponseLowestEsgHolding    `json:"lowestESGHoldings"`
	OverallEsgScore       float64                                                            `json:"overallESGScore"`
	PortfolioID           string                                                             `json:"portfolioId"`
	TopEsgHoldings        []SustainabilityInvestmentAnalyzeImpactResponseTopEsgHolding       `json:"topESGHoldings"`
	JSON                  sustainabilityInvestmentAnalyzeImpactResponseJSON                  `json:"-"`
}

// sustainabilityInvestmentAnalyzeImpactResponseJSON contains the JSON metadata for
// the struct [SustainabilityInvestmentAnalyzeImpactResponse]
type sustainabilityInvestmentAnalyzeImpactResponseJSON struct {
	AIRecommendations     apijson.Field
	BenchmarkEsgScore     apijson.Field
	BreakdownByEsgFactors apijson.Field
	LowestEsgHoldings     apijson.Field
	OverallEsgScore       apijson.Field
	PortfolioID           apijson.Field
	TopEsgHoldings        apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SustainabilityInvestmentAnalyzeImpactResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityInvestmentAnalyzeImpactResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactors struct {
	EnvironmentalScore float64                                                                `json:"environmentalScore"`
	GovernanceScore    float64                                                                `json:"governanceScore"`
	SocialScore        float64                                                                `json:"socialScore"`
	JSON               sustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactorsJSON `json:"-"`
}

// sustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactorsJSON contains
// the JSON metadata for the struct
// [SustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactors]
type sustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactorsJSON struct {
	EnvironmentalScore apijson.Field
	GovernanceScore    apijson.Field
	SocialScore        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityInvestmentAnalyzeImpactResponseBreakdownByEsgFactorsJSON) RawJSON() string {
	return r.raw
}

type SustainabilityInvestmentAnalyzeImpactResponseLowestEsgHolding struct {
	AssetName   string                                                            `json:"assetName"`
	AssetSymbol string                                                            `json:"assetSymbol"`
	EsgScore    float64                                                           `json:"esgScore"`
	JSON        sustainabilityInvestmentAnalyzeImpactResponseLowestEsgHoldingJSON `json:"-"`
}

// sustainabilityInvestmentAnalyzeImpactResponseLowestEsgHoldingJSON contains the
// JSON metadata for the struct
// [SustainabilityInvestmentAnalyzeImpactResponseLowestEsgHolding]
type sustainabilityInvestmentAnalyzeImpactResponseLowestEsgHoldingJSON struct {
	AssetName   apijson.Field
	AssetSymbol apijson.Field
	EsgScore    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SustainabilityInvestmentAnalyzeImpactResponseLowestEsgHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityInvestmentAnalyzeImpactResponseLowestEsgHoldingJSON) RawJSON() string {
	return r.raw
}

type SustainabilityInvestmentAnalyzeImpactResponseTopEsgHolding struct {
	AssetName   string                                                         `json:"assetName"`
	AssetSymbol string                                                         `json:"assetSymbol"`
	EsgScore    float64                                                        `json:"esgScore"`
	JSON        sustainabilityInvestmentAnalyzeImpactResponseTopEsgHoldingJSON `json:"-"`
}

// sustainabilityInvestmentAnalyzeImpactResponseTopEsgHoldingJSON contains the JSON
// metadata for the struct
// [SustainabilityInvestmentAnalyzeImpactResponseTopEsgHolding]
type sustainabilityInvestmentAnalyzeImpactResponseTopEsgHoldingJSON struct {
	AssetName   apijson.Field
	AssetSymbol apijson.Field
	EsgScore    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SustainabilityInvestmentAnalyzeImpactResponseTopEsgHolding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityInvestmentAnalyzeImpactResponseTopEsgHoldingJSON) RawJSON() string {
	return r.raw
}
