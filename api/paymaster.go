package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"solver/utils"
	"strconv"
	"strings"
)

type GasCompensationRequest struct {
	Token    string `json:"token"`
	Gas      int64  `json:"gas"`
	GasPrice int64  `json:"gas_price"`
}

type GasCompensationResponse map[string]struct {
	Amount string `json:"amount"`
	Value  string `json:"value"`
}

type PriceResponse struct {
	Coins map[string]struct {
		Price    float64 `json:"price"`
		Decimals int     `json:"decimals"`
	} `json:"coins"`
}

const (
	ETHEREUM = "ethereum:0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
)

func fetchTokenPrices(tokens ...string) (map[string]struct {
	Price    float64
	Decimals int
}, error) {
	url := fmt.Sprintf("https://coins.llama.fi/prices/current/%s,%s?searchWidth=4h", ETHEREUM, strings.Join(tokens, ","))
	response, err := utils.MakeHTTPRequest(
		url,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
		},
		nil,
		nil,
		PriceResponse{},
	)
	if err != nil {
		return nil, err
	}

	result := make(map[string]struct {
		Price    float64
		Decimals int
	})

	for _, token := range append(tokens, ETHEREUM) {
		result[token] = struct {
			Price    float64
			Decimals int
		}{
			Price:    response.Coins[token].Price,
			Decimals: response.Coins[token].Decimals,
		}
	}

	return result, nil
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	var req GasCompensationRequest

	gas, err := strconv.ParseInt(r.URL.Query().Get("gas"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid gas", http.StatusBadRequest)
		return
	}
	gasPrice, err := strconv.ParseInt(r.URL.Query().Get("gas_price"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid gas_price", http.StatusBadRequest)
		return
	}
	req = GasCompensationRequest{
		Token:    r.URL.Query().Get("token"),
		Gas:      gas,
		GasPrice: gasPrice,
	}

	prices, err := fetchTokenPrices(req.Token)
	if err != nil {
		http.Error(w, "Error fetching token prices: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ethPrice := prices[ETHEREUM].Price
	tokenPrice := prices[req.Token].Price
	tokenDecimals := prices[req.Token].Decimals

	gasPriceWei := float64(req.GasPrice) * 1e9
	ethAmountWei := float64(req.Gas) * gasPriceWei
	ethAmount := ethAmountWei / 1e18
	ethUSD := ethAmount * ethPrice
	tokenAmount := (ethUSD / tokenPrice) * math.Pow10(tokenDecimals)

	resp := GasCompensationResponse{
		ETHEREUM: {
			Amount: fmt.Sprintf("%.0f", ethAmountWei),
			Value:  fmt.Sprintf("%.2f", ethUSD),
		},
		req.Token: {
			Amount: fmt.Sprintf("%.0f", tokenAmount),
			Value:  fmt.Sprintf("%.2f", ethUSD),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
