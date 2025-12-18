// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
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

// Retrieves an analysis of the user's personal carbon footprint based on their
// transaction history.
func (r *SustainabilityService) GetCarbonFootprint(ctx context.Context, opts ...option.RequestOption) (res *SustainabilityGetCarbonFootprintResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sustainability/carbon-footprint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SustainabilityGetCarbonFootprintResponse struct {
	AIRecommendations    []AIInsight                                                   `json:"aiRecommendations"`
	BreakdownByCategory  []SustainabilityGetCarbonFootprintResponseBreakdownByCategory `json:"breakdownByCategory"`
	Period               string                                                        `json:"period"`
	TotalFootprintKgCo2e float64                                                       `json:"totalFootprintKgCO2e"`
	JSON                 sustainabilityGetCarbonFootprintResponseJSON                  `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseJSON contains the JSON metadata for the
// struct [SustainabilityGetCarbonFootprintResponse]
type sustainabilityGetCarbonFootprintResponseJSON struct {
	AIRecommendations    apijson.Field
	BreakdownByCategory  apijson.Field
	Period               apijson.Field
	TotalFootprintKgCo2e apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseJSON) RawJSON() string {
	return r.raw
}

type SustainabilityGetCarbonFootprintResponseBreakdownByCategory struct {
	Category        string                                                          `json:"category"`
	FootprintKgCo2e float64                                                         `json:"footprintKgCO2e"`
	JSON            sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON `json:"-"`
}

// sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON contains the
// JSON metadata for the struct
// [SustainabilityGetCarbonFootprintResponseBreakdownByCategory]
type sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON struct {
	Category        apijson.Field
	FootprintKgCo2e apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SustainabilityGetCarbonFootprintResponseBreakdownByCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sustainabilityGetCarbonFootprintResponseBreakdownByCategoryJSON) RawJSON() string {
	return r.raw
}
