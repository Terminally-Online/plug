package reads

import (
	"bytes"
	"encoding/json"
	"fmt"
	"solver/internal/actions/morpho/types"
	"solver/internal/utils"
)

const (
	vaultsQuery = `query {
	 	vaults(where: { whitelisted: true, chainId_in: [%d] }, skip: %d) {
			items {
				address
				name
				symbol
				metadata {
					image
					description
				}
				asset {
					address
					decimals
					name
					symbol
					logoURI
				}
				dailyApys {
					apy
					netApy
				}
				state {
					allocation {
						enabled
						market {
							whitelisted
							uniqueKey
							lltv
							oracleAddress
							irmAddress
							loanAsset {
								address
								symbol
								decimals
								logoURI
							}
							collateralAsset {
								address
								symbol
								decimals
								logoURI
							}
							dailyApys {
								borrowApy
								supplyApy
								netBorrowApy
								netSupplyApy
							}
							state {
								price
								borrowShares
								borrowAssets
								borrowApy
								dailyBorrowApy
								supplyShares
								supplyAssets
								supplyApy
								dailySupplyApy
							}
						}
					}
				}
			}
			pageInfo {
				count
				countTotal
				limit
			}
		}
	}`
)

func GetVaults(chainId uint64) ([]types.Vault, error) {
	var vaults []types.Vault
	skip := 0
	limit := 100

	for {
		paginatedQuery := fmt.Sprintf(vaultsQuery, chainId, skip)
		requestBody := struct {
			Query string `json:"query"`
		}{
			Query: paginatedQuery,
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return []types.Vault{}, err
		}

		vaultResponse, err := utils.MakeHTTPRequest(
			morphoApiUrl,
			"POST",
			map[string]string{
				"Content-Type": "application/json",
			},
			nil,
			bytes.NewBuffer(jsonBody),
			struct {
				Data struct {
					Vaults struct {
						Items    []types.Vault `json:"items"`
						PageInfo struct {
							CountTotal int `json:"countTotal"`
							Count      int `json:"count"`
						} `json:"pageInfo"`
					} `json:"vaults"`
				} `json:"data"`
			}{},
		)
		if err != nil {
			return []types.Vault{}, err
		}

		vaults = append(vaults, vaultResponse.Data.Vaults.Items...)

		if vaultResponse.Data.Vaults.PageInfo.Count < limit ||
			skip+vaultResponse.Data.Vaults.PageInfo.Count >= vaultResponse.Data.Vaults.PageInfo.CountTotal {
			break
		}

		skip += limit
	}

	return vaults, nil
}

func GetVault(address string, chainId uint64) (types.Vault, error) {
	vaults, err := GetVaults(chainId)
	if err != nil {
		return types.Vault{}, err
	}
	for _, vault := range vaults {
		if vault.Address == address {
			return vault, nil
		}
	}
	return types.Vault{}, fmt.Errorf("vault not found for address: %s", address)
}
