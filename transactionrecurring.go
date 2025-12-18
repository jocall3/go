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

// TransactionRecurringService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionRecurringService] method instead.
type TransactionRecurringService struct {
	Options []option.RequestOption
}

// NewTransactionRecurringService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionRecurringService(opts ...option.RequestOption) (r *TransactionRecurringService) {
	r = &TransactionRecurringService{}
	r.Options = opts
	return
}

// Defines a new recurring transaction pattern for future tracking and budgeting.
func (r *TransactionRecurringService) New(ctx context.Context, body TransactionRecurringNewParams, opts ...option.RequestOption) (res *RecurringTransaction, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a list of all detected or user-defined recurring transactions, useful
// for budget tracking and subscription management.
func (r *TransactionRecurringService) List(ctx context.Context, query TransactionRecurringListParams, opts ...option.RequestOption) (res *TransactionRecurringListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions/recurring"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RecurringTransaction struct {
	ID                string                        `json:"id"`
	AIConfidenceScore float64                       `json:"aiConfidenceScore"`
	Amount            float64                       `json:"amount"`
	Category          string                        `json:"category"`
	Currency          string                        `json:"currency"`
	Description       string                        `json:"description"`
	Frequency         RecurringTransactionFrequency `json:"frequency"`
	LastPaidDate      time.Time                     `json:"lastPaidDate" format:"date"`
	LinkedAccountID   string                        `json:"linkedAccountId"`
	NextDueDate       time.Time                     `json:"nextDueDate" format:"date"`
	Status            RecurringTransactionStatus    `json:"status"`
	JSON              recurringTransactionJSON      `json:"-"`
}

// recurringTransactionJSON contains the JSON metadata for the struct
// [RecurringTransaction]
type recurringTransactionJSON struct {
	ID                apijson.Field
	AIConfidenceScore apijson.Field
	Amount            apijson.Field
	Category          apijson.Field
	Currency          apijson.Field
	Description       apijson.Field
	Frequency         apijson.Field
	LastPaidDate      apijson.Field
	LinkedAccountID   apijson.Field
	NextDueDate       apijson.Field
	Status            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecurringTransaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recurringTransactionJSON) RawJSON() string {
	return r.raw
}

type RecurringTransactionFrequency string

const (
	RecurringTransactionFrequencyWeekly  RecurringTransactionFrequency = "weekly"
	RecurringTransactionFrequencyMonthly RecurringTransactionFrequency = "monthly"
	RecurringTransactionFrequencyYearly  RecurringTransactionFrequency = "yearly"
)

func (r RecurringTransactionFrequency) IsKnown() bool {
	switch r {
	case RecurringTransactionFrequencyWeekly, RecurringTransactionFrequencyMonthly, RecurringTransactionFrequencyYearly:
		return true
	}
	return false
}

type RecurringTransactionStatus string

const (
	RecurringTransactionStatusActive    RecurringTransactionStatus = "active"
	RecurringTransactionStatusPaused    RecurringTransactionStatus = "paused"
	RecurringTransactionStatusCancelled RecurringTransactionStatus = "cancelled"
)

func (r RecurringTransactionStatus) IsKnown() bool {
	switch r {
	case RecurringTransactionStatusActive, RecurringTransactionStatusPaused, RecurringTransactionStatusCancelled:
		return true
	}
	return false
}

type TransactionRecurringListResponse struct {
	Data []RecurringTransaction               `json:"data"`
	JSON transactionRecurringListResponseJSON `json:"-"`
	PaginatedList
}

// transactionRecurringListResponseJSON contains the JSON metadata for the struct
// [TransactionRecurringListResponse]
type transactionRecurringListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionRecurringListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionRecurringListResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionRecurringNewParams struct {
	Amount          param.Field[float64]                                `json:"amount,required"`
	Category        param.Field[string]                                 `json:"category,required"`
	Currency        param.Field[string]                                 `json:"currency,required"`
	Description     param.Field[string]                                 `json:"description,required"`
	Frequency       param.Field[TransactionRecurringNewParamsFrequency] `json:"frequency,required"`
	StartDate       param.Field[time.Time]                              `json:"startDate,required" format:"date"`
	LinkedAccountID param.Field[string]                                 `json:"linkedAccountId"`
}

func (r TransactionRecurringNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionRecurringNewParamsFrequency string

const (
	TransactionRecurringNewParamsFrequencyWeekly  TransactionRecurringNewParamsFrequency = "weekly"
	TransactionRecurringNewParamsFrequencyMonthly TransactionRecurringNewParamsFrequency = "monthly"
	TransactionRecurringNewParamsFrequencyYearly  TransactionRecurringNewParamsFrequency = "yearly"
)

func (r TransactionRecurringNewParamsFrequency) IsKnown() bool {
	switch r {
	case TransactionRecurringNewParamsFrequencyWeekly, TransactionRecurringNewParamsFrequencyMonthly, TransactionRecurringNewParamsFrequencyYearly:
		return true
	}
	return false
}

type TransactionRecurringListParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [TransactionRecurringListParams]'s query parameters as
// `url.Values`.
func (r TransactionRecurringListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
