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
		Amount: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](5000),
		Beneficiary: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentInternationalInitiateParamsBeneficiary{
			Address:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Hauptstrasse 1, 10115 Berlin, Germany"),
			BankName:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Deutsche Bank"),
			Name:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Maria Schmidt"),
			AccountNumber: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
			Iban:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("DE89370400440532013000"),
			RoutingNumber: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
			SwiftBic:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("DEUTDEFF"),
		}),
		Purpose:         jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Vendor payment for Q2 services."),
		SourceAccountID: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("acc_chase_checking_4567"),
		SourceCurrency:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USD"),
		TargetCurrency:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("EUR"),
		FxRateLock:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
		FxRateProvider:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.PaymentInternationalInitiateParamsFxRateProviderProprietaryAI),
		Reference:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
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
