package types

import (
	"math/big"
)

type Transaction struct {
	Transaction string   `json:"transaction"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Value       *big.Int `json:"value"`
	Gas         uint64   `json:"gas"`
}

type ActionInputs interface {
	Validate() error
	Build(chainId int, from string) (*Transaction, error)
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
