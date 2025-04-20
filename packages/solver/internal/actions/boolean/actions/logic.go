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
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}

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

	var calldata []byte
	if lookup.Inputs.Operation == "isNot" {
		calldata, err = booleanAbi.Pack(lookup.Inputs.Operation, a)
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

		calldata, err = booleanAbi.Pack(lookup.Inputs.Operation, a, b)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", lookup.Inputs.Operation, err)
	}

	booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       booleanContract,
		Data:     calldata,
		Updates:  updates,
	}}, nil
}
