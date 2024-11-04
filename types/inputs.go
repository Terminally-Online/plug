package types

import (
	"math/big"
	"solver/utils"
)

type BaseInputs struct {
	ChainId  int      `json:"chainId"`
	Protocol Protocol `json:"protocol"`
	Target   *string  `json:"target"`
}

type ActionInputs interface {
	Validate() error
	// GetProtocol() Protocol
}

type DepositInputs struct {
	BaseInputs
	TokenIn  string   `json:"tokenIn"`  // Address of the token to send (deposit).
	TokenOut string   `json:"tokenOut"` // Address of the token to receive (withdraw).
	AmountIn *big.Int `json:"amountIn"` // Raw amount to send (deposit).
	Target   *string  `json:"target"`   // Address of smart contract to interact with.
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
	Collateral string   `json:"collateral"` // Address of the collateral token (supplied).
	TokenOut   string   `json:"tokenOut"`   // Address of the token to receive (borrow).
	AmountOut  *big.Int `json:"amountOut"`  // Raw amount of tokens to borrow.
}

func (i *BorrowInputs) Validate() error {
	if !utils.IsAddress(i.Collateral) {
		return utils.ErrInvalidAddress("collateral", i.Collateral)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if i.AmountOut.Cmp(big.NewInt(0)) >= 0 && i.AmountOut.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amountOut", i.AmountOut.String())
	}
	return nil
}

func (i *BorrowInputs) GetProtocol() Protocol {
	return i.Protocol
}

type RedeemInputs struct {
	BaseInputs
	TokenIn  string  `json:"tokenIn"`  // Address of the token to send (redeem).
	TokenOut string  `json:"tokenOut"` // Address of the token to receive (redeeming for).
	AmountIn big.Int `json:"amountIn"` // Raw amount of tokens to send.
	Target   *string `json:"target"`   // Address of smart contract to interact with.
}

func (i *RedeemInputs) Validate() error {
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

func (i *RedeemInputs) GetProtocol() Protocol {
	return i.Protocol
}

type RepayInputs struct {
	BaseInputs
	TokenIn  string  `json:"tokenIn"`  // Address of the token to repay.
	AmountIn big.Int `json:"amountIn"` // Raw amount of tokens to repay.
}

func (i *RepayInputs) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if i.AmountIn.Cmp(big.NewInt(0)) >= 0 && i.AmountIn.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amountIn", i.AmountIn.String())
	}

	return nil
}

func (i *RepayInputs) GetProtocol() Protocol {
	return i.Protocol
}

type HarvestInputs struct { 
	BaseInputs
	Token    string `json:"token"`    // Address of the token to harvest.
}

func (i *HarvestInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	return nil
}

func (i *HarvestInputs) GetProtocol() Protocol {
	return i.Protocol
}
