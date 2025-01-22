package aave_v3

import (
	"math/big"
	"os"
	"solver/bindings/aave_v3_pool"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/internal/references"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

func getReserves(chainId int) ([]aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}
	dataProvider, err := aave_v3_ui_pool_data_provider.NewAaveV3UiPoolDataProvider(
		common.HexToAddress(references.Networks[chainId].References["aave_v3"]["ui_pool_data_provider"]),
		provider,
	)
	if err != nil {
		return nil, err
	}
	reserves, _, err := dataProvider.GetReservesData(
		utils.BuildCallOpts(os.Getenv("SOLVER_ADDRESS"), big.NewInt(0)),
		common.HexToAddress(references.Networks[chainId].References["aave_v3"]["ui_pool_address_provider"]),
	)
	if err != nil {
		return nil, err
	}

	return reserves, nil
}

func getHealthFactor(chainId int, userAddress string) (*big.Int, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}
	pool, err := aave_v3_pool.NewAaveV3Pool(
		common.HexToAddress(references.Networks[chainId].References["aave"]["pool"]),
		provider,
	)
	if err != nil {
		return nil, err
	}

	userAccountData, err := pool.GetUserAccountData(
		utils.BuildCallOpts(os.Getenv("SOLVER_ADDRESS"), big.NewInt(0)),
		common.HexToAddress(userAddress),
	)
	if err != nil {
		return nil, err
	}

	return userAccountData.HealthFactor, nil
}
