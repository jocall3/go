// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/param"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// UserMeBiometricService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserMeBiometricService] method instead.
type UserMeBiometricService struct {
	Options []option.RequestOption
}

// NewUserMeBiometricService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserMeBiometricService(opts ...option.RequestOption) (r *UserMeBiometricService) {
	r = &UserMeBiometricService{}
	r.Options = opts
	return
}

// Removes all enrolled biometric data associated with the user's account for
// security reasons.
func (r *UserMeBiometricService) Deregister(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "users/me/biometrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Initiates the enrollment process for biometric authentication (e.g.,
// fingerprint, facial scan) to enable secure and convenient access to sensitive
// features. Requires a biometric signature for initial proof.
func (r *UserMeBiometricService) Enroll(ctx context.Context, body UserMeBiometricEnrollParams, opts ...option.RequestOption) (res *BiometricStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics/enroll"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the current status of biometric enrollments for the authenticated
// user.
func (r *UserMeBiometricService) Status(ctx context.Context, opts ...option.RequestOption) (res *BiometricStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Performs real-time biometric verification to authorize sensitive actions or
// access protected resources, using a one-time biometric signature.
func (r *UserMeBiometricService) Verify(ctx context.Context, body UserMeBiometricVerifyParams, opts ...option.RequestOption) (res *UserMeBiometricVerifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/me/biometrics/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BiometricStatus struct {
	BiometricsEnrolled bool                               `json:"biometricsEnrolled"`
	EnrolledBiometrics []BiometricStatusEnrolledBiometric `json:"enrolledBiometrics"`
	LastUsed           time.Time                          `json:"lastUsed" format:"date-time"`
	JSON               biometricStatusJSON                `json:"-"`
}

// biometricStatusJSON contains the JSON metadata for the struct [BiometricStatus]
type biometricStatusJSON struct {
	BiometricsEnrolled apijson.Field
	EnrolledBiometrics apijson.Field
	LastUsed           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BiometricStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r biometricStatusJSON) RawJSON() string {
	return r.raw
}

type BiometricStatusEnrolledBiometric struct {
	DeviceID       string                                `json:"deviceId"`
	EnrollmentDate time.Time                             `json:"enrollmentDate" format:"date-time"`
	Type           BiometricStatusEnrolledBiometricsType `json:"type"`
	JSON           biometricStatusEnrolledBiometricJSON  `json:"-"`
}

// biometricStatusEnrolledBiometricJSON contains the JSON metadata for the struct
// [BiometricStatusEnrolledBiometric]
type biometricStatusEnrolledBiometricJSON struct {
	DeviceID       apijson.Field
	EnrollmentDate apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BiometricStatusEnrolledBiometric) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r biometricStatusEnrolledBiometricJSON) RawJSON() string {
	return r.raw
}

type BiometricStatusEnrolledBiometricsType string

const (
	BiometricStatusEnrolledBiometricsTypeFacialRecognition BiometricStatusEnrolledBiometricsType = "facial_recognition"
	BiometricStatusEnrolledBiometricsTypeFingerprint       BiometricStatusEnrolledBiometricsType = "fingerprint"
)

func (r BiometricStatusEnrolledBiometricsType) IsKnown() bool {
	switch r {
	case BiometricStatusEnrolledBiometricsTypeFacialRecognition, BiometricStatusEnrolledBiometricsTypeFingerprint:
		return true
	}
	return false
}

type UserMeBiometricVerifyResponse struct {
	Message            string                                          `json:"message"`
	VerificationStatus UserMeBiometricVerifyResponseVerificationStatus `json:"verificationStatus"`
	JSON               userMeBiometricVerifyResponseJSON               `json:"-"`
}

// userMeBiometricVerifyResponseJSON contains the JSON metadata for the struct
// [UserMeBiometricVerifyResponse]
type userMeBiometricVerifyResponseJSON struct {
	Message            apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserMeBiometricVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeBiometricVerifyResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeBiometricVerifyResponseVerificationStatus string

const (
	UserMeBiometricVerifyResponseVerificationStatusSuccess UserMeBiometricVerifyResponseVerificationStatus = "success"
	UserMeBiometricVerifyResponseVerificationStatusFailed  UserMeBiometricVerifyResponseVerificationStatus = "failed"
)

func (r UserMeBiometricVerifyResponseVerificationStatus) IsKnown() bool {
	switch r {
	case UserMeBiometricVerifyResponseVerificationStatusSuccess, UserMeBiometricVerifyResponseVerificationStatusFailed:
		return true
	}
	return false
}

type UserMeBiometricEnrollParams struct {
	// Base64 encoded biometric template for enrollment.
	BiometricSignature param.Field[string]                                   `json:"biometricSignature,required"`
	BiometricType      param.Field[UserMeBiometricEnrollParamsBiometricType] `json:"biometricType,required"`
	DeviceID           param.Field[string]                                   `json:"deviceId,required"`
	DeviceName         param.Field[string]                                   `json:"deviceName"`
}

func (r UserMeBiometricEnrollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMeBiometricEnrollParamsBiometricType string

const (
	UserMeBiometricEnrollParamsBiometricTypeFacialRecognition UserMeBiometricEnrollParamsBiometricType = "facial_recognition"
	UserMeBiometricEnrollParamsBiometricTypeFingerprint       UserMeBiometricEnrollParamsBiometricType = "fingerprint"
)

func (r UserMeBiometricEnrollParamsBiometricType) IsKnown() bool {
	switch r {
	case UserMeBiometricEnrollParamsBiometricTypeFacialRecognition, UserMeBiometricEnrollParamsBiometricTypeFingerprint:
		return true
	}
	return false
}

type UserMeBiometricVerifyParams struct {
	// Base64 encoded one-time biometric proof for verification.
	BiometricSignature param.Field[string]                                   `json:"biometricSignature,required"`
	BiometricType      param.Field[UserMeBiometricVerifyParamsBiometricType] `json:"biometricType,required"`
	DeviceID           param.Field[string]                                   `json:"deviceId,required"`
}

func (r UserMeBiometricVerifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserMeBiometricVerifyParamsBiometricType string

const (
	UserMeBiometricVerifyParamsBiometricTypeFacialRecognition UserMeBiometricVerifyParamsBiometricType = "facial_recognition"
	UserMeBiometricVerifyParamsBiometricTypeFingerprint       UserMeBiometricVerifyParamsBiometricType = "fingerprint"
)

func (r UserMeBiometricVerifyParamsBiometricType) IsKnown() bool {
	switch r {
	case UserMeBiometricVerifyParamsBiometricTypeFacialRecognition, UserMeBiometricVerifyParamsBiometricTypeFingerprint:
		return true
	}
	return false
}
