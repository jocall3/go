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

func TestLendingApplicationGet(t *testing.T) {
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
	_, err := client.Lending.Applications.Get(context.TODO(), "loan_app_creditflow-123")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLendingApplicationSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Lending.Applications.Submit(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.LendingApplicationSubmitParams{
		LoanAmount:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](10000),
		LoanPurpose:         jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.LendingApplicationSubmitParamsLoanPurposeHomeImprovement),
		RepaymentTermMonths: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](36),
		AdditionalNotes:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Funds needed to replace a broken HVAC system."),
		CoApplicant: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.LendingApplicationSubmitParamsCoApplicant{
			Email:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("jane.doe@example.com"),
			Income: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](75000),
			Name:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Jane Doe"),
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
