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
		Email:    jamesburvelocallaghaniiicitibankdemobusinessinc.F("quantum.visionary@demobank.com"),
		Password: jamesburvelocallaghaniiicitibankdemobusinessinc.F("YourSecurePassword123"),
		MfaCode:  jamesburvelocallaghaniiicitibankdemobusinessinc.F("mfaCode"),
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
		Email:    jamesburvelocallaghaniiicitibankdemobusinessinc.F("alice.w@example.com"),
		Name:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("Alice Wonderland"),
		Password: jamesburvelocallaghaniiicitibankdemobusinessinc.F("SecureP@ssw0rd2024!"),
		Phone:    jamesburvelocallaghaniiicitibankdemobusinessinc.F("+1-555-987-6543"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
