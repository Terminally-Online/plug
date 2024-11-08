package aave_v3

import (
	"math/big"
	"os"
	"solver/bindings/aave_v3_pool"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/utils"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func getReserves(force ...bool) ([]aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData, error) {
	currentTime := time.Now().Unix()
	if !((len(force) > 0 && force[0]) || reservesCache == nil || (currentTime-lastCacheUpdate) >= cacheDuration) {
		return reservesCache, nil
	}
	
	provider, err := utils.GetProvider(1)
	if err != nil {
		return nil, err
	}
	dataProvider, err := aave_v3_ui_pool_data_provider.NewAaveV3UiPoolDataProvider(common.HexToAddress(uiPoolDataProviderAddress), provider)
	if err != nil {
		return nil, err
	}

	reserves, _, err := dataProvider.GetReservesData(
		utils.BuildCallOpts(os.Getenv("SOLVER_ADDRESS"), big.NewInt(0)),
		common.HexToAddress(poolAddressProviderAddress),
	)
	if err != nil {
		return nil, err
	}

	reservesCache = reserves
	lastCacheUpdate = currentTime

	return reserves, nil
}

func getHealthFactor(userAddress string) (*big.Int, error) {
	provider, err := utils.GetProvider(1)
	if err != nil {
		return nil, err
	}
	pool, err := aave_v3_pool.NewAaveV3Pool(common.HexToAddress(poolAddress), provider)
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
