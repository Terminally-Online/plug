package plug

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func TransferOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
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

func SwapOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
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

func PriceOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
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

func BalanceOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
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
