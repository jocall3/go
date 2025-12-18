package web3

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// Common errors in the Web3 domain.
var (
	ErrInvalidAddress  = errors.New("invalid blockchain address format")
	ErrInvalidChain    = errors.New("unsupported or invalid blockchain network")
	ErrInsufficientGas = errors.New("insufficient funds for gas")
	ErrTxFailed        = errors.New("transaction failed on chain")
)

// Chain represents a supported blockchain network.
type Chain string

const (
	ChainEthereum MainnetChain = "ethereum"
	ChainPolygon  MainnetChain = "polygon"
	ChainBSC      MainnetChain = "bsc"
	ChainSolana   MainnetChain = "solana"
	ChainArbitrum MainnetChain = "arbitrum"
	ChainOptimism MainnetChain = "optimism"
)

type MainnetChain Chain

// String returns the string representation of the chain.
func (c Chain) String() string {
	return string(c)
}

// TokenType distinguishes between different token standards.
type TokenType string

const (
	TokenTypeNative  TokenType = "NATIVE"
	TokenTypeERC20   TokenType = "ERC20"
	TokenTypeERC721  TokenType = "ERC721"
	TokenTypeERC1155 TokenType = "ERC1155"
	TokenTypeSPL     TokenType = "SPL" // Solana Program Library
)

// TransactionStatus represents the lifecycle state of a blockchain transaction.
type TransactionStatus string

const (
	TxStatusPending   TransactionStatus = "PENDING"
	TxStatusConfirmed TransactionStatus = "CONFIRMED"
	TxStatusFailed    TransactionStatus = "FAILED"
	TxStatusDropped   TransactionStatus = "DROPPED"
)

// Wallet represents a user's cryptographic wallet on a specific chain.
type Wallet struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Address   string    `json:"address"`
	Chain     Chain     `json:"chain"`
	Label     string    `json:"label,omitempty"`
	IsCustodial bool    `json:"is_custodial"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewWallet creates a new Wallet instance with validation.
func NewWallet(id, userID, address string, chain Chain, custodial bool) (*Wallet, error) {
	w := &Wallet{
		ID:          id,
		UserID:      userID,
		Address:     address,
		Chain:       chain,
		IsCustodial: custodial,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	if err := w.Validate(); err != nil {
		return nil, err
	}
	return w, nil
}

// Validate checks if the wallet address format is valid for the given chain.
func (w *Wallet) Validate() error {
	if w.Address == "" {
		return ErrInvalidAddress
	}
	
	// Basic regex validation for EVM addresses (0x followed by 40 hex chars)
	if w.Chain == Chain(ChainEthereum) || w.Chain == Chain(ChainPolygon) || w.Chain == Chain(ChainBSC) {
		match, _ := regexp.MatchString("^0x[a-fA-F0-9]{40}$", w.Address)
		if !match {
			return ErrInvalidAddress
		}
	}
	
	// Basic length check for Solana (Base58 strings are usually 32-44 chars)
	if w.Chain == Chain(ChainSolana) {
		if len(w.Address) < 32 || len(w.Address) > 44 {
			return ErrInvalidAddress
		}
	}

	return nil
}

// NFT represents a Non-Fungible Token asset.
type NFT struct {
	ID              string                 `json:"id"`
	ContractAddress string                 `json:"contract_address"`
	TokenID         string                 `json:"token_id"` // String to handle large big.Int values
	Chain           Chain                  `json:"chain"`
	Standard        TokenType              `json:"standard"`
	OwnerAddress    string                 `json:"owner_address"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description"`
	ImageURL        string                 `json:"image_url"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	MintedAt        *time.Time             `json:"minted_at,omitempty"`
	LastSyncedAt    time.Time              `json:"last_synced_at"`
}

// Collection represents a grouping of NFTs (Smart Contract).
type Collection struct {
	ContractAddress string    `json:"contract_address"`
	Chain           Chain     `json:"chain"`
	Name            string    `json:"name"`
	Symbol          string    `json:"symbol"`
	Standard        TokenType `json:"standard"`
	Verified        bool      `json:"verified"`
}

// Transaction represents a state change on the blockchain.
type Transaction struct {
	ID            string            `json:"id"`
	Hash          string            `json:"hash"`
	Chain         Chain             `json:"chain"`
	FromAddress   string            `json:"from_address"`
	ToAddress     string            `json:"to_address"`
	Value         string            `json:"value"` // Amount in Wei/Lamports as string
	GasPrice      string            `json:"gas_price,omitempty"`
	GasUsed       uint64            `json:"gas_used,omitempty"`
	Nonce         uint64            `json:"nonce"`
	BlockNumber   uint64            `json:"block_number,omitempty"`
	BlockHash     string            `json:"block_hash,omitempty"`
	Status        TransactionStatus `json:"status"`
	InputData     string            `json:"input_data,omitempty"` // Hex encoded input data
	Confirmations uint64            `json:"confirmations"`
	Timestamp     time.Time         `json:"timestamp"`
}

// IsConfirmed checks if the transaction is considered final based on status.
func (tx *Transaction) IsConfirmed() bool {
	return tx.Status == TxStatusConfirmed
}

// ExplorerURL returns the public block explorer URL for this transaction.
func (tx *Transaction) ExplorerURL() string {
	switch tx.Chain {
	case Chain(ChainEthereum):
		return "https://etherscan.io/tx/" + tx.Hash
	case Chain(ChainPolygon):
		return "https://polygonscan.com/tx/" + tx.Hash
	case Chain(ChainBSC):
		return "https://bscscan.com/tx/" + tx.Hash
	case Chain(ChainSolana):
		return "https://solscan.io/tx/" + tx.Hash
	default:
		return ""
	}
}

// Balance represents a snapshot of an asset holding.
type Balance struct {
	WalletAddress   string    `json:"wallet_address"`
	ContractAddress string    `json:"contract_address,omitempty"` // Empty for native currency
	Chain           Chain     `json:"chain"`
	Amount          string    `json:"amount"` // Raw amount (Wei/Lamports)
	Decimals        int       `json:"decimals"`
	Symbol          string    `json:"symbol"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// IsNative returns true if the balance represents the chain's native currency.
func (b *Balance) IsNative() bool {
	return b.ContractAddress == "" || b.ContractAddress == "0x0000000000000000000000000000000000000000"
}

// NormalizeAddress ensures addresses are stored in a consistent format (lowercase for EVM).
func NormalizeAddress(address string) string {
	return strings.ToLower(strings.TrimSpace(address))
}