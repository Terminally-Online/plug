package actions

import (
	"math/big"
	"solver/protocols/aave_v3"
	"solver/types"
	"solver/utils"
)

type BorrowInputsImpl struct {
	Protocol       string  `json:"protocol"`       // Slug of the protocol to use.
	Collateral     string  `json:"collateral"`     // Address of the collateral token (supplied).
	TokenOut       string  `json:"tokenOut"`       // Address of the token to receive (borrow).
	AmountOut      big.Int `json:"amountOut"`      // Raw amount of tokens to borrow.
}

func (i *BorrowInputsImpl) Validate() error {
	if !utils.IsAddress(i.Collateral) {
		return utils.ErrInvalidAddress("collateral", i.Collateral)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if !utils.IsUint(i.AmountOut.String(), 256) {
		return utils.ErrInvalidUint("amountOut", i.AmountOut.String(), 256)
	}
	return nil
}

func (i *BorrowInputsImpl) Build(chainId int, from string) (*types.Transaction, error) {
	switch i.Protocol {
	case aave_v3.Key:
		return aave_v3.BuildBorrow(i, chainId, from)
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
}

func (i *BorrowInputsImpl) GetProtocol() string       { return i.Protocol }
func (i *BorrowInputsImpl) GetCollateral() string     { return i.Collateral }
func (i *BorrowInputsImpl) GetTokenOut() string       { return i.TokenOut }
func (i *BorrowInputsImpl) GetAmountOut() *big.Int    { return new(big.Int).Set(&i.AmountOut) }
