// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/apiquery"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// UserMeDeviceService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeDeviceService] method instead.
type UserMeDeviceService struct {
	Options []option.RequestOption
}

// NewUserMeDeviceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMeDeviceService(opts ...option.RequestOption) (r *UserMeDeviceService) {
	r = &UserMeDeviceService{}
	r.Options = opts
	return
}

// Retrieves a list of all devices linked to the user's account, including mobile
// phones, tablets, and desktops, indicating their last active status and security
// posture.
func (r *UserMeDeviceService) List(ctx context.Context, query UserMeDeviceListParams, opts ...option.RequestOption) (res *UserMeDeviceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes a specific device from the user's linked devices, revoking its access
// and requiring re-registration for future use. Useful for lost or compromised
// devices.
func (r *UserMeDeviceService) Deregister(ctx context.Context, deviceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if deviceID == "" {
		err = errors.New("missing required deviceId parameter")
		return
	}
	path := fmt.Sprintf("users/me/devices/%s", deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Registers a new device for secure access and multi-factor authentication,
// associating it with the user's profile. This typically initiates a biometric or
// MFA enrollment flow.
func (r *UserMeDeviceService) Register(ctx context.Context, body UserMeDeviceRegisterParams, opts ...option.RequestOption) (res *Device, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/devices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Device struct {
	ID         string           `json:"id"`
	IPAddress  string           `json:"ipAddress"`
	LastActive time.Time        `json:"lastActive" format:"date-time"`
	Model      string           `json:"model"`
	Os         string           `json:"os"`
	PushToken  string           `json:"pushToken"`
	TrustLevel DeviceTrustLevel `json:"trustLevel"`
	Type       DeviceType       `json:"type"`
	JSON       deviceJSON       `json:"-"`
}

// deviceJSON contains the JSON metadata for the struct [Device]
type deviceJSON struct {
	ID          apijson.Field
	IPAddress   apijson.Field
	LastActive  apijson.Field
	Model       apijson.Field
	Os          apijson.Field
	PushToken   apijson.Field
	TrustLevel  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Device) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceJSON) RawJSON() string {
	return r.raw
}

type DeviceTrustLevel string

const (
	DeviceTrustLevelTrusted             DeviceTrustLevel = "trusted"
	DeviceTrustLevelUntrusted           DeviceTrustLevel = "untrusted"
	DeviceTrustLevelPendingVerification DeviceTrustLevel = "pending_verification"
)

func (r DeviceTrustLevel) IsKnown() bool {
	switch r {
	case DeviceTrustLevelTrusted, DeviceTrustLevelUntrusted, DeviceTrustLevelPendingVerification:
		return true
	}
	return false
}

type DeviceType string

const (
	DeviceTypeMobile  DeviceType = "mobile"
	DeviceTypeDesktop DeviceType = "desktop"
	DeviceTypeTablet  DeviceType = "tablet"
)

func (r DeviceType) IsKnown() bool {
	switch r {
	case DeviceTypeMobile, DeviceTypeDesktop, DeviceTypeTablet:
		return true
	}
	return false
}

type PaginatedList struct {
	// The number of items returned in this page.
	Limit int64 `json:"limit,required"`
	// The starting position of the returned items.
	Offset int64 `json:"offset,required"`
	// The total number of items available.
	Total int64 `json:"total,required"`
	// The offset for the next page of results, if available.
	NextOffset int64             `json:"nextOffset"`
	JSON       paginatedListJSON `json:"-"`
}

// paginatedListJSON contains the JSON metadata for the struct [PaginatedList]
type paginatedListJSON struct {
	Limit       apijson.Field
	Offset      apijson.Field
	Total       apijson.Field
	NextOffset  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedListJSON) RawJSON() string {
	return r.raw
}

type UserMeDeviceListResponse struct {
	Data []Device                     `json:"data"`
	JSON userMeDeviceListResponseJSON `json:"-"`
	PaginatedList
}

// userMeDeviceListResponseJSON contains the JSON metadata for the struct
// [UserMeDeviceListResponse]
type userMeDeviceListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMeDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeDeviceListParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [UserMeDeviceListParams]'s query parameters as `url.Values`.
func (r UserMeDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserMeDeviceRegisterParams struct {
	DeviceType param.Field[UserMeDeviceRegisterParamsDeviceType] `json:"deviceType,required"`
	Model      param.Field[string]                               `json:"model,required"`
	Os         param.Field[string]                               `json:"os,required"`
	// Base64 encoded biometric proof for enrollment.
	BiometricSignature param.Field[string] `json:"biometricSignature"`
	DeviceName         param.Field[string] `json:"deviceName"`
}

func (r UserMeDeviceRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMeDeviceRegisterParamsDeviceType string

const (
	UserMeDeviceRegisterParamsDeviceTypeMobile  UserMeDeviceRegisterParamsDeviceType = "mobile"
	UserMeDeviceRegisterParamsDeviceTypeDesktop UserMeDeviceRegisterParamsDeviceType = "desktop"
	UserMeDeviceRegisterParamsDeviceTypeTablet  UserMeDeviceRegisterParamsDeviceType = "tablet"
)

func (r UserMeDeviceRegisterParamsDeviceType) IsKnown() bool {
	switch r {
	case UserMeDeviceRegisterParamsDeviceTypeMobile, UserMeDeviceRegisterParamsDeviceTypeDesktop, UserMeDeviceRegisterParamsDeviceTypeTablet:
		return true
	}
	return false
}
