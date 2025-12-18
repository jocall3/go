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

func TestCorporateRiskFraudRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Risk.Fraud.Rules.New(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleNewParams{
		Action: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleActionParam{
			Details:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Hold payment, notify sender for additional verification, and escalate to compliance."),
			Type:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleActionTypeAutoReview),
			TargetTeam: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Fraud Prevention Team"),
		}),
		Criteria: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaParam{
			AccountInactivityDays:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](90),
			CountryOfOrigin:           jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"US", "CA"}),
			GeographicDistanceKm:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](5000),
			LastLoginDays:             jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](7),
			NoTravelNotification:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			PaymentCountMin:           jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](3),
			RecipientCountryRiskLevel: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaRecipientCountryRiskLevel{jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaRecipientCountryRiskLevelHigh, jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaRecipientCountryRiskLevelVeryHigh}),
			RecipientNew:              jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
			TimeframeHours:            jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](24),
			TransactionAmountMin:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](5000),
			TransactionType:           jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaTransactionTypeDebit),
		}),
		Description: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Detects multiple international payments to new beneficiaries in high-risk countries within a short timeframe."),
		Name:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Suspicious International Payment Pattern"),
		Severity:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleNewParamsSeverityCritical),
		Status:      jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleNewParamsStatusActive),
	})
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorporateRiskFraudRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Risk.Fraud.Rules.Update(
		context.TODO(),
		"fraud_rule_high_value_inactive",
		jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleUpdateParams{
			Action: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleActionParam{
				Details:    jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Flag for manual review only, do not block."),
				Type:       jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleActionTypeBlock),
				TargetTeam: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Fraud Prevention Team"),
			}),
			Criteria: jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaParam{
				AccountInactivityDays:     jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](60),
				CountryOfOrigin:           jamesburvelocallaghaniiicitibankdemobusinessinc.F([]interface{}{"US", "CA"}),
				GeographicDistanceKm:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](5000),
				LastLoginDays:             jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](7),
				NoTravelNotification:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				PaymentCountMin:           jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](3),
				RecipientCountryRiskLevel: jamesburvelocallaghaniiicitibankdemobusinessinc.F([]jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaRecipientCountryRiskLevel{jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaRecipientCountryRiskLevelLow}),
				RecipientNew:              jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](true),
				TimeframeHours:            jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](24),
				TransactionAmountMin:      jamesburvelocallaghaniiicitibankdemobusinessinc.F[any](7500),
				TransactionType:           jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.FraudRuleCriteriaTransactionTypeDebit),
			}),
			Description: jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Revised logic for flagging high value transactions from dormant accounts."),
			Name:        jamesburvelocallaghaniiicitibankdemobusinessinc.F[any]("Revised High Value Transaction Rule"),
			Severity:    jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleUpdateParamsSeverityHigh),
			Status:      jamesburvelocallaghaniiicitibankdemobusinessinc.F(jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleUpdateParamsStatusInactive),
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

func TestCorporateRiskFraudRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Corporate.Risk.Fraud.Rules.List(context.TODO(), jamesburvelocallaghaniiicitibankdemobusinessinc.CorporateRiskFraudRuleListParams{
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

func TestCorporateRiskFraudRuleDelete(t *testing.T) {
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
	err := client.Corporate.Risk.Fraud.Rules.Delete(context.TODO(), "fraud_rule_high_value_inactive")
	if err != nil {
		var apierr *jamesburvelocallaghaniiicitibankdemobusinessinc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
