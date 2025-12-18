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

func TestWeb3TransactionInitiateTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.Web3.Transactions.InitiateTransfer(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.Web3TransactionInitiateTransferParams{
		Amount:            jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](0.1),
		AssetSymbol:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("ETH"),
		BlockchainNetwork: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Ethereum"),
		RecipientAddress:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("0xdef4567890abcdef1234567890abcdef1234567890"),
		SourceWalletID:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("wallet_conn_eth_0xabc123"),
		GasPriceGwei:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](50),
		Memo:              jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Payment for services"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
