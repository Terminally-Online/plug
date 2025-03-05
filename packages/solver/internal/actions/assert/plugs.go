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

type AssertTrueInput struct {
	Condition bool   `json:"condition"`
	Message   string `json:"message"`
}

type AssertFalseInput struct {
	Condition bool   `json:"condition"`
	Message   string `json:"message"`
}

type FailInput struct {
	Message string `json:"message"`
}


func HandleAssertTrue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs AssertTrueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal assert true inputs: %w", err)
	}

	if !inputs.Condition {
		errorMsg := "condition-false"
		if inputs.Message != "" {
			errorMsg = inputs.Message
		}
		return nil, fmt.Errorf("PlugAssert:%s", errorMsg)
	}

	assertContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["assert"])
	assertAbi, err := plug_assert.PlugAssertMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugAssert ABI: %w", err)
	}
	
	var calldata []byte
	if inputs.Message != "" {
		calldata, err = assertAbi.Pack("assertTrue", inputs.Condition, inputs.Message)
	} else {
		calldata, err = assertAbi.Pack("assertTrue", inputs.Condition)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to pack assertTrue calldata: %w", err)
	}

	plug := signature.Plug{
		To:    assertContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandleAssertFalse(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs AssertFalseInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal assert false inputs: %w", err)
	}

	if inputs.Condition {
		errorMsg := "condition-true"
		if inputs.Message != "" {
			errorMsg = inputs.Message
		}
		return nil, fmt.Errorf("PlugAssert:%s", errorMsg)
	}

	assertContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["assert"])
	assertAbi, err := plug_assert.PlugAssertMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugAssert ABI: %w", err)
	}
	
	var calldata []byte
	if inputs.Message != "" {
		calldata, err = assertAbi.Pack("assertFalse", inputs.Condition, inputs.Message)
	} else {
		calldata, err = assertAbi.Pack("assertFalse", inputs.Condition)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to pack assertFalse calldata: %w", err)
	}

	plug := signature.Plug{
		To:    assertContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandleFail(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs FailInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal fail inputs: %w", err)
	}
	
	assertContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["assert"])
	assertAbi, err := plug_assert.PlugAssertMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugAssert ABI: %w", err)
	}
	
	calldata, err := assertAbi.Pack("fail", inputs.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to pack fail calldata: %w", err)
	}

	plug := signature.Plug{
		To:    assertContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}
