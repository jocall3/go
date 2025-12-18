// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/jocall3/1231-go/internal/apiform"
	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
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

// Retrieves the current status of the user's KYC verification.
func (r *IdentityKYCService) GetStatus(ctx context.Context, opts ...option.RequestOption) (res *IdentityKYCGetStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "identity/kyc/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits Know Your Customer (KYC) documentation for identity verification.
func (r *IdentityKYCService) Submit(ctx context.Context, body IdentityKYCSubmitParams, opts ...option.RequestOption) (res *IdentityKYCSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "identity/kyc/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IdentityKYCGetStatusResponse struct {
	Details     string                             `json:"details"`
	LastChecked time.Time                          `json:"lastChecked" format:"date-time"`
	Status      IdentityKYCGetStatusResponseStatus `json:"status"`
	JSON        identityKYCGetStatusResponseJSON   `json:"-"`
}

// identityKYCGetStatusResponseJSON contains the JSON metadata for the struct
// [IdentityKYCGetStatusResponse]
type identityKYCGetStatusResponseJSON struct {
	Details     apijson.Field
	LastChecked apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityKYCGetStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityKYCGetStatusResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityKYCGetStatusResponseStatus string

const (
	IdentityKYCGetStatusResponseStatusNotStarted IdentityKYCGetStatusResponseStatus = "not_started"
	IdentityKYCGetStatusResponseStatusPending    IdentityKYCGetStatusResponseStatus = "pending"
	IdentityKYCGetStatusResponseStatusVerified   IdentityKYCGetStatusResponseStatus = "verified"
	IdentityKYCGetStatusResponseStatusRejected   IdentityKYCGetStatusResponseStatus = "rejected"
)

func (r IdentityKYCGetStatusResponseStatus) IsKnown() bool {
	switch r {
	case IdentityKYCGetStatusResponseStatusNotStarted, IdentityKYCGetStatusResponseStatusPending, IdentityKYCGetStatusResponseStatusVerified, IdentityKYCGetStatusResponseStatusRejected:
		return true
	}
	return false
}

type IdentityKYCSubmitResponse struct {
	Details     string                          `json:"details"`
	LastChecked time.Time                       `json:"lastChecked" format:"date-time"`
	Status      IdentityKYCSubmitResponseStatus `json:"status"`
	JSON        identityKYCSubmitResponseJSON   `json:"-"`
}

// identityKYCSubmitResponseJSON contains the JSON metadata for the struct
// [IdentityKYCSubmitResponse]
type identityKYCSubmitResponseJSON struct {
	Details     apijson.Field
	LastChecked apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityKYCSubmitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityKYCSubmitResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityKYCSubmitResponseStatus string

const (
	IdentityKYCSubmitResponseStatusNotStarted IdentityKYCSubmitResponseStatus = "not_started"
	IdentityKYCSubmitResponseStatusPending    IdentityKYCSubmitResponseStatus = "pending"
	IdentityKYCSubmitResponseStatusVerified   IdentityKYCSubmitResponseStatus = "verified"
	IdentityKYCSubmitResponseStatusRejected   IdentityKYCSubmitResponseStatus = "rejected"
)

func (r IdentityKYCSubmitResponseStatus) IsKnown() bool {
	switch r {
	case IdentityKYCSubmitResponseStatusNotStarted, IdentityKYCSubmitResponseStatusPending, IdentityKYCSubmitResponseStatusVerified, IdentityKYCSubmitResponseStatusRejected:
		return true
	}
	return false
}

type IdentityKYCSubmitParams struct {
	// Front side of the ID document.
	DocumentFront param.Field[io.Reader]                           `json:"documentFront,required" format:"binary"`
	DocumentType  param.Field[IdentityKYCSubmitParamsDocumentType] `json:"documentType,required"`
	// A selfie of the user holding the ID document.
	Selfie param.Field[io.Reader] `json:"selfie,required" format:"binary"`
	// Back side of the ID document (if applicable).
	DocumentBack param.Field[io.Reader] `json:"documentBack" format:"binary"`
}

func (r IdentityKYCSubmitParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type IdentityKYCSubmitParamsDocumentType string

const (
	IdentityKYCSubmitParamsDocumentTypePassport       IdentityKYCSubmitParamsDocumentType = "passport"
	IdentityKYCSubmitParamsDocumentTypeDriversLicense IdentityKYCSubmitParamsDocumentType = "drivers_license"
	IdentityKYCSubmitParamsDocumentTypeNationalID     IdentityKYCSubmitParamsDocumentType = "national_id"
)

func (r IdentityKYCSubmitParamsDocumentType) IsKnown() bool {
	switch r {
	case IdentityKYCSubmitParamsDocumentTypePassport, IdentityKYCSubmitParamsDocumentTypeDriversLicense, IdentityKYCSubmitParamsDocumentTypeNationalID:
		return true
	}
	return false
}
