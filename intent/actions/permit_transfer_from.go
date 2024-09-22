package intent

import (
	"solver/utils"
)

type PermitTransferFromInputs struct {
	Token     string `json:"token"`     // Address of the token to approve.
	Amount    string `json:"amount"`    // Amount of tokens to approve.
	Nonce     string `json:"nonce"`     // Nonce of the permit to prevent signature replays.
	Deadline  string `json:"deadline"`  // Deadline of the permit signature.
	Signature string `json:"signature"` // Signature of the permit.
}

func (i PermitTransferFromInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsUint(i.Amount, 256) {
		return utils.ErrInvalidUint("amount", i.Amount, 256)
	}
	if !utils.IsUint(i.Nonce, 256) {
		return utils.ErrInvalidUint("nonce", i.Nonce, 256)
	}
	if !utils.IsUint(i.Deadline, 256) {
		return utils.ErrInvalidUint("deadline", i.Deadline, 256)
	}
	if !utils.IsHex(i.Signature) {
		return utils.ErrInvalidHex("signature", i.Signature)
	}

	return nil
}
