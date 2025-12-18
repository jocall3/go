// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/1231-go/internal/apijson"
	"github.com/stainless-sdks/1231-go/internal/requestconfig"
	"github.com/stainless-sdks/1231-go/option"
)

// IdentityKYCService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityKYCService] method instead.
type IdentityKYCService struct {
	Options []option.RequestOption
}

// NewIdentityKYCService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityKYCService(opts ...option.RequestOption) (r *IdentityKYCService) {
	r = &IdentityKYCService{}
	r.Options = opts
	return
}

// Retrieves the current status of the user's KYC/AML identity verification.
func (r *IdentityKYCService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *KYCStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "identity/kyc/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type KYCStatus struct {
	LastUpdated time.Time       `json:"lastUpdated" format:"date-time"`
	Message     string          `json:"message"`
	Status      KYCStatusStatus `json:"status"`
	JSON        kycStatusJSON   `json:"-"`
}

// kycStatusJSON contains the JSON metadata for the struct [KYCStatus]
type kycStatusJSON struct {
	LastUpdated apijson.Field
	Message     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KYCStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kycStatusJSON) RawJSON() string {
	return r.raw
}

type KYCStatusStatus string

const (
	KYCStatusStatusNotStarted    KYCStatusStatus = "not_started"
	KYCStatusStatusPendingReview KYCStatusStatus = "pending_review"
	KYCStatusStatusVerified      KYCStatusStatus = "verified"
	KYCStatusStatusRejected      KYCStatusStatus = "rejected"
)

func (r KYCStatusStatus) IsKnown() bool {
	switch r {
	case KYCStatusStatusNotStarted, KYCStatusStatusPendingReview, KYCStatusStatusVerified, KYCStatusStatusRejected:
		return true
	}
	return false
}
