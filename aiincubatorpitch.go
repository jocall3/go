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

// AIIncubatorPitchService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIIncubatorPitchService] method instead.
type AIIncubatorPitchService struct {
	Options []option.RequestOption
}

// NewAIIncubatorPitchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIIncubatorPitchService(opts ...option.RequestOption) (r *AIIncubatorPitchService) {
	r = &AIIncubatorPitchService{}
	r.Options = opts
	return
}

// Retrieves the granular AI-driven analysis, strategic feedback, market validation
// results, and any outstanding questions from Quantum Weaver for a specific
// business pitch.
func (r *AIIncubatorPitchService) GetDetails(ctx context.Context, pitchID string, opts ...option.RequestOption) (res *AIIncubatorPitchGetDetailsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if pitchID == "" {
		err = errors.New("missing required pitchId parameter")
		return
	}
	path := fmt.Sprintf("ai/incubator/pitch/%s/details", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits a detailed business plan to the Quantum Weaver AI for rigorous analysis,
// market validation, and seed funding consideration. This initiates the AI-driven
// incubation journey, aiming to transform innovative ideas into commercially
// successful ventures.
func (r *AIIncubatorPitchService) Submit(ctx context.Context, body AIIncubatorPitchSubmitParams, opts ...option.RequestOption) (res *QuantumWeaverState, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/incubator/pitch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows the entrepreneur to respond to specific questions or provide additional
// details requested by Quantum Weaver, moving the pitch forward in the incubation
// process.
func (r *AIIncubatorPitchService) SubmitFeedback(ctx context.Context, pitchID string, body AIIncubatorPitchSubmitFeedbackParams, opts ...option.RequestOption) (res *QuantumWeaverState, err error) {
	opts = slices.Concat(r.Options, opts)
	if pitchID == "" {
		err = errors.New("missing required pitchId parameter")
		return
	}
	path := fmt.Sprintf("ai/incubator/pitch/%s/feedback", pitchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type QuantumWeaverState struct {
	EstimatedFundingOffer float64                      `json:"estimatedFundingOffer"`
	FeedbackSummary       string                       `json:"feedbackSummary"`
	LastUpdated           time.Time                    `json:"lastUpdated" format:"date-time"`
	NextSteps             string                       `json:"nextSteps"`
	PitchID               string                       `json:"pitchId"`
	Questions             []QuantumWeaverStateQuestion `json:"questions"`
	Stage                 QuantumWeaverStateStage      `json:"stage"`
	StatusMessage         string                       `json:"statusMessage"`
	JSON                  quantumWeaverStateJSON       `json:"-"`
}

// quantumWeaverStateJSON contains the JSON metadata for the struct
// [QuantumWeaverState]
type quantumWeaverStateJSON struct {
	EstimatedFundingOffer apijson.Field
	FeedbackSummary       apijson.Field
	LastUpdated           apijson.Field
	NextSteps             apijson.Field
	PitchID               apijson.Field
	Questions             apijson.Field
	Stage                 apijson.Field
	StatusMessage         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *QuantumWeaverState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quantumWeaverStateJSON) RawJSON() string {
	return r.raw
}

type QuantumWeaverStateQuestion struct {
	ID         string                         `json:"id"`
	Category   string                         `json:"category"`
	IsRequired bool                           `json:"isRequired"`
	Question   string                         `json:"question"`
	JSON       quantumWeaverStateQuestionJSON `json:"-"`
}

// quantumWeaverStateQuestionJSON contains the JSON metadata for the struct
// [QuantumWeaverStateQuestion]
type quantumWeaverStateQuestionJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	IsRequired  apijson.Field
	Question    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuantumWeaverStateQuestion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quantumWeaverStateQuestionJSON) RawJSON() string {
	return r.raw
}

type QuantumWeaverStateStage string

const (
	QuantumWeaverStateStageSubmitted          QuantumWeaverStateStage = "submitted"
	QuantumWeaverStateStageInitialReview      QuantumWeaverStateStage = "initial_review"
	QuantumWeaverStateStageAIAnalysis         QuantumWeaverStateStage = "ai_analysis"
	QuantumWeaverStateStageFeedbackRequired   QuantumWeaverStateStage = "feedback_required"
	QuantumWeaverStateStageTestPhase          QuantumWeaverStateStage = "test_phase"
	QuantumWeaverStateStageFinalReview        QuantumWeaverStateStage = "final_review"
	QuantumWeaverStateStageApprovedForFunding QuantumWeaverStateStage = "approved_for_funding"
	QuantumWeaverStateStageRejected           QuantumWeaverStateStage = "rejected"
	QuantumWeaverStateStageIncubatedGraduated QuantumWeaverStateStage = "incubated_graduated"
)

func (r QuantumWeaverStateStage) IsKnown() bool {
	switch r {
	case QuantumWeaverStateStageSubmitted, QuantumWeaverStateStageInitialReview, QuantumWeaverStateStageAIAnalysis, QuantumWeaverStateStageFeedbackRequired, QuantumWeaverStateStageTestPhase, QuantumWeaverStateStageFinalReview, QuantumWeaverStateStageApprovedForFunding, QuantumWeaverStateStageRejected, QuantumWeaverStateStageIncubatedGraduated:
		return true
	}
	return false
}

type AIIncubatorPitchGetDetailsResponse struct {
	AICoachingPlan     AIIncubatorPitchGetDetailsResponseAICoachingPlan   `json:"aiCoachingPlan"`
	AIFinancialModel   AIIncubatorPitchGetDetailsResponseAIFinancialModel `json:"aiFinancialModel"`
	AIMarketAnalysis   AIIncubatorPitchGetDetailsResponseAIMarketAnalysis `json:"aiMarketAnalysis"`
	AIRiskAssessment   AIIncubatorPitchGetDetailsResponseAIRiskAssessment `json:"aiRiskAssessment"`
	InvestorMatchScore float64                                            `json:"investorMatchScore"`
	JSON               aiIncubatorPitchGetDetailsResponseJSON             `json:"-"`
	QuantumWeaverState
}

// aiIncubatorPitchGetDetailsResponseJSON contains the JSON metadata for the struct
// [AIIncubatorPitchGetDetailsResponse]
type aiIncubatorPitchGetDetailsResponseJSON struct {
	AICoachingPlan     apijson.Field
	AIFinancialModel   apijson.Field
	AIMarketAnalysis   apijson.Field
	AIRiskAssessment   apijson.Field
	InvestorMatchScore apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAICoachingPlan struct {
	Steps   []AIIncubatorPitchGetDetailsResponseAICoachingPlanStep `json:"steps"`
	Summary string                                                 `json:"summary"`
	Title   string                                                 `json:"title"`
	JSON    aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON   `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON contains the JSON metadata
// for the struct [AIIncubatorPitchGetDetailsResponseAICoachingPlan]
type aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON struct {
	Steps       apijson.Field
	Summary     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAICoachingPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAICoachingPlanJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAICoachingPlanStep struct {
	Description string                                                          `json:"description"`
	Resources   []AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource `json:"resources"`
	Status      AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus     `json:"status"`
	Timeline    string                                                          `json:"timeline"`
	Title       string                                                          `json:"title"`
	JSON        aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON        `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAICoachingPlanStep]
type aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON struct {
	Description apijson.Field
	Resources   apijson.Field
	Status      apijson.Field
	Timeline    apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAICoachingPlanStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAICoachingPlanStepJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource struct {
	Name string                                                            `json:"name"`
	URL  string                                                            `json:"url" format:"uri"`
	JSON aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON contains the
// JSON metadata for the struct
// [AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource]
type aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON struct {
	Name        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAICoachingPlanStepsResourceJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus string

const (
	AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusPending    AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus = "pending"
	AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusInProgress AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus = "in_progress"
	AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusCompleted  AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus = "completed"
)

func (r AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatus) IsKnown() bool {
	switch r {
	case AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusPending, AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusInProgress, AIIncubatorPitchGetDetailsResponseAICoachingPlanStepsStatusCompleted:
		return true
	}
	return false
}

type AIIncubatorPitchGetDetailsResponseAIFinancialModel struct {
	BreakevenPoint        string                                                 `json:"breakevenPoint"`
	CapitalRequirements   float64                                                `json:"capitalRequirements"`
	CostStructureAnalysis interface{}                                            `json:"costStructureAnalysis"`
	RevenueBreakdown      interface{}                                            `json:"revenueBreakdown"`
	SensitivityAnalysis   []interface{}                                          `json:"sensitivityAnalysis"`
	JSON                  aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAIFinancialModel]
type aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON struct {
	BreakevenPoint        apijson.Field
	CapitalRequirements   apijson.Field
	CostStructureAnalysis apijson.Field
	RevenueBreakdown      apijson.Field
	SensitivityAnalysis   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIFinancialModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIFinancialModelJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAIMarketAnalysis struct {
	CompetitiveAdvantages []string                                               `json:"competitiveAdvantages"`
	GrowthOpportunities   string                                                 `json:"growthOpportunities"`
	RiskFactors           string                                                 `json:"riskFactors"`
	TargetMarketSize      string                                                 `json:"targetMarketSize"`
	JSON                  aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAIMarketAnalysis]
type aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON struct {
	CompetitiveAdvantages apijson.Field
	GrowthOpportunities   apijson.Field
	RiskFactors           apijson.Field
	TargetMarketSize      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIMarketAnalysis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIMarketAnalysisJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchGetDetailsResponseAIRiskAssessment struct {
	MarketRisk    string                                                 `json:"marketRisk"`
	TeamRisk      string                                                 `json:"teamRisk"`
	TechnicalRisk string                                                 `json:"technicalRisk"`
	JSON          aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON `json:"-"`
}

// aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON contains the JSON
// metadata for the struct [AIIncubatorPitchGetDetailsResponseAIRiskAssessment]
type aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON struct {
	MarketRisk    apijson.Field
	TeamRisk      apijson.Field
	TechnicalRisk apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AIIncubatorPitchGetDetailsResponseAIRiskAssessment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiIncubatorPitchGetDetailsResponseAIRiskAssessmentJSON) RawJSON() string {
	return r.raw
}

type AIIncubatorPitchSubmitParams struct {
	BusinessPlan         param.Field[string]                                           `json:"businessPlan,required"`
	FinancialProjections param.Field[AIIncubatorPitchSubmitParamsFinancialProjections] `json:"financialProjections,required"`
	FoundingTeam         param.Field[[]AIIncubatorPitchSubmitParamsFoundingTeam]       `json:"foundingTeam,required"`
	MarketOpportunity    param.Field[string]                                           `json:"marketOpportunity,required"`
}

func (r AIIncubatorPitchSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorPitchSubmitParamsFinancialProjections struct {
	ProfitabilityEstimate param.Field[string]    `json:"profitabilityEstimate"`
	ProjectionYears       param.Field[int64]     `json:"projectionYears"`
	RevenueForecast       param.Field[[]float64] `json:"revenueForecast"`
	SeedRoundAmount       param.Field[float64]   `json:"seedRoundAmount"`
	ValuationPreMoney     param.Field[float64]   `json:"valuationPreMoney"`
}

func (r AIIncubatorPitchSubmitParamsFinancialProjections) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorPitchSubmitParamsFoundingTeam struct {
	Experience param.Field[string] `json:"experience"`
	Name       param.Field[string] `json:"name"`
	Role       param.Field[string] `json:"role"`
}

func (r AIIncubatorPitchSubmitParamsFoundingTeam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorPitchSubmitFeedbackParams struct {
	Answers  param.Field[[]AIIncubatorPitchSubmitFeedbackParamsAnswer] `json:"answers"`
	Feedback param.Field[string]                                       `json:"feedback"`
}

func (r AIIncubatorPitchSubmitFeedbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIIncubatorPitchSubmitFeedbackParamsAnswer struct {
	Answer     param.Field[string] `json:"answer,required"`
	QuestionID param.Field[string] `json:"questionId,required"`
}

func (r AIIncubatorPitchSubmitFeedbackParamsAnswer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
