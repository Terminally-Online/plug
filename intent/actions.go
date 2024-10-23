package intent

import (
	"encoding/json"
	"fmt"
	"solver/types"
)

func ParseAction(action types.Action) (types.ActionInputs, error) {
	inputs, err := ParseType(action)
	if err != nil { 
		return nil, err
	}

	if err := json.Unmarshal(action.Inputs, inputs); err != nil {
		return nil, fmt.Errorf("failed to parse inputs for action %s: %v", action.Type, err)
	}

	if err := inputs.Validate(); err != nil {
		return nil, fmt.Errorf("invalid inputs for action %s: %v", action.Type, err)
	}

	return inputs, nil
}
