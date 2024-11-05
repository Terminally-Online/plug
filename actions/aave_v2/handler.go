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
			Name:            "Aave V2",
			Icon:            "https://app.aave.com/favicon.ico",
			SupportedChains: []int{1},
		},
	}
	return h.init()
}

func (h *Handler) GetIcon() string {
	return h.Protocol.Icon
}

func (h *Handler) init() *Handler {
	h.schemas[types.ActionDeposit] = types.BaseDepositSchema
	h.schemas[types.ActionBorrow] = types.BaseBorrowSchema
	return h
}

func (h *Handler) SupportedActions() []types.Action {
	return []types.Action{
		types.ActionDeposit,
		types.ActionBorrow,
	}
}

func (h *Handler) SupportedChains() []int {
	return h.Protocol.SupportedChains
}

func (h *Handler) GetSchema(action types.Action) (types.Schema, error) {
	schema, exists := h.schemas[action]
	if !exists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
	return schema, nil
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
			common.HexToAddress(inputs.TokenOut),
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
