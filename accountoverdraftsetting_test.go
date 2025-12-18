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

func TestAccountOverdraftSettingGetOverdraftSettings(t *testing.T) {
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
	_, err := client.Accounts.OverdraftSettings.GetOverdraftSettings(context.TODO(), "acc_chase_checking_4567")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountOverdraftSettingUpdateOverdraftSettingsWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.OverdraftSettings.UpdateOverdraftSettings(
		context.TODO(),
		"acc_chase_checking_4567",
		jamesburvelocallaghaniiicitibankdemobusinessinc.AccountOverdraftSettingUpdateOverdraftSettingsParams{
			Enabled:                jamesburvelocallaghaniiicitibankdemobusinessinc.F(false),
			FeePreference:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceDeclineIfOverLimit),
			LinkedSavingsAccountID: jamesburvelocallaghaniiicitibankdemobusinessinc.F("linkedSavingsAccountId"),
			LinkToSavings:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(false),
			ProtectionLimit:        jamesburvelocallaghaniiicitibankdemobusinessinc.F(0.000000),
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
