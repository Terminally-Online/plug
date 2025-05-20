package options

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func BalanceOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	holdingsIndex := 0
	holdingsOptions, err := options.GetHoldingsOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	addressIndex := 1
	addressOptions, err := options.GetAddressOptions(lookup, addressIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		holdingsIndex: {Simple: holdingsOptions},
		addressIndex:  {Simple: addressOptions},
	}, nil
}
