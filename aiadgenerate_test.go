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
		LengthSeconds:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](15),
		Prompt:               jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("A captivating ad featuring a young entrepreneur using 's AI tools to grow their startup. Focus on innovation and ease of use."),
		Style:                jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParamsStyleCinematic),
		AspectRatio:          jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParamsAspectRatio16_9),
		AudienceTarget:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParamsAudienceTargetCorporate),
		BackgroundMusicGenre: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParamsBackgroundMusicGenreCorporate),
		BrandAssets:          jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"https://demobank.com/assets/corporate_logo.png"}),
		BrandColors:          jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"#0000FF", "#FFD700"}),
		CallToAction: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParamsCallToAction{
			DisplayTimeSeconds: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](5),
			Text:               jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Learn more at DemoBank.com/business"),
			URL:                jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("https://demobank.com/business"),
		}),
		Keywords:       jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"innovation", "fintech", "startup"}),
		VoiceoverStyle: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.AIAdGenerateAdvancedParamsVoiceoverStyleMaleProfessional),
		VoiceoverText:  jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](": Your business, powered by intelligent finance."),
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
			LengthSeconds: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](15),
			Prompt:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("A captivating ad featuring a young entrepreneur using 's AI tools to grow their startup. Focus on innovation and ease of use."),
			Style:         jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.GenerateVideoStyleCinematic),
			AspectRatio:   jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.GenerateVideoAspectRatio16_9),
			BrandColors:   jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"#0000FF", "#FFD700"}),
			Keywords:      jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"innovation", "fintech", "startup"}),
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
