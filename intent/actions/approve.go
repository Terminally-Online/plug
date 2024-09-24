package intent

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"solver/bindings/erc20"
	"solver/utils"
)

type ApproveInputs struct {
	Token   string  `json:"token"`   // Address of the token to approve.
	Spender string  `json:"spender"` // Address of the spender.
	Amount  big.Int `json:"amount"`  // Amount to approve.
}

func (i ApproveInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Spender) {
		return utils.ErrInvalidAddress("spender", i.Spender)
	}
	if !utils.IsUint(i.Amount.String(), 256) {
		return utils.ErrInvalidUint("amount", i.Amount.String(), 256)
	}
	return nil
}

func (i ApproveInputs) Build(chainId int, from string) (*utils.Transaction, error) {
	ethClient, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	contract, err := erc20.NewErc20(common.HexToAddress(i.Token), ethClient)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	approve, err := contract.Approve(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Spender),
		&i.Amount,
	)
	if err != nil {
		return nil, err
	}

	return &utils.Transaction{
		Transaction: "0x" + hex.EncodeToString(approve.Data()),
		From:        from,
		To:          approve.To().Hex(),
		Value:       approve.Value(),
		Gas:         approve.Gas(),
	}, nil
}
