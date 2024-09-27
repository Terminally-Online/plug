package aave_v2

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"solver/types"
)

var (
	Key = "aave_v2"
)

func BuildBorrow(inputs types.BorrowInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	return nil, nil
}

func BuildDeposit(inputs types.DepositInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	return nil, nil
}

func BuildRedeem(inputs types.RepayInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	return nil, nil
}

func BuildRepay(inputs types.RepayInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	return nil, nil
}
