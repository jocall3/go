package web3

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Common ERC721 Method Selectors
var (
	methodBalanceOf          = crypto.Keccak256([]byte("balanceOf(address)"))[:4]
	methodTokenOfOwnerByIndex = crypto.Keccak256([]byte("tokenOfOwnerByIndex(address,uint256)"))[:4]
	methodTokenURI           = crypto.Keccak256([]byte("tokenURI(uint256)"))[:4]
)

// Service defines the capabilities of the Web3 integration layer.
type Service interface {
	// Authentication
	VerifyWalletSignature(ctx context.Context, address string, message string, signatureHex string) (bool, error)

	// Data Retrieval
	GetNativeBalance(ctx context.Context, address string) (*big.Int, error)
	GetNFTs(ctx context.Context, contractAddress string, ownerAddress string) ([]NFT, error)

	// Transactions
	SendNative(ctx context.Context, privateKeyHex string, toAddress string, amountWei *big.Int) (string, error)
	ExecuteContractInteraction(ctx context.Context, privateKeyHex string, contractAddress string, data []byte, value *big.Int) (string, error)
	
	// Lifecycle
	Close()
}

// NFT represents a simplified view of a Non-Fungible Token.
type NFT struct {
	ContractAddress string   `json:"contract_address"`
	TokenID         *big.Int `json:"token_id"`
	URI             string   `json:"uri"`
	Owner           string   `json:"owner"`
}

type service struct {
	client  *ethclient.Client
	chainID *big.Int
}

// NewService initializes a new Web3 service with the provided RPC URL.
func NewService(rpcURL string) (Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum client: %w", err)
	}

	// Validate connection and get Chain ID
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve chain id: %w", err)
	}

	return &service{
		client:  client,
		chainID: chainID,
	}, nil
}

func (s *service) Close() {
	if s.client != nil {
		s.client.Close()
	}
}

// VerifyWalletSignature checks if a given signature was signed by the owner of the address
// for the provided message, adhering to EIP-191.
func (s *service) VerifyWalletSignature(ctx context.Context, address string, message string, signatureHex string) (bool, error) {
	if !common.IsHexAddress(address) {
		return false, errors.New("invalid address format")
	}

	// Decode the signature
	sig, err := hexutil.Decode(signatureHex)
	if err != nil {
		return false, fmt.Errorf("invalid signature hex: %w", err)
	}

	// Handle signature V value normalization (27/28 -> 0/1)
	if len(sig) == 65 {
		if sig[64] == 27 || sig[64] == 28 {
			sig[64] -= 27
		}
	} else {
		return false, errors.New("invalid signature length")
	}

	// Construct the EIP-191 signed message
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	data := []byte(prefix + message)
	hash := crypto.Keccak256Hash(data)

	// Recover the public key from the signature
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %w", err)
	}

	// Convert public key to address
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// Compare addresses (case-insensitive)
	return strings.EqualFold(recoveredAddr.Hex(), address), nil
}

// GetNativeBalance retrieves the native token balance (ETH, MATIC, etc.) for an address.
func (s *service) GetNativeBalance(ctx context.Context, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := s.client.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve balance: %w", err)
	}
	return balance, nil
}

// GetNFTs attempts to fetch NFTs owned by a specific address from an ERC721 contract.
// It relies on the ERC721Enumerable extension (tokenOfOwnerByIndex).
func (s *service) GetNFTs(ctx context.Context, contractAddress string, ownerAddress string) ([]NFT, error) {
	contract := common.HexToAddress(contractAddress)
	owner := common.HexToAddress(ownerAddress)

	// 1. Get Balance
	balance, err := s.erc721BalanceOf(ctx, contract, owner)
	if err != nil {
		return nil, err
	}

	if balance.Cmp(big.NewInt(0)) == 0 {
		return []NFT{}, nil
	}

	// Limit the number of NFTs to fetch to prevent timeouts on large collections
	maxFetch := int(balance.Int64())
	if maxFetch > 50 {
		maxFetch = 50 // Cap at 50 for this implementation
	}

	var nfts []NFT

	// 2. Iterate tokens using tokenOfOwnerByIndex
	for i := 0; i < maxFetch; i++ {
		tokenID, err := s.erc721TokenOfOwnerByIndex(ctx, contract, owner, i)
		if err != nil {
			// If enumeration fails, we might stop or continue. Here we stop.
			// This likely means the contract is not Enumerable.
			return nfts, fmt.Errorf("contract does not support enumeration or call failed: %w", err)
		}

		uri, err := s.erc721TokenURI(ctx, contract, tokenID)
		if err != nil {
			uri = "" // Allow failure on URI fetch
		}

		nfts = append(nfts, NFT{
			ContractAddress: contractAddress,
			TokenID:         tokenID,
			URI:             uri,
			Owner:           ownerAddress,
		})
	}

	return nfts, nil
}

// SendNative sends native currency (ETH) from a private key to a destination address.
func (s *service) SendNative(ctx context.Context, privateKeyHex string, toAddress string, amountWei *big.Int) (string, error) {
	return s.ExecuteContractInteraction(ctx, privateKeyHex, toAddress, nil, amountWei)
}

// ExecuteContractInteraction handles generic transaction creation, signing, and broadcasting.
// It supports EIP-1559 dynamic fee transactions.
func (s *service) ExecuteContractInteraction(ctx context.Context, privateKeyHex string, toAddressStr string, data []byte, value *big.Int) (string, error) {
	// 1. Parse Private Key
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return "", fmt.Errorf("invalid private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress(toAddressStr)

	// 2. Get Nonce
	nonce, err := s.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %w", err)
	}

	// 3. Estimate Gas
	msg := ethereum.CallMsg{
		From:  fromAddress,
		To:    &toAddress,
		Gas:   0,
		Price: nil,
		Value: value,
		Data:  data,
	}
	gasLimit, err := s.client.EstimateGas(ctx, msg)
	if err != nil {
		// Fallback gas limit if estimation fails (risky, but sometimes necessary)
		gasLimit = 300000 
	} else {
		// Add a buffer to the estimated gas
		gasLimit = gasLimit + (gasLimit / 10)
	}

	// 4. Get Gas Tip and Fee Cap (EIP-1559)
	tipCap, err := s.client.SuggestGasTipCap(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get gas tip cap: %w", err)
	}

	head, err := s.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get latest header: %w", err)
	}
	
	// BaseFee * 2 + TipCap
	baseFee := head.BaseFee
	gasFeeCap := new(big.Int).Add(
		new(big.Int).Mul(baseFee, big.NewInt(2)),
		tipCap,
	)

	// 5. Create Transaction
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.chainID,
		Nonce:     nonce,
		GasTipCap: tipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	})

	// 6. Sign Transaction
	signer := types.LatestSignerForChainID(s.chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %w", err)
	}

	// 7. Broadcast
	err = s.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	return signedTx.Hash().Hex(), nil
}

// --- Internal ERC721 Helpers (Raw Calls) ---

func (s *service) erc721BalanceOf(ctx context.Context, contract, owner common.Address) (*big.Int, error) {
	// Construct calldata: balanceOf(address)
	// Selector: 0x70a08231
	// Arg: address (padded to 32 bytes)
	data := make([]byte, 0)
	data = append(data, methodBalanceOf...)
	data = append(data, common.LeftPadBytes(owner.Bytes(), 32)...)

	result, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("balanceOf call failed: %w", err)
	}

	if len(result) == 0 {
		return nil, errors.New("balanceOf returned no data")
	}

	return new(big.Int).SetBytes(result), nil
}

func (s *service) erc721TokenOfOwnerByIndex(ctx context.Context, contract, owner common.Address, index int) (*big.Int, error) {
	// Construct calldata: tokenOfOwnerByIndex(address,uint256)
	// Selector: 0x2f745c59
	data := make([]byte, 0)
	data = append(data, methodTokenOfOwnerByIndex...)
	data = append(data, common.LeftPadBytes(owner.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(big.NewInt(int64(index)).Bytes(), 32)...)

	result, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetBytes(result), nil
}

func (s *service) erc721TokenURI(ctx context.Context, contract common.Address, tokenID *big.Int) (string, error) {
	// Construct calldata: tokenURI(uint256)
	// Selector: 0xc87b56dd
	data := make([]byte, 0)
	data = append(data, methodTokenURI...)
	data = append(data, common.LeftPadBytes(tokenID.Bytes(), 32)...)

	result, err := s.client.CallContract(ctx, ethereum.CallMsg{
		To:   &contract,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	// Decode String (ABI encoding: offset, length, data)
	// This is a simplified decoder for standard ABI strings
	if len(result) < 64 {
		return "", errors.New("invalid string data returned")
	}

	// Read length (usually at offset 32 if dynamic offset is 32)
	// Standard ABI return for string: [32 bytes offset][32 bytes length][bytes data...]
	// We assume the offset points to the next word (32).
	
	// Skip offset (first 32 bytes)
	lengthBytes := result[32:64]
	length := new(big.Int).SetBytes(lengthBytes).Uint64()

	if uint64(len(result)) < 64+length {
		return "", errors.New("incomplete string data")
	}

	strBytes := result[64 : 64+length]
	return string(strBytes), nil
}