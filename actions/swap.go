package actions

import (
	"math/big"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/ethclient"
)

type SwapInputsImpl struct {
	Protocol string  `json:"protocol"` // Slug of the protocol to use.
	TokenIn  string  `json:"tokenIn"`  // Address of the token to swap (sell).
	TokenOut string  `json:"tokenOut"` // Address of the token to swap (buy).
	AmountIn big.Int `json:"amountIn"` // Raw amount of tokens to swap (sell).
	Slippage big.Int `json:"slippage"` // Slippage tolerance when executing the swap.
}

func (i *SwapInputsImpl) Validate() error {
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

func (i *SwapInputsImpl) Build(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	switch i.Protocol {
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
}

func (i *SwapInputsImpl) GetProtocol() string   { return i.Protocol }
func (i *SwapInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *SwapInputsImpl) GetTokenOut() string   { return i.TokenOut }
func (i *SwapInputsImpl) GetAmountIn() *big.Int { return new(big.Int).Set(&i.AmountIn) }
