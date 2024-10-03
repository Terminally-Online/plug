package intent

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
)

// ParseAction parses the Action struct and returns the specific ActionInputs
func ParseAction(action types.Action) (types.ActionInputs, error) {
	var inputs types.ActionInputs

	switch action.Type {
	// Standard based actions that do not require a protocol
	case "approve":
		inputs = &actions.ApproveInputsImpl{}
	case "transfer":
		inputs = &actions.TransferInputsImpl{}
	case "transfer_from":
		inputs = &actions.TransferFromInputsImpl{}
	case "route":
		inputs = &actions.RouteInputsImpl{}
	// Contract based actions that require a protocol
	case "borrow":
		inputs = &actions.BorrowInputsImpl{}
	case "deposit":
		inputs = &actions.DepositInputsImpl{}
	case "harvest":
		inputs = &actions.HarvestInputsImpl{}
	case "redeem":
		inputs = &actions.RedeemInputsImpl{}
	case "repay":
		inputs = &actions.RepayInputsImpl{}
	case "swap":
		inputs = &actions.SwapInputsImpl{}
	default:
		return nil, fmt.Errorf("unknown action type: %s", action.Type)
	}

	if err := json.Unmarshal(action.Inputs, inputs); err != nil {
		return nil, fmt.Errorf("failed to parse inputs for action %s: %v", action.Type, err)
	}

	if err := inputs.Validate(); err != nil {
		return nil, fmt.Errorf("invalid inputs for action %s: %v", action.Type, err)
	}

	return inputs, nil
}
