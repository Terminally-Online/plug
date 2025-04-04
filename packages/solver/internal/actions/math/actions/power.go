package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_math"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type PowerRequest struct {
	Base     coil.CoilInput[*big.Int, *big.Int] `json:"base"`
	Exponent *big.Int                           `json:"exponent"`
}

var PowerFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_math.PlugMathMetaData,
	FunctionName: "power",
}

func Power(lookup *actions.SchemaLookup[PowerRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.Base.GetValue() == nil || lookup.Inputs.Exponent == nil {
		return nil, fmt.Errorf("base and exponent values must be provided")
	}

	var baseUpdates []coil.Update
	base, baseUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Base,
		lookup.Inputs.Base.GetValueWithError,
		&PowerFunc,
		"base",
		baseUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get base value: %w", err)
	}

	if lookup.Inputs.Exponent.Cmp(big.NewInt(1000)) > 0 {
		return nil, fmt.Errorf("exponent too large: maximum allowed is 1000")
	}

	calldata, err := PowerFunc.GetCalldata(
		base,
		lookup.Inputs.Exponent,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get calldata: %w", err)
	}

	return []signature.Plug{{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"]),
		Data:    calldata,
		Value:   nil,
		Updates: baseUpdates,
	}}, nil
}
