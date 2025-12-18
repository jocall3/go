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

func TestWeb3WalletListWithOptionalParams(t *testing.T) {
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
	_, err := client.Web3.Wallets.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.Web3WalletListParams{
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

func TestWeb3WalletConnectWithOptionalParams(t *testing.T) {
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
	_, err := client.Web3.Wallets.Connect(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.Web3WalletConnectParams{
		BlockchainNetwork:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Ethereum"),
		SignedMessage:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"),
		WalletAddress:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("0x123abc456def7890..."),
		WalletProvider:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("MetaMask"),
		RequestWriteAccess: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWeb3WalletGetBalancesWithOptionalParams(t *testing.T) {
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
	_, err := client.Web3.Wallets.GetBalances(
		context.TODO(),
		"wallet_conn_eth_0xabc123",
		jamesburvelocallaghaniiicitibankdemobusinessinc.Web3WalletGetBalancesParams{
			Limit:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
			Offset: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
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
