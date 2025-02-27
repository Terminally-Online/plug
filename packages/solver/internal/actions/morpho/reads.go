package morpho

import (
	"bytes"
	"encoding/json"
	"fmt"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

const (
	morphoApiUrl = "https://blue-api.morpho.org/graphql"
	marketsQuery = `query {
		markets(where: { whitelisted: true, chainId_in: [%d] }, skip: %d) {
			items {
				whitelisted
				uniqueKey
				lltv
				oracleAddress
				irmAddress
				loanAsset {
					address
					symbol
					name
					decimals
					logoURI
				}
				collateralAsset {
					address
					symbol
					name
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
			pageInfo {
				countTotal
				limit
				count
			}
		}
	}`
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
	rewardsApiUrl = "https://rewards.morpho.org/v1/users/%s/distributions?chain_id=%d"
)

func GetVaults(chainId uint64) ([]Vault, error) {
	var vaults []Vault
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
			return []Vault{}, err
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
						Items    []Vault `json:"items"`
						PageInfo struct {
							CountTotal int `json:"countTotal"`
							Count      int `json:"count"`
						} `json:"pageInfo"`
					} `json:"vaults"`
				} `json:"data"`
			}{},
		)
		if err != nil {
			return []Vault{}, err
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

func GetVault(address string, chainId uint64) (Vault, error) {
	vaults, err := GetVaults(chainId)
	if err != nil {
		return Vault{}, err
	}
	for _, vault := range vaults {
		if vault.Address == address {
			return vault, nil
		}
	}
	return Vault{}, fmt.Errorf("vault not found for address: %s", address)
}

func GetMarkets(chainId uint64) ([]Market, error) {
	var markets []Market
	skip := 0
	limit := 100

	for {
		paginatedQuery := fmt.Sprintf(marketsQuery, chainId, skip)
		requestBody := struct {
			Query string `json:"query"`
		}{
			Query: paginatedQuery,
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return []Market{}, err
		}

		marketResponse, err := utils.MakeHTTPRequest(
			morphoApiUrl,
			"POST",
			map[string]string{
				"Content-Type": "application/json",
			},
			nil,
			bytes.NewBuffer(jsonBody),
			struct {
				Data struct {
					Markets Markets `json:"markets"`
				} `json:"data"`
			}{},
		)
		if err != nil {
			return []Market{}, err
		}

		markets = append(markets, marketResponse.Data.Markets.Items...)

		if marketResponse.Data.Markets.PageInfo.Count < limit ||
			skip+marketResponse.Data.Markets.PageInfo.Count >= marketResponse.Data.Markets.PageInfo.CountTotal {
			break
		}

		skip += limit
	}

	return markets, nil
}

func GetMarket(uniqueKey string, chainId uint64) (Market, error) {
	markets, err := GetMarkets(chainId)
	if err != nil {
		return Market{}, err
	}
	for _, market := range markets {
		if market.UniqueKey == uniqueKey {
			return market, nil
		}
	}
	return Market{}, fmt.Errorf("market not found for unique key: %s", uniqueKey)
}

func GetDistributions(address common.Address, chainId uint64) ([]Distribution, error) {
	url := fmt.Sprintf(rewardsApiUrl, address, chainId)

	response, err := utils.MakeHTTPRequest(
		url,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
		},
		nil,
		nil,
		DistributionResponse{},
	)
	if err != nil {
		return []Distribution{}, err
	}

	return response.Data, nil
}
