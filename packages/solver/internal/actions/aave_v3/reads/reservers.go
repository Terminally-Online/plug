package reads

import (
	"os"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/internal/bindings/references"
	"solver/internal/client"

	"github.com/ethereum/go-ethereum/common"
)

func GetReserves(chainId uint64) ([]aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData, error) {
	client, err := client.New(chainId)
	if err != nil {
		return nil, err
	}
	dataProvider, err := aave_v3_ui_pool_data_provider.NewAaveV3UiPoolDataProvider(
		common.HexToAddress(references.Networks[chainId].References["aave_v3"]["ui_pool_data_provider"]),
		client,
	)
	if err != nil {
		return nil, err
	}
	reserves, _, err := dataProvider.GetReservesData(
		client.ReadOptions(common.HexToAddress(os.Getenv("SOLVER_ADDRESS"))),
		common.HexToAddress(references.Networks[chainId].References["aave_v3"]["ui_pool_address_provider"]),
	)
	if err != nil {
		return nil, err
	}

	return reserves, nil
}
