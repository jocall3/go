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

func TestAIOracleSimulateRunAdvanced(t *testing.T) {
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
		Prompt: jamesburvelocallaghaniiicitibankdemobusinessinc.F("Evaluate the long-term impact of a sudden job loss combined with a variable market downturn, analyzing worst-case and best-case recovery scenarios over a decade."),
		Scenarios: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsScenario{{
			DurationYears: jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(10)),
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
			Name: jamesburvelocallaghaniiicitibankdemobusinessinc.F("Job Loss & Mild Market Recovery"),
			SensitivityAnalysisParams: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.AIOracleSimulateRunAdvancedParamsScenariosSensitivityAnalysisParam{{
				Max:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(0.070000),
				Min:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(0.030000),
				ParamName: jamesburvelocallaghaniiicitibankdemobusinessinc.F("marketRecoveryRate"),
				Step:      jamesburvelocallaghaniiicitibankdemobusinessinc.F(0.010000),
			}}),
		}}),
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
		Prompt: jamesburvelocallaghaniiicitibankdemobusinessinc.F("What if I invest an additional $1,000 per month into my aggressive growth portfolio for the next 5 years?"),
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
