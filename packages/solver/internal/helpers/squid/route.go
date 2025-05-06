package squid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"solver/bindings/erc_20"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type SquidRouteRequest struct {
	FromAddress common.Address `json:"fromAddress"`
	FromChain   uint64         `json:"fromChain"`
	FromToken   common.Address `json:"fromToken"`
	FromAmount  *big.Int       `json:"fromAmount"`
	ToChain     uint64         `json:"toChain"`
	ToToken     common.Address `json:"toToken"`
	ToAddress   common.Address `json:"toAddress"`
}

type SquidAction struct {
	Type                 string     `json:"type"`
	ChainType            string     `json:"chainType"`
	Data                 ActionData `json:"data"`
	FromChain            string     `json:"fromChain"`
	ToChain              string     `json:"toChain"`
	FromToken            Token      `json:"fromToken"`
	ToToken              Token      `json:"toToken"`
	FromAmount           string     `json:"fromAmount"`
	ToAmount             string     `json:"toAmount"`
	ToAmountMin          string     `json:"toAmountMin"`
	ExchangeRate         string     `json:"exchangeRate"`
	PriceImpact          string     `json:"priceImpact"`
	Stage                int        `json:"stage"`
	Provider             string     `json:"provider"`
	LogoURI              string     `json:"logoURI"`
	Description          string     `json:"description"`
	EstimatedDuration    int        `json:"estimatedDuration,omitempty"`
	AggregatedVolatility float64    `json:"aggregatedVolatility"`
}

type ActionData struct {
	Address           string   `json:"address,omitempty"`
	ChainId           string   `json:"chainId,omitempty"`
	CoinAddresses     []string `json:"coinAddresses,omitempty"`
	Dex               string   `json:"dex,omitempty"`
	Enabled           bool     `json:"enabled"`
	Path              []string `json:"path,omitempty"`
	Slippage          float64  `json:"slippage,omitempty"`
	AggregateSlippage float64  `json:"aggregateSlippage,omitempty"`
	Target            string   `json:"target,omitempty"`
	Name              string   `json:"name,omitempty"`
	Provider          string   `json:"provider,omitempty"`
	Type              string   `json:"type,omitempty"`
	LogoURI           string   `json:"logoURI,omitempty"`
}

type Token struct {
	ID                  string   `json:"id"`
	Symbol              string   `json:"symbol"`
	Address             string   `json:"address"`
	ChainId             string   `json:"chainId"`
	Name                string   `json:"name"`
	Decimals            int      `json:"decimals"`
	CoingeckoId         string   `json:"coingeckoId"`
	Type                string   `json:"type"`
	LogoURI             string   `json:"logoURI"`
	AxelarNetworkSymbol string   `json:"axelarNetworkSymbol"`
	SubGraphOnly        bool     `json:"subGraphOnly"`
	SubGraphIds         []string `json:"subGraphIds"`
	Enabled             bool     `json:"enabled"`
	Active              bool     `json:"active"`
	Visible             bool     `json:"visible"`
	UsdPrice            float64  `json:"usdPrice"`
	InterchainTokenId   *string  `json:"interchainTokenId,omitempty"`
	Volatility          int      `json:"volatility,omitempty"`
}

type FeeCost struct {
	Amount        string   `json:"amount"`
	AmountUsd     string   `json:"amountUsd"`
	Description   string   `json:"description"`
	GasLimit      string   `json:"gasLimit,omitempty"`
	GasMultiplier float64  `json:"gasMultiplier,omitempty"`
	Name          string   `json:"name"`
	Token         Token    `json:"token"`
	LogoURI       string   `json:"logoURI"`
	Data          *FeeData `json:"data,omitempty"`
}

type FeeData struct {
	AxelarFeeData  *AxelarFeeData  `json:"axelarFeeData,omitempty"`
	ToChainFeeData *ToChainFeeData `json:"toChainFeeData,omitempty"`
}

type AxelarFeeData struct {
	BaseFee           float64 `json:"baseFee"`
	ExpressFee        float64 `json:"expressFee"`
	ExecuteMultiplier float64 `json:"executeMultiplier"`
	ExpressMultiplier float64 `json:"expressMultiplier"`
	ExpressSupported  bool    `json:"expressSupported"`
}

type ToChainFeeData struct {
	LastBaseFeePerGas    string `json:"lastBaseFeePerGas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	GasPrice             string `json:"gasPrice"`
}

type GasCost struct {
	Type      string `json:"type"`
	Token     Token  `json:"token"`
	Amount    string `json:"amount"`
	GasLimit  string `json:"gasLimit"`
	AmountUsd string `json:"amountUsd"`
}

type RouteEstimate struct {
	Actions                []SquidAction `json:"actions"`
	FromAmount             string        `json:"fromAmount"`
	ToAmount               string        `json:"toAmount"`
	ToAmountMin            string        `json:"toAmountMin"`
	ExchangeRate           string        `json:"exchangeRate"`
	AggregatePriceImpact   string        `json:"aggregatePriceImpact"`
	FromAmountUSD          string        `json:"fromAmountUSD"`
	ToAmountUSD            string        `json:"toAmountUSD"`
	ToAmountMinUSD         string        `json:"toAmountMinUSD"`
	AggregateSlippage      float64       `json:"aggregateSlippage"`
	FromToken              Token         `json:"fromToken"`
	ToToken                Token         `json:"toToken"`
	IsBoostSupported       bool          `json:"isBoostSupported"`
	FeeCosts               []FeeCost     `json:"feeCosts"`
	GasCosts               []GasCost     `json:"gasCosts"`
	EstimatedRouteDuration int           `json:"estimatedRouteDuration"`
}

type TransactionRequest struct {
	Type                 string `json:"type"`
	Target               string `json:"target"`
	Data                 string `json:"data"`
	Value                string `json:"value"`
	GasLimit             string `json:"gasLimit"`
	LastBaseFeePerGas    string `json:"lastBaseFeePerGas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	GasPrice             string `json:"gasPrice"`
	RequestId            string `json:"requestId"`
	Expiry               string `json:"expiry"`
	ExpiryOffset         string `json:"expiryOffset"`
}

type Route struct {
	Estimate           RouteEstimate      `json:"estimate"`
	TransactionRequest TransactionRequest `json:"transactionRequest"`
	Params             SquidRouteRequest  `json:"params"`
}

type SquidRouteResponse struct {
	Route     Route  `json:"route"`
	RequestId string `json:"requestId"`
}

func GetRoute(request SquidRouteRequest) (*SquidRouteResponse, error) {
	url := "https://v2.api.squidrouter.com/v2/route"

	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	headers := map[string]string{
		"Content-Type":    "application/json",
		"x-integrator-id": os.Getenv("SQUID_INTEGRATOR_ID"),
	}

	response, err := utils.MakeHTTPRequest(
		url,
		"POST",
		headers,
		nil,
		bytes.NewBuffer(requestBody),
		SquidRouteResponse{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch routes: %w", err)
	}

	return &response, nil
}

func GetPlugs(request SquidRouteRequest) ([]signature.Plug, error) {
	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}

	route, err := GetRoute(request)
	if err != nil {
		return nil, err
	}
	value, ok := new(big.Int).SetString(route.Route.TransactionRequest.Value, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse transaction value: %s", route.Route.TransactionRequest.Value)
	}

	transactions := []signature.Plug{{
		To:    common.HexToAddress(route.Route.TransactionRequest.Target),
		Data:  []byte(route.Route.TransactionRequest.Data),
		Value: value,
	}}

	if request.FromToken == utils.NativeTokenAddress {
		return transactions, nil
	}

	approveCalldata, err := erc20Abi.Pack("approve",
		common.HexToAddress(route.Route.TransactionRequest.Target),
		request.FromAmount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return append([]signature.Plug{{
		To:   request.FromToken,
		Data: approveCalldata,
	}}, transactions...), nil
}
