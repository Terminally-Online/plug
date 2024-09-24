package intent

import (
	"solver/utils"
)

type RedeemInputs struct {
	TokenIn        string `json:"tokenIn"`        // Address of the token to send (redeem).
	TokenOut       string `json:"tokenOut"`       // Address of the token to receive (redeeming for).
	AmountIn       string `json:"amountIn"`       // Raw amount of tokens to send.
	PrimaryAddress string `json:"primaryAddress"` // Address of the smart contract to interact with.
}

func (i RedeemInputs) Validate() error {
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

func (i RedeemInputs) Build(chainId int, from string) (*utils.Transaction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}

	return nil, utils.ErrNotImplemented("RedeemInputs.Build")
}
