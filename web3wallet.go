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

// Web3WalletService contains methods and other services that help with interacting
// with the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3WalletService] method instead.
type Web3WalletService struct {
	Options []option.RequestOption
}

// NewWeb3WalletService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWeb3WalletService(opts ...option.RequestOption) (r *Web3WalletService) {
	r = &Web3WalletService{}
	r.Options = opts
	return
}

// Retrieves a list of all securely linked cryptocurrency wallets (e.g., MetaMask,
// Ledger integration), showing their addresses, associated networks, and
// verification status.
func (r *Web3WalletService) List(ctx context.Context, query Web3WalletListParams, opts ...option.RequestOption) (res *Web3WalletListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Initiates the process to securely connect a new cryptocurrency wallet to the
// user's profile, typically involving a signed message or OAuth flow from the
// wallet provider.
func (r *Web3WalletService) Connect(ctx context.Context, body Web3WalletConnectParams, opts ...option.RequestOption) (res *CryptoWalletConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "web3/wallets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the current balances of all recognized crypto assets within a specific
// connected wallet.
func (r *Web3WalletService) GetBalances(ctx context.Context, walletID string, query Web3WalletGetBalancesParams, opts ...option.RequestOption) (res *Web3WalletGetBalancesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required walletId parameter")
		return
	}
	path := fmt.Sprintf("web3/wallets/%s/balances", walletID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CryptoWalletConnection struct {
	ID                 string                       `json:"id"`
	BlockchainNetwork  string                       `json:"blockchainNetwork"`
	LastSynced         time.Time                    `json:"lastSynced" format:"date-time"`
	ReadAccessGranted  bool                         `json:"readAccessGranted"`
	Status             CryptoWalletConnectionStatus `json:"status"`
	WalletAddress      string                       `json:"walletAddress"`
	WalletProvider     string                       `json:"walletProvider"`
	WriteAccessGranted bool                         `json:"writeAccessGranted"`
	JSON               cryptoWalletConnectionJSON   `json:"-"`
}

// cryptoWalletConnectionJSON contains the JSON metadata for the struct
// [CryptoWalletConnection]
type cryptoWalletConnectionJSON struct {
	ID                 apijson.Field
	BlockchainNetwork  apijson.Field
	LastSynced         apijson.Field
	ReadAccessGranted  apijson.Field
	Status             apijson.Field
	WalletAddress      apijson.Field
	WalletProvider     apijson.Field
	WriteAccessGranted apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CryptoWalletConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cryptoWalletConnectionJSON) RawJSON() string {
	return r.raw
}

type CryptoWalletConnectionStatus string

const (
	CryptoWalletConnectionStatusConnected           CryptoWalletConnectionStatus = "connected"
	CryptoWalletConnectionStatusDisconnected        CryptoWalletConnectionStatus = "disconnected"
	CryptoWalletConnectionStatusPendingVerification CryptoWalletConnectionStatus = "pending_verification"
)

func (r CryptoWalletConnectionStatus) IsKnown() bool {
	switch r {
	case CryptoWalletConnectionStatusConnected, CryptoWalletConnectionStatusDisconnected, CryptoWalletConnectionStatusPendingVerification:
		return true
	}
	return false
}

type Web3WalletListResponse struct {
	Data []CryptoWalletConnection   `json:"data"`
	JSON web3WalletListResponseJSON `json:"-"`
	PaginatedList
}

// web3WalletListResponseJSON contains the JSON metadata for the struct
// [Web3WalletListResponse]
type web3WalletListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3WalletListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletListResponseJSON) RawJSON() string {
	return r.raw
}

type Web3WalletGetBalancesResponse struct {
	Data []Web3WalletGetBalancesResponseData `json:"data"`
	JSON web3WalletGetBalancesResponseJSON   `json:"-"`
	PaginatedList
}

// web3WalletGetBalancesResponseJSON contains the JSON metadata for the struct
// [Web3WalletGetBalancesResponse]
type web3WalletGetBalancesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Web3WalletGetBalancesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletGetBalancesResponseJSON) RawJSON() string {
	return r.raw
}

type Web3WalletGetBalancesResponseData struct {
	AssetName       string                                `json:"assetName"`
	AssetSymbol     string                                `json:"assetSymbol"`
	Balance         float64                               `json:"balance"`
	ContractAddress string                                `json:"contractAddress"`
	UsdValue        float64                               `json:"usdValue"`
	JSON            web3WalletGetBalancesResponseDataJSON `json:"-"`
}

// web3WalletGetBalancesResponseDataJSON contains the JSON metadata for the struct
// [Web3WalletGetBalancesResponseData]
type web3WalletGetBalancesResponseDataJSON struct {
	AssetName       apijson.Field
	AssetSymbol     apijson.Field
	Balance         apijson.Field
	ContractAddress apijson.Field
	UsdValue        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Web3WalletGetBalancesResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r web3WalletGetBalancesResponseDataJSON) RawJSON() string {
	return r.raw
}

type Web3WalletListParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [Web3WalletListParams]'s query parameters as `url.Values`.
func (r Web3WalletListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type Web3WalletConnectParams struct {
	BlockchainNetwork param.Field[string] `json:"blockchainNetwork,required"`
	SignedMessage     param.Field[string] `json:"signedMessage,required"`
	WalletAddress     param.Field[string] `json:"walletAddress,required"`
	WalletProvider    param.Field[string] `json:"walletProvider,required"`
}

func (r Web3WalletConnectParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Web3WalletGetBalancesParams struct {
	// The maximum number of items to return.
	Limit param.Field[int64] `query:"limit"`
	// The number of items to skip before starting to collect the result set.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [Web3WalletGetBalancesParams]'s query parameters as
// `url.Values`.
func (r Web3WalletGetBalancesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
