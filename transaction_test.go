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
		Category:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Groceries"),
		EndDate:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-12-31"),
		Limit:       jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
		MaxAmount:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](100),
		MinAmount:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](20),
		Offset:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
		SearchQuery: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Starbucks"),
		StartDate:   jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("2024-01-01"),
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
			Category:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Home > Groceries"),
			ApplyToFuture: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			Notes:         jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Bulk purchase for party"),
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
			Details:             jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("I did not authorize this purchase. My card may have been compromised and I was traveling internationally on this date."),
			Reason:              jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.TransactionDisputeParamsReasonUnauthorized),
			SupportingDocuments: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"https://demobank.com/uploads/flight_ticket.png"}),
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
			Notes: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("This was a special coffee for a client meeting."),
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
