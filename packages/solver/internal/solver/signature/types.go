package signature

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"solver/bindings/plug_router"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type Transaction struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Data  hexutil.Bytes  `json:"data"`
	Value *big.Int       `json:"value"`
	Gas   *big.Int       `json:"gas"`
}

type EIP712Domain struct {
	Name              string         `json:"name"`
	Version           string         `json:"version"`
	ChainId           *big.Int       `json:"chainId"`
	VerifyingContract common.Address `json:"verifyingContract"`
}

type Selector uint8

var (
	Call          Selector = 0
	DelegateCall  Selector = 1
	CallWithValue Selector = 2
	StaticCall    Selector = 3
	ForwardedCall Selector = 4
)

// Plug represents a single transaction to be executed as part of a bundle.
// It includes all necessary data for contract interaction and dynamic data updates.
type Plug struct {
	Selector Selector       `json:"selector"`
	To       common.Address `json:"to"`
	Data     hexutil.Bytes  `json:"data"`
	Value    *big.Int       `json:"value"`
	Updates  []coil.Update  `json:"updates,omitempty"`
	Meta     any            `json:"meta,omitempty"`
}

func (p Plug) Wrap() (*plug_router.PlugTypesLibPlug, error) {
	updates := make([]plug_router.PlugTypesLibUpdate, len(p.Updates))
	for index, update := range p.Updates {
		updates[index] = update.Wrap()
	}

	data, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 8}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
		{Type: abi.Type{T: abi.BytesTy}},
	}.Pack(p.Selector, p.To, p.Value, []byte(p.Data))
	if err != nil {
		return nil, err
	}

	return &plug_router.PlugTypesLibPlug{
		Data:    data,
		Updates: updates,
	}, nil
}

type Plugs struct {
	Socket common.Address `json:"socket"`
	Plugs  []Plug         `json:"plugs"`
	Solver []byte         `json:"solver"`
	Salt   []byte         `json:"salt"`
}

func (p Plugs) Wrap() (*plug_router.PlugTypesLibPlugs, error) {
	var plugs []plug_router.PlugTypesLibPlug
	for _, plug := range p.Plugs {
		plug, err := plug.Wrap()
		if err != nil {
			return nil, err
		}
		plugs = append(plugs, *plug)
	}

	return &plug_router.PlugTypesLibPlugs{
		Socket: p.Socket,
		Plugs:  plugs,
		Solver: p.Solver,
		Salt:   p.Salt,
	}, nil
}

// LivePlugs is the central model for transaction processing throughout the application.
// It represents a bundle of transactions (Plugs) with associated metadata and signature
// for on-chain execution via the Plug router contract.
//
// This is the single source of truth for all transaction data, handling both storage
// in the database and on-chain execution.
type LivePlugs struct {
	// Database fields (serialized to database)
	Id      string `json:"id,omitempty" gorm:"primaryKey;type:text"` // Changed ID to Id for consistency with GORM conventions
	ChainId uint64 `json:"chainId" gorm:"type:int"`
	From    string `json:"from,omitempty" gorm:"type:text"`
	// Pre-packed transaction data for the router contract (for caching/storage)
	Data string `json:"data,omitempty" gorm:"type:bytea"`
	// Reference to the Intent that created this LivePlugs
	IntentId  string    `json:"intentId,omitempty" gorm:"type:text"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"deletedAt,omitempty" gorm:"index"`

	// Core LivePlugs fields used for on-chain execution
	// Collection of transactions to execute atomically
	Plugs Plugs `json:"plugs" gorm:"serializer:json"`
	// EIP-712 signature authorizing the execution
	Signature []byte `json:"signature" gorm:"type:bytea"`
}

func (l LivePlugs) Wrap() (*plug_router.PlugTypesLibLivePlugs, error) {
	plugs, err := l.Plugs.Wrap()
	if err != nil {
		return nil, err
	}
	return &plug_router.PlugTypesLibLivePlugs{
		Plugs:     *plugs,
		Signature: l.Signature,
	}, nil
}

// Helper method to get router contract address for this chain
func (l *LivePlugs) GetRouterAddress() common.Address {
	if router, ok := references.Networks[l.ChainId].References["plug"]["router"]; ok {
		return common.HexToAddress(router)
	}
	return common.Address{}
}

// Helper method to get packed call data
func (l *LivePlugs) GetCallData() ([]byte, error) {
	routerAbi, err := plug_router.PlugRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get router ABI: %w", err)
	}

	livePlugs, err := l.Wrap()
	if err != nil {
		return nil, err
	}

	livePlugSlice := []plug_router.PlugTypesLibLivePlugs{*livePlugs}

	plugCalldata, err := routerAbi.Pack("plug0", livePlugSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to pack calldata: %w", err)
	}

	// Add identifier for tracing
	identifier := []byte("plug")
	return append(plugCalldata, identifier...), nil
}

// Execute submits the LivePlugs transaction to the blockchain and returns the transaction hash
func (l *LivePlugs) Execute() (string, error) {
	// Get router address for this chain
	routerAddress := l.GetRouterAddress()
	if (routerAddress == common.Address{}) {
		return "", fmt.Errorf("router address not found for chain ID %d", l.ChainId)
	}

	// Get RPC URL for the chain
	rpcURL, err := utils.GetRPCURL(l.ChainId)
	if err != nil {
		return "", fmt.Errorf("failed to get RPC URL for chain %d: %w", l.ChainId, err)
	}

	// Connect to the node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", fmt.Errorf("failed to connect to blockchain: %w", err)
	}
	defer client.Close()

	// Get the operator's account from environment
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if privateKeyHex == "" {
		return "", fmt.Errorf("PRIVATE_KEY environment variable not set")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("failed to load private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}

	// Get the sender address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the latest nonce for the account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get the current gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get gas price: %w", err)
	}

	// Get plug execution calldata
	calldata, err := l.GetCallData()
	if err != nil {
		return "", fmt.Errorf("failed to get calldata: %w", err)
	}

	// Estimate gas for the transaction (with some buffer)
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &routerAddress,
		GasPrice: gasPrice,
		Data:     calldata,
	})
	if err != nil {
		return "", fmt.Errorf("failed to estimate gas: %w", err)
	}
	gasLimit = uint64(float64(gasLimit) * 1.2) // Add 20% buffer

	// Create transaction
	chainID := big.NewInt(int64(l.ChainId))
	tx := types.NewTransaction(
		nonce,
		routerAddress,
		big.NewInt(0), // No value transfer
		gasLimit,
		gasPrice,
		calldata,
	)

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Submit the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}

	// Return the transaction hash
	return signedTx.Hash().Hex(), nil
}

// BeforeCreate is automatically called by GORM before inserting a new record.
// It generates a UUID if one isn't provided and ensures the packed calldata
// field is populated for backwards compatibility with simulation code.
func (l *LivePlugs) BeforeCreate(tx *gorm.DB) error {
	if l.Id == "" {
		l.Id = utils.GenerateUUID()
	}

	// Pre-pack the transaction data for storage
	data, err := l.GetCallData()
	if err != nil {
		return err
	}
	l.Data = hexutil.Bytes(data).String()

	return nil
}

// BeforeSave is automatically called by GORM before any update operation.
// It ensures the packed Data field stays in sync with any changes to the Plugs.
func (l *LivePlugs) BeforeSave(tx *gorm.DB) error {
	data, err := l.GetCallData()
	if err != nil {
		return err
	}
	l.Data = hexutil.Bytes(data).String()

	return nil
}

// AfterFind is a placeholder hook for future deserialization needs.
// Currently, all deserialization happens automatically through GORM tags.
func (l *LivePlugs) AfterFind(tx *gorm.DB) error {
	return nil
}

type Result struct {
	Success bool          `json:"success"`
	Result  hexutil.Bytes `json:"result"`
}
