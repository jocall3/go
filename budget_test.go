// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/1231-go"
	"github.com/stainless-sdks/1231-go/internal/testutil"
	"github.com/stainless-sdks/1231-go/option"
)

func TestBudgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.New(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetNewParams{
		EndDate:        jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
		Name:           jamesburvelocallaghaniiicitibankdemobusinessinc.F("September Living Expenses"),
		Period:         jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetNewParamsPeriodMonthly),
		StartDate:      jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
		TotalAmount:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(2800.000000),
		AIAutoPopulate: jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
		AlertThreshold: jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(75)),
		Categories: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetNewParamsCategory{{
			Allocated: jamesburvelocallaghaniiicitibankdemobusinessinc.F(1500.000000),
			Name:      jamesburvelocallaghaniiicitibankdemobusinessinc.F("Rent"),
		}, {
			Allocated: jamesburvelocallaghaniiicitibankdemobusinessinc.F(400.000000),
			Name:      jamesburvelocallaghaniiicitibankdemobusinessinc.F("Groceries"),
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

func TestBudgetGet(t *testing.T) {
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
	_, err := client.Budgets.Get(context.TODO(), "budget_monthly_aug")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.Update(
		context.TODO(),
		"budget_monthly_aug",
		jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetUpdateParams{
			AlertThreshold: jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(85)),
			Name:           jamesburvelocallaghaniiicitibankdemobusinessinc.F("name"),
			TotalAmount:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(3200.000000),
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

func TestBudgetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Budgets.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetListParams{
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

func TestBudgetDelete(t *testing.T) {
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
	err := client.Budgets.Delete(context.TODO(), "budget_monthly_aug")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
