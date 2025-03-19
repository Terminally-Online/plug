package options

import (
	"solver/internal/actions"
	"solver/internal/actions/nouns/reads"
)

func GetTraitOptions() ([]actions.Option, map[string][]actions.Option, error) {
	traitTypes, err := GetTraitTypeOptions()
	if err != nil {
		return nil, nil, err
	}

	traits, err := reads.GetTraits()
	if err != nil {
		return nil, nil, err
	}

	traitOptions := make(map[string][]actions.Option)
	for _, traitType := range traitTypes {
		traitOptions[traitType.Value] = make([]actions.Option, 0)
	}

	for _, traitType := range traitTypes {
		for _, trait := range traits[traitType.Value] {
			traitOptions[traitType.Value] = append(traitOptions[traitType.Value], actions.Option{
				Name:  trait.Name,
				Label: trait.Label,
				Value: trait.Value,
				Icon:  &actions.OptionIcon{Default: trait.Icon},
			})
		}
	}

	return traitTypes, traitOptions, nil
}

func HasTraitOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	traitTypeOptions, traitOptions, err := GetTraitOptions()
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: traitTypeOptions},
		1: {Complex: traitOptions},
	}, nil
}
