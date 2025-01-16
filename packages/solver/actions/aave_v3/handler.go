package aave_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	poolAddress      = utils.Mainnet.References["aave_v3"]["pool"]
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
			Name:   "Aave V3",
			Icon:   "https://app.aave.com/favicon.ico",
			Tags:   []string{"lending", "defi"},
			Chains: utils.Mainnet.ChainIds,
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	collateralOptions, borrowOptions, err := GetOptions()
	if err != nil {
		return nil
	}

	h.schemas[types.ActionDeposit] = types.Schema{
		Sentence: "Deposit {0<tokenIn:address>} {1<amountIn:uint256>}.",
		Options: map[int]types.SchemaOptions{
			0: {Simple: collateralOptions},
		},
	}

	h.schemas[types.ActionBorrow] = types.Schema{
		Sentence: "Borrow {0<tokenOut:address>} {1<amountOut:uint256>}.",
		Options: map[int]types.SchemaOptions{
			0: {Simple: borrowOptions},
		},
	}

	h.schemas[types.ActionRepay] = types.Schema{
		Sentence: "Repay {0<tokenIn:address>} {1<amountIn:uint256>}.",
		Options: map[int]types.SchemaOptions{
			0: {Simple: borrowOptions},
		},
	}

	h.schemas[types.ActionWithdraw] = types.Schema{
		Sentence: "Withdraw {0<tokenOut:address>} {1<amountOut:uint256>}.",
		Options: map[int]types.SchemaOptions{
			0: {Simple: collateralOptions},
		},
	}

	h.schemas[types.ConstraintHealthFactor] = types.Schema{
		Sentence: "Health factor is {0<operator:int8>} than {1<threshold:uint256>}.",
		Options: map[int]types.SchemaOptions{
			0: {Simple: types.BaseThresholdFields},
		},
	}

	aggregatedOptions := func() []types.Option {
		seen := make(map[string]bool)
		options := make([]types.Option, 0)
		for _, opt := range append(collateralOptions, borrowOptions...) {
			if !seen[opt.Value] {
				seen[opt.Value] = true
				opt.Info = ""
				options = append(options, opt)
			}
		}
		return options
	}()
	h.schemas[types.ConstraintAPY] = types.Schema{
		Sentence: "{0<direction:int8>} APY of {1<token:address>} is {2<operator:int8>} than {3<threshold:uint256>}%.",
		Options: map[int]types.SchemaOptions{
			0: {
				Simple: []types.Option{
					{Label: "Borrow", Name: "Borrow", Value: "-1"},
					{Label: "Deposit", Name: "Deposit", Value: "1"},
				},
			},
			1: {Simple: aggregatedOptions},
			2: {Simple: types.BaseThresholdFields},
		},
	}

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
	switch action {
	case types.ActionDeposit:
		return HandleActionDeposit(rawInputs, params)
	case types.ActionBorrow:
		return HandleActionBorrow(rawInputs, params)
	case types.ActionRepay:
		return HandleActionRepay(rawInputs, params)
	case types.ActionWithdraw:
		return HandleActionWithdraw(rawInputs, params)
	case types.ConstraintHealthFactor:
		return HandleConstraintHealthFactor(rawInputs, params)
	case types.ConstraintAPY:
		return HandleConstraintAPY(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
