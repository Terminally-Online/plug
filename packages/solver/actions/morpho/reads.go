package morpho

import (
	"bytes"
	"encoding/json"
	"fmt"
	"solver/utils"
)

const (
	morphoApiUrl = "https://blue-api.morpho.org/graphql"
	marketsQuery = `query {
		markets(where:{whitelisted: true}, skip: %d) {
			items {
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
				state {
					borrowApy
					borrowAssets
					borrowAssetsUsd
					supplyApy
					supplyAssets
					supplyAssetsUsd
					fee
					utilization
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

type Asset struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

type MarketState struct {
	BorrowApy       float64 `json:"borrowApy"`
	BorrowAssetsUsd float64 `json:"borrowAssetsUsd"`
	SupplyApy       float64 `json:"supplyApy"`
	SupplyAssetsUsd float64 `json:"supplyAssetsUsd"`
	Fee             float64 `json:"fee"`
	Utilization     float64 `json:"utilization"`
}

type Market struct {
	Whitelisted     bool        `json:"whitelisted"`
	UniqueKey       string      `json:"uniqueKey"`
	LLTV            interface{} `json:"lltv"`
	OracleAddress   string      `json:"oracleAddress"`
	IRMAddress      string      `json:"irmAddress"`
	LoanAsset       Asset       `json:"loanAsset"`
	CollateralAsset Asset       `json:"collateralAsset"`
	State           MarketState `json:"state"`
}

type MarketResponse struct {
	Data struct {
		Markets struct {
			Items    []Market `json:"items"`
			PageInfo struct {
				CountTotal int `json:"countTotal"`
				Count      int `json:"count"`
			} `json:"pageInfo"`
		} `json:"markets"`
	} `json:"data"`
}

func GetMarkets() ([]Market, error) {
	var allMarkets []Market
	skip := 0
	limit := 100

	for {
		paginatedQuery := fmt.Sprintf(marketsQuery, skip)
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
			MarketResponse{},
		)
		if err != nil {
			return []Market{}, err
		}

		allMarkets = append(allMarkets, marketResponse.Data.Markets.Items...)

		if marketResponse.Data.Markets.PageInfo.Count < limit ||
			skip+marketResponse.Data.Markets.PageInfo.Count >= marketResponse.Data.Markets.PageInfo.CountTotal {
			break
		}

		skip += limit
	}

	return allMarkets, nil
}
