package sentio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	SentioBaseURL  = "https://app.sentio.xyz/api/v1/solidity"
	DefaultTimeout = 30 * time.Second
)

type Client struct {
	APIKey       string
	ProjectOwner string
	ProjectSlug  string
	BaseURL      string
	HttpClient   *http.Client
}

func NewSentioClient(apiKey string) *Client {
	// TODO MASON: remove these placeholders for a production env one
	if apiKey == "" {
		apiKey = "6hyo9WIlX1As5wG6NWLfMTu8H6Y5BS9Wr"
	}
	return &Client{
		ProjectOwner: "mason",
		ProjectSlug:  "testing",
		APIKey:       apiKey,
		BaseURL:      SentioBaseURL,
		HttpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func (c *Client) SimulateTransaction(chainID uint64, from, to common.Address, data string, gas string, gasPrice string, blockNumber string, value big.Int) (string, error) {
	// https://app.sentio.xyz/api/v1/solidity/{owner}/{slug}/{chainId}/simulation
	chainIDStr := fmt.Sprintf("%d", chainID)
	url := fmt.Sprintf("%s/%s/%s/%s/simulation", c.BaseURL, c.ProjectOwner, c.ProjectSlug, chainIDStr)

	reqBody := map[string]interface{}{
		"simulation": map[string]interface{}{
			"networkId": "8453",
			"chainId":   chainIDStr,
			"chainSpec": map[string]interface{}{
				"chainId": chainIDStr,
			},
			"to":               to,
			"from":             from,
			"input":            data,
			"blockNumber":      "latest",
			"gas":              gas,
			"gasPrice":         gasPrice,
			"value":            hexutil.EncodeBig(&value),
			"transactionIndex": "0", // this should not matter unless we're hitting existing blocks I believe
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal simulation request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create simulation request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("simulation request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read simulation response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("simulation failed with status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Printf("Simulation response: %s\n", string(body))

	var simResponse SimulationResponse
	if err := json.Unmarshal(body, &simResponse); err != nil {
		return "", fmt.Errorf("failed to parse simulation response: %v", err)
	}
	fmt.Printf("Simulation response parsed: %+v\n", simResponse)

	// Check for errors
	if simResponse.Result.TransactionReceipt.Error != "" {
		return "", fmt.Errorf("simulation failed: %s", simResponse.Result.TransactionReceipt.Error)
	}

	return simResponse.Simulation.Id, nil
}
