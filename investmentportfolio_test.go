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

func TestInvestmentPortfolioNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.New(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioNewParams{
		Currency:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USD"),
		InitialInvestment: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](10000),
		Name:              jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("My First Growth Portfolio"),
		RiskTolerance:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioNewParamsRiskToleranceConservative),
		Type:              jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioNewParamsTypeDiversified),
		AIAutoAllocate:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
		LinkedAccountID:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("acc_chase_checking_4567"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentPortfolioGet(t *testing.T) {
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
	_, err := client.Investments.Portfolios.Get(context.TODO(), "portfolio_equity_growth")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentPortfolioUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.Update(
		context.TODO(),
		"portfolio_equity_growth",
		jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioUpdateParams{
			AIRebalancingFrequency: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioUpdateParamsAIRebalancingFrequencyQuarterly),
			Name:                   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Revised Growth Portfolio"),
			RiskTolerance:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioUpdateParamsRiskToleranceConservative),
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

func TestInvestmentPortfolioListWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioListParams{
		Limit:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
		Offset: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInvestmentPortfolioRebalanceWithOptionalParams(t *testing.T) {
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
	_, err := client.Investments.Portfolios.Rebalance(
		context.TODO(),
		"portfolio_equity_growth",
		jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioRebalanceParams{
			TargetRiskTolerance:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioRebalanceParamsTargetRiskToleranceConservative),
			ConfirmationRequired: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			DryRun:               jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
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
