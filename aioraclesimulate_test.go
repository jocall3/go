// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/jocall3/1231-go"
	"github.com/jocall3/1231-go/internal/testutil"
	"github.com/jocall3/1231-go/option"
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
				Details: jamesburvelocallaghaniiicitibankdemobusinessinc.F(map[string]interface{}{
					"durationMonths":       "bar",
					"severanceAmount":      "bar",
					"unemploymentBenefits": "bar",
				}),
				Type: jamesburvelocallaghaniiicitibankdemobusinessinc.F("job_loss"),
			}, {
				Details: jamesburvelocallaghaniiicitibankdemobusinessinc.F(map[string]interface{}{
					"impactPercentage": "bar",
					"recoveryYears":    "bar",
				}),
				Type: jamesburvelocallaghaniiicitibankdemobusinessinc.F("market_downturn"),
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
		Parameters: jamesburvelocallaghaniiicitibankdemobusinessinc.F(map[string]interface{}{
			"durationYears":           "bar",
			"monthlyInvestmentAmount": "bar",
			"riskTolerance":           "bar",
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
