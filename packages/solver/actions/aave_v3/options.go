package aave_v3

import (
	"fmt"
	"math/big"
	"solver/types"
	"solver/utils"
)

var (
	uiPoolDataProviderAddress  = utils.Mainnet.References["aave_v3"]["ui_pool_data_provider"]
	poolAddressProviderAddress = "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e"
)

func GetCollateralAssetOptions(chainId int) ([]types.Option, error) {
	reserves, err := getReserves(chainId)
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
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", 1, reserve.UnderlyingAsset.String()),
			Label: reserve.Symbol,
			Name:  reserve.Name,
			Info:  rate,
			Value: reserve.UnderlyingAsset.String(),
		})
	}

	return options, nil
}

func GetBorrowAssetOptions(chainId int) ([]types.Option, error) {
	reserves, err := getReserves(chainId)
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
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", 1, reserve.UnderlyingAsset.String()),
			Label: reserve.Symbol,
			Name:  reserve.Name,
			Info:  rate,
			Value: reserve.UnderlyingAsset.String(),
		})
	}

	return options, nil
}

func GetOptions(chainId int) ([]types.Option, []types.Option, error) {
	collateralOptions, err := GetCollateralAssetOptions(chainId)
	if err != nil {
		return nil, nil, err
	}
	debtOptions, err := GetBorrowAssetOptions(chainId)
	if err != nil {
		return nil, nil, err
	}
	return collateralOptions, debtOptions, nil
}
