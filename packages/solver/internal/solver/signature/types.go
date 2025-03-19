package signature

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_router"
	"solver/internal/bindings/references"
	"solver/internal/solver/coil"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm"
)

type Transaction struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Data  []byte         `json:"data"`
	Value *big.Int       `json:"value"`
	Gas   *big.Int       `json:"gas"`
}

type EIP712Domain struct {
	Name              string         `json:"name"`
	Version           string         `json:"version"`
	ChainId           *big.Int       `json:"chainId"`
	VerifyingContract common.Address `json:"verifyingContract"`
}

type MinimalPlug struct {
	To    common.Address `json:"to"`
	Data  []byte         `json:"data"`
	Value *big.Int       `json:"value"`
}

// Plug represents a single transaction to be executed as part of a bundle.
// It includes all necessary data for contract interaction and dynamic data updates.
type Plug struct {
	// Selector determines call type: 0 for standard call, 1 for delegatecall, 2 call with value, 3 static call
	Selector uint8          `json:"selector"`
	To       common.Address `json:"to"`
	Data     []byte         `json:"data"`
	Value    *big.Int       `json:"value"`
	// Updates contains dynamic data modifications to be applied at execution time
	Updates []coil.Update `json:"updates"`

	// Meta contains additional protocol-specific data (not used for execution)
	Meta any `json:"meta,omitempty"`
}

func (p Plug) Wrap() plug_router.PlugTypesLibPlug {
	updates := make([]plug_router.PlugTypesLibUpdate, len(p.Updates))
	for index, update := range p.Updates {
		updates[index] = update.Wrap()
	}

	return plug_router.PlugTypesLibPlug{
		Selector: p.Selector,
		To:       p.To,
		Data:     p.Data,
		Value:    p.Value,
		Updates:  updates,
	}
}

func (p Plug) Minify() *MinimalPlug {
	return &MinimalPlug{
		To:    p.To,
		Data:  p.Data,
		Value: p.Value,
	}
}

type Plugs struct {
	Socket common.Address `json:"socket"`
	Plugs  []Plug         `json:"plugs"`
	Solver []byte         `json:"solver"`
	Salt   []byte         `json:"salt"`
}

func (p Plugs) Wrap() plug_router.PlugTypesLibPlugs {
	var plugs []plug_router.PlugTypesLibPlug
	for _, plug := range p.Plugs {
		plugs = append(plugs, plug.Wrap())
	}

	return plug_router.PlugTypesLibPlugs{
		Socket: p.Socket,
		Plugs:  plugs,
		Solver: p.Solver,
		Salt:   p.Salt,
	}
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

func (l LivePlugs) Wrap() plug_router.PlugTypesLibLivePlugs {
	return plug_router.PlugTypesLibLivePlugs{
		Plugs:     l.Plugs.Wrap(),
		Signature: l.Signature,
	}
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

	plugCalldata, err := routerAbi.Pack("plug0", l.Wrap())
	if err != nil {
		return nil, fmt.Errorf("failed to pack calldata: %w", err)
	}

	// Add identifier for tracing
	identifier := []byte("plug")
	return append(plugCalldata, identifier...), nil
}

// GetRawPlugs returns the individual transactions that would be executed
func (l *LivePlugs) GetRawPlugs() []Transaction {
	txs := make([]Transaction, len(l.Plugs.Plugs))
	identifier := []byte("plug")

	for idx, plug := range l.Plugs.Plugs {
		data := append(plug.Data, identifier...)
		txs[idx] = Transaction{
			From:  common.HexToAddress(l.From),
			To:    plug.To,
			Data:  data,
			Value: plug.Value,
			Gas:   nil, // Will be estimated during simulation
		}
	}

	return txs
}

type Result struct {
	Success bool   `json:"success"`
	Result  []byte `json:"result"`
}

// GORM lifecycle hooks for LivePlugs

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
