package intent

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
	Type      int     `json:"type"`
	Token     string  `json:"token"`     // Address of the token to transfer.
	Recipient string  `json:"recipient"` // Address of the recipient.
	Amount    big.Int `json:"amount"`    // Raw amount of tokens to transfer.
}

func (i TransferInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Recipient) {
		return utils.ErrInvalidAddress("recipient", i.Recipient)
	}

	if i.Amount.Cmp(big.NewInt(0)) <= 0 {
		return utils.ErrInvalidField("amount", "must be greater than 0 for token transfers")
	}

	if !utils.IsUint(i.Amount.String(), 256) {
		return utils.ErrInvalidUint("amount", i.Amount.String(), 256)
	}

	return nil
}

func (i TransferInputs) Build(chainId int, from string) (*utils.Transaction, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	var transfer *types.Transaction
	switch i.Type {
	case 0:
		transfer, err = i.BuildNativeTransfer(provider, from)
	case 20:
		transfer, err = i.BuildERC20Transfer(provider, from)
	default:
		return nil, utils.ErrInvalidTokenStandard("type", i.Type)
	}
	if err != nil {
		return nil, err
	}

	return &utils.Transaction{
		Transaction: "0x" + hex.EncodeToString(transfer.Data()),
		From:        from,
		To:          transfer.To().Hex(),
		Value:       transfer.Value(),
		Gas:         transfer.Gas(),
	}, nil
}

func (i TransferInputs) BuildNativeTransfer(provider *ethclient.Client, from string) (*types.Transaction, error) {
	return types.NewTransaction(
		0,
		common.HexToAddress(i.Recipient),
		&i.Amount,
		utils.NativeTransferGas,
		big.NewInt(0),
		nil,
	), nil
}

func (i TransferInputs) BuildERC20Transfer(provider *ethclient.Client, from string) (*types.Transaction, error) {
	contract, err := erc20.NewErc20(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.Transfer(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Recipient),
		&i.Amount,
	)
}
