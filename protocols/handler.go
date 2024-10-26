package protocols

import (
	"solver/types"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ProtocolHandler interface {
	SupportedActions() []types.Action
	SupportedChains() []int

	HandleGetDeposit() types.ActionSchema
	HandlePostDeposit(inputs *types.DepositInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error)

	HandleGetBorrow() types.ActionSchema
	HandlePostBorrow(inputs *types.BorrowInputs, provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error)
}
