package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"solver/types"
	"solver/utils"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	baseAPIURL          = "https://api.enso.finance/api/v1/shortcuts/route"
	disableRFQs         = "false"
	authTokenKey        = "ENSO_API_KEY"
	routeFeeBPSKey      = "ROUTE_FEE_BPS"
	routeFeeReceiverKey = "ROUTE_FEE_RECEIVER"
	httpTimeout         = 10 * time.Second
)

type RouteInputsImpl struct {
	TokenIn  string  `json:"tokenIn"`  // Address of the token to send (sell).
	TokenOut string  `json:"tokenOut"` // Address of the token to receive (buy).
	AmountIn big.Int `json:"amountIn"` // Raw amount of tokens to send (sell).
	Slippage big.Int `json:"slippage"` // Slippage tolerance when executing the swap.
}

type Transaction struct {
	Data  string `json:"data"`
	To    string `json:"to"`
	From  string `json:"from"`
	Value string `json:"value"`
}

type RouteStep struct {
	Action   string   `json:"action"`
	Protocol string   `json:"protocol"`
	TokenIn  []string `json:"tokenIn"`
	TokenOut []string `json:"tokenOut"`
}

type APIResponse struct {
	Gas         string      `json:"gas"`
	AmountOut   string      `json:"amountOut"`
	PriceImpact int         `json:"priceImpact"`
	CreatedAt   int64       `json:"createdAt"`
	Tx          Transaction `json:"tx"`
	Route       []RouteStep `json:"route"`
}

func (i *RouteInputsImpl) Validate() error {
	if !utils.IsAddress(i.TokenIn) {
		return utils.ErrInvalidAddress("tokenIn", i.TokenIn)
	}
	if !utils.IsAddress(i.TokenOut) {
		return utils.ErrInvalidAddress("tokenOut", i.TokenOut)
	}
	if !utils.IsUint(i.AmountIn.String(), 256) {
		return utils.ErrInvalidUint("amountIn", i.AmountIn.String(), 256)
	}
	if i.AmountIn.Cmp(big.NewInt(0)) >= 0 && i.AmountIn.Cmp(utils.Uint256Max) > 0 {
		return utils.ErrInvalidField("amountIn", i.AmountIn.String())
	}

	return nil
}

func (i *RouteInputsImpl) Get(provider *ethclient.Client, chainId int) (*types.ActionSchema, error) {
	return nil, utils.ErrNotImplemented("RouteInputsImpl.Get")
}

func (i *RouteInputsImpl) Post(provider *ethclient.Client, chainId int, from string) ([]*types.Transaction, error) {
	baseURL, err := url.Parse(baseAPIURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	params := url.Values{
		"chainId":     {fmt.Sprintf("%d", chainId)},
		"fromAddress": {from},
		"receiver":    {from},
		"spender":     {from},
		"amountIn":    {i.AmountIn.String()},
		"amountOut":   {"0"},
		"slippage":    {i.Slippage.String()},
		"disableRFQs": {disableRFQs},
		"tokenIn":     {i.TokenIn},
		"tokenOut":    {i.TokenOut},
	}

	routeFeeBps := os.Getenv(routeFeeBPSKey)
	routeFeeReceiver := os.Getenv(routeFeeReceiverKey)
	if routeFeeBps != "" {
		if routeFeeReceiver == "" {
			return nil, utils.ErrEnvironmentVarNotSet(routeFeeReceiverKey)
		}
		params.Set("fee", routeFeeBps)
		params.Set("feeReceiver", routeFeeReceiver)
	}

	baseURL.RawQuery = params.Encode()

	authToken := os.Getenv(authTokenKey)
	if authToken == "" {
		return nil, fmt.Errorf("authentication token not set in environment variables")
	}

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+authToken)

	client := &http.Client{Timeout: httpTimeout}
	resp, err := client.Get(baseURL.String())
	if err != nil {
		return nil, utils.ErrNetworkRequestFailed(fmt.Sprintf("failed to perform GET request: %s", err.Error()))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.ErrAPIRequestFailed(fmt.Sprintf("request failed with status code %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	value, ok := new(big.Int).SetString(apiResp.Tx.Value, 10)
	if !ok {
		return nil, fmt.Errorf("invalid value: %s", apiResp.Tx.Value)
	}

	return []*types.Transaction{{
		Transaction: apiResp.Tx.Data,
		To:          apiResp.Tx.To,
		Value:       value,
	}}, nil
}

func (i *RouteInputsImpl) GetTokenIn() string    { return i.TokenIn }
func (i *RouteInputsImpl) GetTokenOut() string   { return i.TokenOut }
func (i *RouteInputsImpl) GetAmountIn() *big.Int { return &i.AmountIn }
func (i *RouteInputsImpl) GetSlippage() *big.Int { return &i.Slippage }
