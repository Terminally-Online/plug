package actions

import (
	"math/big"
	"solver/protocols/aave_v3"
	"solver/types"
	"solver/utils"
)

type RepayInputsImpl struct {
	Protocol string  `json:"protocol"` // Slug of the protocol to use.
	TokenIn  string  `json:"tokenIn"`  // Address of the token to repay.
	AmountIn big.Int `json:"amountIn"` // Raw amount of tokens to repay.
}

func (i *RepayInputsImpl) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsUint(i.AmountIn.String(), 256) {
		return utils.ErrInvalidUint("amountIn", i.AmountIn.String(), 256)
	}

	return nil
}

func (i *RepayInputsImpl) Build(chainId int, from string) (*types.Transaction, error) {
	switch i.Protocol {
	case aave_v3.Key:
		return aave_v3.BuildRepay(i, chainId, from)
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
}

func (i *RepayInputsImpl) GetProtocol() string   { return i.Protocol }
func (i *RepayInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *RepayInputsImpl) GetAmountIn() *big.Int { return new(big.Int).Set(&i.AmountIn) }
