package actions

import (
	"encoding/hex"
	"math/big"
	"solver/protocols/aave_v2"
	"solver/protocols/aave_v3"
	"solver/protocols/yearn_v3"
	"solver/types"
	"solver/utils"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type DepositInputsImpl struct {
	Protocol string  `json:"protocol"` // Slug of the protocol to use.
	TokenIn  string  `json:"tokenIn"`  // Address of the token to send (deposit).
	TokenOut string  `json:"tokenOut"` // Address of the token to receive (withdraw).
	AmountIn big.Int `json:"amountIn"` // Raw amount to send (deposit).
	Target   *string `json:"target"`   // Address of smart contract to interact with.
}

func (i *DepositInputsImpl) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if (i.AmountIn.Cmp(big.NewInt(0)) >= 0 && i.AmountIn.Cmp(utils.Uint256Max) > 0) { 
		return utils.ErrInvalidField("amountIn", i.AmountIn.String())
	}
	return nil
}

func (i *DepositInputsImpl) Build(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	var deposit []*ethtypes.Transaction
	var err error
	switch i.Protocol {
	case aave_v2.Key:
		deposit, err = aave_v2.BuildDeposit(i, provider, chainId, from)
	case aave_v3.Key:
		deposit, err = aave_v3.BuildDeposit(i, provider, chainId, from)
	case yearn_v3.Key:
		deposit, err = yearn_v3.BuildDeposit(i, provider, chainId, from)
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
	if err != nil {
		return nil, err
	}

	var txs []*types.Transaction
	for _, tx := range deposit {
		txs = append(txs, &types.Transaction{
			Transaction: "0x" + hex.EncodeToString(tx.Data()),
			To:          tx.To().Hex(),
			Value:       tx.Value(),
		})
	}
	return txs, nil
}

func (i *DepositInputsImpl) GetProtocol() string   { return i.Protocol }
func (i *DepositInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *DepositInputsImpl) GetTokenOut() string   { return i.TokenOut }
func (i *DepositInputsImpl) GetAmountIn() *big.Int { return &i.AmountIn }
func (i *DepositInputsImpl) GetTarget() *string    { return i.Target }
