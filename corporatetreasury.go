// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"github.com/stainless-sdks/1231-go/option"
)

// CorporateTreasuryService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorporateTreasuryService] method instead.
type CorporateTreasuryService struct {
	Options  []option.RequestOption
	CashFlow *CorporateTreasuryCashFlowService
}

// NewCorporateTreasuryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCorporateTreasuryService(opts ...option.RequestOption) (r *CorporateTreasuryService) {
	r = &CorporateTreasuryService{}
	r.Options = opts
	r.CashFlow = NewCorporateTreasuryCashFlowService(opts...)
	return
}
