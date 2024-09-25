package actions

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"solver/bindings/erc1155"
	"solver/bindings/erc20"
	"solver/bindings/erc721"
	"solver/types"
	"solver/utils"
)

/*
ApproveInputs represents the inputs for the Approve action which supports ERC20, ERC721 and ERC1155.

An Approve action approves the spender to transfer tokens on behalf of the sender.

Based on the optional fields being provided or not, the proper function call is determined and built
without explicit definition of the function to call. All approvals are handled by the contract, so
the transaction will be a simple call to the contract.
*/
type ApproveInputsImpl struct {
	Type     int      `json:"type"`     // Type (numerical standard) of the token to approve.
	Token    string   `json:"token"`    // Address of the token to approve.
	Spender  string   `json:"spender"`  // Address of the spender.
	TokenId  *big.Int `json:"tokenId"`  // Token ID of the token to transfer.
	Amount   *big.Int `json:"amount"`   // Raw amount of tokens to transfer.
	Approved *bool    `json:"approved"` // Whether to approve or revoke approval.
}

func (i *ApproveInputsImpl) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.Spender) {
		return utils.ErrInvalidAddress("spender", i.Spender)
	}

	switch i.Type {
	case 20:
		if i.TokenId != nil {
			return utils.ErrInvalidField("tokenId", "must not be provided for ERC20 tokens")
		}
		if i.Amount == nil {
			return utils.ErrInvalidField("amount", "must be provided for ERC20 tokens")
		}
		if i.Amount.Cmp(big.NewInt(0)) < 0 {
			return utils.ErrInvalidField("amount", "must be non-negative for ERC20 tokens")
		}
		if i.Approved != nil {
			return utils.ErrInvalidField("approved", "must not be provided for ERC20 tokens")
		}

	case 721:
		if i.Amount != nil {
			return utils.ErrInvalidField("amount", "must not be provided for ERC721 tokens")
		}
		if i.TokenId == nil && i.Approved == nil {
			return utils.ErrInvalidField("tokenId or approved", "either tokenId or approved must be provided for ERC721 tokens")
		}
		if i.TokenId != nil && i.Approved != nil {
			return utils.ErrInvalidField("tokenId and approved", "only one of tokenId or approved should be provided for ERC721 tokens")
		}
		if i.TokenId != nil && i.TokenId.Cmp(big.NewInt(0)) <= 0 {
			return utils.ErrInvalidField("tokenId", "must be greater than 0 for ERC721 tokens")
		}

	case 1155:
		if i.TokenId != nil {
			return utils.ErrInvalidField("tokenId", "must not be provided for ERC1155 tokens")
		}
		if i.Amount != nil {
			return utils.ErrInvalidField("amount", "must not be provided for ERC1155 tokens")
		}
		if i.Approved == nil {
			return utils.ErrInvalidField("approved", "must be provided for ERC1155 tokens")
		}

	default:
		return utils.ErrInvalidTokenStandard("type", i.Type)
	}

	if i.Amount != nil && !utils.IsUint(i.Amount.String(), 256) {
		return utils.ErrInvalidUint("amount", i.Amount.String(), 256)
	}
	if i.TokenId != nil && !utils.IsUint(i.TokenId.String(), 256) {
		return utils.ErrInvalidUint("tokenId", i.TokenId.String(), 256)
	}

	return nil
}

func (i *ApproveInputsImpl) Build(chainId int, from string) (*types.Transaction, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	var approve *ethtypes.Transaction
	switch i.Type {
	case 20:
		approve, err = i.BuildERC20Approve(provider, from)
	case 721:
		approve, err = i.BuildERC721Approve(provider, from)
	case 1155:
		approve, err = i.BuildERC1155Approve(provider, from)
	default:
		return nil, utils.ErrInvalidTokenStandard("type", i.Type)
	}
	if err != nil {
		return nil, err
	}

	return &types.Transaction{
		Transaction: "0x" + hex.EncodeToString(approve.Data()),
		From:        from,
		To:          approve.To().Hex(),
		Value:       approve.Value(),
		Gas:         approve.Gas(),
	}, nil
}

func (i *ApproveInputsImpl) BuildERC20Approve(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	contract, err := erc20.NewErc20(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.Approve(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Spender),
		i.Amount,
	)
}

/*
BuildERC721Approve builds an approve transaction for ERC721 tokens based on the inputs provided without
explicit definition of the function to call. As the standard supports both singular token approvals and
batched token approvals, this function will build the appropriate transaction based on the inputs.

When a tokenId is not provided, the transaction will be a SetApprovalForAll transaction -- thus becoming
reliant on `Approved` provided in the inputs. Otherwise, it will be an Approve transaction.
*/
func (i *ApproveInputsImpl) BuildERC721Approve(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	contract, err := erc721.NewErc721(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	if i.TokenId == nil {
		return contract.SetApprovalForAll(
			utils.DummyTransactOpts(from, big.NewInt(0)),
			common.HexToAddress(i.Spender),
			*i.Approved,
		)
	}

	return contract.Approve(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Spender),
		i.TokenId,
	)
}

func (i *ApproveInputsImpl) BuildERC1155Approve(provider *ethclient.Client, from string) (*ethtypes.Transaction, error) {
	contract, err := erc1155.NewErc1155(common.HexToAddress(i.Token), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	return contract.SetApprovalForAll(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.Spender),
		*i.Approved,
	)
}

func (b *ApproveInputsImpl) GetType() int         { return b.Type }
func (b *ApproveInputsImpl) GetToken() string     { return b.Token }
func (b *ApproveInputsImpl) GetSpender() string   { return b.Spender }
func (b *ApproveInputsImpl) GetTokenId() *big.Int { return b.TokenId }
func (b *ApproveInputsImpl) GetAmount() *big.Int  { return b.Amount }
func (b *ApproveInputsImpl) GetApproved() *bool   { return b.Approved }
