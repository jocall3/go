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

// AccountService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options           []option.RequestOption
	Transactions      *AccountTransactionService
	OverdraftSettings *AccountOverdraftSettingService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Transactions = NewAccountTransactionService(opts...)
	r.OverdraftSettings = NewAccountOverdraftSettingService(opts...)
	return
}

// Begins the secure process of linking a new external financial institution (e.g.,
// another bank, investment platform) to the user's profile, typically involving a
// third-party tokenized flow.
func (r *AccountService) Link(ctx context.Context, body AccountLinkParams, opts ...option.RequestOption) (res *AccountLinkResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/link"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves comprehensive analytics for a specific financial account, including
// historical balance trends, projected cash flow, and AI-driven insights into
// spending patterns.
func (r *AccountService) GetDetails(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/details", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches a comprehensive, real-time list of all external financial accounts
// linked to the user's profile, including consolidated balances and institutional
// details.
func (r *AccountService) GetMe(ctx context.Context, query AccountGetMeParams, opts ...option.RequestOption) (res *AccountGetMeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "accounts/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetches digital statements for a specific account, allowing filtering by date
// range and format.
func (r *AccountService) GetStatements(ctx context.Context, accountID string, query AccountGetStatementsParams, opts ...option.RequestOption) (res *AccountGetStatementsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/statements", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LinkedAccount struct {
	ID               string            `json:"id"`
	AvailableBalance float64           `json:"availableBalance"`
	Currency         string            `json:"currency"`
	CurrentBalance   float64           `json:"currentBalance"`
	ExternalID       string            `json:"externalId"`
	InstitutionName  string            `json:"institutionName"`
	LastUpdated      time.Time         `json:"lastUpdated" format:"date-time"`
	Mask             string            `json:"mask"`
	Name             string            `json:"name"`
	Subtype          string            `json:"subtype"`
	Type             string            `json:"type"`
	JSON             linkedAccountJSON `json:"-"`
}

// linkedAccountJSON contains the JSON metadata for the struct [LinkedAccount]
type linkedAccountJSON struct {
	ID               apijson.Field
	AvailableBalance apijson.Field
	Currency         apijson.Field
	CurrentBalance   apijson.Field
	ExternalID       apijson.Field
	InstitutionName  apijson.Field
	LastUpdated      apijson.Field
	Mask             apijson.Field
	Name             apijson.Field
	Subtype          apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LinkedAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r linkedAccountJSON) RawJSON() string {
	return r.raw
}

type AccountLinkResponse struct {
	AuthUri       string                  `json:"authUri" format:"uri"`
	LinkSessionID string                  `json:"linkSessionId"`
	Message       string                  `json:"message"`
	Status        string                  `json:"status"`
	JSON          accountLinkResponseJSON `json:"-"`
}

// accountLinkResponseJSON contains the JSON metadata for the struct
// [AccountLinkResponse]
type accountLinkResponseJSON struct {
	AuthUri       apijson.Field
	LinkSessionID apijson.Field
	Message       apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLinkResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetDetailsResponse struct {
	AccountHolder     string                                     `json:"accountHolder"`
	BalanceHistory    []AccountGetDetailsResponseBalanceHistory  `json:"balanceHistory"`
	InterestRate      float64                                    `json:"interestRate"`
	OpenedDate        time.Time                                  `json:"openedDate" format:"date"`
	ProjectedCashFlow AccountGetDetailsResponseProjectedCashFlow `json:"projectedCashFlow"`
	TransactionsCount int64                                      `json:"transactionsCount"`
	JSON              accountGetDetailsResponseJSON              `json:"-"`
	LinkedAccount
}

// accountGetDetailsResponseJSON contains the JSON metadata for the struct
// [AccountGetDetailsResponse]
type accountGetDetailsResponseJSON struct {
	AccountHolder     apijson.Field
	BalanceHistory    apijson.Field
	InterestRate      apijson.Field
	OpenedDate        apijson.Field
	ProjectedCashFlow apijson.Field
	TransactionsCount apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetDetailsResponseBalanceHistory struct {
	Balance float64                                     `json:"balance"`
	Date    time.Time                                   `json:"date" format:"date"`
	JSON    accountGetDetailsResponseBalanceHistoryJSON `json:"-"`
}

// accountGetDetailsResponseBalanceHistoryJSON contains the JSON metadata for the
// struct [AccountGetDetailsResponseBalanceHistory]
type accountGetDetailsResponseBalanceHistoryJSON struct {
	Balance     apijson.Field
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetDetailsResponseBalanceHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetDetailsResponseBalanceHistoryJSON) RawJSON() string {
	return r.raw
}

type AccountGetDetailsResponseProjectedCashFlow struct {
	ConfidenceScore int64                                          `json:"confidenceScore"`
	Days30          float64                                        `json:"days30"`
	Days90          float64                                        `json:"days90"`
	JSON            accountGetDetailsResponseProjectedCashFlowJSON `json:"-"`
}

// accountGetDetailsResponseProjectedCashFlowJSON contains the JSON metadata for
// the struct [AccountGetDetailsResponseProjectedCashFlow]
type accountGetDetailsResponseProjectedCashFlowJSON struct {
	ConfidenceScore apijson.Field
	Days30          apijson.Field
	Days90          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountGetDetailsResponseProjectedCashFlow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetDetailsResponseProjectedCashFlowJSON) RawJSON() string {
	return r.raw
}

type AccountGetMeResponse struct {
	Data []LinkedAccount          `json:"data"`
	JSON accountGetMeResponseJSON `json:"-"`
	PaginatedList
}

// accountGetMeResponseJSON contains the JSON metadata for the struct
// [AccountGetMeResponse]
type accountGetMeResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetMeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetMeResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetStatementsResponse struct {
	AccountID    string                                   `json:"accountId"`
	DownloadURLs AccountGetStatementsResponseDownloadURLs `json:"downloadUrls"`
	Period       string                                   `json:"period"`
	StatementID  string                                   `json:"statementId"`
	JSON         accountGetStatementsResponseJSON         `json:"-"`
}

// accountGetStatementsResponseJSON contains the JSON metadata for the struct
// [AccountGetStatementsResponse]
type accountGetStatementsResponseJSON struct {
	AccountID    apijson.Field
	DownloadURLs apijson.Field
	Period       apijson.Field
	StatementID  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountGetStatementsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetStatementsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetStatementsResponseDownloadURLs struct {
	Csv  string                                       `json:"csv" format:"uri"`
	Pdf  string                                       `json:"pdf" format:"uri"`
	JSON accountGetStatementsResponseDownloadURLsJSON `json:"-"`
}

// accountGetStatementsResponseDownloadURLsJSON contains the JSON metadata for the
// struct [AccountGetStatementsResponseDownloadURLs]
type accountGetStatementsResponseDownloadURLsJSON struct {
	Csv         apijson.Field
	Pdf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetStatementsResponseDownloadURLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetStatementsResponseDownloadURLsJSON) RawJSON() string {
	return r.raw
}

type AccountLinkParams struct {
	CountryCode     param.Field[string] `json:"countryCode,required"`
	InstitutionName param.Field[string] `json:"institutionName,required"`
}

func (r AccountLinkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGetMeParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AccountGetMeParams]'s query parameters as `url.Values`.
func (r AccountGetMeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountGetStatementsParams struct {
	// Month for the statement (1-12).
	Month param.Field[int64] `query:"month,required"`
	// Year for the statement.
	Year param.Field[int64] `query:"year,required"`
	// Desired format for the statement. Use 'application/json' Accept header for
	// download links.
	Format param.Field[AccountGetStatementsParamsFormat] `query:"format"`
}

// URLQuery serializes [AccountGetStatementsParams]'s query parameters as
// `url.Values`.
func (r AccountGetStatementsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Desired format for the statement. Use 'application/json' Accept header for
// download links.
type AccountGetStatementsParamsFormat string

const (
	AccountGetStatementsParamsFormatPdf AccountGetStatementsParamsFormat = "pdf"
	AccountGetStatementsParamsFormatCsv AccountGetStatementsParamsFormat = "csv"
)

func (r AccountGetStatementsParamsFormat) IsKnown() bool {
	switch r {
	case AccountGetStatementsParamsFormatPdf, AccountGetStatementsParamsFormatCsv:
		return true
	}
	return false
}
