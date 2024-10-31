package solver

import (
    "fmt"
    "solver/protocols"
    "solver/protocols/aave_v2"
    "solver/types"
    "solver/utils"
)

type Solver struct {
    protocols map[types.Protocol]protocols.BaseProtocolHandler
}

func New() *Solver {
    return &Solver{
        protocols: map[types.Protocol]protocols.BaseProtocolHandler{
            types.ProtocolAaveV2: aave_v2.New(),
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

func (s *Solver) BuildTransaction(action types.Action, inputs types.ActionInputs, chainId int, from string) ([]*types.Transaction, error) {
    if err := inputs.Validate(); err != nil {
        return nil, fmt.Errorf("invalid inputs: %w", err)
    }

    protocol := inputs.GetProtocol()
    handler, exists := s.protocols[protocol]
    if !exists {
        return nil, fmt.Errorf("unsupported protocol: %s", protocol)
    }

    // Check chain support
    supported := false
    for _, supportedChain := range handler.SupportedChains() {
        if supportedChain == chainId {
            supported = true
            break
        }
    }
    if !supported {
        return nil, fmt.Errorf("chain %d not supported by protocol %s", chainId, protocol)
    }

    supported = false
    for _, a := range handler.SupportedActions() {
        if a == action {
            supported = true
            break
        }
    }
    if !supported {
        return nil, fmt.Errorf("action %s not supported by protocol %s", action, protocol)
    }

    provider, err := utils.GetProvider(chainId)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
    }

    switch action {
    case types.ActionDeposit:
        depositInputs, ok := inputs.(*types.DepositInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for deposit action")
        }
        depositHandler, ok := handler.(protocols.DepositHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement deposit handler")
        }
        return depositHandler.HandlePostDeposit(depositInputs, provider, chainId, from)

    case types.ActionBorrow:
        borrowInputs, ok := inputs.(*types.BorrowInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for borrow action")
        }
        borrowHandler, ok := handler.(protocols.BorrowHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement borrow handler")
        }
        return borrowHandler.HandlePostBorrow(borrowInputs, provider, chainId, from)

    case types.ActionRedeem:
        redeemInputs, ok := inputs.(*types.RedeemInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for borrow action")
        }
        redeemHandler, ok := handler.(protocols.RedeemHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement borrow handler")
        }
        return redeemHandler.HandlePostRedeem(redeemInputs, provider, chainId, from)

    case types.ActionRepay:
        repayInputs, ok := inputs.(*types.RepayInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for borrow action")
        }
        repayHandler, ok := handler.(protocols.RepayHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement borrow handler")
        }
        return repayHandler.HandlePostRepay(repayInputs, provider, chainId, from)

    default:
        return nil, fmt.Errorf("unsupported action: %s", action)
    }
}

func (s *Solver) GetProtocolHandler(protocol types.Protocol) (protocols.BaseProtocolHandler, bool) {
    handler, exists := s.protocols[protocol]
    return handler, exists
}

func (s *Solver) SupportsAction(handler protocols.BaseProtocolHandler, action types.Action) bool {
    for _, a := range handler.SupportedActions() {
        if a == action {
            return true
        }
    }
    return false
}
