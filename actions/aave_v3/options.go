package aave_v3

import (
	"fmt"
	"math/big"
	"os"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/types"
	"solver/utils"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

var (
	uiPoolDataProviderAddress  = utils.Mainnet.References["aave_v3"]["ui_pool_data_provider"]
	poolAddressProviderAddress = "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e"
	
	reservesCache     []aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData
	lastCacheUpdate   int64
	cacheDuration     int64 = 300
)

func getReserves() ([]aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData, error) {
	currentTime := time.Now().Unix()
	if reservesCache != nil && (currentTime - lastCacheUpdate) < cacheDuration {
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

func GetCollateralAssetOptions() ([]types.Option, error) {
	reserves, err := getReserves()
	if err != nil {
		return nil, err
	}

	options := make([]types.Option, 0)
	for _, reserve := range reserves {
		// TODO: (#12) Does not include 'Isolated' assets as optional collateral due to the nuance of
		// variable allowance. For this to be supported, we need arrow function support across
		// multiple actions because the depositing of a collateral asset that is isolated cannot
		// be used across all forms of borrowable assets -- The exposure is limited.
		//
		// NOTE: Right now they are filtered out by checking the `DebtCeiling` value because isolated
		// assets are the only ones that have a non-zero value. This is a bit of a hack and should
		// be revisited in the future.
		//
		// NOTE: Realistically, this is not something we will probably ever support though so the only
		// other option is to just return the isolated collateral assets and let a user
		// figure it out themselves which seems less ideal. The isolated assets however are the
		// exogenous bases while non-isolated tend to be native and stable assets.
		if !reserve.UsageAsCollateralEnabled || reserve.DebtCeiling.Cmp(big.NewInt(0)) > 0 {
			continue
		}

		rateFloat := new(big.Float).Quo(
			new(big.Float).SetInt(reserve.LiquidityRate),
			new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)),
		)
		
		var rate string
		if rateFloat.Cmp(big.NewFloat(0)) > 0 && rateFloat.Cmp(big.NewFloat(0.01)) < 0 {
			rate = "<0.01%"
		} else {
			rate = rateFloat.Text('f', 2) + "%"
		}
		options = append(options, types.Option{
			Value: reserve.UnderlyingAsset.String(),
			Label: reserve.Symbol,
			Name:  reserve.Name,
			Info:  rate,
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", 1, reserve.UnderlyingAsset.String()),
		})
	}

	return options, nil
}

func GetBorrowAssetOptions() ([]types.Option, error) {
	reserves, err := getReserves()
	if err != nil {
		return nil, err
	}

	options := make([]types.Option, 0)
	for _, reserve := range reserves {
		if !reserve.BorrowingEnabled {
			continue
		}

		rateFloat := new(big.Float).Quo(
			new(big.Float).SetInt(reserve.VariableBorrowRate),
			new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)),
		)
		
		var rate string
		if rateFloat.Cmp(big.NewFloat(0)) > 0 && rateFloat.Cmp(big.NewFloat(0.01)) < 0 {
			rate = "<0.01%"
		} else {
			rate = rateFloat.Text('f', 2) + "%"
		}

		options = append(options, types.Option{
			Value: reserve.UnderlyingAsset.String(),
			Label: reserve.Symbol,
			Name:  reserve.Name,
			Info:  rate,
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", 1, reserve.UnderlyingAsset.String()),
		})
	}

	return options, nil
}

func GetOptions() ([]types.Option, []types.Option, error) {
	collateralOptions, err := GetCollateralAssetOptions()
	if err != nil {
		return nil, nil, err
	}
	debtOptions, err := GetBorrowAssetOptions()
	if err != nil {
		return nil, nil, err
	}
	return collateralOptions, debtOptions, nil
}
