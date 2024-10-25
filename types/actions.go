package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ActionSchema struct {
	Protocol string               `json:"protocol"`
	Type     string               `json:"type"`
	Values   []ActionSchemaValues `json:"values"`
}

type ActionSchemaValues struct {
	Label   string                     `json:"label"`
	Type    string                     `json:"type"`
	Options *[]ActionSchemaValueOption `json:"options"`
}

type ActionSchemaValueOption struct {
	Label     string  `json:"label"`
	Value     string  `json:"value"`
	Icon      *string `json:"icon"`
	Connector *int    `json:"connector"`
}

type Action struct {
	Type   string          `json:"type"`
	Inputs json.RawMessage `json:"inputs"`
}

type Transaction struct {
	Transaction string   `json:"transaction"`
	To          string   `json:"to"`
	Value       *big.Int `json:"value"`
}

type ActionInputs interface {
	Validate() error
	Get(provider *ethclient.Client, chainId int) (*ActionSchema, error)
	Post(provider *ethclient.Client, chainId int, from string) ([]*Transaction, error)
}

type ApproveInputs interface {
	ActionInputs
	GetType() int
	GetToken() string
	GetSpender() string
	GetTokenId() *big.Int
	GetAmount() *big.Int
	GetApproved() *bool
}

type BorrowInputs interface {
	ActionInputs
	GetProtocol() string
	GetCollateral() string
	GetTokenOut() string
	GetAmountOut() *big.Int
}

type DepositInputs interface {
	ActionInputs
	GetProtocol() string
	GetTokenIn() string
	GetTokenOut() string
	GetAmountIn() *big.Int
	GetTarget() *string
}

type HarvestInputs interface {
	ActionInputs
	GetProtocol() string
	GetToken() string
}

type RedeemInputs interface {
	ActionInputs
	GetProtocol() string
	GetTokenIn() string
	GetTokenOut() string
	GetAmountIn() *big.Int
	GetTarget() *string
}

type RepayInputs interface {
	ActionInputs
	GetProtocol() string
	GetTokenIn() string
	GetAmountIn() *big.Int
}

type RouteInputs interface {
	ActionInputs
	GetTokenIn() string
	GetTokenOut() string
	GetAmountIn() *big.Int
	GetSlippage() *big.Int
}

type SwapInputs interface {
	ActionInputs
	GetProtocol() string
	GetTokenIn() string
	GetTokenOut() string
	GetAmountIn() *big.Int
	GetSlippage() *big.Int
}

type TransferInputs interface {
	ActionInputs
	GetType() int
	GetToken() string
	GetRecipient() string
	GetAmount() *big.Int
}

type TransferFromInputs interface {
	ActionInputs
	GetType() int
	GetToken() string
	GetSender() string
	GetRecipient() string
	GetTokenId() *big.Int
	GetAmount() *big.Int
}
