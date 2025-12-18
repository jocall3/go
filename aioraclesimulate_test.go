// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/1231-go"
	"github.com/stainless-sdks/1231-go/internal/testutil"
	"github.com/stainless-sdks/1231-go/option"
)

func TestAIOracleSimulateRunAdvancedWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.AI.Oracle.Simulate.RunAdvanced(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParams{
		Prompt: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Evaluate the long-term impact of a sudden job loss combined with a variable market downturn, analyzing worst-case and best-case recovery scenarios over a decade."),
		Scenarios: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsScenario{{
			DurationYears: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](10),
			Events: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsScenariosEvent{{
				Details: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{
					"durationMonths":       6,
					"severanceAmount":      10000,
					"unemploymentBenefits": 2000,
				}),
				Type: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsScenariosEventsTypeJobLoss),
			}, {
				Details: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{
					"impactPercentage": 0.15,
					"recoveryYears":    3,
				}),
				Type: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsScenariosEventsTypeMarketDownturn),
			}}),
			Name: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Job Loss & Mild Market Recovery"),
			SensitivityAnalysisParams: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam{{
				Max:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](0.07),
				Min:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](0.03),
				ParamName: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("marketRecoveryRate"),
				Step:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](0.01),
			}}),
		}}),
		GlobalEconomicFactors: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsGlobalEconomicFactors{
			InflationRate:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](0.03),
			InterestRateBaseline: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](0.05),
		}),
		PersonalAssumptions: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsPersonalAssumptions{
			AnnualSavingsRate: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](0.15),
			RiskTolerance:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsPersonalAssumptionsRiskToleranceAggressive),
		}),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIOracleSimulateRunStandardWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := jamesburvelocallaghaniiicitibankdemobusinessinc.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.AI.Oracle.Simulate.RunStandard(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunStandardParams{
		Prompt: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("What if I invest an additional $1,000 per month into my aggressive growth portfolio for the next 5 years?"),
		Parameters: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{
			"durationYears":           5,
			"monthlyInvestmentAmount": 1000,
			"riskTolerance":           "aggressive",
		}),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
