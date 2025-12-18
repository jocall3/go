// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/param"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// SustainabilityService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityService] method instead.
type SustainabilityService struct {
	Options     []option.RequestOption
	Investments *SustainabilityInvestmentService
}

// NewSustainabilityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSustainabilityService(opts ...option.RequestOption) (r *SustainabilityService) {
	r = &SustainabilityService{}
	r.Options = opts
	r.Investments = NewSustainabilityInvestmentService(opts...)
	return
}

// Allows users to purchase carbon offsets to neutralize their estimated carbon
// footprint, supporting environmental initiatives.
func (r *SustainabilityService) PurchaseCarbonOffsets(ctx context.Context, body SustainabilityPurchaseCarbonOffsetsParams, opts ...option.RequestOption) (res *SustainabilityPurchaseCarbonOffsetsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/carbon-offsets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generates a detailed report of the user's estimated carbon footprint based on
// transaction data, lifestyle choices, and AI-driven impact assessments, offering
// insights and reduction strategies.
func (r *SustainabilityService) GetCarbonFootprint(ctx context.Context, opts ...option.RequestOption) (res *SustainabilityGetCarbonFootprintResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/carbon-footprint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SustainabilityPurchaseCarbonOffsetsResponse struct {
	AmountOffsetKgCo2e float64                                         `json:"amountOffsetKgCO2e"`
	CertificateURL     string                                          `json:"certificateUrl" format:"uri"`
	ProjectSupported   string                                          `json:"projectSupported"`
	PurchaseDate       time.Time                                       `json:"purchaseDate" format:"date-time"`
	PurchaseID         string                                          `json:"purchaseId"`
	TotalCostUsd       float64                                         `json:"totalCostUSD"`
	TransactionID      string                                          `json:"transactionId"`
	JSON               sustainabilityPurchaseCarbonOffsetsResponseJSON `json:"-"`
}

// sustainabilityPurchaseCarbonOffsetsResponseJSON contains the JSON metadata for
// the struct [SustainabilityPurchaseCarbonOffsetsResponse]
type sustainabilityPurchaseCarbonOffsetsResponseJSON struct {
	AmountOffsetKgCo2e apijson.Field
	CertificateURL     apijson.Field
	ProjectSupported   apijson.Field
	PurchaseDate       apijson.Field
	PurchaseID         apijson.Field
	TotalCostUsd       apijson.Field
	TransactionID      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SustainabilityPurchaseCarbonOffsetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityPurchaseCarbonOffsetsResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponse struct {
	AIInsights                 []AIInsight                                                    `json:"aiInsights"`
	BreakdownByCategory        []SustainabilityGetCarbonFootprintResponseBreakdownByCategory  `json:"breakdownByCategory"`
	OffsetRecommendations      []SustainabilityGetCarbonFootprintResponseOffsetRecommendation `json:"offsetRecommendations"`
	Period                     string                                                         `json:"period"`
	ReportID                   string                                                         `json:"reportId"`
	TotalCarbonFootprintKgCo2e float64                                                        `json:"totalCarbonFootprintKgCO2e"`
	JSON                       sustainabilityGetCarbonFootprintResponseJSON                   `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseJSON contains the JSON metadata for the
// struct [SustainabilityGetCarbonFootprintResponse]
type sustainabilityGetCarbonFootprintResponseJSON struct {
	AIInsights                 apijson.Field
	BreakdownByCategory        apijson.Field
	OffsetRecommendations      apijson.Field
	Period                     apijson.Field
	ReportID                   apijson.Field
	TotalCarbonFootprintKgCo2e apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponseBreakdownByCategory struct {
	CarbonFootprintKgCo2e float64                                                         `json:"carbonFootprintKgCO2e"`
	Category              string                                                          `json:"category"`
	Percentage            float64                                                         `json:"percentage"`
	JSON                  sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON contains the
// JSON metadata for the struct
// [SustainabilityGetCarbonFootprintResponseBreakdownByCategory]
type sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON struct {
	CarbonFootprintKgCo2e apijson.Field
	Category              apijson.Field
	Percentage            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponseBreakdownByCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponseOffsetRecommendation struct {
	CostPerTonUsd      float64                                                          `json:"costPerTonUSD"`
	OffsetAmountKgCo2e float64                                                          `json:"offsetAmountKgCO2e"`
	Project            string                                                           `json:"project"`
	TotalCostUsd       float64                                                          `json:"totalCostUSD"`
	JSON               sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON contains the
// JSON metadata for the struct
// [SustainabilityGetCarbonFootprintResponseOffsetRecommendation]
type sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON struct {
	CostPerTonUsd      apijson.Field
	OffsetAmountKgCo2e apijson.Field
	Project            apijson.Field
	TotalCostUsd       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponseOffsetRecommendation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseOffsetRecommendationJSON) RawJSON() string {
	return r.raw
}

type SustainabilityPurchaseCarbonOffsetsParams struct {
	AmountKgCo2e     param.Field[float64] `json:"amountKgCO2e,required"`
	PaymentAccountID param.Field[string]  `json:"paymentAccountId,required"`
	OffsetProject    param.Field[string]  `json:"offsetProject"`
}

func (r SustainabilityPurchaseCarbonOffsetsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
