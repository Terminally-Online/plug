package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type HealthFactorRequest struct{}

// TODO MASON: reimplement this
func HealthFactor(lookup *actions.SchemaLookup[HealthFactorRequest]) ([]signature.Plug, error) {
	// NOTE: Aave v3 uses 18 decimals for their health factor.
	//      https://github.com/aave/aave-v3-core/blob/782f51917056a53a2c228701058a6c3fb233684a/test-suites/emode.spec.ts#L555
	// healthFactor, err := reads.GetHealthFactor(lookup.ChainId, lookup.From)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get health factor: %w", err)
	// }

	return nil, nil
}
