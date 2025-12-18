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

func TestTransactionGet(t *testing.T) {
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
	_, err := client.Transactions.Get(context.TODO(), "txn_quantum-2024-07-21-A7B8C9")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionListParams{
		Category:    jamesburvelocallaghaniiicitibankdemobusinessinc.F("Groceries"),
		EndDate:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
		Limit:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(20)),
		MaxAmount:   jamesburvelocallaghaniiicitibankdemobusinessinc.F(100.000000),
		MinAmount:   jamesburvelocallaghaniiicitibankdemobusinessinc.F(20.000000),
		Offset:      jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(0)),
		SearchQuery: jamesburvelocallaghaniiicitibankdemobusinessinc.F("Starbucks"),
		StartDate:   jamesburvelocallaghaniiicitibankdemobusinessinc.F(time.Now()),
		Type:        jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionListParamsTypeExpense),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionCategorizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Categorize(
		context.TODO(),
		"txn_quantum-2024-07-21-A7B8C9",
		jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionCategorizeParams{
			Category:      jamesburvelocallaghaniiicitibankdemobusinessinc.F("Home > Groceries"),
			ApplyToFuture: jamesburvelocallaghaniiicitibankdemobusinessinc.F(true),
			Notes:         jamesburvelocallaghaniiicitibankdemobusinessinc.F("Bulk purchase for party"),
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

func TestTransactionDisputeWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.Dispute(
		context.TODO(),
		"txn_quantum-2024-07-21-A7B8C9",
		jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionDisputeParams{
			Details:             jamesburvelocallaghaniiicitibankdemobusinessinc.F("I did not authorize this purchase. My card may have been compromised and I was traveling internationally on this date."),
			Reason:              jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionDisputeParamsReasonUnauthorized),
			SupportingDocuments: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"https://demobank.com/uploads/flight_ticket.png"}),
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

func TestTransactionUpdateNotes(t *testing.T) {
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
	_, err := client.Transactions.UpdateNotes(
		context.TODO(),
		"txn_quantum-2024-07-21-A7B8C9",
		jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionUpdateNotesParams{
			Notes: jamesburvelocallaghaniiicitibankdemobusinessinc.F("This was a special coffee for a client meeting."),
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
