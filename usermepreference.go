// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/param"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// UserMePreferenceService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMePreferenceService] method instead.
type UserMePreferenceService struct {
	Options []option.RequestOption
}

// NewUserMePreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserMePreferenceService(opts ...option.RequestOption) (r *UserMePreferenceService) {
	r = &UserMePreferenceService{}
	r.Options = opts
	return
}

// Retrieves the user's deep personalization preferences, including AI
// customization settings, notification channel priorities, thematic choices, and
// data sharing consents.
func (r *UserMePreferenceService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserPreferences, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the user's deep personalization preferences, allowing dynamic control
// over AI behavior, notification delivery, thematic choices, and data privacy
// settings.
func (r *UserMePreferenceService) Update(ctx context.Context, body UserMePreferenceUpdateParams, opts ...option.RequestOption) (res *UserPreferences, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type UserPreferences struct {
	AIInteractionMode    UserPreferencesAIInteractionMode    `json:"aiInteractionMode"`
	DataSharingConsent   bool                                `json:"dataSharingConsent"`
	NotificationChannels UserPreferencesNotificationChannels `json:"notificationChannels"`
	PreferredLanguage    string                              `json:"preferredLanguage"`
	Theme                string                              `json:"theme"`
	TransactionGrouping  UserPreferencesTransactionGrouping  `json:"transactionGrouping"`
	JSON                 userPreferencesJSON                 `json:"-"`
}

// userPreferencesJSON contains the JSON metadata for the struct [UserPreferences]
type userPreferencesJSON struct {
	AIInteractionMode    apijson.Field
	DataSharingConsent   apijson.Field
	NotificationChannels apijson.Field
	PreferredLanguage    apijson.Field
	Theme                apijson.Field
	TransactionGrouping  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UserPreferences) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferencesJSON) RawJSON() string {
	return r.raw
}

type UserPreferencesAIInteractionMode string

const (
	UserPreferencesAIInteractionModeProactive UserPreferencesAIInteractionMode = "proactive"
	UserPreferencesAIInteractionModeBalanced  UserPreferencesAIInteractionMode = "balanced"
	UserPreferencesAIInteractionModeReactive  UserPreferencesAIInteractionMode = "reactive"
)

func (r UserPreferencesAIInteractionMode) IsKnown() bool {
	switch r {
	case UserPreferencesAIInteractionModeProactive, UserPreferencesAIInteractionModeBalanced, UserPreferencesAIInteractionModeReactive:
		return true
	}
	return false
}

type UserPreferencesNotificationChannels struct {
	Email bool                                    `json:"email"`
	InApp bool                                    `json:"inApp"`
	Push  bool                                    `json:"push"`
	SMS   bool                                    `json:"sms"`
	JSON  userPreferencesNotificationChannelsJSON `json:"-"`
}

// userPreferencesNotificationChannelsJSON contains the JSON metadata for the
// struct [UserPreferencesNotificationChannels]
type userPreferencesNotificationChannelsJSON struct {
	Email       apijson.Field
	InApp       apijson.Field
	Push        apijson.Field
	SMS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPreferencesNotificationChannels) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPreferencesNotificationChannelsJSON) RawJSON() string {
	return r.raw
}

type UserPreferencesTransactionGrouping string

const (
	UserPreferencesTransactionGroupingCategory UserPreferencesTransactionGrouping = "category"
	UserPreferencesTransactionGroupingMerchant UserPreferencesTransactionGrouping = "merchant"
	UserPreferencesTransactionGroupingDate     UserPreferencesTransactionGrouping = "date"
)

func (r UserPreferencesTransactionGrouping) IsKnown() bool {
	switch r {
	case UserPreferencesTransactionGroupingCategory, UserPreferencesTransactionGroupingMerchant, UserPreferencesTransactionGroupingDate:
		return true
	}
	return false
}

type UserPreferencesParam struct {
	AIInteractionMode    param.Field[UserPreferencesAIInteractionMode]         `json:"aiInteractionMode"`
	DataSharingConsent   param.Field[bool]                                     `json:"dataSharingConsent"`
	NotificationChannels param.Field[UserPreferencesNotificationChannelsParam] `json:"notificationChannels"`
	PreferredLanguage    param.Field[string]                                   `json:"preferredLanguage"`
	Theme                param.Field[string]                                   `json:"theme"`
	TransactionGrouping  param.Field[UserPreferencesTransactionGrouping]       `json:"transactionGrouping"`
}

func (r UserPreferencesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserPreferencesNotificationChannelsParam struct {
	Email param.Field[bool] `json:"email"`
	InApp param.Field[bool] `json:"inApp"`
	Push  param.Field[bool] `json:"push"`
	SMS   param.Field[bool] `json:"sms"`
}

func (r UserPreferencesNotificationChannelsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMePreferenceUpdateParams struct {
	UserPreferences UserPreferencesParam `json:"user_preferences,required"`
}

func (r UserMePreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UserPreferences)
}
