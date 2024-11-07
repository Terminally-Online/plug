package aave_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

var (
	poolAddress          = utils.Mainnet.References["aave_v3"]["pool"]
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
	collateralOptions, borrowOptions, err := GetOptions()
	if err != nil {
		return nil
	}

	h.schemas[types.ActionDeposit] = types.Schema{
		Sentence: "Deposit {0} {1}.",
		Fields: []types.SchemaField{
			{
				Name: "tokenIn",
				Type: "address",
				Options: collateralOptions,
			},
			{
				Name: "amountIn",
				Type: "uint256",
			},
		},
	}

	h.schemas[types.ActionBorrow] = types.Schema{
		Sentence: "Borrow {0} {1}.",
		Fields: []types.SchemaField{
			{
				Name: "tokenOut",
				Type: "address",
				Options: borrowOptions,
			},
			{
				Name: "amountOut",
				Type: "uint256",
			},
		},
	}

	h.schemas[types.ActionRepay] = types.Schema{
		Sentence: "Repay {0} {1}.",
		Fields: []types.SchemaField{
			{
				Name: "tokenIn",
				Type: "address",
				Options: borrowOptions,
			},
			{
				Name: "amountIn",
				Type: "uint256",
			},
		},
	}

	h.schemas[types.ActionWithdraw] = types.Schema{
		Sentence: "Withdraw {0} {1}.",
		Fields: []types.SchemaField{
			{
				Name: "tokenOut",
				Type: "address",
				Options: collateralOptions,
			},
			{
				Name: "amountOut",
				Type: "uint256",
			},
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
	var calldata []byte
	var err error
	switch action {
	case types.ActionDeposit:
		calldata, err = HandleDeposit(rawInputs, params)
	case types.ActionBorrow:
		calldata, err = HandleBorrow(rawInputs, params)
	case types.ActionRepay:
		calldata, err = HandleRepay(rawInputs, params)
	case types.ActionWithdraw:
		calldata, err = HandleWithdraw(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   poolAddress,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}
