// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/apiquery"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
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
func (r *AIOracleSimulationService) Get(ctx context.Context, simulationID interface{}, opts ...option.RequestOption) (res *SimulationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("ai/oracle/simulations/%v", simulationID)
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
func (r *AIOracleSimulationService) Delete(ctx context.Context, simulationID interface{}, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("ai/oracle/simulations/%v", simulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
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
	// Timestamp when the simulation was initiated.
	CreationDate interface{} `json:"creationDate,required"`
	// Timestamp when the simulation status or results were last updated.
	LastUpdated interface{} `json:"lastUpdated,required"`
	// Unique identifier for the simulation.
	SimulationID interface{} `json:"simulationId,required"`
	// Current status of the simulation.
	Status AIOracleSimulationListResponseDataStatus `json:"status,required"`
	// A brief summary of what the simulation evaluated.
	Summary interface{} `json:"summary,required"`
	// A user-friendly title for the simulation.
	Title interface{}                            `json:"title,required"`
	JSON  aiOracleSimulationListResponseDataJSON `json:"-"`
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

// Current status of the simulation.
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
	// Maximum number of items to return in a single page.
	Limit param.Field[interface{}] `query:"limit"`
	// Number of items to skip before starting to collect the result set.
	Offset param.Field[interface{}] `query:"offset"`
}

// URLQuery serializes [AIOracleSimulationListParams]'s query parameters as
// `url.Values`.
func (r AIOracleSimulationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
