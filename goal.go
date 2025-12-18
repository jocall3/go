// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/param"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// GoalService contains methods and other services that help with interacting with
// the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGoalService] method instead.
type GoalService struct {
	Options []option.RequestOption
}

// NewGoalService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGoalService(opts ...option.RequestOption) (r *GoalService) {
	r = &GoalService{}
	r.Options = opts
	return
}

// Creates a new long-term financial goal, with options for AI to generate a
// strategic plan to achieve it.
func (r *GoalService) New(ctx context.Context, body GoalNewParams, opts ...option.RequestOption) (res *FinancialGoal, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "goals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FinancialGoal struct {
	ID                      string              `json:"id"`
	AIStrategicPlan         string              `json:"aiStrategicPlan"`
	CurrentAmount           float64             `json:"currentAmount"`
	Name                    string              `json:"name"`
	ProjectedCompletionDate time.Time           `json:"projectedCompletionDate" format:"date"`
	Status                  FinancialGoalStatus `json:"status"`
	TargetAmount            float64             `json:"targetAmount"`
	TargetDate              time.Time           `json:"targetDate" format:"date"`
	JSON                    financialGoalJSON   `json:"-"`
}

// financialGoalJSON contains the JSON metadata for the struct [FinancialGoal]
type financialGoalJSON struct {
	ID                      apijson.Field
	AIStrategicPlan         apijson.Field
	CurrentAmount           apijson.Field
	Name                    apijson.Field
	ProjectedCompletionDate apijson.Field
	Status                  apijson.Field
	TargetAmount            apijson.Field
	TargetDate              apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *FinancialGoal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r financialGoalJSON) RawJSON() string {
	return r.raw
}

type FinancialGoalStatus string

const (
	FinancialGoalStatusActive    FinancialGoalStatus = "active"
	FinancialGoalStatusCompleted FinancialGoalStatus = "completed"
	FinancialGoalStatusArchived  FinancialGoalStatus = "archived"
)

func (r FinancialGoalStatus) IsKnown() bool {
	switch r {
	case FinancialGoalStatusActive, FinancialGoalStatusCompleted, FinancialGoalStatusArchived:
		return true
	}
	return false
}

type GoalNewParams struct {
	Name           param.Field[string]    `json:"name,required"`
	TargetAmount   param.Field[float64]   `json:"targetAmount,required"`
	TargetDate     param.Field[time.Time] `json:"targetDate,required" format:"date"`
	AIGeneratePlan param.Field[bool]      `json:"aiGeneratePlan"`
}

func (r GoalNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
