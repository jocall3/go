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
		Limit:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
		Offset: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
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
			AtmWithdrawals:               jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](false),
			ContactlessPayments:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](false),
			DailyLimit:                   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](500),
			InternationalTransactions:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](false),
			MerchantCategoryRestrictions: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"Advertising"}),
			MonthlyLimit:                 jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](1000),
			OnlineTransactions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			SingleTransactionLimit:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](200),
			VendorRestrictions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"Facebook Ads", "Google Ads"}),
		}),
		ExpirationDate:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2025-12-31"),
		HolderName:           jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Marketing Campaign Q4"),
		Purpose:              jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Online advertising for Q4 campaigns"),
		AssociatedEmployeeID: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("emp_marketing_01"),
		SpendingPolicyID:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("policy_marketing_fixed"),
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
			Freeze: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
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
			EndDate:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-12-31"),
			Limit:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
			Offset:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
			StartDate: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-01-01"),
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
				AtmWithdrawals:               jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				ContactlessPayments:          jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				DailyLimit:                   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](750),
				InternationalTransactions:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				MerchantCategoryRestrictions: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"Software Subscriptions", "Conferences"}),
				MonthlyLimit:                 jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](3000),
				OnlineTransactions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				SingleTransactionLimit:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](1000),
				VendorRestrictions:           jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"Amazon", "Uber"}),
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
