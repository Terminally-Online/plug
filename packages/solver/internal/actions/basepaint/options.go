package basepaint

import (
	"solver/internal/actions"
	"solver/internal/actions/options"
)

func MintLatestOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
	return nil, nil
}

func TransferOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
	recipientIndex := 1
	recipientOptions, err := options.GetAddressOptions(lookup, recipientIndex)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		recipientIndex: {Simple: recipientOptions},
	}, nil
}
