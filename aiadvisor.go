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

// AIAdvisorService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdvisorService] method instead.
type AIAdvisorService struct {
	Options []option.RequestOption
	Chat    *AIAdvisorChatService
}

// NewAIAdvisorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdvisorService(opts ...option.RequestOption) (r *AIAdvisorService) {
	r = &AIAdvisorService{}
	r.Options = opts
	r.Chat = NewAIAdvisorChatService(opts...)
	return
}

// Retrieves a dynamic manifest of all integrated AI tools that Quantum can invoke
// and execute, providing details on their capabilities, parameters, and access
// requirements.
func (r *AIAdvisorService) ListTools(ctx context.Context, query AIAdvisorListToolsParams, opts ...option.RequestOption) (res *AIAdvisorListToolsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AIAdvisorListToolsResponse struct {
	Data []AIAdvisorListToolsResponseData `json:"data"`
	JSON aiAdvisorListToolsResponseJSON   `json:"-"`
	PaginatedList
}

// aiAdvisorListToolsResponseJSON contains the JSON metadata for the struct
// [AIAdvisorListToolsResponse]
type aiAdvisorListToolsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorListToolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorListToolsResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorListToolsResponseData struct {
	AccessScope string                             `json:"accessScope"`
	Description string                             `json:"description"`
	Name        string                             `json:"name"`
	Parameters  map[string]interface{}             `json:"parameters"`
	JSON        aiAdvisorListToolsResponseDataJSON `json:"-"`
}

// aiAdvisorListToolsResponseDataJSON contains the JSON metadata for the struct
// [AIAdvisorListToolsResponseData]
type aiAdvisorListToolsResponseDataJSON struct {
	AccessScope apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Parameters  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorListToolsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorListToolsResponseDataJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorListToolsParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AIAdvisorListToolsParams]'s query parameters as
// `url.Values`.
func (r AIAdvisorListToolsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
