package squid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"solver/internal/utils"
)

type SquidRouteRequest struct {
	FromAddress string `json:"fromAddress"`
	FromChain   string `json:"fromChain"`
	FromToken   string `json:"fromToken"`
	FromAmount  string `json:"fromAmount"`
	ToChain     string `json:"toChain"`
	ToToken     string `json:"toToken"`
	ToAddress   string `json:"toAddress"`
}

type SquidRouteResponse struct {
	Routes []SquidRoute `json:"routes"`
}

type SquidRoute struct {
	Route string `json:"route"`
}

func GetRoutes(request SquidRouteRequest) (*SquidRouteResponse, error) {
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
