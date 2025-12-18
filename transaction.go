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

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/apiquery"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// TransactionService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options   []option.RequestOption
	Recurring *TransactionRecurringService
	Insights  *TransactionInsightService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	r.Recurring = NewTransactionRecurringService(opts...)
	r.Insights = NewTransactionInsightService(opts...)
	return
}

// Retrieves granular information for a single transaction by its unique ID,
// including AI categorization confidence, merchant details, and associated carbon
// footprint.
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a paginated list of the user's transactions, with extensive options
// for filtering by type, category, date range, amount, and intelligent AI-driven
// sorting and search capabilities.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *PaginatedTransactions, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Allows the user to override or refine the AI's categorization for a transaction,
// improving future AI accuracy and personal financial reporting.
func (r *TransactionService) Categorize(ctx context.Context, transactionID string, body TransactionCategorizeParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/categorize", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Begins the process of disputing a specific transaction, providing details and
// supporting documentation for review by our compliance team and AI.
func (r *TransactionService) Dispute(ctx context.Context, transactionID string, body TransactionDisputeParams, opts ...option.RequestOption) (res *TransactionDisputeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/dispute", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows the user to add or update personal notes for a specific transaction.
func (r *TransactionService) UpdateNotes(ctx context.Context, transactionID string, body TransactionUpdateNotesParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s/notes", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type PaginatedTransactions struct {
	Data []Transaction             `json:"data"`
	JSON paginatedTransactionsJSON `json:"-"`
	PaginatedList
}

// paginatedTransactionsJSON contains the JSON metadata for the struct
// [PaginatedTransactions]
type paginatedTransactionsJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedTransactions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedTransactionsJSON) RawJSON() string {
	return r.raw
}

type Transaction struct {
	ID                   string                     `json:"id"`
	AccountID            string                     `json:"accountId"`
	AICategoryConfidence float64                    `json:"aiCategoryConfidence"`
	Amount               float64                    `json:"amount"`
	CarbonFootprint      float64                    `json:"carbonFootprint"`
	Category             string                     `json:"category"`
	Currency             string                     `json:"currency"`
	Date                 time.Time                  `json:"date" format:"date"`
	Description          string                     `json:"description"`
	DisputeStatus        TransactionDisputeStatus   `json:"disputeStatus"`
	Location             TransactionLocation        `json:"location"`
	MerchantDetails      TransactionMerchantDetails `json:"merchantDetails"`
	Notes                string                     `json:"notes"`
	PaymentChannel       TransactionPaymentChannel  `json:"paymentChannel"`
	PostedDate           time.Time                  `json:"postedDate" format:"date"`
	ReceiptURL           string                     `json:"receiptUrl" format:"uri"`
	Tags                 []string                   `json:"tags"`
	Type                 TransactionType            `json:"type"`
	JSON                 transactionJSON            `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	AICategoryConfidence apijson.Field
	Amount               apijson.Field
	CarbonFootprint      apijson.Field
	Category             apijson.Field
	Currency             apijson.Field
	Date                 apijson.Field
	Description          apijson.Field
	DisputeStatus        apijson.Field
	Location             apijson.Field
	MerchantDetails      apijson.Field
	Notes                apijson.Field
	PaymentChannel       apijson.Field
	PostedDate           apijson.Field
	ReceiptURL           apijson.Field
	Tags                 apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}

type TransactionDisputeStatus string

const (
	TransactionDisputeStatusNone     TransactionDisputeStatus = "none"
	TransactionDisputeStatusPending  TransactionDisputeStatus = "pending"
	TransactionDisputeStatusResolved TransactionDisputeStatus = "resolved"
	TransactionDisputeStatusDenied   TransactionDisputeStatus = "denied"
)

func (r TransactionDisputeStatus) IsKnown() bool {
	switch r {
	case TransactionDisputeStatusNone, TransactionDisputeStatusPending, TransactionDisputeStatusResolved, TransactionDisputeStatusDenied:
		return true
	}
	return false
}

type TransactionLocation struct {
	City      string                  `json:"city"`
	Latitude  float64                 `json:"latitude"`
	Longitude float64                 `json:"longitude"`
	JSON      transactionLocationJSON `json:"-"`
}

// transactionLocationJSON contains the JSON metadata for the struct
// [TransactionLocation]
type transactionLocationJSON struct {
	City        apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionLocationJSON) RawJSON() string {
	return r.raw
}

type TransactionMerchantDetails struct {
	Address TransactionMerchantDetailsAddress `json:"address"`
	LogoURL string                            `json:"logoUrl" format:"uri"`
	Name    string                            `json:"name"`
	Website string                            `json:"website" format:"uri"`
	JSON    transactionMerchantDetailsJSON    `json:"-"`
}

// transactionMerchantDetailsJSON contains the JSON metadata for the struct
// [TransactionMerchantDetails]
type transactionMerchantDetailsJSON struct {
	Address     apijson.Field
	LogoURL     apijson.Field
	Name        apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionMerchantDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionMerchantDetailsJSON) RawJSON() string {
	return r.raw
}

type TransactionMerchantDetailsAddress struct {
	City  string                                `json:"city"`
	State string                                `json:"state"`
	Zip   string                                `json:"zip"`
	JSON  transactionMerchantDetailsAddressJSON `json:"-"`
}

// transactionMerchantDetailsAddressJSON contains the JSON metadata for the struct
// [TransactionMerchantDetailsAddress]
type transactionMerchantDetailsAddressJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionMerchantDetailsAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionMerchantDetailsAddressJSON) RawJSON() string {
	return r.raw
}

type TransactionPaymentChannel string

const (
	TransactionPaymentChannelInStore     TransactionPaymentChannel = "in_store"
	TransactionPaymentChannelOnline      TransactionPaymentChannel = "online"
	TransactionPaymentChannelBillPayment TransactionPaymentChannel = "bill_payment"
	TransactionPaymentChannelTransfer    TransactionPaymentChannel = "transfer"
)

func (r TransactionPaymentChannel) IsKnown() bool {
	switch r {
	case TransactionPaymentChannelInStore, TransactionPaymentChannelOnline, TransactionPaymentChannelBillPayment, TransactionPaymentChannelTransfer:
		return true
	}
	return false
}

type TransactionType string

const (
	TransactionTypeIncome      TransactionType = "income"
	TransactionTypeExpense     TransactionType = "expense"
	TransactionTypeTransfer    TransactionType = "transfer"
	TransactionTypeInvestment  TransactionType = "investment"
	TransactionTypeRefund      TransactionType = "refund"
	TransactionTypeBillPayment TransactionType = "bill_payment"
)

func (r TransactionType) IsKnown() bool {
	switch r {
	case TransactionTypeIncome, TransactionTypeExpense, TransactionTypeTransfer, TransactionTypeInvestment, TransactionTypeRefund, TransactionTypeBillPayment:
		return true
	}
	return false
}

type TransactionDisputeResponse struct {
	DisputeID   string                           `json:"disputeId"`
	LastUpdated time.Time                        `json:"lastUpdated" format:"date-time"`
	NextSteps   string                           `json:"nextSteps"`
	Status      TransactionDisputeResponseStatus `json:"status"`
	JSON        transactionDisputeResponseJSON   `json:"-"`
}

// transactionDisputeResponseJSON contains the JSON metadata for the struct
// [TransactionDisputeResponse]
type transactionDisputeResponseJSON struct {
	DisputeID   apijson.Field
	LastUpdated apijson.Field
	NextSteps   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionDisputeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionDisputeResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionDisputeResponseStatus string

const (
	TransactionDisputeResponseStatusPending     TransactionDisputeResponseStatus = "pending"
	TransactionDisputeResponseStatusUnderReview TransactionDisputeResponseStatus = "under_review"
	TransactionDisputeResponseStatusResolved    TransactionDisputeResponseStatus = "resolved"
	TransactionDisputeResponseStatusDenied      TransactionDisputeResponseStatus = "denied"
)

func (r TransactionDisputeResponseStatus) IsKnown() bool {
	switch r {
	case TransactionDisputeResponseStatusPending, TransactionDisputeResponseStatusUnderReview, TransactionDisputeResponseStatusResolved, TransactionDisputeResponseStatusDenied:
		return true
	}
	return false
}

type TransactionListParams struct {
	// Filter transactions by their AI-assigned or user-defined category.
	Category param.Field[string] `query:"category"`
	// Retrieve transactions up to this date (inclusive).
	EndDate param.Field[time.Time] `query:"endDate" format:"date"`
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// Filter for transactions with an amount less than or equal to this value.
	MaxAmount param.Field[float64] `query:"maxAmount"`
	// Filter for transactions with an amount greater than or equal to this value.
	MinAmount param.Field[float64] `query:"minAmount"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// Free-text search across transaction descriptions, merchants, and notes.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// Retrieve transactions from this date (inclusive).
	StartDate param.Field[time.Time] `query:"startDate" format:"date"`
	// Filter transactions by type (e.g., income, expense, transfer).
	Type param.Field[TransactionListParamsType] `query:"type"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter transactions by type (e.g., income, expense, transfer).
type TransactionListParamsType string

const (
	TransactionListParamsTypeIncome      TransactionListParamsType = "income"
	TransactionListParamsTypeExpense     TransactionListParamsType = "expense"
	TransactionListParamsTypeTransfer    TransactionListParamsType = "transfer"
	TransactionListParamsTypeInvestment  TransactionListParamsType = "investment"
	TransactionListParamsTypeRefund      TransactionListParamsType = "refund"
	TransactionListParamsTypeBillPayment TransactionListParamsType = "bill_payment"
)

func (r TransactionListParamsType) IsKnown() bool {
	switch r {
	case TransactionListParamsTypeIncome, TransactionListParamsTypeExpense, TransactionListParamsTypeTransfer, TransactionListParamsTypeInvestment, TransactionListParamsTypeRefund, TransactionListParamsTypeBillPayment:
		return true
	}
	return false
}

type TransactionCategorizeParams struct {
	Category      param.Field[string] `json:"category,required"`
	ApplyToFuture param.Field[bool]   `json:"applyToFuture"`
	Notes         param.Field[string] `json:"notes"`
}

func (r TransactionCategorizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionDisputeParams struct {
	Details             param.Field[string]                         `json:"details,required"`
	Reason              param.Field[TransactionDisputeParamsReason] `json:"reason,required"`
	SupportingDocuments param.Field[[]string]                       `json:"supportingDocuments" format:"uri"`
}

func (r TransactionDisputeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionDisputeParamsReason string

const (
	TransactionDisputeParamsReasonUnauthorized       TransactionDisputeParamsReason = "unauthorized"
	TransactionDisputeParamsReasonProductNotReceived TransactionDisputeParamsReason = "product_not_received"
	TransactionDisputeParamsReasonIncorrectAmount    TransactionDisputeParamsReason = "incorrect_amount"
	TransactionDisputeParamsReasonOther              TransactionDisputeParamsReason = "other"
)

func (r TransactionDisputeParamsReason) IsKnown() bool {
	switch r {
	case TransactionDisputeParamsReasonUnauthorized, TransactionDisputeParamsReasonProductNotReceived, TransactionDisputeParamsReasonIncorrectAmount, TransactionDisputeParamsReasonOther:
		return true
	}
	return false
}

type TransactionUpdateNotesParams struct {
	Notes param.Field[string] `json:"notes,required"`
}

func (r TransactionUpdateNotesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
