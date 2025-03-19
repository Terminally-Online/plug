package options

import "solver/internal/actions"

func HealthFactorOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	marketOptions, _, err := GetMarketAndVaultOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: marketOptions},
		1: {Simple: actions.BaseThresholdFields},
	}, nil
}
