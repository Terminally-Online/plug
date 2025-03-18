package options

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func SwapOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	fungiblesOutIndex := 1
	fungiblesOutOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesOutIndex)
	if err != nil {
		return nil, err
	}

	fungiblesInIndex := 2
	fungiblesInOptions, err := options.GetFungiblesAndFungiblesHeldOptions(lookup, fungiblesInIndex)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		fungiblesOutIndex: {Simple: fungiblesOutOptions},
		fungiblesInIndex:  {Simple: fungiblesInOptions},
	}, nil
}
