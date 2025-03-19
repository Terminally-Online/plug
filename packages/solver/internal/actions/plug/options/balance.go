package options

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func BalanceOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	fungiblesIndex := 0
	fungiblesOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesIndex)
	if err != nil {
		return nil, err
	}

	addressIndex := 1
	addressOptions, err := options.GetAddressOptions(lookup, addressIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		fungiblesIndex: {Simple: fungiblesOptions},
		addressIndex:   {Simple: addressOptions},
		2:              {Simple: actions.BaseThresholdFields},
	}, nil
}
