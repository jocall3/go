// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/apiquery"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// AIAdvisorChatService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAdvisorChatService] method instead.
type AIAdvisorChatService struct {
	Options []option.RequestOption
}

// NewAIAdvisorChatService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIAdvisorChatService(opts ...option.RequestOption) (r *AIAdvisorChatService) {
	r = &AIAdvisorChatService{}
	r.Options = opts
	return
}

// Fetches the full conversation history with the Quantum AI Advisor for a given
// session or user.
func (r *AIAdvisorChatService) GetHistory(ctx context.Context, query AIAdvisorChatGetHistoryParams, opts ...option.RequestOption) (res *AIAdvisorChatGetHistoryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/chat/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Initiates or continues a sophisticated conversation with Quantum, the AI
// Advisor. Quantum can provide advanced financial insights, execute complex tasks
// via an expanding suite of intelligent tools, and learn from user interactions to
// offer hyper-personalized guidance.
func (r *AIAdvisorChatService) SendMessage(ctx context.Context, body AIAdvisorChatSendMessageParams, opts ...option.RequestOption) (res *AIAdvisorChatSendMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/advisor/chat"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIAdvisorChatGetHistoryResponse struct {
	Data []AIAdvisorChatGetHistoryResponseData `json:"data"`
	JSON aiAdvisorChatGetHistoryResponseJSON   `json:"-"`
	PaginatedList
}

// aiAdvisorChatGetHistoryResponseJSON contains the JSON metadata for the struct
// [AIAdvisorChatGetHistoryResponse]
type aiAdvisorChatGetHistoryResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorChatGetHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatGetHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatGetHistoryResponseData struct {
	Content   string                                  `json:"content"`
	Role      AIAdvisorChatGetHistoryResponseDataRole `json:"role"`
	Timestamp time.Time                               `json:"timestamp" format:"date-time"`
	JSON      aiAdvisorChatGetHistoryResponseDataJSON `json:"-"`
}

// aiAdvisorChatGetHistoryResponseDataJSON contains the JSON metadata for the
// struct [AIAdvisorChatGetHistoryResponseData]
type aiAdvisorChatGetHistoryResponseDataJSON struct {
	Content     apijson.Field
	Role        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorChatGetHistoryResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatGetHistoryResponseDataJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatGetHistoryResponseDataRole string

const (
	AIAdvisorChatGetHistoryResponseDataRoleUser      AIAdvisorChatGetHistoryResponseDataRole = "user"
	AIAdvisorChatGetHistoryResponseDataRoleAssistant AIAdvisorChatGetHistoryResponseDataRole = "assistant"
)

func (r AIAdvisorChatGetHistoryResponseDataRole) IsKnown() bool {
	switch r {
	case AIAdvisorChatGetHistoryResponseDataRoleUser, AIAdvisorChatGetHistoryResponseDataRoleAssistant:
		return true
	}
	return false
}

type AIAdvisorChatSendMessageResponse struct {
	FunctionCalls     []AIAdvisorChatSendMessageResponseFunctionCall `json:"functionCalls"`
	ProactiveInsights []AIInsight                                    `json:"proactiveInsights"`
	SessionID         string                                         `json:"sessionId"`
	Text              string                                         `json:"text"`
	JSON              aiAdvisorChatSendMessageResponseJSON           `json:"-"`
}

// aiAdvisorChatSendMessageResponseJSON contains the JSON metadata for the struct
// [AIAdvisorChatSendMessageResponse]
type aiAdvisorChatSendMessageResponseJSON struct {
	FunctionCalls     apijson.Field
	ProactiveInsights apijson.Field
	SessionID         apijson.Field
	Text              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AIAdvisorChatSendMessageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatSendMessageResponseJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatSendMessageResponseFunctionCall struct {
	ID   string                                           `json:"id"`
	Args map[string]interface{}                           `json:"args"`
	Name string                                           `json:"name"`
	JSON aiAdvisorChatSendMessageResponseFunctionCallJSON `json:"-"`
}

// aiAdvisorChatSendMessageResponseFunctionCallJSON contains the JSON metadata for
// the struct [AIAdvisorChatSendMessageResponseFunctionCall]
type aiAdvisorChatSendMessageResponseFunctionCallJSON struct {
	ID          apijson.Field
	Args        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIAdvisorChatSendMessageResponseFunctionCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiAdvisorChatSendMessageResponseFunctionCallJSON) RawJSON() string {
	return r.raw
}

type AIAdvisorChatGetHistoryParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// Optional: Filter history by a specific session ID. If omitted, recent
	// conversations will be returned.
	SessionID param.Field[string] `query:"sessionId"`
}

// URLQuery serializes [AIAdvisorChatGetHistoryParams]'s query parameters as
// `url.Values`.
func (r AIAdvisorChatGetHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAdvisorChatSendMessageParams struct {
	FunctionResponse param.Field[AIAdvisorChatSendMessageParamsFunctionResponse] `json:"functionResponse"`
	Message          param.Field[string]                                         `json:"message"`
	SessionID        param.Field[string]                                         `json:"sessionId"`
}

func (r AIAdvisorChatSendMessageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIAdvisorChatSendMessageParamsFunctionResponse struct {
	Name     param.Field[string]                 `json:"name"`
	Response param.Field[map[string]interface{}] `json:"response"`
}

func (r AIAdvisorChatSendMessageParamsFunctionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
