package solver

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/actions/aave_v2"
	"solver/types"
	"solver/utils"
)

type Solver struct {
	protocols map[types.Protocol]actions.BaseProtocolHandler
}

func New() *Solver {
	return &Solver{
		protocols: map[types.Protocol]actions.BaseProtocolHandler{
			types.ProtocolAaveV2: aave_v2.New(),
		},
	}
}

func (s *Solver) GetSupportedProtocols(action types.Action) []types.Protocol {
	supported := make([]types.Protocol, 0)
	for protocol, handler := range s.protocols {
		for _, supportedAction := range handler.GetActions() {
			if supportedAction == action {
				supported = append(supported, protocol)
				break
			}
		}
	}
	return supported
}

func (s *Solver) GetTransaction(action types.Action, rawInputs json.RawMessage, chainId int, from string) ([]*types.Transaction, error) {
	var baseInputs types.BaseInputs
	if err := json.Unmarshal(rawInputs, &baseInputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base inputs: %v", err)
	}

	handler, exists := s.protocols[baseInputs.Protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", baseInputs.Protocol)
	}

	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %v", err)
	}

	params := actions.HandlerParams{
		Provider: provider,
		ChainId:  chainId,
		From:     from,
	}

	return handler.GetTransaction(action, rawInputs, params)
}

func (s *Solver) GetProtocolHandler(protocol types.Protocol) (actions.BaseProtocolHandler, bool) {
	handler, exists := s.protocols[protocol]
	return handler, exists
}

func (s *Solver) GetProtocols() map[types.Protocol]actions.BaseProtocolHandler {
	return s.protocols
}

func (s *Solver) SupportsAction(handler actions.BaseProtocolHandler, action types.Action) bool {
	for _, a := range handler.GetActions() {
		if a == action {
			return true
		}
	}
	return false
}
