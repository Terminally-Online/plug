package options

import "solver/internal/actions"

func GetTraitTypeOptions() ([]actions.Option, error) {
	fields := []string{"background", "body", "accessory", "head", "glasses"}

	options := make([]actions.Option, len(fields))
	for i, field := range fields {
		options[i] = actions.Option{
			Value: field,
			Name:  field,
			Label: field,
		}
	}

	return options, nil
}
