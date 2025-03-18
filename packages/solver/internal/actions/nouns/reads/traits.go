package reads

import (
	"encoding/json"
	"fmt"
	"os"
)

type Trait struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
	Icon  string `json:"icon"`
}
type Traits map[string][]Trait

func GetTraits() (Traits, error) {
	traitData, err := os.ReadFile("actions/nouns/data/traits.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read traits.json: %v", err)
	}

	var traits Traits
	if err := json.Unmarshal(traitData, &traits); err != nil {
		return nil, fmt.Errorf("failed to unmarshal traits: %v", err)
	}

	return traits, nil
}
