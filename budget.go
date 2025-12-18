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

// BudgetService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBudgetService] method instead.
type BudgetService struct {
	Options []option.RequestOption
}

// NewBudgetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBudgetService(opts ...option.RequestOption) (r *BudgetService) {
	r = &BudgetService{}
	r.Options = opts
	return
}

// Creates a new financial budget for the user, with optional AI auto-population of
// categories and amounts.
func (r *BudgetService) New(ctx context.Context, body BudgetNewParams, opts ...option.RequestOption) (res *Budget, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves detailed information for a specific budget, including current
// spending, remaining amounts, and AI recommendations.
func (r *BudgetService) Get(ctx context.Context, budgetID string, opts ...option.RequestOption) (res *Budget, err error) {
	opts = slices.Concat(r.Options, opts)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return
	}
	path := fmt.Sprintf("budgets/%s", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the parameters of an existing budget, such as total amount, dates, or
// categories.
func (r *BudgetService) Update(ctx context.Context, budgetID string, body BudgetUpdateParams, opts ...option.RequestOption) (res *Budget, err error) {
	opts = slices.Concat(r.Options, opts)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return
	}
	path := fmt.Sprintf("budgets/%s", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of all active and historical budgets for the authenticated
// user.
func (r *BudgetService) List(ctx context.Context, query BudgetListParams, opts ...option.RequestOption) (res *BudgetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a specific budget from the user's profile.
func (r *BudgetService) Delete(ctx context.Context, budgetID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return
	}
	path := fmt.Sprintf("budgets/%s", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Budget struct {
	ID                string      `json:"id"`
	AIRecommendations []AIInsight `json:"aiRecommendations"`
	// Percentage threshold to trigger an alert (e.g., 80 for 80%).
	AlertThreshold  int64            `json:"alertThreshold"`
	Categories      []BudgetCategory `json:"categories"`
	EndDate         time.Time        `json:"endDate" format:"date"`
	Name            string           `json:"name"`
	Period          BudgetPeriod     `json:"period"`
	RemainingAmount float64          `json:"remainingAmount"`
	SpentAmount     float64          `json:"spentAmount"`
	StartDate       time.Time        `json:"startDate" format:"date"`
	Status          BudgetStatus     `json:"status"`
	TotalAmount     float64          `json:"totalAmount"`
	JSON            budgetJSON       `json:"-"`
}

// budgetJSON contains the JSON metadata for the struct [Budget]
type budgetJSON struct {
	ID                apijson.Field
	AIRecommendations apijson.Field
	AlertThreshold    apijson.Field
	Categories        apijson.Field
	EndDate           apijson.Field
	Name              apijson.Field
	Period            apijson.Field
	RemainingAmount   apijson.Field
	SpentAmount       apijson.Field
	StartDate         apijson.Field
	Status            apijson.Field
	TotalAmount       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Budget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetJSON) RawJSON() string {
	return r.raw
}

type BudgetCategory struct {
	Allocated float64            `json:"allocated"`
	Name      string             `json:"name"`
	Remaining float64            `json:"remaining"`
	Spent     float64            `json:"spent"`
	JSON      budgetCategoryJSON `json:"-"`
}

// budgetCategoryJSON contains the JSON metadata for the struct [BudgetCategory]
type budgetCategoryJSON struct {
	Allocated   apijson.Field
	Name        apijson.Field
	Remaining   apijson.Field
	Spent       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetCategoryJSON) RawJSON() string {
	return r.raw
}

type BudgetPeriod string

const (
	BudgetPeriodWeekly    BudgetPeriod = "weekly"
	BudgetPeriodMonthly   BudgetPeriod = "monthly"
	BudgetPeriodQuarterly BudgetPeriod = "quarterly"
	BudgetPeriodYearly    BudgetPeriod = "yearly"
)

func (r BudgetPeriod) IsKnown() bool {
	switch r {
	case BudgetPeriodWeekly, BudgetPeriodMonthly, BudgetPeriodQuarterly, BudgetPeriodYearly:
		return true
	}
	return false
}

type BudgetStatus string

const (
	BudgetStatusActive   BudgetStatus = "active"
	BudgetStatusArchived BudgetStatus = "archived"
)

func (r BudgetStatus) IsKnown() bool {
	switch r {
	case BudgetStatusActive, BudgetStatusArchived:
		return true
	}
	return false
}

type BudgetListResponse struct {
	Data []Budget               `json:"data"`
	JSON budgetListResponseJSON `json:"-"`
	PaginatedList
}

// budgetListResponseJSON contains the JSON metadata for the struct
// [BudgetListResponse]
type budgetListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BudgetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r budgetListResponseJSON) RawJSON() string {
	return r.raw
}

type BudgetNewParams struct {
	EndDate     param.Field[time.Time]             `json:"endDate,required" format:"date"`
	Name        param.Field[string]                `json:"name,required"`
	Period      param.Field[BudgetNewParamsPeriod] `json:"period,required"`
	StartDate   param.Field[time.Time]             `json:"startDate,required" format:"date"`
	TotalAmount param.Field[float64]               `json:"totalAmount,required"`
	// If true, AI will populate categories based on past spending.
	AIAutoPopulate param.Field[bool]                      `json:"aiAutoPopulate"`
	AlertThreshold param.Field[int64]                     `json:"alertThreshold"`
	Categories     param.Field[[]BudgetNewParamsCategory] `json:"categories"`
}

func (r BudgetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetNewParamsPeriod string

const (
	BudgetNewParamsPeriodWeekly    BudgetNewParamsPeriod = "weekly"
	BudgetNewParamsPeriodMonthly   BudgetNewParamsPeriod = "monthly"
	BudgetNewParamsPeriodQuarterly BudgetNewParamsPeriod = "quarterly"
	BudgetNewParamsPeriodYearly    BudgetNewParamsPeriod = "yearly"
)

func (r BudgetNewParamsPeriod) IsKnown() bool {
	switch r {
	case BudgetNewParamsPeriodWeekly, BudgetNewParamsPeriodMonthly, BudgetNewParamsPeriodQuarterly, BudgetNewParamsPeriodYearly:
		return true
	}
	return false
}

type BudgetNewParamsCategory struct {
	Allocated param.Field[float64] `json:"allocated"`
	Name      param.Field[string]  `json:"name"`
}

func (r BudgetNewParamsCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetUpdateParams struct {
	AlertThreshold param.Field[int64]   `json:"alertThreshold"`
	Name           param.Field[string]  `json:"name"`
	TotalAmount    param.Field[float64] `json:"totalAmount"`
}

func (r BudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetListParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BudgetListParams]'s query parameters as `url.Values`.
func (r BudgetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
