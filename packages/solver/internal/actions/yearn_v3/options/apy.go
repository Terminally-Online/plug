package options

import (
	"solver/internal/actions"
)

func APYOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	vaultOptions, err := GetVaultOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: vaultOptions},
		1: {Simple: actions.BaseThresholdFields},
	}, nil
}
