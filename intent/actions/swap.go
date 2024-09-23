package intent

import (
	"solver/utils"
)

type SwapInputs struct {
	TokenIn        string `json:"tokenIn"`        // Address of the token to swap (sell).
	TokenOut       string `json:"tokenOut"`       // Address of the token to swap (buy).
	AmountIn       string `json:"amountIn"`       // Raw amount of tokens to swap (sell).
	Slippage       string `json:"slippage"`       // Slippage tolerance when executing the swap.
	PrimaryAddress string `json:"primaryAddress"` // Address of the smart contract to interact with.
}

func (i SwapInputs) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if !utils.IsUint(i.AmountIn, 256) {
		return utils.ErrInvalidUint("amountIn", i.AmountIn, 256)
	}
	if !utils.IsUint(i.Slippage, 256) {
		return utils.ErrInvalidUint("slippage", i.Slippage, 256)
	}
	if !utils.IsAddress(i.PrimaryAddress) {
		return utils.ErrInvalidAddress("primaryAddress", i.PrimaryAddress)
	}

	return nil
}

func (i SwapInputs) Build(from string) (*utils.Transaction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}

	return nil, utils.ErrNotImplemented("SwapInputs.Build")
}
