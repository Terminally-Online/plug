package options

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func MintLatestOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	recipientIndex := 1
	recipientOptions, err := options.GetAddressOptions(lookup, recipientIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		recipientIndex: {Simple: recipientOptions},
	}, nil
}
