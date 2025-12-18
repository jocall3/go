// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"github.com/jocall3/1231-go/option"
)

// SustainabilityInvestmentService contains methods and other services that help
// with interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSustainabilityInvestmentService] method instead.
type SustainabilityInvestmentService struct {
	Options []option.RequestOption
}

// NewSustainabilityInvestmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSustainabilityInvestmentService(opts ...option.RequestOption) (r *SustainabilityInvestmentService) {
	r = &SustainabilityInvestmentService{}
	r.Options = opts
	return
}
