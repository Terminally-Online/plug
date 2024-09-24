package intent

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"solver/bindings/erc20"
	"solver/utils"
)

/*
TransferFromInputs represents the inputs for the TransferFrom action.

A TransferFrom action transfers tokens from the sender to a target recipient.

This is of use when an ethereum account is transferring tokens for themselves
in an external context or someone else is transferring tokens on the behalf
of the sender. In this case, approval of the tokens is requried beforehand.
*/
type TransferFromInputs struct {
	Token     string `json:"token"`     // Address of the token to transfer.
	Sender    string `json:"sender"`    // Address of the sender.
	Recipient string `json:"recipient"` // Address of the recipient.
	Amount    string `json:"amount"`    // Raw amount of tokens to transfer.
}

func (i TransferFromInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Sender) {
		return utils.ErrInvalidAddress("sender", i.Sender)
	}
	if !utils.IsAddress(i.Recipient) {
		return utils.ErrInvalidAddress("recipient", i.Recipient)
	}
	if !utils.IsUint(i.Amount, 256) {
		return utils.ErrInvalidUint("amount", i.Amount, 256)
	}

	return nil
}

func (i TransferFromInputs) Build(chainId int, from string) (*utils.Transaction, error) {
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

	transferFrom, err := contract.TransferFrom(
		utils.DummyTransactOpts(),
		common.HexToAddress(i.Sender),
		common.HexToAddress(i.Recipient),
		amount,
	)
	if err != nil {
		return nil, err
	}

	return &utils.Transaction{
		Transaction: "0x" + hex.EncodeToString(transferFrom.Data()),
		From:        from,
		To:          transferFrom.To().Hex(),
		Value:       big.NewInt(0),
	}, nil
}
