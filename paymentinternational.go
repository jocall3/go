// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/param"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// PaymentInternationalService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentInternationalService] method instead.
type PaymentInternationalService struct {
	Options []option.RequestOption
}

// NewPaymentInternationalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPaymentInternationalService(opts ...option.RequestOption) (r *PaymentInternationalService) {
	r = &PaymentInternationalService{}
	r.Options = opts
	return
}

// Facilitates the secure initiation of an international wire transfer to a
// beneficiary in another country and currency, leveraging optimal FX rates and
// tracking capabilities.
func (r *PaymentInternationalService) Initiate(ctx context.Context, body PaymentInternationalInitiateParams, opts ...option.RequestOption) (res *InternationalPaymentStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "payments/international/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the current processing status and details of an initiated
// international payment.
func (r *PaymentInternationalService) GetStatus(ctx context.Context, paymentID string, opts ...option.RequestOption) (res *InternationalPaymentStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if paymentID == "" {
		err = errors.New("missing required paymentId parameter")
		return
	}
	path := fmt.Sprintf("payments/international/%s/status", paymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type InternationalPaymentStatus struct {
	EstimatedCompletionTime time.Time                        `json:"estimatedCompletionTime" format:"date-time"`
	FeesApplied             float64                          `json:"feesApplied"`
	FxRateApplied           float64                          `json:"fxRateApplied"`
	Message                 string                           `json:"message"`
	PaymentID               string                           `json:"paymentId"`
	SourceAmount            float64                          `json:"sourceAmount"`
	SourceCurrency          string                           `json:"sourceCurrency"`
	Status                  InternationalPaymentStatusStatus `json:"status"`
	TargetAmount            float64                          `json:"targetAmount"`
	TargetCurrency          string                           `json:"targetCurrency"`
	TrackingURL             string                           `json:"trackingUrl" format:"uri"`
	JSON                    internationalPaymentStatusJSON   `json:"-"`
}

// internationalPaymentStatusJSON contains the JSON metadata for the struct
// [InternationalPaymentStatus]
type internationalPaymentStatusJSON struct {
	EstimatedCompletionTime apijson.Field
	FeesApplied             apijson.Field
	FxRateApplied           apijson.Field
	Message                 apijson.Field
	PaymentID               apijson.Field
	SourceAmount            apijson.Field
	SourceCurrency          apijson.Field
	Status                  apijson.Field
	TargetAmount            apijson.Field
	TargetCurrency          apijson.Field
	TrackingURL             apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InternationalPaymentStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r internationalPaymentStatusJSON) RawJSON() string {
	return r.raw
}

type InternationalPaymentStatusStatus string

const (
	InternationalPaymentStatusStatusInProgress    InternationalPaymentStatusStatus = "in_progress"
	InternationalPaymentStatusStatusCompleted     InternationalPaymentStatusStatus = "completed"
	InternationalPaymentStatusStatusFailed        InternationalPaymentStatusStatus = "failed"
	InternationalPaymentStatusStatusHeldForReview InternationalPaymentStatusStatus = "held_for_review"
)

func (r InternationalPaymentStatusStatus) IsKnown() bool {
	switch r {
	case InternationalPaymentStatusStatusInProgress, InternationalPaymentStatusStatusCompleted, InternationalPaymentStatusStatusFailed, InternationalPaymentStatusStatusHeldForReview:
		return true
	}
	return false
}

type PaymentInternationalInitiateParams struct {
	Amount          param.Field[float64]                                       `json:"amount,required"`
	Beneficiary     param.Field[PaymentInternationalInitiateParamsBeneficiary] `json:"beneficiary,required"`
	SourceAccountID param.Field[string]                                        `json:"sourceAccountId,required"`
	SourceCurrency  param.Field[string]                                        `json:"sourceCurrency,required"`
	TargetCurrency  param.Field[string]                                        `json:"targetCurrency,required"`
	FxRateLock      param.Field[bool]                                          `json:"fxRateLock"`
	FxRateProvider  param.Field[string]                                        `json:"fxRateProvider"`
	Purpose         param.Field[string]                                        `json:"purpose"`
}

func (r PaymentInternationalInitiateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PaymentInternationalInitiateParamsBeneficiary struct {
	Address  param.Field[string] `json:"address"`
	BankName param.Field[string] `json:"bankName"`
	Iban     param.Field[string] `json:"iban"`
	Name     param.Field[string] `json:"name"`
	SwiftBic param.Field[string] `json:"swiftBic"`
}

func (r PaymentInternationalInitiateParamsBeneficiary) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
