package intent

import (
	"solver/utils"
)

type ApproveInputs struct {
	// Address of the token to approve.
	Token   string `json:"token"`
	// Address of the spender.
	Spender string `json:"spender"`
	// Amount to approve.
	Amount  string `json:"amount"`
}

func (i ApproveInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Spender) {
		return utils.ErrInvalidAddress("spender", i.Spender)
	}
	if !utils.IsUint(i.Amount, 256) {
		return utils.ErrInvalidUint("amount", i.Amount, 256)
	}
	return nil
}
