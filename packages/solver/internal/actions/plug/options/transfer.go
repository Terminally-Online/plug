package options

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func TransferOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	fungiblesIndex := 1
	fungiblesOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesIndex)
	if err != nil {
		return nil, err
	}

	recipientIndex := 2
	recipientOptions, err := options.GetAddressOptions(lookup, recipientIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		fungiblesIndex: {Simple: fungiblesOptions},
		recipientIndex: {Simple: recipientOptions},
	}, nil
}
