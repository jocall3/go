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

func TestPaymentFxConvertWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Fx.Convert(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentFxConvertParams{
		SourceAccountID: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("acc_chase_checking_4567"),
		SourceAmount:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](1000),
		SourceCurrency:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USD"),
		TargetCurrency:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("EUR"),
		FxRateLock:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
		TargetAccountID: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("acc_euro_savings_9876"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentFxGetRatesWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Fx.GetRates(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentFxGetRatesParams{
		BaseCurrency:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USD"),
		TargetCurrency: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("EUR"),
		ForecastDays:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](7),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
