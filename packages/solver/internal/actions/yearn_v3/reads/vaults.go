package reads

import (
	"fmt"
	"solver/internal/actions/yearn_v3/types"
	"solver/internal/utils"
	"strings"
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

func GetVault(chainId uint64, vaultAddress string) (*types.YearnVault, error) {
	vaults, err := GetVaults(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}

	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, vaultAddress) {
			return &vault, nil
		}
	}

	return nil, fmt.Errorf("vault not found: %s", vaultAddress)
}
