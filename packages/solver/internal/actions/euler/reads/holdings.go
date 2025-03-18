package reads

import (
	"fmt"
	"solver/internal/helpers/zerion"
	"solver/internal/utils"
	"time"

	"strings"

	"github.com/ethereum/go-ethereum/common"
)


func GetMainAddressVaultHoldings(address common.Address, chainId uint64) ([]zerion.ZerionPosition, error) {
	cacheKey := fmt.Sprintf("euler:mainPositions:%s:%d", address, chainId)
	res, err := utils.WithCache(cacheKey, []time.Duration{5 * time.Minute}, true, func() ([]zerion.ZerionPosition, error) {
		vaults, err := GetVerifiedVaults(chainId)
		if err != nil {
			return nil, err
		}
		vaultAddresses := make(map[string]bool)
		for _, vault := range vaults {
			vaultAddresses[strings.ToLower(vault.Vault.String())] = true
		}

		positions, err := zerion.GetFungiblePositions([]string{"base"}, address, address, "")
		if err != nil {
			return nil, err
		}

		zerionPositions := make([]zerion.ZerionPosition, 0)
		for _, position := range positions {
			for _, impl := range position.Attributes.FungibleInfo.Implementations {
				if impl.ChainID == fmt.Sprintf("%d", chainId) {
					// Check if this implementation's address matches any vault
					if vaultAddresses[strings.ToLower(impl.Address)] {
						zerionPositions = append(zerionPositions, position)
						break // Found a match, no need to check other implementations
					}
				}
			}
		}

		return zerionPositions, nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
