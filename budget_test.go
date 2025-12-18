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
		EndDate:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-09-30"),
		Name:           jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("September Living Expenses"),
		Period:         jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetNewParamsPeriodMonthly),
		StartDate:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-09-01"),
		TotalAmount:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](2800),
		AIAutoPopulate: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
		AlertThreshold: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](75),
		Categories: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetNewParamsCategory{{
			Allocated: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](1500),
			Name:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Rent"),
		}, {
			Allocated: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](400),
			Name:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Groceries"),
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
			AlertThreshold: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](85),
			Categories: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetUpdateParamsCategory{{
				Allocated: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](550),
				Name:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Groceries"),
			}}),
			EndDate:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-08-31"),
			Name:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("August 2024 Revised Household Budget"),
			StartDate:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-08-01"),
			Status:      jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.BudgetUpdateParamsStatusActive),
			TotalAmount: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](3200),
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
