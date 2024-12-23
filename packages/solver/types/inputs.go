package types

import (
	"math/big"
	"solver/utils"
	"strconv"
)

type BaseInputs struct {
	Protocol Protocol `json:"protocol"`
	Action   Action   `json:"action"`
}

func (i *BaseInputs) Validate() error { 
	if i.Protocol == "" { 
		return utils.ErrInvalidField("protocol", "must be non-empty string")
	}
	if i.Action == "" { 
		return utils.ErrInvalidField("action", "must be non-empty string")
	}

	return nil
}

type ActionInputs interface {
	Validate() error
	GetProtocol() Protocol
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

type WithdrawInputs struct {
	BaseInputs
	TokenOut  string  `json:"tokenOut"` // Address of the token to receive (redeeming for).
	AmountOut big.Int `json:"amountIn"` // Raw amount of tokens to send.
}

func (i *WithdrawInputs) Validate() error {
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if i.AmountOut.Cmp(big.NewInt(0)) >= 0 && i.AmountOut.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amountIn", i.AmountOut.String())
	}

	return nil
}

func (i *WithdrawInputs) GetProtocol() Protocol {
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
	Token string `json:"token"` // Address of the token to harvest.
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

type ThresholdInputs struct {
	BaseInputs
	Operator  int        `json:"operator"`  // The operator to use for the threshold comparison.
	Threshold *big.Float `json:"threshold"` // The threshold value to compare against.
}

func (i *ThresholdInputs) Validate() error {
	if i.Operator != -1 && i.Operator != 1 {
		return utils.ErrInvalidField("operator", strconv.Itoa(i.Operator))
	}
	if i.Threshold.Cmp(big.NewFloat(0)) >= 0 && i.Threshold.Cmp(new(big.Float).SetInt(utils.Uint256Max)) > 0 {
		return utils.ErrInvalidField("threshold", i.Threshold.String())
	}
	return nil
}

func (i *ThresholdInputs) GetProtocol() Protocol {
	return i.Protocol
}
