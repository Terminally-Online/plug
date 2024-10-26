package protocols

import (
	"solver/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProtocolHandler interface {
	SupportedActions() []types.Action
	SupportedChains() []int

	HandleGetDeposit() types.ActionSchema
	HandlePostDeposit(inputs *types.DepositInputs, provider *ethclient.Client, chainId int, from string) ([]*ethtypes.Transaction, error)

	HandleGetBorrow() types.ActionSchema
	HandlePostBorrow(inputs *types.BorrowInputs, provider *ethclient.Client, chainId int, from string) ([]*ethtypes.Transaction, error)
}
