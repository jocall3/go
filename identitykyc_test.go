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

func TestIdentityKYCGetStatus(t *testing.T) {
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
	_, err := client.Identity.KYC.GetStatus(context.TODO())
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIdentityKYCSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Identity.KYC.Submit(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.IdentityKYCSubmitParams{
		CountryOfIssue:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("US"),
		DocumentNumber:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("ABC12345"),
		DocumentType:        jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.IdentityKYCSubmitParamsDocumentTypeDriversLicense),
		ExpirationDate:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2030-01-01"),
		IssueDate:           jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2020-01-01"),
		AdditionalDocuments: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{map[string]interface{}{}}),
		DocumentBackImage:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("base64encoded_image_of_drivers_license_back"),
		DocumentFrontImage:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("base64encoded_image_of_drivers_license_front"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
