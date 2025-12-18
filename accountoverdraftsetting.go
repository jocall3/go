// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// AccountOverdraftSettingService contains methods and other services that help
// with interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountOverdraftSettingService] method instead.
type AccountOverdraftSettingService struct {
	Options []option.RequestOption
}

// NewAccountOverdraftSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountOverdraftSettingService(opts ...option.RequestOption) (r *AccountOverdraftSettingService) {
	r = &AccountOverdraftSettingService{}
	r.Options = opts
	return
}

// Retrieves the current overdraft protection settings for a specific account.
func (r *AccountOverdraftSettingService) GetOverdraftSettings(ctx context.Context, accountID string, opts ...option.RequestOption) (res *OverdraftSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the overdraft protection settings for a specific account, enabling or
// disabling protection and configuring preferences.
func (r *AccountOverdraftSettingService) UpdateOverdraftSettings(ctx context.Context, accountID string, body AccountOverdraftSettingUpdateOverdraftSettingsParams, opts ...option.RequestOption) (res *OverdraftSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/overdraft-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type OverdraftSettings struct {
	AccountID              string                         `json:"accountId"`
	Enabled                bool                           `json:"enabled"`
	FeePreference          OverdraftSettingsFeePreference `json:"feePreference"`
	LinkedSavingsAccountID string                         `json:"linkedSavingsAccountId"`
	LinkToSavings          bool                           `json:"linkToSavings"`
	ProtectionLimit        float64                        `json:"protectionLimit"`
	JSON                   overdraftSettingsJSON          `json:"-"`
}

// overdraftSettingsJSON contains the JSON metadata for the struct
// [OverdraftSettings]
type overdraftSettingsJSON struct {
	AccountID              apijson.Field
	Enabled                apijson.Field
	FeePreference          apijson.Field
	LinkedSavingsAccountID apijson.Field
	LinkToSavings          apijson.Field
	ProtectionLimit        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *OverdraftSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overdraftSettingsJSON) RawJSON() string {
	return r.raw
}

type OverdraftSettingsFeePreference string

const (
	OverdraftSettingsFeePreferenceAlwaysPay          OverdraftSettingsFeePreference = "always_pay"
	OverdraftSettingsFeePreferenceDeclineIfOverLimit OverdraftSettingsFeePreference = "decline_if_over_limit"
)

func (r OverdraftSettingsFeePreference) IsKnown() bool {
	switch r {
	case OverdraftSettingsFeePreferenceAlwaysPay, OverdraftSettingsFeePreferenceDeclineIfOverLimit:
		return true
	}
	return false
}

type AccountOverdraftSettingUpdateOverdraftSettingsParams struct {
	Enabled       param.Field[bool]                                                              `json:"enabled"`
	FeePreference param.Field[AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference] `json:"feePreference"`
	LinkToSavings param.Field[bool]                                                              `json:"linkToSavings"`
}

func (r AccountOverdraftSettingUpdateOverdraftSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference string

const (
	AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceAlwaysPay          AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference = "always_pay"
	AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceDeclineIfOverLimit AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference = "decline_if_over_limit"
)

func (r AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreference) IsKnown() bool {
	switch r {
	case AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceAlwaysPay, AccountOverdraftSettingUpdateOverdraftSettingsParamsFeePreferenceDeclineIfOverLimit:
		return true
	}
	return false
}
