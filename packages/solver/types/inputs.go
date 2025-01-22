package types

import (
	"fmt"
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
	AmountIn string  `json:"amountIn"` // Raw amount to send (deposit).
	Target   *string  `json:"target"`   // Address of smart contract to interact with.
}

func (i *DepositInputs) Validate() error {
	if _, _, err := utils.ParseAddressAndDecimals(i.TokenOut); err != nil {
		return utils.ErrInvalidField("tokenOut", err.Error())
	}

	_, decimals, err := utils.ParseAddressAndDecimals(i.TokenIn);
	if err != nil {
		return utils.ErrInvalidField("tokenIn", err.Error())
	}

	if _, err := utils.StringToUint(i.AmountIn, decimals); err != nil {
		return utils.ErrInvalidField("amountIn", err.Error())
	}
	if !utils.IsAddress(*i.Target) {
		return utils.ErrInvalidAddress("target", *i.Target)
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
	AmountOut  string   `json:"amountOut"`  // Raw amount of tokens to borrow.
}

func (i *BorrowInputs) Validate() error {
	if _, _, err := utils.ParseAddressAndDecimals(i.Collateral); err != nil {
		return utils.ErrInvalidField("collateral", err.Error())
	}

	_, decimals, err := utils.ParseAddressAndDecimals(i.TokenOut)
	if err != nil {
		return utils.ErrInvalidField("tokenOut", err.Error())
	}

	if _, err := utils.StringToUint(i.AmountOut, decimals); err != nil {
		return utils.ErrInvalidField("amountOut", err.Error())
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
	AmountIn string  `json:"amountIn"` // Raw amount of tokens to send.
	Target   *string `json:"target"`   // Address of smart contract to interact with.
}

func (i *RedeemInputs) Validate() error {
	if _, _, err := utils.ParseAddressAndDecimals(i.TokenOut); err != nil {
		return utils.ErrInvalidField("tokenOut", err.Error())
	}
	
	_, decimals, err := utils.ParseAddressAndDecimals(i.TokenIn)
	if err != nil {
		return utils.ErrInvalidField("tokenIn", err.Error())
	}

    if _, err := utils.StringToUint(i.AmountIn, decimals); err != nil {
		return utils.ErrInvalidField("amountIn", err.Error())
	}

	return nil
}

func (i *RedeemInputs) GetProtocol() Protocol {
	return i.Protocol
}

type WithdrawInputs struct {
	BaseInputs
	TokenOut  string  `json:"tokenOut"` // Address of the token to receive (redeeming for).
	AmountOut string  `json:"amountIn"` // Raw amount of tokens to send.
}

func (i *WithdrawInputs) Validate() error {
	_, decimals, err := utils.ParseAddressAndDecimals(i.TokenOut)
	if err != nil {
		return utils.ErrInvalidField("tokenOut", err.Error())
	}

	if _, err := utils.StringToUint(i.AmountOut, decimals); err != nil {
		return utils.ErrInvalidField("amountOut", err.Error())
	}
    
	return nil
}

func (i *WithdrawInputs) GetProtocol() Protocol {
	return i.Protocol
}

type RepayInputs struct {
	BaseInputs
	TokenIn  string  `json:"tokenIn"`  // Address of the token to repay.
	AmountIn string  `json:"amountIn"` // Raw amount of tokens to repay.
}

func (i *RepayInputs) Validate() error {
	_, decimals, err := utils.ParseAddressAndDecimals(i.TokenIn)
	if err != nil {
		return utils.ErrInvalidField("tokenIn", err.Error())
	}
	
	if _, err := utils.StringToUint(i.AmountIn, decimals); err != nil {
		return utils.ErrInvalidField("amountIn", err.Error())
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
	Operator  int    `json:"operator"`  // The operator to use for the threshold comparison.
	Threshold string `json:"threshold"` // The threshold value to compare against.
}

func (i *ThresholdInputs) Validate() error {
	if i.Operator != -1 && i.Operator != 1 {
		return utils.ErrInvalidField("operator", strconv.Itoa(i.Operator))
	}

	threshold, err := strconv.ParseFloat(i.Threshold, 64)
    if err != nil {
        return utils.ErrInvalidField("threshold", "must be a valid number")
    }
    
    if threshold < 0 || threshold > float64(utils.Uint256Max.Int64()) {
        return utils.ErrInvalidField("threshold", fmt.Sprintf("%f", threshold))
    }
	return nil
}

func (i *ThresholdInputs) GetProtocol() Protocol {
	return i.Protocol
}
