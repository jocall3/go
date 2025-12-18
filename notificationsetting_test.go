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

func TestNotificationSettingGet(t *testing.T) {
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
	_, err := client.Notifications.Settings.Get(context.TODO())
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNotificationSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Notifications.Settings.Update(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.NotificationSettingUpdateParams{
		ChannelPreferences: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.NotificationSettingUpdateParamsChannelPreferences{
			Email: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			InApp: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			Push:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			SMS:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
		}),
		EventPreferences: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.NotificationSettingUpdateParamsEventPreferences{
			AIInsights:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			BudgetAlerts:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			PromotionalOffers: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](false),
			SecurityAlerts:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			TransactionAlerts: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
		}),
		QuietHours: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.NotificationSettingUpdateParamsQuietHours{
			Enabled:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			EndTime:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("08:00"),
			StartTime: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("22:00"),
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
