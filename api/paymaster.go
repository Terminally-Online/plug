package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

type GasCompensationRequest struct {
	Token    string `json:"token"`
	Gas      int64  `json:"gas"`
	GasPrice int64  `json:"gas_price"`
}

type GasCompensationResponse map[string]GasDetails

type GasDetails struct {
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
	NATIVE_TOKEN_ADDRESS = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
)

func fetchTokenPrices(token string) (float64, float64, int, error) {
	ethAddress := fmt.Sprintf("ethereum:%s", NATIVE_TOKEN_ADDRESS)
	url := fmt.Sprintf("https://coins.llama.fi/prices/current/%s,%s?searchWidth=4h", ethAddress, token)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, 0, err
	}

	var priceResp PriceResponse
	err = json.Unmarshal(body, &priceResp)
	if err != nil {
		return 0, 0, 0, err
	}

	ethPrice := priceResp.Coins[ethAddress].Price
	tokenPrice := priceResp.Coins[token].Price
	tokenDecimals := priceResp.Coins[token].Decimals

	return ethPrice, tokenPrice, tokenDecimals, nil
}

func GetPayment(w http.ResponseWriter, r *http.Request) {
	var req GasCompensationRequest

	if r.Method == "GET" {
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
	} else {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
	}

	ethPrice, tokenPrice, tokenDecimals, err := fetchTokenPrices(req.Token)
	if err != nil {
		http.Error(w, "Error fetching token prices: "+err.Error(), http.StatusInternalServerError)
		return
	}

	gasPriceWei := float64(req.GasPrice) * 1e9
	ethAmountWei := float64(req.Gas) * gasPriceWei
	ethAmount := ethAmountWei / 1e18
	ethUSD := ethAmount * ethPrice
	tokenAmount := (ethUSD / tokenPrice) * math.Pow10(tokenDecimals)

	resp := GasCompensationResponse{
		fmt.Sprintf("ethereum:%s", NATIVE_TOKEN_ADDRESS): {
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
