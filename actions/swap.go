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
	if i.AmountIn.Cmp(big.NewInt(0)) >= 0 && i.AmountIn.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amountIn", i.AmountIn.String())
	}
	if i.Slippage.Cmp(big.NewInt(0)) >= 0 && i.Slippage.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("slippage", i.Slippage.String())
	}

	return nil
}

func (i *SwapInputsImpl) Get(provider *ethclient.Client, chainId int) (*types.ActionSchema, error) {
	switch i.Protocol {
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
}

func (i *SwapInputsImpl) Post(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	switch i.Protocol {
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
}

func (i *SwapInputsImpl) GetProtocol() string   { return i.Protocol }
func (i *SwapInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *SwapInputsImpl) GetTokenOut() string   { return i.TokenOut }
func (i *SwapInputsImpl) GetAmountIn() *big.Int { return new(big.Int).Set(&i.AmountIn) }
