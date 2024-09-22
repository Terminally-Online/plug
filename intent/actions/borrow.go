package intent

import (
	"solver/utils"
)

type BorrowInputs struct {
	Collateral     string `json:"collateral"`
	TokenOut       string `json:"tokenOut"`
	AmountOut      string `json:"amountOut"`
	PrimaryAddress string `json:"primaryAddress"`
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
