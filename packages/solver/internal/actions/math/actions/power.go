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

type PowerRequest struct {
	Base     *big.Int `json:"base"`
	Exponent *big.Int `json:"exponent"`
}

func Power(lookup *actions.SchemaLookup[PowerRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.Base == nil || lookup.Inputs.Exponent == nil {
		return nil, fmt.Errorf("base and exponent values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}

	if lookup.Inputs.Exponent.Cmp(big.NewInt(1000)) > 0 {
		return nil, fmt.Errorf("exponent too large: maximum allowed is 1000")
	}

	calldata, err := mathAbi.Pack("power", lookup.Inputs.Base, lookup.Inputs.Exponent)
	if err != nil {
		return nil, fmt.Errorf("failed to pack power calldata: %w", err)
	}

	return []signature.Plug{{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}}, nil
}
