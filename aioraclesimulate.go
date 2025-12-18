// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// AIOracleSimulateService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIOracleSimulateService] method instead.
type AIOracleSimulateService struct {
	Options []option.RequestOption
}

// NewAIOracleSimulateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIOracleSimulateService(opts ...option.RequestOption) (r *AIOracleSimulateService) {
	r = &AIOracleSimulateService{}
	r.Options = opts
	return
}

// Engages the Quantum Oracle for highly complex, multi-variable simulations,
// allowing precise control over numerous financial parameters, market conditions,
// and personal events to generate deep, predictive insights and sensitivity
// analysis.
func (r *AIOracleSimulateService) RunAdvanced(ctx context.Context, body AIOracleSimulateRunAdvancedParams, opts ...option.RequestOption) (res *AdvancedSimulationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulate/advanced"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submits a hypothetical scenario to the Quantum Oracle AI for standard financial
// impact analysis. The AI simulates the effect on the user's current financial
// state and provides a summary.
func (r *AIOracleSimulateService) RunStandard(ctx context.Context, body AIOracleSimulateRunStandardParams, opts ...option.RequestOption) (res *SimulationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/oracle/simulate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AdvancedSimulationResponse struct {
	// Discriminator field for oneOf.
	SimulationType           AdvancedSimulationResponseSimulationType   `json:"simulationType,required"`
	OverallSummary           string                                     `json:"overallSummary"`
	ScenarioResults          []AdvancedSimulationResponseScenarioResult `json:"scenarioResults"`
	SimulationID             string                                     `json:"simulationId"`
	StrategicRecommendations []AIInsight                                `json:"strategicRecommendations"`
	JSON                     advancedSimulationResponseJSON             `json:"-"`
}

// advancedSimulationResponseJSON contains the JSON metadata for the struct
// [AdvancedSimulationResponse]
type advancedSimulationResponseJSON struct {
	SimulationType           apijson.Field
	OverallSummary           apijson.Field
	ScenarioResults          apijson.Field
	SimulationID             apijson.Field
	StrategicRecommendations apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AdvancedSimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseJSON) RawJSON() string {
	return r.raw
}

func (r AdvancedSimulationResponse) implementsAIOracleSimulationGetResponse() {}

// Discriminator field for oneOf.
type AdvancedSimulationResponseSimulationType string

const (
	AdvancedSimulationResponseSimulationTypeAdvanced AdvancedSimulationResponseSimulationType = "advanced"
)

func (r AdvancedSimulationResponseSimulationType) IsKnown() bool {
	switch r {
	case AdvancedSimulationResponseSimulationTypeAdvanced:
		return true
	}
	return false
}

type AdvancedSimulationResponseScenarioResult struct {
	FinalNetWorthProjected    float64                                                             `json:"finalNetWorthProjected"`
	LiquidityMetrics          AdvancedSimulationResponseScenarioResultsLiquidityMetrics           `json:"liquidityMetrics"`
	NarrativeSummary          string                                                              `json:"narrativeSummary"`
	ScenarioName              string                                                              `json:"scenarioName"`
	SensitivityAnalysisGraphs []AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph `json:"sensitivityAnalysisGraphs"`
	JSON                      advancedSimulationResponseScenarioResultJSON                        `json:"-"`
}

// advancedSimulationResponseScenarioResultJSON contains the JSON metadata for the
// struct [AdvancedSimulationResponseScenarioResult]
type advancedSimulationResponseScenarioResultJSON struct {
	FinalNetWorthProjected    apijson.Field
	LiquidityMetrics          apijson.Field
	NarrativeSummary          apijson.Field
	ScenarioName              apijson.Field
	SensitivityAnalysisGraphs apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultJSON) RawJSON() string {
	return r.raw
}

type AdvancedSimulationResponseScenarioResultsLiquidityMetrics struct {
	MinCashBalance     float64                                                       `json:"minCashBalance"`
	RecoveryTimeMonths int64                                                         `json:"recoveryTimeMonths"`
	JSON               advancedSimulationResponseScenarioResultsLiquidityMetricsJSON `json:"-"`
}

// advancedSimulationResponseScenarioResultsLiquidityMetricsJSON contains the JSON
// metadata for the struct
// [AdvancedSimulationResponseScenarioResultsLiquidityMetrics]
type advancedSimulationResponseScenarioResultsLiquidityMetricsJSON struct {
	MinCashBalance     apijson.Field
	RecoveryTimeMonths apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResultsLiquidityMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultsLiquidityMetricsJSON) RawJSON() string {
	return r.raw
}

type AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph struct {
	Data      []AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData `json:"data"`
	ParamName string                                                                   `json:"paramName"`
	JSON      advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON    `json:"-"`
}

// advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON contains
// the JSON metadata for the struct
// [AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph]
type advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON struct {
	Data        apijson.Field
	ParamName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraph) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphJSON) RawJSON() string {
	return r.raw
}

type AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData struct {
	OutcomeValue float64                                                                    `json:"outcomeValue"`
	ParamValue   float64                                                                    `json:"paramValue"`
	JSON         advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON `json:"-"`
}

// advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON
// contains the JSON metadata for the struct
// [AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData]
type advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON struct {
	OutcomeValue apijson.Field
	ParamValue   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AdvancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedSimulationResponseScenarioResultsSensitivityAnalysisGraphsDataJSON) RawJSON() string {
	return r.raw
}

type SimulationResponse struct {
	// Discriminator field for oneOf.
	SimulationType   SimulationResponseSimulationType   `json:"simulationType,required"`
	KeyImpacts       []SimulationResponseKeyImpact      `json:"keyImpacts"`
	NarrativeSummary string                             `json:"narrativeSummary"`
	Recommendations  []SimulationResponseRecommendation `json:"recommendations"`
	RiskAnalysis     SimulationResponseRiskAnalysis     `json:"riskAnalysis"`
	SimulationID     string                             `json:"simulationId"`
	JSON             simulationResponseJSON             `json:"-"`
}

// simulationResponseJSON contains the JSON metadata for the struct
// [SimulationResponse]
type simulationResponseJSON struct {
	SimulationType   apijson.Field
	KeyImpacts       apijson.Field
	NarrativeSummary apijson.Field
	Recommendations  apijson.Field
	RiskAnalysis     apijson.Field
	SimulationID     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseJSON) RawJSON() string {
	return r.raw
}

func (r SimulationResponse) implementsAIOracleSimulationGetResponse() {}

// Discriminator field for oneOf.
type SimulationResponseSimulationType string

const (
	SimulationResponseSimulationTypeStandard SimulationResponseSimulationType = "standard"
)

func (r SimulationResponseSimulationType) IsKnown() bool {
	switch r {
	case SimulationResponseSimulationTypeStandard:
		return true
	}
	return false
}

type SimulationResponseKeyImpact struct {
	Metric   string                               `json:"metric"`
	Severity SimulationResponseKeyImpactsSeverity `json:"severity"`
	Value    string                               `json:"value"`
	JSON     simulationResponseKeyImpactJSON      `json:"-"`
}

// simulationResponseKeyImpactJSON contains the JSON metadata for the struct
// [SimulationResponseKeyImpact]
type simulationResponseKeyImpactJSON struct {
	Metric      apijson.Field
	Severity    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SimulationResponseKeyImpact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseKeyImpactJSON) RawJSON() string {
	return r.raw
}

type SimulationResponseKeyImpactsSeverity string

const (
	SimulationResponseKeyImpactsSeverityLow    SimulationResponseKeyImpactsSeverity = "low"
	SimulationResponseKeyImpactsSeverityMedium SimulationResponseKeyImpactsSeverity = "medium"
	SimulationResponseKeyImpactsSeverityHigh   SimulationResponseKeyImpactsSeverity = "high"
)

func (r SimulationResponseKeyImpactsSeverity) IsKnown() bool {
	switch r {
	case SimulationResponseKeyImpactsSeverityLow, SimulationResponseKeyImpactsSeverityMedium, SimulationResponseKeyImpactsSeverityHigh:
		return true
	}
	return false
}

type SimulationResponseRecommendation struct {
	ActionTrigger string                               `json:"actionTrigger"`
	Description   string                               `json:"description"`
	Title         string                               `json:"title"`
	JSON          simulationResponseRecommendationJSON `json:"-"`
}

// simulationResponseRecommendationJSON contains the JSON metadata for the struct
// [SimulationResponseRecommendation]
type simulationResponseRecommendationJSON struct {
	ActionTrigger apijson.Field
	Description   apijson.Field
	Title         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SimulationResponseRecommendation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseRecommendationJSON) RawJSON() string {
	return r.raw
}

type SimulationResponseRiskAnalysis struct {
	MaxDrawdown     float64                            `json:"maxDrawdown"`
	VolatilityIndex float64                            `json:"volatilityIndex"`
	JSON            simulationResponseRiskAnalysisJSON `json:"-"`
}

// simulationResponseRiskAnalysisJSON contains the JSON metadata for the struct
// [SimulationResponseRiskAnalysis]
type simulationResponseRiskAnalysisJSON struct {
	MaxDrawdown     apijson.Field
	VolatilityIndex apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SimulationResponseRiskAnalysis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r simulationResponseRiskAnalysisJSON) RawJSON() string {
	return r.raw
}

type AIOracleSimulateRunAdvancedParams struct {
	Prompt    param.Field[string]                                      `json:"prompt,required"`
	Scenarios param.Field[[]AIOracleSimulateRunAdvancedParamsScenario] `json:"scenarios,required"`
}

func (r AIOracleSimulateRunAdvancedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsScenario struct {
	DurationYears             param.Field[int64]                                                                `json:"durationYears"`
	Events                    param.Field[[]AIOracleSimulateRunAdvancedParamsScenariosEvent]                    `json:"events"`
	Name                      param.Field[string]                                                               `json:"name"`
	SensitivityAnalysisParams param.Field[[]AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam] `json:"sensitivityAnalysisParams"`
}

func (r AIOracleSimulateRunAdvancedParamsScenario) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsScenariosEvent struct {
	Details param.Field[map[string]interface{}] `json:"details"`
	Type    param.Field[string]                 `json:"type"`
}

func (r AIOracleSimulateRunAdvancedParamsScenariosEvent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam struct {
	Max       param.Field[float64] `json:"max"`
	Min       param.Field[float64] `json:"min"`
	ParamName param.Field[string]  `json:"paramName"`
	Step      param.Field[float64] `json:"step"`
}

func (r AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIOracleSimulateRunStandardParams struct {
	Prompt     param.Field[string]                 `json:"prompt,required"`
	Parameters param.Field[map[string]interface{}] `json:"parameters"`
}

func (r AIOracleSimulateRunStandardParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
