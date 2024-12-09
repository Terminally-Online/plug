package morpho

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
)

var (
	ActionSupplyCollateral      = "supply_collateral"
	ActionWithdrawCollateral    = "withdraw_collateral"
	ActionWithdrawMaxCollateral = "withdraw_max_collateral"
	ActionBorrow                = "borrow"
	ActionRepay                 = "repay"
	ActionRepayMax              = "repay_max"
	ActionClaimRewards          = "claim_rewards"
	ConstraintHealthFactor      = "health_factor"
	ConstraintAPY               = "apy"
)

type Handler struct {
	schemas map[types.Action]types.Schema
	actions.Protocol
}

func New() actions.BaseProtocolHandler {
	h := &Handler{
		schemas: make(map[types.Action]types.Schema),
		Protocol: actions.Protocol{
			Name:   "Morpho",
			Icon:   "https://morpho.org/butterfly-light.png",
			Tags:   []string{"lending", "defi"},
			Chains: []int{1},
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	marketOptions, err := GetMarketOptions()
	if err != nil {
		return nil
	}

	h.schemas[types.Action(ActionSupplyCollateral)] = types.Schema{
		Sentence: "Supply {0<tokenIn:address>} {1<amountIn:uint256>}.",
	}
	h.schemas[types.Action(ActionWithdrawCollateral)] = types.Schema{
		Sentence: "Withdraw {0<amount:uint256>} collateral from {1<vault:address>}.",
	}
	h.schemas[types.Action(ActionWithdrawMaxCollateral)] = types.Schema{
		Sentence: "Withdraw max collateral from {0<vault:address>}.",
	}
	h.schemas[types.Action(ActionBorrow)] = types.Schema{
		Sentence: "Supply {0<tokenIn:address>} {1<amountIn:uint256>}.",
	}
	h.schemas[types.Action(ActionRepay)] = types.Schema{
		Sentence: "Repay {0<amount:uint256>} of {1<vault:address>}.",
	}
	h.schemas[types.Action(ActionRepayMax)] = types.Schema{
		Sentence: "Repay max of {0<vault:address>}.",
	}
	h.schemas[types.Action(ActionClaimRewards)] = types.Schema{
		Sentence: "Claim rewards from {0<vault:address>}.",
	}
	h.schemas[types.Action(ConstraintHealthFactor)] = types.Schema{
		Sentence: "Check health factor of {0<vault:address>}.",
	}
	h.schemas[types.Action(ConstraintAPY)] = types.Schema{
		Sentence: "{0<action:int8>} APY in {1<target:string>} is {2<operator:int8>} than {3<threshold:uint256>}%.",
		Options: map[int]types.SchemaOptions{
			0: {
				Simple: []types.Option{
					{Label: "Borrow", Name: "Borrow", Value: "-1"},
					{Label: "Deposit", Name: "Deposit", Value: "1"},
				},
			},
			1: {Simple: marketOptions},
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
	case types.ConstraintAPY:
		return HandleConstraintAPY(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
