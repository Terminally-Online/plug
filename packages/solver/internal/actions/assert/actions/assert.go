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

func HandleAssert(lookup *actions.SchemaLookup[AssertRequest]) ([]signature.Plug, error) {
	assertContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["assert"])
	assertAbi, err := plug_assert.PlugAssertMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugAssert ABI: %w", err)
	}

	functionName := "assertTrue"
	if lookup.Inputs.Assertion == "false" {
		functionName = "assertFalse"
	}

	calldata, err := assertAbi.Pack(functionName, lookup.Inputs.Condition)
	if err != nil {
		return nil, fmt.Errorf("failed to pack assertTrue calldata: %w", err)
	}

	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       assertContract,
		Data:     calldata,
	}}, nil
}
