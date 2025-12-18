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

// CorporateCardService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateCardService] method instead.
type CorporateCardService struct {
	Options []option.RequestOption
}

// NewCorporateCardService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorporateCardService(opts ...option.RequestOption) (r *CorporateCardService) {
	r = &CorporateCardService{}
	r.Options = opts
	return
}

// Retrieves a comprehensive list of all physical and virtual corporate cards
// associated with the user's organization, including their status, assigned
// holder, and current spending controls.
func (r *CorporateCardService) List(ctx context.Context, query CorporateCardListParams, opts ...option.RequestOption) (res *CorporateCardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Creates and issues a new virtual corporate card with specified spending limits,
// merchant restrictions, and expiration dates, ideal for secure online purchases
// and temporary projects.
func (r *CorporateCardService) NewVirtual(ctx context.Context, body CorporateCardNewVirtualParams, opts ...option.RequestOption) (res *CorporateCard, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "corporate/cards/virtual"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Immediately changes the frozen status of a corporate card, preventing or
// allowing transactions in real-time, critical for security and expense
// management.
func (r *CorporateCardService) Freeze(ctx context.Context, cardID string, body CorporateCardFreezeParams, opts ...option.RequestOption) (res *CorporateCard, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("corporate/cards/%s/freeze", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of transactions made with a specific corporate card,
// including AI categorization and compliance flags.
func (r *CorporateCardService) ListTransactions(ctx context.Context, cardID string, query CorporateCardListTransactionsParams, opts ...option.RequestOption) (res *PaginatedTransactions, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("corporate/cards/%s/transactions", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the sophisticated spending controls, limits, and policy overrides for a
// specific corporate card, enabling real-time adjustments for security and budget
// adherence.
func (r *CorporateCardService) UpdateControls(ctx context.Context, cardID string, body CorporateCardUpdateControlsParams, opts ...option.RequestOption) (res *CorporateCard, err error) {
	opts = slices.Concat(r.Options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return
	}
	path := fmt.Sprintf("corporate/cards/%s/controls", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CorporateCard struct {
	ID                   string                `json:"id"`
	AssociatedEmployeeID string                `json:"associatedEmployeeId"`
	CardNumberMask       string                `json:"cardNumberMask"`
	CardType             CorporateCardCardType `json:"cardType"`
	Controls             CorporateCardControls `json:"controls"`
	CreatedDate          time.Time             `json:"createdDate" format:"date-time"`
	Currency             string                `json:"currency"`
	ExpirationDate       time.Time             `json:"expirationDate" format:"date"`
	Frozen               bool                  `json:"frozen"`
	HolderName           string                `json:"holderName"`
	SpendingPolicyID     string                `json:"spendingPolicyId"`
	Status               CorporateCardStatus   `json:"status"`
	JSON                 corporateCardJSON     `json:"-"`
}

// corporateCardJSON contains the JSON metadata for the struct [CorporateCard]
type corporateCardJSON struct {
	ID                   apijson.Field
	AssociatedEmployeeID apijson.Field
	CardNumberMask       apijson.Field
	CardType             apijson.Field
	Controls             apijson.Field
	CreatedDate          apijson.Field
	Currency             apijson.Field
	ExpirationDate       apijson.Field
	Frozen               apijson.Field
	HolderName           apijson.Field
	SpendingPolicyID     apijson.Field
	Status               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CorporateCard) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardJSON) RawJSON() string {
	return r.raw
}

type CorporateCardCardType string

const (
	CorporateCardCardTypePhysical CorporateCardCardType = "physical"
	CorporateCardCardTypeVirtual  CorporateCardCardType = "virtual"
)

func (r CorporateCardCardType) IsKnown() bool {
	switch r {
	case CorporateCardCardTypePhysical, CorporateCardCardTypeVirtual:
		return true
	}
	return false
}

type CorporateCardStatus string

const (
	CorporateCardStatusActive    CorporateCardStatus = "Active"
	CorporateCardStatusSuspended CorporateCardStatus = "Suspended"
	CorporateCardStatusCancelled CorporateCardStatus = "Cancelled"
)

func (r CorporateCardStatus) IsKnown() bool {
	switch r {
	case CorporateCardStatusActive, CorporateCardStatusSuspended, CorporateCardStatusCancelled:
		return true
	}
	return false
}

type CorporateCardControls struct {
	AtmWithdrawals               bool                      `json:"atmWithdrawals"`
	ContactlessPayments          bool                      `json:"contactlessPayments"`
	DailyLimit                   float64                   `json:"dailyLimit"`
	InternationalTransactions    bool                      `json:"internationalTransactions"`
	MerchantCategoryRestrictions []string                  `json:"merchantCategoryRestrictions"`
	MonthlyLimit                 float64                   `json:"monthlyLimit"`
	OnlineTransactions           bool                      `json:"onlineTransactions"`
	SingleTransactionLimit       float64                   `json:"singleTransactionLimit"`
	VendorRestrictions           []string                  `json:"vendorRestrictions"`
	JSON                         corporateCardControlsJSON `json:"-"`
}

// corporateCardControlsJSON contains the JSON metadata for the struct
// [CorporateCardControls]
type corporateCardControlsJSON struct {
	AtmWithdrawals               apijson.Field
	ContactlessPayments          apijson.Field
	DailyLimit                   apijson.Field
	InternationalTransactions    apijson.Field
	MerchantCategoryRestrictions apijson.Field
	MonthlyLimit                 apijson.Field
	OnlineTransactions           apijson.Field
	SingleTransactionLimit       apijson.Field
	VendorRestrictions           apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CorporateCardControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardControlsJSON) RawJSON() string {
	return r.raw
}

type CorporateCardControlsParam struct {
	AtmWithdrawals               param.Field[bool]     `json:"atmWithdrawals"`
	ContactlessPayments          param.Field[bool]     `json:"contactlessPayments"`
	DailyLimit                   param.Field[float64]  `json:"dailyLimit"`
	InternationalTransactions    param.Field[bool]     `json:"internationalTransactions"`
	MerchantCategoryRestrictions param.Field[[]string] `json:"merchantCategoryRestrictions"`
	MonthlyLimit                 param.Field[float64]  `json:"monthlyLimit"`
	OnlineTransactions           param.Field[bool]     `json:"onlineTransactions"`
	SingleTransactionLimit       param.Field[float64]  `json:"singleTransactionLimit"`
	VendorRestrictions           param.Field[[]string] `json:"vendorRestrictions"`
}

func (r CorporateCardControlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardListResponse struct {
	Data []CorporateCard               `json:"data"`
	JSON corporateCardListResponseJSON `json:"-"`
	PaginatedList
}

// corporateCardListResponseJSON contains the JSON metadata for the struct
// [CorporateCardListResponse]
type corporateCardListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CorporateCardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r corporateCardListResponseJSON) RawJSON() string {
	return r.raw
}

type CorporateCardListParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CorporateCardListParams]'s query parameters as
// `url.Values`.
func (r CorporateCardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorporateCardNewVirtualParams struct {
	Controls             param.Field[CorporateCardControlsParam] `json:"controls,required"`
	HolderName           param.Field[string]                     `json:"holderName,required"`
	AssociatedEmployeeID param.Field[string]                     `json:"associatedEmployeeId"`
	ExpirationDate       param.Field[time.Time]                  `json:"expirationDate" format:"date"`
	Purpose              param.Field[string]                     `json:"purpose"`
}

func (r CorporateCardNewVirtualParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardFreezeParams struct {
	Freeze param.Field[bool] `json:"freeze,required"`
}

func (r CorporateCardFreezeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CorporateCardListTransactionsParams struct {
	// The end date for the date range filter (inclusive).
	EndDate param.Field[time.Time] `query:"endDate" format:"date"`
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
	// The start date for the date range filter (inclusive).
	StartDate param.Field[time.Time] `query:"startDate" format:"date"`
}

// URLQuery serializes [CorporateCardListTransactionsParams]'s query parameters as
// `url.Values`.
func (r CorporateCardListTransactionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorporateCardUpdateControlsParams struct {
	CorporateCardControls CorporateCardControlsParam `json:"corporate_card_controls,required"`
}

func (r CorporateCardUpdateControlsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CorporateCardControls)
}
