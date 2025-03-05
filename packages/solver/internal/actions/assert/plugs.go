package assert

import (
	"encoding/json"
	"fmt"
	"solver/internal/actions"
	"solver/internal/solver/signature"
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

// HandleAssertTrue implements the assertTrue function from PlugAssert.sol
func HandleAssertTrue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs AssertTrueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal assert true inputs: %w", err)
	}

	// Local implementation (this will be replaced with contract call after ABI generation)
	if !inputs.Condition {
		errorMsg := "condition-false"
		if inputs.Message != "" {
			errorMsg = inputs.Message
		}
		return nil, fmt.Errorf("PlugAssert:%s", errorMsg)
	}

	// When the assertion passes, we return an empty plug
	return []signature.Plug{{}}, nil
}

// HandleAssertFalse implements the assertFalse function from PlugAssert.sol
func HandleAssertFalse(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs AssertFalseInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal assert false inputs: %w", err)
	}

	// Local implementation (this will be replaced with contract call after ABI generation)
	if inputs.Condition {
		errorMsg := "condition-true"
		if inputs.Message != "" {
			errorMsg = inputs.Message
		}
		return nil, fmt.Errorf("PlugAssert:%s", errorMsg)
	}

	// When the assertion passes, we return an empty plug
	return []signature.Plug{{}}, nil
}

// HandleFail implements the fail function from PlugAssert.sol
func HandleFail(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs FailInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal fail inputs: %w", err)
	}

	// Always fail with the provided message or a default
	errorMsg := "explicit-fail"
	if inputs.Message != "" {
		errorMsg = inputs.Message
	}
	return nil, fmt.Errorf("PlugAssert:%s", errorMsg)
}