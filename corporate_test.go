// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/jocall3/1231-go"
	"github.com/jocall3/1231-go/internal/testutil"
	"github.com/jocall3/1231-go/option"
)

func TestCorporatePerformSanctionScreeningWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.PerformSanctionScreening(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.CorporatePerformSanctionScreeningParams{
		Country:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("US"),
		EntityType:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporatePerformSanctionScreeningParamsEntityTypeIndividual),
		Name:        jamesburvelocallaghaniiicitibankdemobusinessinc.F("John Doe"),
		DateOfBirth: jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
