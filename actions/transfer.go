package actions

import (
	"encoding/hex"
	"math/big"
	"solver/bindings/erc_20"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
TransferInputsImpl represents the inputs for the Transfer action.

A Transfer action transfers tokens from the local sender (source) of the transaction
to a target recipient. In essence, the sender is not provided because the source is
defined at runtime of the transaction by the caller.

This is of use when a smart account is transferring tokens directly from their own
account to another destination. As the call is being routed through the smart account,
the sender does not need to be provided as well as approvals are not required.
*/
type TransferInputsImpl struct {
	Type      int     `json:"type"`
	Token     string  `json:"token"`     // Address of the token to transfer.
	Recipient string  `json:"recipient"` // Address of the recipient.
	Amount    big.Int `json:"amount"`    // Raw amount of tokens to transfer.
}

func (i *TransferInputsImpl) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Recipient) {
		return utils.ErrInvalidAddress("recipient", i.Recipient)
	}

	if i.Amount.Cmp(big.NewInt(0)) <= 0 {
		return utils.ErrInvalidField("amount", "must be greater than 0 for token transfers")
	}
	if i.Amount.Cmp(big.NewInt(0)) >= 0 && i.Amount.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amount", i.Amount.String())
	}

	return nil
}

func (i *TransferInputsImpl) Get(provider *ethclient.Client, chainId int) (*types.ActionSchema, error) {
	return nil, utils.ErrNotImplemented("TransferInputsImpl.Get")
}

func (i *TransferInputsImpl) Post(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	var transfer *ethtypes.Transaction
	var err error
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

	return []*types.Transaction{{
		Transaction: "0x" + hex.EncodeToString(transfer.Data()),
		To:          transfer.To().Hex(),
		Value:       transfer.Value(),
	}}, nil
}

func (i *TransferInputsImpl) BuildNativeTransfer(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	return ethtypes.NewTransaction(
		0,
		common.HexToAddress(i.Recipient),
		&i.Amount,
		utils.NativeTransferGas,
		big.NewInt(0),
		nil,
	), nil
}

func (i *TransferInputsImpl) BuildERC20Transfer(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	contract, err := erc_20.NewErc20(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.Transfer(
		utils.BuildTransactionOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Recipient),
		&i.Amount,
	)
}

func (i *TransferInputsImpl) GetType() uint64      { return uint64(i.Type) }
func (i *TransferInputsImpl) GetToken() string     { return i.Token }
func (i *TransferInputsImpl) GetRecipient() string { return i.Recipient }
func (i *TransferInputsImpl) GetAmount() *big.Int  { return new(big.Int).Set(&i.Amount) }
