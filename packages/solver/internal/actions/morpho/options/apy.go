package options

import "solver/internal/actions"

func APYOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	_, marketAndVaultOptions, err := GetMarketAndVaultOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: actions.BaseLendActionTypeFields},
		1: {Simple: marketAndVaultOptions},
		2: {Simple: actions.BaseThresholdFields},
	}, nil
}
