package options

import (
	"slices"
	"solver/internal/actions"
)

/**
 * GetHoldingsOptions returns a list of options to be used that include held fungibles and collectibles, as well as non-held fungibles.
 */
func GetHoldingsOptions[T any](lookup *actions.SchemaLookup[T], index int) ([]actions.Option, error) {
	holdingsIndex := 0
	heldFungiblesOptions, err := GetFungiblesHeldOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	otherFungiblesOptions, err := GetFungiblesOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	heldFungiblesMap := make(map[string]actions.Option, len(heldFungiblesOptions))
	for _, option := range heldFungiblesOptions {
		heldFungiblesMap[option.Value] = option
	}

	heldCollectiblesOptions, err := GetCollectiblesHeldOptions(lookup, holdingsIndex)
	if err != nil {
		return nil, err
	}

	// We append held fungibles and collectibles options before including non held fungibles for better display order.
	holdingsOptions := append(heldFungiblesOptions, heldCollectiblesOptions...)

	allHoldingsOptions := make([]actions.Option, 0, len(holdingsOptions)+len(otherFungiblesOptions))
	allHoldingsOptions = append(allHoldingsOptions, holdingsOptions...)

	return append(allHoldingsOptions, slices.DeleteFunc(otherFungiblesOptions, func(o actions.Option) bool {
		_, exists := heldFungiblesMap[o.Value]
		return exists
	})...), nil
}
