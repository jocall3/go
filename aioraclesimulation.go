// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/apiquery"
	"github.com/stainless-sdks/1231-go/internal/param"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
	"github.com/tidwall/gjson"
)

// AIOracleSimulationService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOracleSimulationService] method instead.
type AIOracleSimulationService struct {
	Options []option.RequestOption
}

// NewAIOracleSimulationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOracleSimulationService(opts ...option.RequestOption) (r *AIOracleSimulationService) {
	r = &AIOracleSimulationService{}
	r.Options = opts
	return
}

// Retrieves the full, detailed results of a specific financial simulation by its
// ID.
func (r *AIOracleSimulationService) Get(ctx context.Context, simulationID string, opts ...option.RequestOption) (res *AIOracleSimulationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if simulationID == "" {
		err = errors.New("missing required simulationId parameter")
		return
	}
	path := fmt.Sprintf("ai/oracle/simulations/%s", simulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a list of all financial simulations previously run by the user,
// including their status and summaries.
func (r *AIOracleSimulationService) List(ctx context.Context, query AIOracleSimulationListParams, opts ...option.RequestOption) (res *AIOracleSimulationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a previously run financial simulation and its results.
func (r *AIOracleSimulationService) Delete(ctx context.Context, simulationID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if simulationID == "" {
		err = errors.New("missing required simulationId parameter")
		return
	}
	path := fmt.Sprintf("ai/oracle/simulations/%s", simulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AIOracleSimulationGetResponse struct {
	// This field can have the runtime type of [[]SimulationResponseKeyImpact].
	KeyImpacts       interface{} `json:"keyImpacts"`
	NarrativeSummary string      `json:"narrativeSummary"`
	OverallSummary   string      `json:"overallSummary"`
	// This field can have the runtime type of [[]SimulationResponseRecommendation].
	Recommendations interface{} `json:"recommendations"`
	// This field can have the runtime type of [SimulationResponseRiskAnalysis].
	RiskAnalysis interface{} `json:"riskAnalysis"`
	// This field can have the runtime type of
	// [[]AdvancedSimulationResponseScenarioResult].
	ScenarioResults interface{} `json:"scenarioResults"`
	SimulationID    string      `json:"simulationId"`
	// This field can have the runtime type of [[]AIInsight].
	StrategicRecommendations interface{}                       `json:"strategicRecommendations"`
	JSON                     aiOracleSimulationGetResponseJSON `json:"-"`
	union                    AIOracleSimulationGetResponseUnion
}

// aiOracleSimulationGetResponseJSON contains the JSON metadata for the struct
// [AIOracleSimulationGetResponse]
type aiOracleSimulationGetResponseJSON struct {
	KeyImpacts               apijson.Field
	NarrativeSummary         apijson.Field
	OverallSummary           apijson.Field
	Recommendations          apijson.Field
	RiskAnalysis             apijson.Field
	ScenarioResults          apijson.Field
	SimulationID             apijson.Field
	StrategicRecommendations apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r aiOracleSimulationGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AIOracleSimulationGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AIOracleSimulationGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AIOracleSimulationGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SimulationResponse],
// [AdvancedSimulationResponse].
func (r AIOracleSimulationGetResponse) AsUnion() AIOracleSimulationGetResponseUnion {
	return r.union
}

// Union satisfied by [SimulationResponse] or [AdvancedSimulationResponse].
type AIOracleSimulationGetResponseUnion interface {
	implementsAIOracleSimulationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIOracleSimulationGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SimulationResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AdvancedSimulationResponse{}),
		},
	)
}

type AIOracleSimulationListResponse struct {
	Data []AIOracleSimulationListResponseData `json:"data"`
	JSON aiOracleSimulationListResponseJSON   `json:"-"`
	PaginatedList
}

// aiOracleSimulationListResponseJSON contains the JSON metadata for the struct
// [AIOracleSimulationListResponse]
type aiOracleSimulationListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIOracleSimulationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOracleSimulationListResponseJSON) RawJSON() string {
	return r.raw
}

type AIOracleSimulationListResponseData struct {
	CreationDate time.Time                                `json:"creationDate" format:"date-time"`
	LastUpdated  time.Time                                `json:"lastUpdated" format:"date-time"`
	SimulationID string                                   `json:"simulationId"`
	Status       AIOracleSimulationListResponseDataStatus `json:"status"`
	Summary      string                                   `json:"summary"`
	Title        string                                   `json:"title"`
	JSON         aiOracleSimulationListResponseDataJSON   `json:"-"`
}

// aiOracleSimulationListResponseDataJSON contains the JSON metadata for the struct
// [AIOracleSimulationListResponseData]
type aiOracleSimulationListResponseDataJSON struct {
	CreationDate apijson.Field
	LastUpdated  apijson.Field
	SimulationID apijson.Field
	Status       apijson.Field
	Summary      apijson.Field
	Title        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AIOracleSimulationListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiOracleSimulationListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AIOracleSimulationListResponseDataStatus string

const (
	AIOracleSimulationListResponseDataStatusProcessing AIOracleSimulationListResponseDataStatus = "processing"
	AIOracleSimulationListResponseDataStatusCompleted  AIOracleSimulationListResponseDataStatus = "completed"
	AIOracleSimulationListResponseDataStatusFailed     AIOracleSimulationListResponseDataStatus = "failed"
)

func (r AIOracleSimulationListResponseDataStatus) IsKnown() bool {
	switch r {
	case AIOracleSimulationListResponseDataStatusProcessing, AIOracleSimulationListResponseDataStatusCompleted, AIOracleSimulationListResponseDataStatusFailed:
		return true
	}
	return false
}

type AIOracleSimulationListParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AIOracleSimulationListParams]'s query parameters as
// `url.Values`.
func (r AIOracleSimulationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
