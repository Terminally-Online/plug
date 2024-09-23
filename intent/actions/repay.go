package intent

import (
	"solver/utils"
)

type RepayInputs struct {
	TokenIn        string `json:"tokenIn"`        // Address of the token to repay.
	AmountIn       string `json:"amountIn"`       // Raw amount of tokens to repay.
	PrimaryAddress string `json:"primaryAddress"` // Address of the smart contract to interact with.
}

func (i RepayInputs) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsUint(i.AmountIn, 256) {
		return utils.ErrInvalidUint("amountIn", i.AmountIn, 256)
	}
	if !utils.IsAddress(i.PrimaryAddress) {
		return utils.ErrInvalidAddress("primaryAddress", i.PrimaryAddress)
	}

	return nil
}

func (i RepayInputs) Build(from string) (*utils.Transaction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}

	return nil, utils.ErrNotImplemented("RepayInputs.Build")
}
