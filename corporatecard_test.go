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

func TestCorporateCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardListParams{
		Limit:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(20)),
		Offset: jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(0)),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateCardNewVirtualWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.NewVirtual(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardNewVirtualParams{
		Controls: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardControlsParam{
			AtmWithdrawals:               jamesburvelocallaghaniiicitibankdemobusinessinc.F(false),
			ContactlessPayments:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(false),
			DailyLimit:                   jamesburvelocallaghaniiicitibankdemobusinessinc.F(500.000000),
			InternationalTransactions:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(false),
			MerchantCategoryRestrictions: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"Advertising"}),
			MonthlyLimit:                 jamesburvelocallaghaniiicitibankdemobusinessinc.F(1000.000000),
			OnlineTransactions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
			SingleTransactionLimit:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(200.000000),
			VendorRestrictions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"Facebook Ads", "Google Ads"}),
		}),
		HolderName:           jamesburvelocallaghaniiicitibankdemobusinessinc.F("Marketing Campaign Q4"),
		AssociatedEmployeeID: jamesburvelocallaghaniiicitibankdemobusinessinc.F("emp_marketing_01"),
		ExpirationDate:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
		Purpose:              jamesburvelocallaghaniiicitibankdemobusinessinc.F("Online advertising for Q4 campaigns"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateCardFreeze(t *testing.T) {
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
	_, err := client.Corporate.Cards.Freeze(
		context.TODO(),
		"corp_card_xyz987654",
		jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardFreezeParams{
			Freeze: jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
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

func TestCorporateCardListTransactionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.ListTransactions(
		context.TODO(),
		"corp_card_xyz987654",
		jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardListTransactionsParams{
			EndDate:   jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
			Limit:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(20)),
			Offset:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(0)),
			StartDate: jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
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

func TestCorporateCardUpdateControlsWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Cards.UpdateControls(
		context.TODO(),
		"corp_card_xyz987654",
		jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardUpdateControlsParams{
			CorporateCardControls: jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateCardControlsParam{
				AtmWithdrawals:               jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
				ContactlessPayments:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
				DailyLimit:                   jamesburvelocallaghaniiicitibankdemobusinessinc.F(750.000000),
				InternationalTransactions:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
				MerchantCategoryRestrictions: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"Software Subscriptions", "Conferences"}),
				MonthlyLimit:                 jamesburvelocallaghaniiicitibankdemobusinessinc.F(3000.000000),
				OnlineTransactions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
				SingleTransactionLimit:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(0.000000),
				VendorRestrictions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"string"}),
			},
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
