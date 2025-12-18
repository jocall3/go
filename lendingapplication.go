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

// LendingApplicationService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLendingApplicationService] method instead.
type LendingApplicationService struct {
	Options []option.RequestOption
}

// NewLendingApplicationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLendingApplicationService(opts ...option.RequestOption) (r *LendingApplicationService) {
	r = &LendingApplicationService{}
	r.Options = opts
	return
}

// Retrieves the current status and detailed information for a submitted loan
// application, including AI underwriting outcomes, approved terms, and next steps.
func (r *LendingApplicationService) Get(ctx context.Context, applicationID string, opts ...option.RequestOption) (res *LoanApplicationStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationID == "" {
		err = errors.New("missing required applicationId parameter")
		return
	}
	path := fmt.Sprintf("lending/applications/%s", applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits a new loan application, which is instantly processed and underwritten by
// our Quantum AI, providing rapid decisions and personalized loan offers based on
// real-time financial health data.
func (r *LendingApplicationService) Submit(ctx context.Context, body LendingApplicationSubmitParams, opts ...option.RequestOption) (res *LoanApplicationStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "lending/applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LoanApplicationStatus struct {
	AIUnderwritingResult LoanApplicationStatusAIUnderwritingResult `json:"aiUnderwritingResult"`
	ApplicationDate      time.Time                                 `json:"applicationDate" format:"date-time"`
	ApplicationID        string                                    `json:"applicationId"`
	LoanAmountRequested  float64                                   `json:"loanAmountRequested"`
	LoanPurpose          string                                    `json:"loanPurpose"`
	NextSteps            string                                    `json:"nextSteps"`
	OfferDetails         LoanOffer                                 `json:"offerDetails"`
	Status               LoanApplicationStatusStatus               `json:"status"`
	JSON                 loanApplicationStatusJSON                 `json:"-"`
}

// loanApplicationStatusJSON contains the JSON metadata for the struct
// [LoanApplicationStatus]
type loanApplicationStatusJSON struct {
	AIUnderwritingResult apijson.Field
	ApplicationDate      apijson.Field
	ApplicationID        apijson.Field
	LoanAmountRequested  apijson.Field
	LoanPurpose          apijson.Field
	NextSteps            apijson.Field
	OfferDetails         apijson.Field
	Status               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoanApplicationStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanApplicationStatusJSON) RawJSON() string {
	return r.raw
}

type LoanApplicationStatusAIUnderwritingResult struct {
	AIConfidence            float64                                           `json:"aiConfidence"`
	Decision                LoanApplicationStatusAIUnderwritingResultDecision `json:"decision"`
	MaxApprovedAmount       float64                                           `json:"maxApprovedAmount"`
	Reason                  string                                            `json:"reason"`
	RecommendedInterestRate float64                                           `json:"recommendedInterestRate"`
	JSON                    loanApplicationStatusAIUnderwritingResultJSON     `json:"-"`
}

// loanApplicationStatusAIUnderwritingResultJSON contains the JSON metadata for the
// struct [LoanApplicationStatusAIUnderwritingResult]
type loanApplicationStatusAIUnderwritingResultJSON struct {
	AIConfidence            apijson.Field
	Decision                apijson.Field
	MaxApprovedAmount       apijson.Field
	Reason                  apijson.Field
	RecommendedInterestRate apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LoanApplicationStatusAIUnderwritingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loanApplicationStatusAIUnderwritingResultJSON) RawJSON() string {
	return r.raw
}

type LoanApplicationStatusAIUnderwritingResultDecision string

const (
	LoanApplicationStatusAIUnderwritingResultDecisionApproved            LoanApplicationStatusAIUnderwritingResultDecision = "approved"
	LoanApplicationStatusAIUnderwritingResultDecisionDeclined            LoanApplicationStatusAIUnderwritingResultDecision = "declined"
	LoanApplicationStatusAIUnderwritingResultDecisionPendingManualReview LoanApplicationStatusAIUnderwritingResultDecision = "pending_manual_review"
)

func (r LoanApplicationStatusAIUnderwritingResultDecision) IsKnown() bool {
	switch r {
	case LoanApplicationStatusAIUnderwritingResultDecisionApproved, LoanApplicationStatusAIUnderwritingResultDecisionDeclined, LoanApplicationStatusAIUnderwritingResultDecisionPendingManualReview:
		return true
	}
	return false
}

type LoanApplicationStatusStatus string

const (
	LoanApplicationStatusStatusUnderwriting      LoanApplicationStatusStatus = "underwriting"
	LoanApplicationStatusStatusApproved          LoanApplicationStatusStatus = "approved"
	LoanApplicationStatusStatusDeclined          LoanApplicationStatusStatus = "declined"
	LoanApplicationStatusStatusPendingDocuments  LoanApplicationStatusStatus = "pending_documents"
	LoanApplicationStatusStatusFundingInProgress LoanApplicationStatusStatus = "funding_in_progress"
	LoanApplicationStatusStatusFunded            LoanApplicationStatusStatus = "funded"
)

func (r LoanApplicationStatusStatus) IsKnown() bool {
	switch r {
	case LoanApplicationStatusStatusUnderwriting, LoanApplicationStatusStatusApproved, LoanApplicationStatusStatusDeclined, LoanApplicationStatusStatusPendingDocuments, LoanApplicationStatusStatusFundingInProgress, LoanApplicationStatusStatusFunded:
		return true
	}
	return false
}

type LendingApplicationSubmitParams struct {
	LoanAmount          param.Field[float64]                                   `json:"loanAmount,required"`
	LoanPurpose         param.Field[string]                                    `json:"loanPurpose,required"`
	RepaymentTermMonths param.Field[int64]                                     `json:"repaymentTermMonths,required"`
	AdditionalNotes     param.Field[string]                                    `json:"additionalNotes"`
	CoApplicant         param.Field[LendingApplicationSubmitParamsCoApplicant] `json:"coApplicant"`
}

func (r LendingApplicationSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LendingApplicationSubmitParamsCoApplicant struct {
	Email  param.Field[string]  `json:"email" format:"email"`
	Income param.Field[float64] `json:"income"`
	Name   param.Field[string]  `json:"name"`
}

func (r LendingApplicationSubmitParamsCoApplicant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
