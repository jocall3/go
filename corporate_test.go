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
		Country:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("US"),
		EntityType: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporatePerformSanctionScreeningParamsEntityTypeIndividual),
		Name:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("John Doe"),
		Address: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AddressParam{
			City:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Anytown"),
			Country: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USA"),
			State:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("CA"),
			Street:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("123 Main St"),
			Zip:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("90210"),
		}),
		DateOfBirth:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("1970-01-01"),
		IdentificationNumber: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
