package intent

import (
	"encoding/json"
	"fmt"
	"solver/intent/actions"
	"solver/utils"
)

/*
Action represents a generic action with its type and inputs. This is the
standard shape for a new action to be added to the intent that will be parsed,
built and sent back to the caller of the endpoint.
*/
type Action struct {
	Type   string          `json:"type"`
	Inputs json.RawMessage `json:"inputs"`
}

// ActionInputs is an interface that all specific action input structs should implement
type ActionInputs interface {
	Validate() error
	Build(chainId int, from string) (*utils.Transaction, error)
}

// ParseAction parses the Action struct and returns the specific ActionInputs
func ParseAction(action Action) (ActionInputs, error) {
	var inputs ActionInputs

	switch action.Type {
	case "approve":
		inputs = &intent.ApproveInputs{}
	case "borrow":
		inputs = &intent.BorrowInputs{}
	case "deposit":
		inputs = &intent.DepositInputs{}
	case "harvest":
		inputs = &intent.HarvestInputs{}
	case "redeem":
		inputs = &intent.RedeemInputs{}
	case "repay":
		inputs = &intent.RepayInputs{}
	case "swap":
		inputs = &intent.SwapInputs{}
	case "transfer":
		inputs = &intent.TransferInputs{}
	case "transfer_from":
		inputs = &intent.TransferFromInputs{}
	case "route":
		inputs = &intent.RouteInputs{}
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
