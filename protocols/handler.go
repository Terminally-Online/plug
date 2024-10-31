package protocols

import (
	"solver/types"

	"github.com/ethereum/go-ethereum/ethclient"
)

type BaseProtocolHandler interface {
    SupportedActions() []types.Action
    SupportedChains() []int
}

type DepositHandler interface {
    HandleGetDeposit() types.ActionSchema
    HandlePostDeposit(inputs *types.DepositInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error)
}

type BorrowHandler interface {
    HandleGetBorrow() types.ActionSchema
    HandlePostBorrow(inputs *types.BorrowInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error)
}

type RedeemHandler interface {
    HandleGetRedeem() types.ActionSchema
    HandlePostRedeem(inputs *types.RedeemInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error)
}

type RepayHandler interface {
    HandleGetRepay() types.ActionSchema
    HandlePostRepay(inputs *types.RepayInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error)
}

type Protocol struct {
    Name string
    SupportedChains []int
}
