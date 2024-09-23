package intent

import (
	"solver/utils"
)

type RouteInputs struct {
	TokenIn  string `json:"tokenIn"`  // Address of the token to send (sell).
	TokenOut string `json:"tokenOut"` // Address of the token to receive (buy).
	AmountIn string `json:"amountIn"` // Raw amount of tokens to send (sell).
	Slippage string `json:"slippage"` // Slippage tolerance when executing the swap.
}

func (i RouteInputs) Validate() error {
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

	return nil
}

func (i RouteInputs) Build(from string) (*utils.Transaction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}

	return nil, utils.ErrNotImplemented("RouteInputs.Build")
}
