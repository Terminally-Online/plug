package aave_v3

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"solver/bindings/aave_v3_pool"
	"solver/types"
	"solver/utils"
)

var (
	Key = "aave_v3"

	address          = utils.Mainnet.References[Key]["pool"]
	hexAddress       = common.HexToAddress(address)
	interestRateMode = new(big.Int).SetUint64(2)
)

func BuildDeposit(i types.DepositInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	contract, err := aave_v3_pool.NewAaveV3Pool(hexAddress, provider)
	if err != nil {
		return nil, utils.ErrContractFailed(address)
	}

	return contract.Supply(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.GetTokenOut()),
		i.GetAmountIn(),
		common.HexToAddress(from),
		uint16(0),
	)
}

func BuildBorrow(i types.BorrowInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	contract, err := aave_v3_pool.NewAaveV3Pool(hexAddress, provider)
	if err != nil {
		return nil, utils.ErrContractFailed(address)
	}

	return contract.Borrow(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.GetTokenOut()),
		i.GetAmountOut(),
		interestRateMode,
		uint16(0),
		common.HexToAddress(from),
	)
}

func BuildRedeem(i types.RepayInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	contract, err := aave_v3_pool.NewAaveV3Pool(hexAddress, provider)
	if err != nil {
		return nil, utils.ErrContractFailed(address)
	}

	return contract.Withdraw(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.GetTokenIn()),
		i.GetAmountIn(),
		common.HexToAddress(from),
	)
}

func BuildRepay(i types.RepayInputs, provider *ethclient.Client, chainId int, from string) (*ethtypes.Transaction, error) {
	contract, err := aave_v3_pool.NewAaveV3Pool(hexAddress, provider)
	if err != nil {
		return nil, utils.ErrContractFailed(address)
	}

	return contract.Repay(
		utils.DummyTransactOpts(from, big.NewInt(0)),
		common.HexToAddress(i.GetTokenIn()),
		i.GetAmountIn(),
		interestRateMode,
		common.HexToAddress(from),
	)
}
