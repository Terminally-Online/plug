package options

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func TransferOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	holdingsIndex := 1
	holdingsOptions, err := options.GetHoldingsOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	recipientIndex := 2
	recipientOptions, err := options.GetAddressOptions(lookup, recipientIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		holdingsIndex:  {Simple: holdingsOptions},
		recipientIndex: {Simple: recipientOptions},
	}, nil
}
