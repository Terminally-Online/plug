package reads

import (
	"math/big"
	"solver/bindings/aave_v3_pool"
	"solver/internal/bindings/references"
	"solver/internal/client"

	"github.com/ethereum/go-ethereum/common"
)

func GetHealthFactor(chainId uint64, userAddress common.Address) (*big.Int, error) {
	client, err := client.New(chainId)
	if err != nil {
		return nil, err
	}
	pool, err := aave_v3_pool.NewAaveV3Pool(
		common.HexToAddress(references.Networks[chainId].References["aave_v3"]["pool"]),
		client,
	)
	if err != nil {
		return nil, err
	}

	userAccountData, err := pool.GetUserAccountData(
		client.SolverReadOptions(),
		userAddress,
	)
	if err != nil {
		return nil, err
	}

	return userAccountData.HealthFactor, nil
}
