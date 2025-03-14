package assert

import (
	"encoding/json"
	"fmt"
	"solver/bindings/plug_assert"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type AssertInput struct {
	Condition bool `json:"condition"`
	Assertion bool `json:"assertion"`
}

func HandleAssert(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs AssertInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal assert inputs: %w", err)
	}

	assertContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["assert"])
	assertAbi, err := plug_assert.PlugAssertMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugAssert ABI: %w", err)
	}

	calldata, err := assertAbi.Pack("assertTrue", inputs.Condition)
	if err != nil {
		return nil, fmt.Errorf("failed to pack assertTrue calldata: %w", err)
	}

	return []signature.Plug{{
		To:    assertContract,
		Data:  calldata,
		Value: nil,
	}}, nil
}
