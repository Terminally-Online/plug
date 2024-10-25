package intent

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
)

func GetActionInputs(actionInputsType string) (types.ActionInputs, error) {
	var inputs types.ActionInputs
	switch actionInputsType {
	case "approve":
		inputs = &actions.ApproveInputsImpl{}
	case "transfer":
		inputs = &actions.TransferInputsImpl{}
	case "transfer_from":
		inputs = &actions.TransferFromInputsImpl{}
	case "route":
		inputs = &actions.RouteInputsImpl{}
	case "swap":
		inputs = &actions.SwapInputsImpl{}
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
	default:
		return nil, fmt.Errorf("unknown action type: %s", actionInputsType)
	}
	return inputs, nil
}

func ValidateActionInputs(action types.Action) (types.ActionInputs, error) {
	inputsReference, err := GetActionInputs(action.Type)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(action.Inputs, inputsReference); err != nil {
		return nil, fmt.Errorf("failed to parse inputs for action %s: %v", action.Type, err)
	}

	if err := inputsReference.Validate(); err != nil {
		return nil, fmt.Errorf("invalid inputs for action %s: %v", action.Type, err)
	}

	return inputsReference, nil
}
