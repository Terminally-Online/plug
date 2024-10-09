package actions

import (
	"encoding/hex"
	"math/big"
	"solver/bindings/erc_1155"
	"solver/bindings/erc_20"
	"solver/bindings/erc_721"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var tokenTypeToABI = map[int]func() (*abi.ABI, error){
	20:   erc_20.Erc20MetaData.GetAbi,
	721:  erc_721.Erc721MetaData.GetAbi,
	1155: erc_1155.Erc1155MetaData.GetAbi,
}

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
		fallthrough
	case 1155:
		if i.Amount != nil {
			return utils.ErrInvalidField("amount", "must not be provided for ERC721/ERC1155 tokens")
		}
		if i.Type == 721 {
			if i.TokenId == nil && i.Approved == nil {
				return utils.ErrInvalidField("tokenId or approved", "either tokenId or approved must be provided for ERC721 tokens")
			}
			if i.TokenId != nil && i.Approved != nil {
				return utils.ErrInvalidField("tokenId and approved", "only one of tokenId or approved should be provided for ERC721 tokens")
			}
			if i.TokenId != nil && i.TokenId.Cmp(big.NewInt(0)) <= 0 {
				return utils.ErrInvalidField("tokenId", "must be greater than 0 for ERC721 tokens")
			}
		} else {
			if i.TokenId != nil {
				return utils.ErrInvalidField("tokenId", "must not be provided for ERC1155 tokens")
			}
			if i.Approved == nil {
				return utils.ErrInvalidField("approved", "must be provided for ERC1155 tokens")
			}
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

func (i *ApproveInputsImpl) Build(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	tokenAbi, err := tokenTypeToABI[i.Type]()
	if err != nil {
		return nil, utils.ErrContractFailed(i.Token)
	}

	var methodName string
	var args []interface{}

	switch i.Type {
	case 20:
		methodName = "approve"
		args = []interface{}{common.HexToAddress(i.Spender), i.Amount}
	case 721:
		if i.TokenId == nil {
			methodName = "setApprovalForAll"
			args = []interface{}{common.HexToAddress(i.Spender), *i.Approved}
		} else {
			methodName = "approve"
			args = []interface{}{common.HexToAddress(i.Spender), i.TokenId}
		}
	case 1155:
		methodName = "setApprovalForAll"
		args = []interface{}{common.HexToAddress(i.Spender), *i.Approved}
	default:
		return nil, utils.ErrInvalidTokenStandard("type", i.Type)
	}

	data, err := tokenAbi.Pack(methodName, args...)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		Transaction: "0x" + hex.EncodeToString(data),
		To:          i.Token,
		Value:       new(big.Int).SetUint64(0),
	}}, nil
}

func (b *ApproveInputsImpl) GetType() int         { return b.Type }
func (b *ApproveInputsImpl) GetToken() string     { return b.Token }
func (b *ApproveInputsImpl) GetSpender() string   { return b.Spender }
func (b *ApproveInputsImpl) GetTokenId() *big.Int { return b.TokenId }
func (b *ApproveInputsImpl) GetAmount() *big.Int  { return b.Amount }
func (b *ApproveInputsImpl) GetApproved() *bool   { return b.Approved }
