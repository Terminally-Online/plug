package actions

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3/reads"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type HealthFactorRequest struct{}

var HealthFactorFunc = actions.ActionOnchainFunctionResponse{
	Arguments: &abi.Arguments{
		{Name: "health_factor", Type: abi.Type{T: abi.UintTy, Size: 256}},
	},
}

func HealthFactor(lookup *actions.SchemaLookup[HealthFactorRequest]) ([]signature.Plug, error) {
	// NOTE: Aave v3 uses 18 decimals for their health factor.
	//       https://github.com/aave/aave-v3-core/blob/782f51917056a53a2c228701058a6c3fb233684a/test-suites/emode.spec.ts#L555
	healthFactor, err := reads.GetHealthFactor(lookup.ChainId, lookup.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get health factor: %w", err)
	}

	healthFactorCalldata, err := HealthFactorFunc.Arguments.Pack(healthFactor)
	if err != nil {
		return nil, fmt.Errorf("failed to pack price data: %w", err)
	}

	return []signature.Plug{{
		Selector: signature.ForwardedCall,
		Data:     healthFactorCalldata,
	}}, nil
}
