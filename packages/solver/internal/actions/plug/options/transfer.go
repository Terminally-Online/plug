package options

import (
	"slices"
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func TransferOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	holdingsIndex := 1
	heldFungiblesOptions, err := options.GetFungiblesHeldOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	otherFungiblesOptions, err := options.GetFungiblesOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	heldFungiblesMap := make(map[string]actions.Option, len(heldFungiblesOptions))
	for _, option := range heldFungiblesOptions {
		heldFungiblesMap[option.Value] = option
	}

	// TODO MASON: what search index should I be using?
	heldNonFungiblesOptions, err := options.GetCollectiblesHeldOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	holdingsOptions := append(heldFungiblesOptions, heldNonFungiblesOptions...)
	allHoldingsOptions := make([]actions.Option, 0, len(holdingsOptions)+len(otherFungiblesOptions))
	allHoldingsOptions = append(allHoldingsOptions, holdingsOptions...)
	allHoldingsOptions = append(allHoldingsOptions, slices.DeleteFunc(otherFungiblesOptions, func(o actions.Option) bool {
		_, exists := heldFungiblesMap[o.Value]
		return exists
	})...)

	recipientIndex := 2
	recipientOptions, err := options.GetAddressOptions(lookup, recipientIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		holdingsIndex:  {Simple: allHoldingsOptions},
		recipientIndex: {Simple: recipientOptions},
	}, nil
}
