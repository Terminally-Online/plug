package actions

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"solver/types"
	"solver/utils"
)

type RouteInputsImpl struct {
	TokenIn  string  `json:"tokenIn"`  // Address of the token to send (sell).
	TokenOut string  `json:"tokenOut"` // Address of the token to receive (buy).
	AmountIn big.Int `json:"amountIn"` // Raw amount of tokens to send (sell).
	Slippage big.Int `json:"slippage"` // Slippage tolerance when executing the swap.
}

func (i *RouteInputsImpl) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if !utils.IsUint(i.AmountIn.String(), 256) {
		return utils.ErrInvalidUint("amountIn", i.AmountIn.String(), 256)
	}
	if !utils.IsUint(i.Slippage.String(), 256) {
		return utils.ErrInvalidUint("slippage", i.Slippage.String(), 256)
	}

	return nil
}

func (i *RouteInputsImpl) Build(provider *ethclient.Client, chainId int, from string) (*types.Transaction, error) {
	return nil, utils.ErrNotImplemented("RouteInputsImpl.Build")
}

func (i *RouteInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *RouteInputsImpl) GetTokenOut() string   { return i.TokenOut }
func (i *RouteInputsImpl) GetAmountIn() *big.Int { return &i.AmountIn }
func (i *RouteInputsImpl) GetSlippage() *big.Int { return &i.Slippage }
