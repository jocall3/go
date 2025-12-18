// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/apiquery"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// InvestmentAssetService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestmentAssetService] method instead.
type InvestmentAssetService struct {
	Options []option.RequestOption
}

// NewInvestmentAssetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestmentAssetService(opts ...option.RequestOption) (r *InvestmentAssetService) {
	r = &InvestmentAssetService{}
	r.Options = opts
	return
}

// Searches for available investment assets (stocks, ETFs, mutual funds) and
// returns their ESG impact scores.
func (r *InvestmentAssetService) Search(ctx context.Context, query InvestmentAssetSearchParams, opts ...option.RequestOption) (res *InvestmentAssetSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "investments/assets/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type InvestmentAssetSearchResponse struct {
	Data []InvestmentAssetSearchResponseData `json:"data"`
	JSON investmentAssetSearchResponseJSON   `json:"-"`
	PaginatedList
}

// investmentAssetSearchResponseJSON contains the JSON metadata for the struct
// [InvestmentAssetSearchResponse]
type investmentAssetSearchResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestmentAssetSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentAssetSearchResponseJSON) RawJSON() string {
	return r.raw
}

type InvestmentAssetSearchResponseData struct {
	AIEsgInsight       string                                     `json:"aiESGInsight"`
	AssetName          string                                     `json:"assetName"`
	AssetSymbol        string                                     `json:"assetSymbol"`
	AssetType          InvestmentAssetSearchResponseDataAssetType `json:"assetType"`
	Currency           string                                     `json:"currency"`
	CurrentPrice       float64                                    `json:"currentPrice"`
	EnvironmentalScore float64                                    `json:"environmentalScore"`
	EsgControversies   []string                                   `json:"esgControversies"`
	EsgRatingProvider  string                                     `json:"esgRatingProvider"`
	GovernanceScore    float64                                    `json:"governanceScore"`
	OverallEsgScore    float64                                    `json:"overallESGScore"`
	SocialScore        float64                                    `json:"socialScore"`
	JSON               investmentAssetSearchResponseDataJSON      `json:"-"`
}

// investmentAssetSearchResponseDataJSON contains the JSON metadata for the struct
// [InvestmentAssetSearchResponseData]
type investmentAssetSearchResponseDataJSON struct {
	AIEsgInsight       apijson.Field
	AssetName          apijson.Field
	AssetSymbol        apijson.Field
	AssetType          apijson.Field
	Currency           apijson.Field
	CurrentPrice       apijson.Field
	EnvironmentalScore apijson.Field
	EsgControversies   apijson.Field
	EsgRatingProvider  apijson.Field
	GovernanceScore    apijson.Field
	OverallEsgScore    apijson.Field
	SocialScore        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InvestmentAssetSearchResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investmentAssetSearchResponseDataJSON) RawJSON() string {
	return r.raw
}

type InvestmentAssetSearchResponseDataAssetType string

const (
	InvestmentAssetSearchResponseDataAssetTypeStock      InvestmentAssetSearchResponseDataAssetType = "stock"
	InvestmentAssetSearchResponseDataAssetTypeEtf        InvestmentAssetSearchResponseDataAssetType = "etf"
	InvestmentAssetSearchResponseDataAssetTypeMutualFund InvestmentAssetSearchResponseDataAssetType = "mutual_fund"
)

func (r InvestmentAssetSearchResponseDataAssetType) IsKnown() bool {
	switch r {
	case InvestmentAssetSearchResponseDataAssetTypeStock, InvestmentAssetSearchResponseDataAssetTypeEtf, InvestmentAssetSearchResponseDataAssetTypeMutualFund:
		return true
	}
	return false
}

type InvestmentAssetSearchParams struct {
	// Search query for asset name or symbol.
	Query param.Field[string] `query:"query,required"`
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// Minimum desired ESG score (0-10).
	MinEsgScore param.Field[float64] `query:"minESGScore"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [InvestmentAssetSearchParams]'s query parameters as
// `url.Values`.
func (r InvestmentAssetSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
