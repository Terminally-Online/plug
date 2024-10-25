package actions

import (
	"encoding/hex"
	"math/big"
	"solver/bindings/erc_1155"
	"solver/bindings/erc_20"
	"solver/bindings/erc_721"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
TransferFromInputsImpl represents the inputs for the TransferFrom action.

A TransferFrom action transfers tokens from the sender to a target recipient.

This is of use when an ethereum account is transferring tokens for themselves
in an external context or someone else is transferring tokens on the behalf
of the sender. In this case, approval of the tokens is requried beforehand.

This function does not support the transfering of native tokens (ETH). To
transfer native tokens, use the Transfer action instead.
*/
type TransferFromInputsImpl struct {
	Type      int      `json:"type"`      // Type (numerical standard) of the token to transfer.
	Token     string   `json:"token"`     // Address of the token to transfer.
	Sender    string   `json:"sender"`    // Address of the sender.
	Recipient string   `json:"recipient"` // Address of the recipient.
	TokenId   *big.Int `json:"tokenId"`   // Token ID of the token to transfer.
	Amount    *big.Int `json:"amount"`    // Raw amount of tokens to transfer.
}

func (i *TransferFromInputsImpl) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Sender) {
		return utils.ErrInvalidAddress("sender", i.Sender)
	}
	if !utils.IsAddress(i.Recipient) {
		return utils.ErrInvalidAddress("recipient", i.Recipient)
	}

	switch i.Type {
	case 20:
		if i.TokenId != nil {
			return utils.ErrInvalidField("tokenId", "must not be provided for ERC20 tokens")
		}
		if i.Amount == nil {
			return utils.ErrInvalidField("amount", "must be provided for ERC20 tokens")
		}
		if i.Amount.Cmp(big.NewInt(0)) <= 0 {
			return utils.ErrInvalidField("amount", "must be greater than 0 for ERC20 tokens")
		}

	case 721:
		if i.TokenId == nil {
			return utils.ErrInvalidField("tokenId", "must be provided for ERC721 tokens")
		}
		if i.TokenId.Cmp(big.NewInt(0)) <= 0 {
			return utils.ErrInvalidField("tokenId", "must be greater than 0 for ERC721 tokens")
		}
		if i.Amount != nil {
			return utils.ErrInvalidField("amount", "must not be provided for ERC721 tokens")
		}

	case 1155:
		if i.TokenId == nil {
			return utils.ErrInvalidField("tokenId", "must be provided for ERC1155 tokens")
		}
		if i.TokenId.Cmp(big.NewInt(0)) < 0 {
			return utils.ErrInvalidField("tokenId", "must be 0 or greater for ERC1155 tokens")
		}
		if i.Amount == nil {
			return utils.ErrInvalidField("amount", "must be provided for ERC1155 tokens")
		}
		if i.Amount.Cmp(big.NewInt(0)) <= 0 {
			return utils.ErrInvalidField("amount", "must be greater than 0 for ERC1155 tokens")
		}

	default:
		return utils.ErrInvalidTokenStandard("type", i.Type)
	}

	if i.Amount != nil && i.Amount.Cmp(big.NewInt(0)) >= 0 && i.Amount.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amount", i.Amount.String())
	}
	if i.TokenId != nil && i.TokenId.Cmp(big.NewInt(0)) >= 0 && i.TokenId.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("tokenId", i.TokenId.String())
	}

	return nil
}

func (i *TransferFromInputsImpl) Get(provider *ethclient.Client, chainId int) (*types.ActionSchema, error) {
	return nil, utils.ErrNotImplemented("TransferFromInputsImpl.Get")
}

func (i *TransferFromInputsImpl) Post(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	var transferFrom *ethtypes.Transaction
	var err error
	switch i.Type {
	case 20:
		transferFrom, err = i.BuildERC20TransferFrom(provider, from)
	case 721:
		transferFrom, err = i.BuildERC721TransferFrom(provider, from)
	case 1155:
		transferFrom, err = i.BuildERC1155TransferFrom(provider, from)
	default:
		return nil, utils.ErrInvalidTokenStandard("type", i.Type)
	}
	if err != nil {
		return nil, err
	}

	return []*types.Transaction{{
		Transaction: "0x" + hex.EncodeToString(transferFrom.Data()),
		To:          transferFrom.To().Hex(),
		Value:       transferFrom.Value(),
	}}, nil
}

func (i *TransferFromInputsImpl) BuildERC20TransferFrom(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	contract, err := erc_20.NewErc20(common.HexToAddress(i.Token), nil)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.TransferFrom(
		utils.BuildTransactionOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Sender),
		common.HexToAddress(i.Recipient),
		i.Amount,
	)
}

func (i *TransferFromInputsImpl) BuildERC721TransferFrom(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	contract, err := erc_721.NewErc721(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.TransferFrom(
		utils.BuildTransactionOpts(from, big.NewInt(0)),
		common.HexToAddress(from),
		common.HexToAddress(i.Recipient),
		i.TokenId,
	)
}

func (i *TransferFromInputsImpl) BuildERC1155TransferFrom(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	contract, err := erc_1155.NewErc1155(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.SafeTransferFrom(
		utils.BuildTransactionOpts(from, big.NewInt(0)),
		common.HexToAddress(from),
		common.HexToAddress(i.Recipient),
		i.TokenId,
		i.Amount,
		[]byte{},
	)
}

func (i *TransferFromInputsImpl) GetType() uint64      { return uint64(i.Type) }
func (i *TransferFromInputsImpl) GetToken() string     { return i.Token }
func (i *TransferFromInputsImpl) GetSender() string    { return i.Sender }
func (i *TransferFromInputsImpl) GetRecipient() string { return i.Recipient }
func (i *TransferFromInputsImpl) GetTokenId() *big.Int { return new(big.Int).Set(i.TokenId) }
func (i *TransferFromInputsImpl) GetAmount() *big.Int  { return new(big.Int).Set(i.Amount) }
