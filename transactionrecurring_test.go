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

func TestTransactionRecurringNew(t *testing.T) {
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
	_, err := client.Transactions.Recurring.New(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionRecurringNewParams{
		Amount:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](55.5),
		Category:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Health & Fitness"),
		Currency:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USD"),
		Description:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("New Gym Membership"),
		Frequency:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionRecurringNewParamsFrequencyMonthly),
		LinkedAccountID: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("acc_chase_checking_4567"),
		StartDate:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-09-01"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionRecurringListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Recurring.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionRecurringListParams{
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
