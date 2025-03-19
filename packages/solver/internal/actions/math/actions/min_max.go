package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_math"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type MinMaxRequest struct {
	A *big.Int `json:"a"`
	B *big.Int `json:"b"`
}

func Min(lookup *actions.SchemaLookup[MinMaxRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.A == nil || lookup.Inputs.B == nil {
		return nil, fmt.Errorf("a and b values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}

	calldata, err := mathAbi.Pack("min", lookup.Inputs.A, lookup.Inputs.B)
	if err != nil {
		return nil, fmt.Errorf("failed to pack min calldata: %w", err)
	}

	return []signature.Plug{{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}}, nil
}

func Max(lookup *actions.SchemaLookup[MinMaxRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.A == nil || lookup.Inputs.B == nil {
		return nil, fmt.Errorf("a and b values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}

	calldata, err := mathAbi.Pack("max", lookup.Inputs.A, lookup.Inputs.B)
	if err != nil {
		return nil, fmt.Errorf("failed to pack max calldata: %w", err)
	}

	return []signature.Plug{{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}}, nil
}

