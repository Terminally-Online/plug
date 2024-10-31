package aave_v2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"solver/bindings/aave_v2_pool"
	"solver/protocols"
	"solver/types"
	"solver/utils"
)

var (
	address          = utils.Mainnet.References["aave_v2"]["pool"]
	interestRateMode = new(big.Int).SetUint64(2)
)

type Handler struct {
	depositSchema types.ActionSchema
	borrowSchema  types.ActionSchema
	redeemSchema  types.ActionSchema
    repaySchema   types.ActionSchema
	protocols.Protocol
}

func New() protocols.BaseProtocolHandler {
	h := &Handler{}
	return h.init()
}

func (h *Handler) init() *Handler {
	h.Protocol = protocols.Protocol{
		SupportedChains: []int{1},
	}

	h.depositSchema = types.ActionSchema{
		Protocol: types.ProtocolAaveV2,
		Schema: types.Schema{
			Fields: []types.SchemaField{
				{
					Name:        "tokenIn",
					Type:        "address",
					Description: "Token to deposit",
					Options: []types.Option{
						{
							Value: "0x6b175474e89094c44da98b954eedeac495271d0f",
							Label: "DAI",
						},
						{
							Value: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
							Label: "WETH",
						},
					},
				},
				{
					Name:        "amountIn",
					Type:        "uint256",
					Description: "Amount to deposit in wei",
				},
			},
			Required: []string{"tokenIn", "amountIn"},
		},
	}

	h.borrowSchema = types.ActionSchema{
		Protocol: types.ProtocolAaveV2,
		Schema: types.Schema{
			Fields: []types.SchemaField{
				{
					Name:        "tokenToBorrow",
					Type:        "address",
					Description: "Token to borrow",
					Options: []types.Option{
						{
							Value: "0x6b175474e89094c44da98b954eedeac495271d0f",
							Label: "DAI",
						},
					},
				},
			},
			Required: []string{"tokenToBorrow", "amount"},
		},
	}

	h.redeemSchema = types.ActionSchema{
		Protocol: types.ProtocolAaveV2,
		Schema: types.Schema{
			Fields: []types.SchemaField{
				{
					Name:        "tokenToBorrow",
					Type:        "address",
					Description: "Token to borrow",
					Options: []types.Option{
						{
							Value: "0x6b175474e89094c44da98b954eedeac495271d0f",
							Label: "DAI",
						},
					},
				},
			},
			Required: []string{"tokenToBorrow", "amount"},
		},
	}

	h.repaySchema = types.ActionSchema{
		Protocol: types.ProtocolAaveV2,
		Schema: types.Schema{
			Fields: []types.SchemaField{
				{
					Name:        "tokenToBorrow",
					Type:        "address",
					Description: "Token to borrow",
					Options: []types.Option{
						{
							Value: "0x6b175474e89094c44da98b954eedeac495271d0f",
							Label: "DAI",
						},
					},
				},
			},
			Required: []string{"tokenToBorrow", "amount"},
		},
	}

	return h
}

func (h *Handler) SupportedActions() []types.Action {
	return []types.Action{types.ActionDeposit, types.ActionBorrow, types.ActionRedeem, types.ActionRepay}
}

func (h *Handler) SupportedChains() []int {
	return h.Protocol.SupportedChains
}

func (h *Handler) HandleGetDeposit() types.ActionSchema {
	return h.depositSchema
}

func (h *Handler) HandlePostDeposit(inputs *types.DepositInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	poolAbi, err := aave_v2_pool.AaveV2PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}

	calldata, err := poolAbi.Pack("deposit",
		common.HexToAddress(inputs.TokenOut),
		inputs.AmountIn,
		common.HexToAddress(from),
		uint16(0))
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   address,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}

func (h *Handler) HandleGetBorrow() types.ActionSchema {
	return h.borrowSchema
}

func (h *Handler) HandlePostBorrow(inputs *types.BorrowInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	poolAbi, err := aave_v2_pool.AaveV2PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}

	calldata, err := poolAbi.Pack("borrow",
		common.HexToAddress(inputs.TokenOut),
		inputs.AmountOut,
		interestRateMode,
		uint16(0),
		from)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   address,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, err
}

func (h *Handler) HandleGetRedeem() types.ActionSchema {
	return h.redeemSchema
}

func (h *Handler) HandlePostRedeem(inputs *types.RedeemInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	poolAbi, err := aave_v2_pool.AaveV2PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}

	calldata, err := poolAbi.Pack("withdraw",
		common.HexToAddress(inputs.TokenIn),
		inputs.AmountIn,
		common.HexToAddress(from),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   address,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}

func (h *Handler) HandleGetRepay() types.ActionSchema {
	return h.repaySchema
}

func (h *Handler) HandlePostRepay(inputs *types.RepayInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	poolAbi, err := aave_v2_pool.AaveV2PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}

	calldata, err := poolAbi.Pack("repay",
		common.HexToAddress(inputs.TokenIn),
		inputs.AmountIn,
		interestRateMode,
		common.HexToAddress(from),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   address,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}
