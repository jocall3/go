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

func TestPaymentInternationalInitiateWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.International.Initiate(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentInternationalInitiateParams{
		Amount: jamesburvelocallaghaniiicitibankdemobusinessinc.F(5000.000000),
		Beneficiary: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentInternationalInitiateParamsBeneficiary{
			Address:  jamesburvelocallaghaniiicitibankdemobusinessinc.F("Hauptstrasse 1, 10115 Berlin, Germany"),
			BankName: jamesburvelocallaghaniiicitibankdemobusinessinc.F("Deutsche Bank"),
			Iban:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("DE89370400440532013000"),
			Name:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("Maria Schmidt"),
			SwiftBic: jamesburvelocallaghaniiicitibankdemobusinessinc.F("DEUTDEFF"),
		}),
		SourceAccountID: jamesburvelocallaghaniiicitibankdemobusinessinc.F("acc_chase_checking_4567"),
		SourceCurrency:  jamesburvelocallaghaniiicitibankdemobusinessinc.F("USD"),
		TargetCurrency:  jamesburvelocallaghaniiicitibankdemobusinessinc.F("EUR"),
		FxRateLock:      jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
		FxRateProvider:  jamesburvelocallaghaniiicitibankdemobusinessinc.F("proprietary_ai"),
		Purpose:         jamesburvelocallaghaniiicitibankdemobusinessinc.F("Vendor payment for Q2 services."),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentInternationalGetStatus(t *testing.T) {
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
	_, err := client.Payments.International.GetStatus(context.TODO(), "int_pmt_xyz7890")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
