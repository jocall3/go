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

func TestUserMeGet(t *testing.T) {
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
	_, err := client.Users.Me.Get(context.TODO())
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserMeUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Me.Update(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.UserMeUpdateParams{
		Address: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AddressParam{
			City:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Anytown"),
			Country: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("USA"),
			State:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("CA"),
			Street:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("123 Main St"),
			Zip:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("90210"),
		}),
		Name:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Quantum Visionary Pro"),
		Phone: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("+1-555-999-0000"),
		Preferences: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.UserPreferencesParam{
			AIInteractionMode:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.UserPreferencesAIInteractionModeBalanced),
			DataSharingConsent: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			NotificationChannels: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.UserPreferencesNotificationChannelsParam{
				Email: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				InApp: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				Push:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				SMS:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](false),
			}),
			PreferredLanguage:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("en-US"),
			Theme:               jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Dark-Quantum"),
			TransactionGrouping: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.UserPreferencesTransactionGroupingCategory),
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
