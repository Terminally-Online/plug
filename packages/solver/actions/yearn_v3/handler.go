package yearn_v3

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

type Handler struct {
	schemas map[types.Action]types.Schema
	actions.Protocol
}

func New() actions.BaseProtocolHandler {
	h := &Handler{
		schemas: make(map[types.Action]types.Schema),
		Protocol: actions.Protocol{
			Name:   "Yearn V3",
			Icon:   "https://onplug.io/protocols/yearn.png",
			Tags:   []string{"yield", "defi"},
			Chains: utils.Mainnet.ChainIds,
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	underlyingAssetOptions, err := GetUnderlyingAssetOptions()
	if err != nil {
		return nil
	}
	underlyingAssetToVaultOptions, err := GetUnderlyingAssetToVaultOptions()
	if err != nil {
		return nil
	}
	availableStakingGuageOptions, err := GetAvailableStakingGuageOptions()
	if err != nil {
		return nil
	}
	vaultOptions, err := GetVaultOptions()
	if err != nil {
		return nil
	}

	h.schemas[types.ActionDeposit] = types.Schema{
		Sentence: "Deposit {0<amount:uint256>} {1<token:address>} into {1=>2<vault:address>}.",
		Options: map[int]types.SchemaOptions{
			1: {
				Simple: underlyingAssetOptions,
			},
			2: {
				Complex: underlyingAssetToVaultOptions,
			},
		},
	}

	h.schemas[types.ActionWithdraw] = types.Schema{
		Sentence: "Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<vault:address>}.",
		Options: map[int]types.SchemaOptions{
			1: {
				Simple: underlyingAssetOptions,
			},
			2: {
				Complex: underlyingAssetToVaultOptions,
			},
		},
	}

	h.schemas[types.ActionWithdrawMax] = types.Schema{
		Sentence: "Withdraw max {0<token:address>} from {0=>1<vault:address>}",
		Options: map[int]types.SchemaOptions{
			0: {
				Simple: underlyingAssetOptions,
			},
			1: {
				Complex: underlyingAssetToVaultOptions,
			},
		},
	}

	h.schemas[types.ActionStake] = types.Schema{
		Sentence: "Stake {0<amount:uint256>} {1<token:address>}",
		Options: map[int]types.SchemaOptions{
			1: {Simple: availableStakingGuageOptions},
		},
	}

	h.schemas[types.ActionStakeMax] = types.Schema{
		Sentence: "Stake max {0<token:address>}",
		Options: map[int]types.SchemaOptions{
			0: {Simple: availableStakingGuageOptions},
		},
	}

	h.schemas[types.ActionRedeem] = types.Schema{
		Sentence: "Redeem {0<amount:uint256>} {1<token:address>}",
		Options: map[int]types.SchemaOptions{
			1: {Simple: availableStakingGuageOptions},
		},
	}

	h.schemas[types.ActionRedeemMax] = types.Schema{
		Sentence: "Redeem max staking rewards for {0<token:address>}",
		Options: map[int]types.SchemaOptions{
			0: {Simple: availableStakingGuageOptions},
		},
	}

	h.schemas[types.ConstraintAPY] = types.Schema{
		Sentence: "APY of {0<vault:address>} is {1<operator:int8>} than {2<threshold:uint256>}%.",
		Options: map[int]types.SchemaOptions{
			0: {Simple: vaultOptions},
			1: {Simple: types.BaseThresholdFields},
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
	case types.ActionWithdraw:
		return HandleActionWithdraw(rawInputs, params)
	case types.ActionStake:
		return HandleActionStake(rawInputs, params)
	case types.ActionStakeMax:
		return HandleActionStakeMax(rawInputs, params)
	case types.ActionRedeem:
		return HandleActionRedeem(rawInputs, params)
	case types.ActionRedeemMax:
		return HandleActionRedeemMax(rawInputs, params)
	case types.ConstraintAPY:
		return HandleConstraintAPY(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
