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

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Login(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.UserLoginParams{
		Email:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("quantum.visionary@demobank.com"),
		Password: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("YourSecurePassword123"),
		MfaCode:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("123456"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserRegisterWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Register(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.UserRegisterParams{
		Email:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("alice.w@example.com"),
		Name:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Alice Wonderland"),
		Password: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("SecureP@ssw0rd2024!"),
		Address: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AddressParam{
			City:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Anytown"),
			Country: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USA"),
			State:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("CA"),
			Street:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("123 Main St"),
			Zip:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("90210"),
		}),
		DateOfBirth: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("1990-05-10"),
		Phone:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("+1-555-987-6543"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
