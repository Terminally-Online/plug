package options

import "solver/internal/actions"


func APYOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	collateralOptions, err := GetCollateralOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	borrowOptions, err := GetBorrowOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	aggregatedOptions := func() []actions.Option {
		seen := make(map[string]bool)
		options := make([]actions.Option, 0)
		for _, opt := range append(collateralOptions, borrowOptions...) {
			if !seen[opt.Value] {
				seen[opt.Value] = true
				opt.Info = nil
				options = append(options, opt)
			}
		}
		return options
	}()

	return map[int]actions.Options{
		0: {Simple: actions.BaseLendActionTypeFields},
		1: {Simple: aggregatedOptions},
		2: {Simple: actions.BaseThresholdFields},
	}, nil
}
