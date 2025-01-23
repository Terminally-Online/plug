package nouns

import (
	"encoding/json"
	"fmt"
	"os"
	"solver/internal/actions"
)

type NounsOptionsProvider struct{}

func (p *NounsOptionsProvider) GetOptions(chainId int, action string) (map[int]actions.Options, error) {
	traitTypeOptions, traitOptions, err := GetTraitOptions()
	if err != nil {
		return nil, err
	}

	switch action {
	case HasTrait:
		return map[int]actions.Options{
			0: {Simple: traitTypeOptions},
			1: {Complex: traitOptions},
		}, nil
	default:
		return nil, nil // Most actions don't have options
	}
}

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

func GetTraitOptions() ([]actions.Option, map[string][]actions.Option, error) {
	traitTypes, err := GetTraitTypeOptions()
	if err != nil {
		return nil, nil, err
	}

	traitData, err := os.ReadFile("actions/nouns/resources/traits.json")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read traits.json: %v", err)
	}

	var rawTraits map[string][]struct {
		Name  string `json:"name"`
		Label string `json:"label"`
		Value string `json:"value"`
		Icon  string `json:"icon"`
	}
	if err := json.Unmarshal(traitData, &rawTraits); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal traits: %v", err)
	}

	traitOptions := make(map[string][]actions.Option)
	for _, traitType := range traitTypes {
		traitOptions[traitType.Value] = make([]actions.Option, 0)
	}

	for _, traitType := range traitTypes {
		for _, trait := range rawTraits[traitType.Value] {
			traitOptions[traitType.Value] = append(traitOptions[traitType.Value], actions.Option{
				Name:  trait.Name,
				Label: trait.Label,
				Value: trait.Value,
				Icon:  trait.Icon,
			})
		}
	}

	return traitTypes, traitOptions, nil
}