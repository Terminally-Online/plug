package options

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)


func PriceOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	fungiblesIndex := 0
	fungiblesOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesIndex)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		fungiblesIndex: {Simple: fungiblesOptions},
		1:              {Simple: actions.BaseThresholdFields},
	}, nil
}

