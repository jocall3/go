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

func TestAIAdvisorChatGetHistoryWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Advisor.Chat.GetHistory(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdvisorChatGetHistoryParams{
		Limit:     jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(1)),
		Offset:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(0)),
		SessionID: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("session-quantum-xyz-789-alpha"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAdvisorChatSendMessageWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Advisor.Chat.SendMessage(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdvisorChatSendMessageParams{
		FunctionResponse: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdvisorChatSendMessageParamsFunctionResponse{
			Name:     jamesburvelocallaghaniiicitibankdemobusinessinc.F("name"),
			Response: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](map[string]interface{}{}),
		}),
		Message:   jamesburvelocallaghaniiicitibankdemobusinessinc.F("Can you analyze my recent spending patterns and suggest areas for saving, focusing on my dining expenses?"),
		SessionID: jamesburvelocallaghaniiicitibankdemobusinessinc.F("session-quantum-xyz-789-alpha"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
