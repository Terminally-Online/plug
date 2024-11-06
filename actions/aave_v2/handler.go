package aave_v2

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/bindings/aave_v2_pool"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

var (
	address          = utils.Mainnet.References["aave_v2"]["pool"]
	interestRateMode = new(big.Int).SetUint64(2)
)

type Handler struct {
	schemas map[types.Action]types.Schema
	actions.Protocol
}

func New() actions.BaseProtocolHandler {
	h := &Handler{
		schemas: make(map[types.Action]types.Schema),
		Protocol: actions.Protocol{
			Name:   "Aave V2",
			Icon:   "https://app.aave.com/favicon.ico",
			Chains: []int{1},
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	h.schemas[types.ActionDeposit] = types.Schema{
		Sentence: "Deposit {0} {1}.",
		Fields: []types.SchemaField{
			{
				Name: "tokenIn",
				Type: "address",
				Options: []types.Option{
					{
						Value: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
						Label: "WETH",
						Icon:  "https://tokens.1inch.io/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2.png",
					},
					{
						Value: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
						Label: "DAI",
						Icon:  "https://tokens.1inch.io/0x6b175474e89094c44da98b954eedeac495271d0f.png",
					},
					{
						Value: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
						Label: "USDC",
						Icon:  "https://tokens.1inch.io/0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48.png",
					},
				},
			},
			{
				Name: "amountIn",
				Type: "uint256",
			},
		},
	}

	h.schemas[types.ActionBorrow] = types.BaseBorrowSchema
	return h
}

func (h *Handler) GetSchemas() map[types.Action]types.Schema {
	return h.schemas
}

func (h *Handler) GetSchema(action types.Action) (*types.Schema, error) {
	schema, exists := h.schemas[action]
	if !exists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
	return &schema, nil
}

func (h *Handler) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	poolAbi, err := aave_v2_pool.AaveV2PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}

	var calldata []byte

	switch action {
	case types.ActionDeposit:
		var inputs types.DepositInputs
		if err := json.Unmarshal(rawInputs, &inputs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
		}
		if err := inputs.Validate(); err != nil {
			return nil, err
		}

		calldata, err = poolAbi.Pack("deposit",
			common.HexToAddress(inputs.TokenIn),
			inputs.AmountIn,
			common.HexToAddress(params.From),
			uint16(0))

	case types.ActionBorrow:
		var inputs types.BorrowInputs
		if err := json.Unmarshal(rawInputs, &inputs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal borrow inputs: %v", err)
		}
		if err := inputs.Validate(); err != nil {
			return nil, err
		}

		calldata, err = poolAbi.Pack("borrow",
			common.HexToAddress(inputs.TokenOut),
			inputs.AmountOut,
			interestRateMode,
			uint16(0),
			common.HexToAddress(params.From))

	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}

	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   address,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}
