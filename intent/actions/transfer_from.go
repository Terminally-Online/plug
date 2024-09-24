package intent

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"solver/bindings/erc1155"
	"solver/bindings/erc20"
	"solver/bindings/erc721"
	"solver/utils"
)

/*
TransferFromInputs represents the inputs for the TransferFrom action.

A TransferFrom action transfers tokens from the sender to a target recipient.

This is of use when an ethereum account is transferring tokens for themselves
in an external context or someone else is transferring tokens on the behalf
of the sender. In this case, approval of the tokens is requried beforehand.

This function does not support the transfering of native tokens (ETH). To
transfer native tokens, use the Transfer action instead.
*/
type TransferFromInputs struct {
	Type      int      `json:"type"`      // Type (numerical standard) of the token to transfer.
	Token     string   `json:"token"`     // Address of the token to transfer.
	Sender    string   `json:"sender"`    // Address of the sender.
	Recipient string   `json:"recipient"` // Address of the recipient.
	TokenId   *big.Int `json:"tokenId"`   // Token ID of the token to transfer.
	Amount    *big.Int  `json:"amount"`    // Raw amount of tokens to transfer.
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

	return nil
}
func (i TransferFromInputs) Build(chainId int, from string) (*utils.Transaction, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	var transferFrom *types.Transaction
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

	return &utils.Transaction{
		Transaction: "0x" + hex.EncodeToString(transferFrom.Data()),
		From:        from,
		To:          transferFrom.To().Hex(),
		Value:       transferFrom.Value(),
		Gas:         transferFrom.Gas(),
	}, nil
}

func (i TransferFromInputs) BuildERC20TransferFrom(provider *ethclient.Client, from string) (*types.Transaction, error) {
	contract, err := erc20.NewErc20(common.HexToAddress(i.Token), nil)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.TransferFrom(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Sender),
		common.HexToAddress(i.Recipient),
		i.Amount,
	)
}

func (i TransferFromInputs) BuildERC721TransferFrom(provider *ethclient.Client, from string) (*types.Transaction, error) {
	contract, err := erc721.NewErc721(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.TransferFrom(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(from),
		common.HexToAddress(i.Recipient),
		i.TokenId,
	)
}

func (i TransferFromInputs) BuildERC1155TransferFrom(provider *ethclient.Client, from string) (*types.Transaction, error) {
	contract, err := erc1155.NewErc1155(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.SafeTransferFrom(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(from),
		common.HexToAddress(i.Recipient),
		i.TokenId,
		i.Amount,
		[]byte{},
	)
}
