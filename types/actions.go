package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

/*
Action represents a generic action with its type and inputs. This is the
standard shape for a new action to be added to the intent that will be parsed,
built and sent back to the caller of the endpoint.
*/
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
	Build(provider *ethclient.Client, chainId int, from string) ([]*Transaction, error)
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
