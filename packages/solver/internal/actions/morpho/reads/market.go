package reads

import (
	"bytes"
	"encoding/json"
	"fmt"
	"solver/internal/actions/morpho/types"
	"solver/internal/utils"
)

var (
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
)

func GetMarkets(chainId uint64) ([]types.Market, error) {
	var markets []types.Market
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
			return []types.Market{}, err
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
					Markets types.Markets `json:"markets"`
				} `json:"data"`
			}{},
		)
		if err != nil {
			return []types.Market{}, err
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

func GetMarket(uniqueKey string, chainId uint64) (types.Market, error) {
	markets, err := GetMarkets(chainId)
	if err != nil {
		return types.Market{}, err
	}
	for _, market := range markets {
		if market.UniqueKey == uniqueKey {
			return market, nil
		}
	}
	return types.Market{}, fmt.Errorf("market not found for unique key: %s", uniqueKey)
}
