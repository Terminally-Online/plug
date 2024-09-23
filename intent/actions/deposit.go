package intent

import (
	"solver/utils"
)

type DepositInputs struct {
	TokenIn        string `json:"tokenIn"`        // Address of the token to send (deposit).
	TokenOut       string `json:"tokenOut"`       // Address of the token to receive (withdraw).
	AmountIn       string `json:"amountIn"`       // Raw amount to send (deposit).
	PrimaryAddress string `json:"primaryAddress"` // Address of the smart contract to interact with.
}

func (i DepositInputs) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if !utils.IsUint(i.AmountIn, 256) {
		return utils.ErrInvalidUint("amountIn", i.AmountIn, 256)
	}
	if !utils.IsAddress(i.PrimaryAddress) {
		return utils.ErrInvalidAddress("primaryAddress", i.PrimaryAddress)
	}
	return nil
}

func (i DepositInputs) Build() (*utils.Transaction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}

	return nil, utils.ErrNotImplemented("DepositInputs.Build")
}
