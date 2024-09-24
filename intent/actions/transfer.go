package intent

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"solver/bindings/erc20"
	"solver/utils"
)

/*
TransferInputs represents the inputs for the Transfer action.

A Transfer action transfers tokens from the local sender (source) of the transaction
to a target recipient. In essence, the sender is not provided because the source is
defined at runtime of the transaction by the caller.

This is of use when a smart account is transferring tokens directly from their own
account to another destination. As the call is being routed through the smart account,
the sender does not need to be provided as well as approvals are not required.
*/
type TransferInputs struct {
	Token     string `json:"token"`     // Address of the token to transfer.
	Recipient string `json:"recipient"` // Address of the recipient.
	Amount    string `json:"amount"`    // Raw amount of tokens to transfer.
}

func (i TransferInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Recipient) {
		return utils.ErrInvalidAddress("recipient", i.Recipient)
	}
	if !utils.IsUint(i.Amount, 256) {
		return utils.ErrInvalidUint("amount", i.Amount, 256)
	}

	return nil
}

func (i TransferInputs) Build(chainId int, from string) (*utils.Transaction, error) {
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

	transfer, err := contract.Transfer(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Recipient),
		amount,
	)
	if err != nil {
		return nil, err
	}

	return &utils.Transaction{
		Transaction: "0x" + hex.EncodeToString(transfer.Data()),
		From:        from,
		To:          transfer.To().Hex(),
		Value:       big.NewInt(0),
	}, nil
}
