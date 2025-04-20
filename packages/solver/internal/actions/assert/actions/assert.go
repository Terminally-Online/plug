package assert

import (
	"fmt"
	"solver/bindings/plug_assert"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type AssertRequest struct {
	Condition coil.CoilInput[bool, bool] `json:"condition"`
	Assertion string                     `json:"assertion"`
}

var AssertFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_assert.PlugAssertMetaData,
	FunctionName: "assertTrue",
}

func HandleAssert(lookup *actions.SchemaLookup[AssertRequest]) ([]signature.Plug, error) {
	condition, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Condition,
		lookup.Inputs.Condition.GetValueWithError,
		&AssertFunc,
		"y",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get y value: %w", err)
	}

	AssertFunc.FunctionName = lookup.Inputs.Assertion
	calldata, err := AssertFunc.GetCalldata(condition)
	if err != nil {
		return nil, fmt.Errorf("failed to pack assertTrue calldata: %w", err)
	}

	assertContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["assert"])
	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       assertContract,
		Data:     calldata,
		Updates:  updates,
	}}, nil
}
