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

func TestCorporateAnomalyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Anomalies.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyListParams{
		EndDate:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-12-31"),
		EntityType: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyListParamsEntityTypeTransaction),
		Limit:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
		Offset:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
		Severity:   jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyListParamsSeverityCritical),
		StartDate:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-01-01"),
		Status:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyListParamsStatusNew),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateAnomalyUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Anomalies.UpdateStatus(
		context.TODO(),
		"anom_risk-2024-07-21-D1E2F3",
		jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyUpdateStatusParams{
			Status:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateAnomalyUpdateStatusParamsStatusResolved),
			ResolutionNotes: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Confirmed legitimate transaction after contacting vendor. Marked as resolved."),
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
