// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
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

// PaymentFxService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentFxService] method instead.
type PaymentFxService struct {
	Options []option.RequestOption
}

// NewPaymentFxService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaymentFxService(opts ...option.RequestOption) (r *PaymentFxService) {
	r = &PaymentFxService{}
	r.Options = opts
	return
}

// Executes an instant currency conversion between two currencies, either from a
// balance or into a specified account.
func (r *PaymentFxService) Convert(ctx context.Context, body PaymentFxConvertParams, opts ...option.RequestOption) (res *PaymentFxConvertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments/fx/convert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves current and AI-predicted future foreign exchange rates for a specified
// currency pair, including bid/ask spreads and historical volatility data for
// informed decisions.
func (r *PaymentFxService) GetRates(ctx context.Context, query PaymentFxGetRatesParams, opts ...option.RequestOption) (res *PaymentFxGetRatesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments/fx/rates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaymentFxConvertResponse struct {
	ConversionID        string                         `json:"conversionId"`
	ConversionTimestamp time.Time                      `json:"conversionTimestamp" format:"date-time"`
	FeesApplied         float64                        `json:"feesApplied"`
	FxRateApplied       float64                        `json:"fxRateApplied"`
	SourceAmount        float64                        `json:"sourceAmount"`
	SourceCurrency      string                         `json:"sourceCurrency"`
	Status              PaymentFxConvertResponseStatus `json:"status"`
	TargetAmount        float64                        `json:"targetAmount"`
	TransactionID       string                         `json:"transactionId"`
	JSON                paymentFxConvertResponseJSON   `json:"-"`
}

// paymentFxConvertResponseJSON contains the JSON metadata for the struct
// [PaymentFxConvertResponse]
type paymentFxConvertResponseJSON struct {
	ConversionID        apijson.Field
	ConversionTimestamp apijson.Field
	FeesApplied         apijson.Field
	FxRateApplied       apijson.Field
	SourceAmount        apijson.Field
	SourceCurrency      apijson.Field
	Status              apijson.Field
	TargetAmount        apijson.Field
	TransactionID       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PaymentFxConvertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxConvertResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentFxConvertResponseStatus string

const (
	PaymentFxConvertResponseStatusCompleted PaymentFxConvertResponseStatus = "completed"
	PaymentFxConvertResponseStatusFailed    PaymentFxConvertResponseStatus = "failed"
)

func (r PaymentFxConvertResponseStatus) IsKnown() bool {
	switch r {
	case PaymentFxConvertResponseStatusCompleted, PaymentFxConvertResponseStatusFailed:
		return true
	}
	return false
}

type PaymentFxGetRatesResponse struct {
	BaseCurrency         string                                        `json:"baseCurrency"`
	CurrentRate          PaymentFxGetRatesResponseCurrentRate          `json:"currentRate"`
	HistoricalVolatility PaymentFxGetRatesResponseHistoricalVolatility `json:"historicalVolatility"`
	PredictiveRates      []PaymentFxGetRatesResponsePredictiveRate     `json:"predictiveRates"`
	TargetCurrency       string                                        `json:"targetCurrency"`
	JSON                 paymentFxGetRatesResponseJSON                 `json:"-"`
}

// paymentFxGetRatesResponseJSON contains the JSON metadata for the struct
// [PaymentFxGetRatesResponse]
type paymentFxGetRatesResponseJSON struct {
	BaseCurrency         apijson.Field
	CurrentRate          apijson.Field
	HistoricalVolatility apijson.Field
	PredictiveRates      apijson.Field
	TargetCurrency       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponseJSON) RawJSON() string {
	return r.raw
}

type PaymentFxGetRatesResponseCurrentRate struct {
	Ask       float64                                  `json:"ask"`
	Bid       float64                                  `json:"bid"`
	Mid       float64                                  `json:"mid"`
	Timestamp time.Time                                `json:"timestamp" format:"date-time"`
	JSON      paymentFxGetRatesResponseCurrentRateJSON `json:"-"`
}

// paymentFxGetRatesResponseCurrentRateJSON contains the JSON metadata for the
// struct [PaymentFxGetRatesResponseCurrentRate]
type paymentFxGetRatesResponseCurrentRateJSON struct {
	Ask         apijson.Field
	Bid         apijson.Field
	Mid         apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponseCurrentRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponseCurrentRateJSON) RawJSON() string {
	return r.raw
}

type PaymentFxGetRatesResponseHistoricalVolatility struct {
	Past30Days float64                                           `json:"past30Days"`
	Past7Days  float64                                           `json:"past7Days"`
	JSON       paymentFxGetRatesResponseHistoricalVolatilityJSON `json:"-"`
}

// paymentFxGetRatesResponseHistoricalVolatilityJSON contains the JSON metadata for
// the struct [PaymentFxGetRatesResponseHistoricalVolatility]
type paymentFxGetRatesResponseHistoricalVolatilityJSON struct {
	Past30Days  apijson.Field
	Past7Days   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponseHistoricalVolatility) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponseHistoricalVolatilityJSON) RawJSON() string {
	return r.raw
}

type PaymentFxGetRatesResponsePredictiveRate struct {
	AIModelConfidence       float64                                     `json:"aiModelConfidence"`
	ConfidenceIntervalLower float64                                     `json:"confidenceIntervalLower"`
	ConfidenceIntervalUpper float64                                     `json:"confidenceIntervalUpper"`
	Date                    time.Time                                   `json:"date" format:"date"`
	PredictedMidRate        float64                                     `json:"predictedMidRate"`
	JSON                    paymentFxGetRatesResponsePredictiveRateJSON `json:"-"`
}

// paymentFxGetRatesResponsePredictiveRateJSON contains the JSON metadata for the
// struct [PaymentFxGetRatesResponsePredictiveRate]
type paymentFxGetRatesResponsePredictiveRateJSON struct {
	AIModelConfidence       apijson.Field
	ConfidenceIntervalLower apijson.Field
	ConfidenceIntervalUpper apijson.Field
	Date                    apijson.Field
	PredictedMidRate        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PaymentFxGetRatesResponsePredictiveRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paymentFxGetRatesResponsePredictiveRateJSON) RawJSON() string {
	return r.raw
}

type PaymentFxConvertParams struct {
	SourceAccountID param.Field[string]  `json:"sourceAccountId,required"`
	SourceAmount    param.Field[float64] `json:"sourceAmount,required"`
	SourceCurrency  param.Field[string]  `json:"sourceCurrency,required"`
	TargetCurrency  param.Field[string]  `json:"targetCurrency,required"`
	FxRateLock      param.Field[bool]    `json:"fxRateLock"`
	TargetAccountID param.Field[string]  `json:"targetAccountId"`
}

func (r PaymentFxConvertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentFxGetRatesParams struct {
	// The base currency code (e.g., USD).
	BaseCurrency param.Field[interface{}] `query:"baseCurrency,required"`
	// The target currency code (e.g., EUR).
	TargetCurrency param.Field[interface{}] `query:"targetCurrency,required"`
	// Number of days into the future to provide an AI-driven prediction.
	ForecastDays param.Field[interface{}] `query:"forecastDays"`
}

// URLQuery serializes [PaymentFxGetRatesParams]'s query parameters as
// `url.Values`.
func (r PaymentFxGetRatesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
