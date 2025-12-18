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
		Currency:          jamesburvelocallaghaniiicitibankdemobusinessinc.F("USD"),
		Name:              jamesburvelocallaghaniiicitibankdemobusinessinc.F("My First Growth Portfolio"),
		RiskTolerance:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioNewParamsRiskToleranceMedium),
		Type:              jamesburvelocallaghaniiicitibankdemobusinessinc.F("diversified"),
		AIAutoAllocate:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
		InitialInvestment: jamesburvelocallaghaniiicitibankdemobusinessinc.F(10000.000000),
		LinkedAccountID:   jamesburvelocallaghaniiicitibankdemobusinessinc.F("acc_chase_checking_4567"),
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
			Name:                   jamesburvelocallaghaniiicitibankdemobusinessinc.F("name"),
			RiskTolerance:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioUpdateParamsRiskToleranceMedium),
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
		Limit:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(1)),
		Offset: jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(0)),
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
			TargetRiskTolerance:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.InvestmentPortfolioRebalanceParamsTargetRiskToleranceMedium),
			ConfirmationRequired: jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
			DryRun:               jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
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
