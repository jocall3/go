// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/1231-go"
	"github.com/stainless-sdks/1231-go/internal/testutil"
	"github.com/stainless-sdks/1231-go/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	user, err := client.Users.Register(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.UserRegisterParams{
		Email:    jamesburvelocallaghaniiicitibankdemobusinessinc.F("alice.w@example.com"),
		Name:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("Alice Wonderland"),
		Password: jamesburvelocallaghaniiicitibankdemobusinessinc.F("SecureP@ssw0rd2024!"),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", user.ID)
}
