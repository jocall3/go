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

// CorporateAnomalyService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateAnomalyService] method instead.
type CorporateAnomalyService struct {
	Options []option.RequestOption
}

// NewCorporateAnomalyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateAnomalyService(opts ...option.RequestOption) (r *CorporateAnomalyService) {
	r = &CorporateAnomalyService{}
	r.Options = opts
	return
}

// Retrieves a comprehensive list of AI-detected financial anomalies across
// transactions, payments, and corporate cards that require immediate review and
// potential action to mitigate risk and ensure compliance.
func (r *CorporateAnomalyService) List(ctx context.Context, query CorporateAnomalyListParams, opts ...option.RequestOption) (res *CorporateAnomalyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the review status of a specific financial anomaly, allowing compliance
// officers to mark it as dismissed, resolved, or escalate for further
// investigation after thorough AI-assisted and human review.
func (r *CorporateAnomalyService) UpdateStatus(ctx context.Context, anomalyID string, body CorporateAnomalyUpdateStatusParams, opts ...option.RequestOption) (res *FinancialAnomaly, err error) {
	opts = slices.Concat(r.Options, opts)
	if anomalyID == "" {
		err = errors.New("missing required anomalyId parameter")
		return
	}
	path := fmt.Sprintf("corporate/anomalies/%s/status", anomalyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FinancialAnomaly struct {
	ID                  string                   `json:"id"`
	AIConfidenceScore   float64                  `json:"aiConfidenceScore"`
	Description         string                   `json:"description"`
	Details             string                   `json:"details"`
	EntityID            string                   `json:"entityId"`
	EntityType          string                   `json:"entityType"`
	RecommendedAction   string                   `json:"recommendedAction"`
	RelatedTransactions []string                 `json:"relatedTransactions"`
	ResolutionNotes     string                   `json:"resolutionNotes"`
	RiskScore           int64                    `json:"riskScore"`
	Severity            FinancialAnomalySeverity `json:"severity"`
	Status              FinancialAnomalyStatus   `json:"status"`
	Timestamp           time.Time                `json:"timestamp" format:"date-time"`
	JSON                financialAnomalyJSON     `json:"-"`
}

// financialAnomalyJSON contains the JSON metadata for the struct
// [FinancialAnomaly]
type financialAnomalyJSON struct {
	ID                  apijson.Field
	AIConfidenceScore   apijson.Field
	Description         apijson.Field
	Details             apijson.Field
	EntityID            apijson.Field
	EntityType          apijson.Field
	RecommendedAction   apijson.Field
	RelatedTransactions apijson.Field
	ResolutionNotes     apijson.Field
	RiskScore           apijson.Field
	Severity            apijson.Field
	Status              apijson.Field
	Timestamp           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *FinancialAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialAnomalyJSON) RawJSON() string {
	return r.raw
}

type FinancialAnomalySeverity string

const (
	FinancialAnomalySeverityLow      FinancialAnomalySeverity = "Low"
	FinancialAnomalySeverityMedium   FinancialAnomalySeverity = "Medium"
	FinancialAnomalySeverityHigh     FinancialAnomalySeverity = "High"
	FinancialAnomalySeverityCritical FinancialAnomalySeverity = "Critical"
)

func (r FinancialAnomalySeverity) IsKnown() bool {
	switch r {
	case FinancialAnomalySeverityLow, FinancialAnomalySeverityMedium, FinancialAnomalySeverityHigh, FinancialAnomalySeverityCritical:
		return true
	}
	return false
}

type FinancialAnomalyStatus string

const (
	FinancialAnomalyStatusNew         FinancialAnomalyStatus = "New"
	FinancialAnomalyStatusUnderReview FinancialAnomalyStatus = "Under Review"
	FinancialAnomalyStatusEscalated   FinancialAnomalyStatus = "Escalated"
	FinancialAnomalyStatusDismissed   FinancialAnomalyStatus = "Dismissed"
	FinancialAnomalyStatusResolved    FinancialAnomalyStatus = "Resolved"
)

func (r FinancialAnomalyStatus) IsKnown() bool {
	switch r {
	case FinancialAnomalyStatusNew, FinancialAnomalyStatusUnderReview, FinancialAnomalyStatusEscalated, FinancialAnomalyStatusDismissed, FinancialAnomalyStatusResolved:
		return true
	}
	return false
}

type CorporateAnomalyListResponse struct {
	Data []FinancialAnomaly               `json:"data"`
	JSON corporateAnomalyListResponseJSON `json:"-"`
	PaginatedList
}

// corporateAnomalyListResponseJSON contains the JSON metadata for the struct
// [CorporateAnomalyListResponse]
type corporateAnomalyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateAnomalyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateAnomalyListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateAnomalyListParams struct {
	// The end date for the query range (inclusive).
	EndDate param.Field[time.Time] `query:"endDate" format:"date"`
	// Filter anomalies by the type of financial entity they are related to.
	EntityType param.Field[CorporateAnomalyListParamsEntityType] `query:"entityType"`
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// Filter anomalies by their AI-assessed severity level.
	Severity param.Field[CorporateAnomalyListParamsSeverity] `query:"severity"`
	// The start date for the query range (inclusive).
	StartDate param.Field[time.Time] `query:"startDate" format:"date"`
	// Filter anomalies by their current review status.
	Status param.Field[CorporateAnomalyListParamsStatus] `query:"status"`
}

// URLQuery serializes [CorporateAnomalyListParams]'s query parameters as
// `url.Values`.
func (r CorporateAnomalyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter anomalies by the type of financial entity they are related to.
type CorporateAnomalyListParamsEntityType string

const (
	CorporateAnomalyListParamsEntityTypePaymentOrder  CorporateAnomalyListParamsEntityType = "PaymentOrder"
	CorporateAnomalyListParamsEntityTypeTransaction   CorporateAnomalyListParamsEntityType = "Transaction"
	CorporateAnomalyListParamsEntityTypeCounterparty  CorporateAnomalyListParamsEntityType = "Counterparty"
	CorporateAnomalyListParamsEntityTypeCorporateCard CorporateAnomalyListParamsEntityType = "CorporateCard"
	CorporateAnomalyListParamsEntityTypeInvoice       CorporateAnomalyListParamsEntityType = "Invoice"
)

func (r CorporateAnomalyListParamsEntityType) IsKnown() bool {
	switch r {
	case CorporateAnomalyListParamsEntityTypePaymentOrder, CorporateAnomalyListParamsEntityTypeTransaction, CorporateAnomalyListParamsEntityTypeCounterparty, CorporateAnomalyListParamsEntityTypeCorporateCard, CorporateAnomalyListParamsEntityTypeInvoice:
		return true
	}
	return false
}

// Filter anomalies by their AI-assessed severity level.
type CorporateAnomalyListParamsSeverity string

const (
	CorporateAnomalyListParamsSeverityLow      CorporateAnomalyListParamsSeverity = "Low"
	CorporateAnomalyListParamsSeverityMedium   CorporateAnomalyListParamsSeverity = "Medium"
	CorporateAnomalyListParamsSeverityHigh     CorporateAnomalyListParamsSeverity = "High"
	CorporateAnomalyListParamsSeverityCritical CorporateAnomalyListParamsSeverity = "Critical"
)

func (r CorporateAnomalyListParamsSeverity) IsKnown() bool {
	switch r {
	case CorporateAnomalyListParamsSeverityLow, CorporateAnomalyListParamsSeverityMedium, CorporateAnomalyListParamsSeverityHigh, CorporateAnomalyListParamsSeverityCritical:
		return true
	}
	return false
}

// Filter anomalies by their current review status.
type CorporateAnomalyListParamsStatus string

const (
	CorporateAnomalyListParamsStatusNew         CorporateAnomalyListParamsStatus = "New"
	CorporateAnomalyListParamsStatusUnderReview CorporateAnomalyListParamsStatus = "Under Review"
	CorporateAnomalyListParamsStatusEscalated   CorporateAnomalyListParamsStatus = "Escalated"
	CorporateAnomalyListParamsStatusDismissed   CorporateAnomalyListParamsStatus = "Dismissed"
	CorporateAnomalyListParamsStatusResolved    CorporateAnomalyListParamsStatus = "Resolved"
)

func (r CorporateAnomalyListParamsStatus) IsKnown() bool {
	switch r {
	case CorporateAnomalyListParamsStatusNew, CorporateAnomalyListParamsStatusUnderReview, CorporateAnomalyListParamsStatusEscalated, CorporateAnomalyListParamsStatusDismissed, CorporateAnomalyListParamsStatusResolved:
		return true
	}
	return false
}

type CorporateAnomalyUpdateStatusParams struct {
	Status          param.Field[CorporateAnomalyUpdateStatusParamsStatus] `json:"status,required"`
	ResolutionNotes param.Field[string]                                   `json:"resolutionNotes"`
}

func (r CorporateAnomalyUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateAnomalyUpdateStatusParamsStatus string

const (
	CorporateAnomalyUpdateStatusParamsStatusUnderReview CorporateAnomalyUpdateStatusParamsStatus = "Under Review"
	CorporateAnomalyUpdateStatusParamsStatusEscalated   CorporateAnomalyUpdateStatusParamsStatus = "Escalated"
	CorporateAnomalyUpdateStatusParamsStatusDismissed   CorporateAnomalyUpdateStatusParamsStatus = "Dismissed"
	CorporateAnomalyUpdateStatusParamsStatusResolved    CorporateAnomalyUpdateStatusParamsStatus = "Resolved"
)

func (r CorporateAnomalyUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case CorporateAnomalyUpdateStatusParamsStatusUnderReview, CorporateAnomalyUpdateStatusParamsStatusEscalated, CorporateAnomalyUpdateStatusParamsStatusDismissed, CorporateAnomalyUpdateStatusParamsStatusResolved:
		return true
	}
	return false
}
