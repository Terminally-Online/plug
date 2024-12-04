package solver

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/actions/aave_v3"
	"solver/actions/ens"
	"solver/actions/plug"
	"solver/actions/yearn_v3"
	"solver/types"
	"solver/utils"
)

type Solver struct {
	protocols map[types.Protocol]actions.BaseProtocolHandler
}

func New() *Solver {
	return &Solver{
		protocols: map[types.Protocol]actions.BaseProtocolHandler{
			types.ProtocolPlug:    plug.New(),
			types.ProtocolAaveV3:  aave_v3.New(),
			types.ProtocolYearnV3: yearn_v3.New(),
			types.ProtocolENS:     ens.New(),
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

func (s *Solver) GetTransaction(rawInputs json.RawMessage, chainId int, from string) ([]*types.Transaction, error) {
	var inputs types.BaseInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal base inputs: %v", err)
	}

	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	handler, exists := s.protocols[inputs.Protocol]
	if !exists {
		return nil, fmt.Errorf("unsupported protocol: %s", inputs.Protocol)
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

	return handler.GetTransaction(inputs.Action, rawInputs, params)
}

func (s *Solver) GetProtocolHandler(protocol types.Protocol) (actions.BaseProtocolHandler, bool) {
	handler, exists := s.protocols[protocol]
	return handler, exists
}

func (s *Solver) GetProtocols() map[types.Protocol]actions.BaseProtocolHandler {
	return s.protocols
}
