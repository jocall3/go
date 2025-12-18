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

// CorporateService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateService] method instead.
type CorporateService struct {
	Options    []option.RequestOption
	Cards      *CorporateCardService
	Anomalies  *CorporateAnomalyService
	Compliance *CorporateComplianceService
	Treasury   *CorporateTreasuryService
	Risk       *CorporateRiskService
}

// NewCorporateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorporateService(opts ...option.RequestOption) (r *CorporateService) {
	r = &CorporateService{}
	r.Options = opts
	r.Cards = NewCorporateCardService(opts...)
	r.Anomalies = NewCorporateAnomalyService(opts...)
	r.Compliance = NewCorporateComplianceService(opts...)
	r.Treasury = NewCorporateTreasuryService(opts...)
	r.Risk = NewCorporateRiskService(opts...)
	return
}

// Executes a real-time screening of an individual or entity against global
// sanction lists and watchlists.
func (r *CorporateService) PerformSanctionScreening(ctx context.Context, body CorporatePerformSanctionScreeningParams, opts ...option.RequestOption) (res *CorporatePerformSanctionScreeningResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/sanction-screening"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CorporatePerformSanctionScreeningResponse struct {
	MatchDetails       []CorporatePerformSanctionScreeningResponseMatchDetail `json:"matchDetails"`
	MatchFound         bool                                                   `json:"matchFound"`
	ScreeningID        string                                                 `json:"screeningId"`
	ScreeningTimestamp time.Time                                              `json:"screeningTimestamp" format:"date-time"`
	Status             CorporatePerformSanctionScreeningResponseStatus        `json:"status"`
	JSON               corporatePerformSanctionScreeningResponseJSON          `json:"-"`
}

// corporatePerformSanctionScreeningResponseJSON contains the JSON metadata for the
// struct [CorporatePerformSanctionScreeningResponse]
type corporatePerformSanctionScreeningResponseJSON struct {
	MatchDetails       apijson.Field
	MatchFound         apijson.Field
	ScreeningID        apijson.Field
	ScreeningTimestamp apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CorporatePerformSanctionScreeningResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporatePerformSanctionScreeningResponseJSON) RawJSON() string {
	return r.raw
}

type CorporatePerformSanctionScreeningResponseMatchDetail struct {
	ListName    string                                                   `json:"listName"`
	MatchedName string                                                   `json:"matchedName"`
	Reason      string                                                   `json:"reason"`
	Score       float64                                                  `json:"score"`
	JSON        corporatePerformSanctionScreeningResponseMatchDetailJSON `json:"-"`
}

// corporatePerformSanctionScreeningResponseMatchDetailJSON contains the JSON
// metadata for the struct [CorporatePerformSanctionScreeningResponseMatchDetail]
type corporatePerformSanctionScreeningResponseMatchDetailJSON struct {
	ListName    apijson.Field
	MatchedName apijson.Field
	Reason      apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporatePerformSanctionScreeningResponseMatchDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporatePerformSanctionScreeningResponseMatchDetailJSON) RawJSON() string {
	return r.raw
}

type CorporatePerformSanctionScreeningResponseStatus string

const (
	CorporatePerformSanctionScreeningResponseStatusClear          CorporatePerformSanctionScreeningResponseStatus = "clear"
	CorporatePerformSanctionScreeningResponseStatusPotentialMatch CorporatePerformSanctionScreeningResponseStatus = "potential_match"
	CorporatePerformSanctionScreeningResponseStatusConfirmedMatch CorporatePerformSanctionScreeningResponseStatus = "confirmed_match"
)

func (r CorporatePerformSanctionScreeningResponseStatus) IsKnown() bool {
	switch r {
	case CorporatePerformSanctionScreeningResponseStatusClear, CorporatePerformSanctionScreeningResponseStatusPotentialMatch, CorporatePerformSanctionScreeningResponseStatusConfirmedMatch:
		return true
	}
	return false
}

type CorporatePerformSanctionScreeningParams struct {
	EntityType  param.Field[CorporatePerformSanctionScreeningParamsEntityType] `json:"entityType,required"`
	Name        param.Field[string]                                            `json:"name,required"`
	Country     param.Field[string]                                            `json:"country"`
	DateOfBirth param.Field[time.Time]                                         `json:"dateOfBirth" format:"date"`
}

func (r CorporatePerformSanctionScreeningParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporatePerformSanctionScreeningParamsEntityType string

const (
	CorporatePerformSanctionScreeningParamsEntityTypeIndividual   CorporatePerformSanctionScreeningParamsEntityType = "individual"
	CorporatePerformSanctionScreeningParamsEntityTypeOrganization CorporatePerformSanctionScreeningParamsEntityType = "organization"
)

func (r CorporatePerformSanctionScreeningParamsEntityType) IsKnown() bool {
	switch r {
	case CorporatePerformSanctionScreeningParamsEntityTypeIndividual, CorporatePerformSanctionScreeningParamsEntityTypeOrganization:
		return true
	}
	return false
}
