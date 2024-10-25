package types

import (
	"math/big"
	"solver/utils"
)

type BaseInputs struct {
	Protocol Protocol `json:"protocol"`
	Target   *string  `json:"target"`
}

type ActionInputs interface {
	Validate() error
	GetProtocol() Protocol
}

type DepositInputs struct {
	BaseInputs
	TokenIn  string  `json:"tokenIn"`
	TokenOut string  `json:"tokenOut"`
	AmountIn big.Int `json:"amountIn"`
}

func (i *DepositInputs) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if i.AmountIn.Cmp(big.NewInt(0)) >= 0 && i.AmountIn.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amountIn", i.AmountIn.String())
	}
	return nil
}

func (i *DepositInputs) GetProtocol() Protocol {
	return i.Protocol
}

type BorrowInputs struct {
	BaseInputs
	TokenToBorrow string  `json:"tokenToBorrow"`
	Amount        big.Int `json:"amount"`
	Collateral    string  `json:"collateral"`
}

func (i *BorrowInputs) Validate() error {
	return nil
}

func (i *BorrowInputs) GetProtocol() Protocol {
	return i.Protocol
}
