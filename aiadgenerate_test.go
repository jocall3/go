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

func TestAIAdGenerateAdvancedWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Ads.Generate.Advanced(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParams{
		Prompt:         jamesburvelocallaghaniiicitibankdemobusinessinc.F("prompt"),
		AspectRatio:    jamesburvelocallaghaniiicitibankdemobusinessinc.F("aspectRatio"),
		AudienceTarget: jamesburvelocallaghaniiicitibankdemobusinessinc.F("audienceTarget"),
		BrandAssets:    jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"https://example.com"}),
		BrandColors:    jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"string"}),
		CallToAction: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParamsCallToAction{
			DisplayTimeSeconds: jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(0)),
			Text:               jamesburvelocallaghaniiicitibankdemobusinessinc.F("text"),
			URL:                jamesburvelocallaghaniiicitibankdemobusinessinc.F("https://example.com"),
		}),
		LengthSeconds:  jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(0)),
		Style:          jamesburvelocallaghaniiicitibankdemobusinessinc.F("style"),
		VoiceoverStyle: jamesburvelocallaghaniiicitibankdemobusinessinc.F("voiceoverStyle"),
		VoiceoverText:  jamesburvelocallaghaniiicitibankdemobusinessinc.F("voiceoverText"),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAdGenerateStandardWithOptionalParams(t *testing.T) {
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
	_, err := client.AI.Ads.Generate.Standard(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateStandardParams{
		GenerateVideo: jamesburvelocallaghaniiicitibankdemobusinessinc.GenerateVideoParam{
			Prompt:        jamesburvelocallaghaniiicitibankdemobusinessinc.F("A captivating ad featuring a young entrepreneur using 's AI tools to grow their startup. Focus on innovation and ease of use."),
			AspectRatio:   jamesburvelocallaghaniiicitibankdemobusinessinc.F("16:9"),
			BrandColors:   jamesburvelocallaghaniiicitibankdemobusinessinc.F([]string{"#0000FF", "#FFD700"}),
			LengthSeconds: jamesburvelocallaghaniiicitibankdemobusinessinc.F(int64(15)),
			Style:         jamesburvelocallaghaniiicitibankdemobusinessinc.F("Cinematic"),
		},
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
