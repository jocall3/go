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

func TestAIIncubatorPitchGetDetails(t *testing.T) {
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
	_, err := client.AI.Incubator.Pitch.GetDetails(context.TODO(), "pitch_qw_synergychain-xyz")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIIncubatorPitchSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Incubator.Pitch.Submit(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.AIIncubatorPitchSubmitParams{
		BusinessPlan: jamesburvelocallaghaniiicitibankdemobusinessinc.F("Quantum-AI powered financial advisor platform leveraging neural networks for predictive analytics and hyper-personalized advice..."),
		FinancialProjections: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIIncubatorPitchSubmitParamsFinancialProjections{
			ProfitabilityEstimate: jamesburvelocallaghaniiicitibankdemobusinessinc.F("Achieve profitability within 18 months."),
			ProjectionYears:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(3)),
			RevenueForecast:       jamesburvelocallaghaniiicitibankdemobusinessinc.F([]float64{500000.000000, 2000000.000000, 6000000.000000}),
			SeedRoundAmount:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(2500000.000000),
			ValuationPreMoney:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(10000000.000000),
		}),
		FoundingTeam: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.AIIncubatorPitchSubmitParamsFoundingTeam{{
			Experience: jamesburvelocallaghaniiicitibankdemobusinessinc.F("15+ years in AI/ML, PhD in Quantum Computing, ex-Google Brain"),
			Name:       jamesburvelocallaghaniiicitibankdemobusinessinc.F("Dr. Eleanor Vance"),
			Role:       jamesburvelocallaghaniiicitibankdemobusinessinc.F("CEO & Lead AI Scientist"),
		}, {
			Experience: jamesburvelocallaghaniiicitibankdemobusinessinc.F("20+ years in Fintech, ex-Goldman Sachs"),
			Name:       jamesburvelocallaghaniiicitibankdemobusinessinc.F("Marcus Thorne"),
			Role:       jamesburvelocallaghaniiicitibankdemobusinessinc.F("COO & Finance Expert"),
		}}),
		MarketOpportunity: jamesburvelocallaghaniiicitibankdemobusinessinc.F("The booming digital finance market coupled with demand for truly personalized, AI-driven financial guidance presents a multi-billion dollar opportunity. Our unique quantum-AI approach provides unparalleled accuracy and foresight."),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIIncubatorPitchSubmitFeedbackWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Incubator.Pitch.SubmitFeedback(
		context.TODO(),
		"pitch_qw_synergychain-xyz",
		jamesburvelocallaghaniiicitibankdemobusinessinc.AIIncubatorPitchSubmitFeedbackParams{
			Answers: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.AIIncubatorPitchSubmitFeedbackParamsAnswer{{
				Answer:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("Our mitigation strategy includes dedicated R&D and new hires with specific expertise."),
				QuestionID: jamesburvelocallaghaniiicitibankdemobusinessinc.F("q_qa-team-001"),
			}, {
				Answer:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("Our CAC projections are based on pilot program results showing $500 per enterprise client with a conversion rate of 10% from trials."),
				QuestionID: jamesburvelocallaghaniiicitibankdemobusinessinc.F("q_qa-market-002"),
			}}),
			Feedback: jamesburvelocallaghaniiicitibankdemobusinessinc.F("Regarding the technical challenges, our team has allocated 3 months for R&D on quantum-resistant cryptography, mitigating the risk. We've also brought in Dr. Elena Petrova, a leading expert in secure multi-party computation."),
		},
	)
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
