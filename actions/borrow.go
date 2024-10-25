package actions

import (
	"encoding/hex"
	"math/big"
	"solver/protocols/aave_v2"
	"solver/types"
	"solver/utils"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BorrowInputsImpl struct {
	Protocol   string  `json:"protocol"`   // Slug of the protocol to use.
	Collateral string  `json:"collateral"` // Address of the collateral token (supplied).
	TokenOut   string  `json:"tokenOut"`   // Address of the token to receive (borrow).
	AmountOut  big.Int `json:"amountOut"`  // Raw amount of tokens to borrow.
}

func (i *BorrowInputsImpl) Validate() error {
	if !utils.IsAddress(i.Collateral) {
		return utils.ErrInvalidAddress("collateral", i.Collateral)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if i.AmountOut.Cmp(big.NewInt(0)) >= 0 && i.AmountOut.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amountOut", i.AmountOut.String())
	}
	return nil
}

func (i *BorrowInputsImpl) Get(provider *ethclient.Client, chainId int) (*types.ActionSchema, error) {
	switch i.Protocol {
	case aave_v2.Key:
		return aave_v2.GetBorrow(i, provider, chainId)
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
}

func (i *BorrowInputsImpl) Post(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	var borrow *ethtypes.Transaction
	var err error
	switch i.Protocol {
	case aave_v2.Key:
		borrow, err = aave_v2.PostBorrow(i, provider, chainId, from)
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
	if err != nil {
		return nil, err
	}

	return []*types.Transaction{{
		Transaction: "0x" + hex.EncodeToString(borrow.Data()),
		To:          borrow.To().Hex(),
		Value:       borrow.Value(),
	}}, nil
}

func (i *BorrowInputsImpl) GetProtocol() string    { return i.Protocol }
func (i *BorrowInputsImpl) GetCollateral() string  { return i.Collateral }
func (i *BorrowInputsImpl) GetTokenOut() string    { return i.TokenOut }
func (i *BorrowInputsImpl) GetAmountOut() *big.Int { return new(big.Int).Set(&i.AmountOut) }
