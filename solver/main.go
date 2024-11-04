package solver

import (
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

    params := actions.HandlerParams{
        Provider: provider,
        ChainId: chainId,
        From: from,
    }

    switch action {
    case types.ActionDeposit:
        depositInputs, ok := inputs.(*types.DepositInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for deposit action")
        }
        depositHandler, ok := handler.(actions.DepositHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement deposit handler")
        }
        return depositHandler.HandlePostDeposit(depositInputs, params)

    case types.ActionBorrow:
        borrowInputs, ok := inputs.(*types.BorrowInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for borrow action")
        }
        borrowHandler, ok := handler.(actions.BorrowHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement borrow handler")
        }
        return borrowHandler.HandlePostBorrow(borrowInputs, params)

    case types.ActionRedeem:
        redeemInputs, ok := inputs.(*types.RedeemInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for redeem action")
        }
        redeemHandler, ok := handler.(actions.RedeemHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement redeem handler")
        }
        return redeemHandler.HandlePostRedeem(redeemInputs, params)

    case types.ActionRepay:
        repayInputs, ok := inputs.(*types.RepayInputs)
        if !ok {
            return nil, fmt.Errorf("invalid input type for repay action")
        }
        repayHandler, ok := handler.(actions.RepayHandler)
        if !ok {
            return nil, fmt.Errorf("protocol does not implement repay handler")
        }
        return repayHandler.HandlePostRepay(repayInputs, params)

    default:
        return nil, fmt.Errorf("unsupported action: %s", action)
    }
}

func (s *Solver) GetProtocolHandler(protocol types.Protocol) (actions.BaseProtocolHandler, bool) {
    handler, exists := s.protocols[protocol]
    return handler, exists
}

func (s *Solver) SupportsAction(handler actions.BaseProtocolHandler, action types.Action) bool {
    for _, a := range handler.SupportedActions() {
        if a == action {
            return true
        }
    }
    return false
}
