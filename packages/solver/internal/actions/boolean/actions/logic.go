package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type LogicOperationRequest struct {
	A         bool   `json:"a"`
	Operation string `json:"operation"`
	B         bool   `json:"b"`
}


func LogicOperation(lookup *actions.SchemaLookup[LogicOperationRequest]) ([]signature.Plug, error) {
	// operation := strings.ToLower(lookup.Inputs.Operation)
	// booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
	// booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	// }

	var calldata []byte
	// if operation == "not" {
	// 	calldata, err = booleanAbi.Pack(lookup.Inputs.Operation, lookup.Inputs.A)
	// } else {
	// 	calldata, err = booleanAbi.Pack(functionName, lookup.Inputs.A, lookup.Inputs.B)
	// }

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to pack %s calldata: %w", functionName, err)
	// }

	plug := signature.Plug{
		// To:    booleanContract,
		Data:  calldata,
		Value: nil,
	}

	return []signature.Plug{plug}, nil
}
