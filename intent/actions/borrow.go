package intent

import (
	"solver/utils"
)

type BorrowInputs struct {
	Collateral     string `json:"collateral"`     // Address of the collateral token (supplied).
	TokenOut       string `json:"tokenOut"`       // Address of the token to receive (borrow).
	AmountOut      string `json:"amountOut"`      // Raw amount of tokens to borrow.
	PrimaryAddress string `json:"primaryAddress"` // Address of the smart contract to interact with.
}

func (i BorrowInputs) Validate() error {
	if !utils.IsAddress(i.Collateral) {
		return utils.ErrInvalidAddress("collateral", i.Collateral)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if !utils.IsUint(i.AmountOut, 256) {
		return utils.ErrInvalidUint("amountOut", i.AmountOut, 256)
	}
	if !utils.IsAddress(i.PrimaryAddress) {
		return utils.ErrInvalidAddress("primaryAddress", i.PrimaryAddress)
	}
	return nil
}
