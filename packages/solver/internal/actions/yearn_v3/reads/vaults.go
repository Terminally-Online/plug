package reads

import (
	"fmt"
	"solver/internal/actions/yearn_v3/types"
	"solver/internal/utils"
)

func GetVaults(chainId uint64, force ...bool) ([]types.YearnVault, error) {
	url := fmt.Sprintf("https://ydaemon.yearn.finance/%d/vaults/all?limit=99999", chainId)
	response, err := utils.MakeHTTPRequest(
		url,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
		},
		nil,
		nil,
		[]types.YearnVault{},
	)
	if err != nil {
		return nil, err
	}

	endorsedVaults := make([]types.YearnVault, 0)
	for _, vault := range response {
		if vault.Endorsed && !vault.EmergencyShutdown && !vault.Details.IsRetired && !vault.Details.IsHidden && uint64(vault.ChainID) == chainId {
			endorsedVaults = append(endorsedVaults, vault)
		}
	}

	return endorsedVaults, nil
}
