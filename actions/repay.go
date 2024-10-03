package actions

import (
	"encoding/hex"
	"math/big"
	"solver/protocols/aave_v2"
	"solver/protocols/aave_v3"
	"solver/types"
	"solver/utils"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
	if (i.AmountIn.Cmp(big.NewInt(0)) >= 0 && i.AmountIn.Cmp(utils.Uint256Max) > 0) { 
		return utils.ErrInvalidField("amountIn", i.AmountIn.String())
	}

	return nil
}

func (i *RepayInputsImpl) Build(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	var repay *ethtypes.Transaction
	var err error
	switch i.Protocol {
	case aave_v2.Key:
		repay, err = aave_v2.BuildRepay(i, provider, chainId, from)
	case aave_v3.Key:
		repay, err = aave_v3.BuildRepay(i, provider, chainId, from)
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
	if err != nil {
		return nil, err
	}

	return []*types.Transaction{{
		Transaction: "0x" + hex.EncodeToString(repay.Data()),
		To:          repay.To().Hex(),
		Value:       repay.Value(),
	}}, nil
}

func (i *RepayInputsImpl) GetProtocol() string   { return i.Protocol }
func (i *RepayInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *RepayInputsImpl) GetAmountIn() *big.Int { return new(big.Int).Set(&i.AmountIn) }
