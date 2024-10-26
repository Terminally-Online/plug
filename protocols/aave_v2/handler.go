package aave_v2

import (
	"math/big"
	"solver/bindings/aave_v2_pool"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	address          = utils.Mainnet.References["aave_v2"]["pool"]
	hexAddress       = common.HexToAddress(address)
	interestRateMode = new(big.Int).SetUint64(2)
)

type Handler struct {
	depositSchema types.ActionSchema
	borrowSchema  types.ActionSchema
}

func New() *Handler {
	h := &Handler{}
	return h.init()
}

func (h *Handler) init() *Handler {
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

	return h
}

func (h *Handler) SupportedActions() []types.Action {
	return []types.Action{types.ActionDeposit, types.ActionBorrow}
}

func (h *Handler) SupportedChains() []int {
	return []int{1}
}

func (h *Handler) HandleGetDeposit() types.ActionSchema {
	return h.depositSchema
}

func (h *Handler) HandlePostDeposit(inputs *types.DepositInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	poolAbi, err := aave_v2_pool.AaveV2PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}

	calldata, err := poolAbi.Pack("deposit", common.HexToAddress(inputs.TokenOut), inputs.AmountIn, common.HexToAddress(from), uint16(0))
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

	calldata, err := poolAbi.Pack("borrow", common.HexToAddress(inputs.TokenOut), inputs.AmountOut, interestRateMode, uint16(0), from)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   address,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, err
}
