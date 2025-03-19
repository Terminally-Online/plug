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

type ClampRequest struct {
	Value *big.Int `json:"value"`
	Min   *big.Int `json:"min"`
	Max   *big.Int `json:"max"`
}

func HandleClamp(lookup *actions.SchemaLookup[ClampRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.Value == nil || lookup.Inputs.Min == nil || lookup.Inputs.Max == nil {
		return nil, fmt.Errorf("value, min, and max values must be provided")
	}

	if lookup.Inputs.Min.Cmp(lookup.Inputs.Max) > 0 {
		return nil, fmt.Errorf("min value exceeds max value")
	}

	mathContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}

	calldata, err := mathAbi.Pack("clamp", lookup.Inputs.Value, lookup.Inputs.Min, lookup.Inputs.Max)
	if err != nil {
		return nil, fmt.Errorf("failed to pack clamp calldata: %w", err)
	}

	plug := signature.Plug{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}

	return []signature.Plug{plug}, nil
}
