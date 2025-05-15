package actions

import (
	"fmt"
	"solver/bindings/plug_boolean"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type LogicOperationRequest struct {
	A         coil.CoilInput[bool, bool] `json:"a"`
	Operation string                     `json:"operation"`
	B         coil.CoilInput[bool, bool] `json:"b"`
}

var NumberLogicFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_boolean.PlugBooleanMetaData,
	FunctionName: "isAnd",
}

func NumberLogic(lookup *actions.SchemaLookup[LogicOperationRequest]) ([]signature.Plug, error) {
	a, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.A,
		lookup.Inputs.A.GetValueWithError,
		&NumberLogicFunc,
		"a",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	NumberLogicFunc.FunctionName = lookup.Inputs.Operation

	var calldata []byte
	if lookup.Inputs.Operation == "isNot" {
		calldata, err = NumberLogicFunc.GetCalldata(a)
	} else {
		var b bool
		b, updates, err = actions.GetAndUpdate(
			&lookup.Inputs.B,
			lookup.Inputs.B.GetValueWithError,
			&NumberLogicFunc,
			"b",
			updates,
			lookup.PreviousActionDefinition,
		)
		if err != nil {
			return nil, err
		}

		calldata, err = NumberLogicFunc.GetCalldata(a, b)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", lookup.Inputs.Operation, err)
	}

	booleanContract := common.HexToAddress(references.Plug["boolean"])
	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       booleanContract,
		Data:     calldata,
		Updates:  updates,
	}}, nil
}
