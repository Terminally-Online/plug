package solver

import (
	"fmt"
	"solver/protocols"
	"solver/protocols/aave_v2"
	"solver/types"
)

type Solver struct {
	protocols map[types.Protocol]protocols.ProtocolHandler
}

func New() *Solver {
	return &Solver{
		protocols: map[types.Protocol]protocols.ProtocolHandler{
			types.ProtocolAave: aave_v2.New(),
			// types.ProtocolCompound: compound.New(),
		},
	}
}

func (s *Solver) GetSupportedProtocols(action types.Action) []types.Protocol {
	supported := make([]types.Protocol, 0)
	for protocol, handler := range s.protocols {
		for _, supportedAction := range handler.SupportedActions() {
			if supportedAction == action {
				supported = append(supported, protocol)
				break
			}
		}
	}
	return supported
}

func (s *Solver) GetActionSchema(action types.Action) ([]types.ActionSchema, error) {
	schemas := make([]types.ActionSchema, 0)

	for _, handler := range s.protocols {
		supported := false
		for _, a := range handler.SupportedActions() {
			if a == action {
				supported = true
				break
			}
		}
		if !supported {
			continue
		}

		var schema types.ActionSchema
		switch action {
		case types.ActionDeposit:
			schema = handler.HandleGetDeposit()
		case types.ActionBorrow:
			schema = handler.HandleGetBorrow()
		default:
			return nil, fmt.Errorf("unsupported action: %s", action)
		}

		schemas = append(schemas, schema)
	}

	if len(schemas) == 0 {
		return nil, fmt.Errorf("no protocols support action: %s", action)
	}

	return schemas, nil
}

func (s *Solver) BuildTransaction(action types.Action, inputs types.ActionInputs) (*types.Transaction, error) {
	// Validate inputs first
	if err := inputs.Validate(); err != nil {
		return nil, fmt.Errorf("invalid inputs: %w", err)
	}

	protocol := inputs.GetProtocol()
	handler, exists := s.protocols[protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", protocol)
	}

	// Verify protocol supports this action
	supported := false
	for _, a := range handler.SupportedActions() {
		if a == action {
			supported = true
			break
		}
	}
	if !supported {
		return nil, fmt.Errorf("action %s not supported by protocol %s", action, protocol)
	}

	switch action {
	case types.ActionDeposit:
		depositInputs, ok := inputs.(*types.DepositInputs)
		if !ok {
			return nil, fmt.Errorf("invalid input type for deposit action")
		}
		return handler.HandlePostDeposit(depositInputs)
	case types.ActionBorrow:
		borrowInputs, ok := inputs.(*types.BorrowInputs)
		if !ok {
			return nil, fmt.Errorf("invalid input type for borrow action")
		}
		return handler.HandlePostBorrow(borrowInputs)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}

func (s *Solver) GetProtocolHandler(protocol types.Protocol) (protocols.ProtocolHandler, bool) {
	handler, exists := s.protocols[protocol]
	return handler, exists
}

// func Example() {
// 	solver := New()

// 	// Protocol must be specified in inputs
// 	inputs := &types.DepositInputs{
// 		BaseInputs: types.BaseInputs{
// 			Protocol: types.ProtocolAave, // Required
// 		},
// 		TokenIn:  "0x...",
// 		TokenOut: "0x...",
// 		AmountIn: *big.NewInt(1000000),
// 	}

// 	_, err := solver.BuildTransaction(types.ActionDeposit, inputs)
// 	if err != nil {
// 		// Handle error
// 	}
// }
