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

// Web3TransactionService contains methods and other services that help with
// interacting with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3TransactionService] method instead.
type Web3TransactionService struct {
	Options []option.RequestOption
}

// NewWeb3TransactionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3TransactionService(opts ...option.RequestOption) (r *Web3TransactionService) {
	r = &Web3TransactionService{}
	r.Options = opts
	return
}

// Prepares and initiates a cryptocurrency transfer from a connected wallet to a
// specified recipient address. Requires user confirmation (e.g., via wallet
// signature).
func (r *Web3TransactionService) InitiateTransfer(ctx context.Context, body Web3TransactionInitiateTransferParams, opts ...option.RequestOption) (res *Web3TransactionInitiateTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/transactions/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Web3TransactionInitiateTransferResponse struct {
	Message    string                                        `json:"message"`
	Status     Web3TransactionInitiateTransferResponseStatus `json:"status"`
	TransferID string                                        `json:"transferId"`
	JSON       web3TransactionInitiateTransferResponseJSON   `json:"-"`
}

// web3TransactionInitiateTransferResponseJSON contains the JSON metadata for the
// struct [Web3TransactionInitiateTransferResponse]
type web3TransactionInitiateTransferResponseJSON struct {
	Message     apijson.Field
	Status      apijson.Field
	TransferID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3TransactionInitiateTransferResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3TransactionInitiateTransferResponseJSON) RawJSON() string {
	return r.raw
}

type Web3TransactionInitiateTransferResponseStatus string

const (
	Web3TransactionInitiateTransferResponseStatusPendingSignature Web3TransactionInitiateTransferResponseStatus = "pending_signature"
	Web3TransactionInitiateTransferResponseStatusSubmitted        Web3TransactionInitiateTransferResponseStatus = "submitted"
	Web3TransactionInitiateTransferResponseStatusConfirmed        Web3TransactionInitiateTransferResponseStatus = "confirmed"
	Web3TransactionInitiateTransferResponseStatusFailed           Web3TransactionInitiateTransferResponseStatus = "failed"
)

func (r Web3TransactionInitiateTransferResponseStatus) IsKnown() bool {
	switch r {
	case Web3TransactionInitiateTransferResponseStatusPendingSignature, Web3TransactionInitiateTransferResponseStatusSubmitted, Web3TransactionInitiateTransferResponseStatusConfirmed, Web3TransactionInitiateTransferResponseStatusFailed:
		return true
	}
	return false
}

type Web3TransactionInitiateTransferParams struct {
	Amount            param.Field[float64] `json:"amount,required"`
	AssetSymbol       param.Field[string]  `json:"assetSymbol,required"`
	BlockchainNetwork param.Field[string]  `json:"blockchainNetwork,required"`
	RecipientAddress  param.Field[string]  `json:"recipientAddress,required"`
	SourceWalletID    param.Field[string]  `json:"sourceWalletId,required"`
	GasPriceGwei      param.Field[float64] `json:"gasPriceGwei"`
	Memo              param.Field[string]  `json:"memo"`
}

func (r Web3TransactionInitiateTransferParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
