package actions

import (
    "solver/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

type HandlerParams struct {
    Provider *ethclient.Client
    ChainId  int
    From     string
}

type BaseProtocolHandler interface {
    SupportedActions() []types.Action
    SupportedChains() []int
}

type DepositHandler interface {
    HandleGetDeposit() types.ActionSchema
    HandlePostDeposit(inputs *types.DepositInputs, params HandlerParams) ([]*types.Transaction, error)
}

type BorrowHandler interface {
    HandleGetBorrow() types.ActionSchema
    HandlePostBorrow(inputs *types.BorrowInputs, params HandlerParams) ([]*types.Transaction, error)
}

type RedeemHandler interface {
    HandleGetRedeem() types.ActionSchema
    HandlePostRedeem(inputs *types.RedeemInputs, params HandlerParams) ([]*types.Transaction, error)
}

type RepayHandler interface {
    HandleGetRepay() types.ActionSchema
    HandlePostRepay(inputs *types.RepayInputs, params HandlerParams) ([]*types.Transaction, error)
}

type HarvestHandler interface {
    HandleGetHarvest() types.ActionSchema
    HandlePostHarvest(inputs *types.HarvestInputs, params HandlerParams) ([]*types.Transaction, error)
}

type Protocol struct {
    Name string
    SupportedChains []int
}
