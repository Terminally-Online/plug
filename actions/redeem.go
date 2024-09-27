package actions

import (
	"encoding/hex"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"solver/protocols/aave_v2"
	"solver/protocols/aave_v3"
	"solver/types"
	"solver/utils"
)

type RedeemInputsImpl struct {
	Protocol string  `json:"protocol"` // Slug of the protocol to use.
	TokenIn  string  `json:"tokenIn"`  // Address of the token to send (redeem).
	TokenOut string  `json:"tokenOut"` // Address of the token to receive (redeeming for).
	AmountIn big.Int `json:"amountIn"` // Raw amount of tokens to send.
}

func (i *RedeemInputsImpl) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if !utils.IsUint(i.AmountIn.String(), 256) {
		return utils.ErrInvalidUint("amountIn", i.AmountIn.String(), 256)
	}

	return nil
}

func (i *RedeemInputsImpl) Build(provider *ethclient.Client, chainId int, from string) (*types.Transaction, error) {
	var redeem *ethtypes.Transaction
	var err error
	switch i.Protocol {
	case aave_v2.Key:
		redeem, err = aave_v2.BuildRedeem(i, provider, chainId, from)
	case aave_v3.Key:
		redeem, err = aave_v3.BuildRedeem(i, provider, chainId, from)
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
	if err != nil {
		return nil, err
	}

	return &types.Transaction{
		Transaction: "0x" + hex.EncodeToString(redeem.Data()),
		From:        from,
		To:          redeem.To().Hex(),
		Value:       redeem.Value(),
		Gas:         redeem.Gas(),
	}, nil
}

func (i *RedeemInputsImpl) GetProtocol() string   { return i.Protocol }
func (i *RedeemInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *RedeemInputsImpl) GetTokenOut() string   { return i.TokenOut }
func (i *RedeemInputsImpl) GetAmountIn() *big.Int { return new(big.Int).Set(big.NewInt(0)) }
