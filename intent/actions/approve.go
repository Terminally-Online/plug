package intent

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"solver/bindings/erc20"
	"solver/utils"
)

type ApproveInputs struct {
	Token   string `json:"token"`   // Address of the token to approve.
	Spender string `json:"spender"` // Address of the spender.
	Amount  string `json:"amount"`  // Amount to approve.
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

func (i ApproveInputs) Build() (*string, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}

	contract, err := erc20.NewErc20(common.HexToAddress(i.Token), nil)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	amount, ok := new(big.Int).SetString(i.Amount, 10)
	if !ok {
		return nil, utils.ErrInvalidUint("amount", i.Amount, 256)
	}

	tx, err := contract.Approve(
		utils.DummyTransactOpts(),
		common.HexToAddress(i.Spender),
		amount,
	)
	if err != nil {
		return nil, err
	}

	data := "0x" + hex.EncodeToString(tx.Data())
	return &data, nil
}
