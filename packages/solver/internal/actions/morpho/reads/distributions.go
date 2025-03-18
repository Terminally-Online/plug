package reads

import (
	"fmt"
	"solver/internal/actions/morpho/types"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

var (
	rewardsApiUrl = "https://rewards.morpho.org/v1/users/%s/distributions?chain_id=%d"
)

func GetDistributions(address common.Address, chainId uint64) ([]types.Distribution, error) {
	url := fmt.Sprintf(rewardsApiUrl, address, chainId)

	response, err := utils.MakeHTTPRequest(
		url,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
		},
		nil,
		nil,
		types.DistributionResponse{},
	)
	if err != nil {
		return []types.Distribution{}, err
	}

	return response.Data, nil
}
