package sentio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	SentioBaseURL  = "https://app.sentio.xyz/api/v1/solidity"
	DefaultTimeout = 30 * time.Second
)

type Client struct {
	APIKey     string
	BaseURL    string
	HttpClient *http.Client
}

func NewClient(apiKey string) *Client {
	// TODO MASON: remove this placeholder test key for a production env one
	apiKey = "6hyo9WIlX1As5wG6NWLfMTu8H6Y5BS9Wr"

	if apiKey == "" {
		apiKey = os.Getenv("SENTIO_API_KEY")
	}

	return &Client{
		APIKey:  apiKey,
		BaseURL: SentioBaseURL,
		HttpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func (c *Client) SimulateTransaction(chainID uint64, from, to common.Address, data []byte, value *big.Int) (string, error) {
	url := fmt.Sprintf("%s/simulate", c.BaseURL)

	// Prepare the request body
	reqBody := map[string]interface{}{
		"networkId": chainID,
		"tx": map[string]interface{}{
			"from": from.Hex(),
			"to":   to.Hex(),
			"data": hexutil.Bytes(data).String(),
		},
	}

	// Add value if provided
	if value != nil && value.Cmp(big.NewInt(0)) != 0 {
		reqBody["tx"].(map[string]interface{})["value"] = hexutil.EncodeBig(value)
	}

	// Marshal request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal simulation request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create simulation request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.APIKey)

	// Execute request
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("simulation request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read simulation response: %v", err)
	}

	// Check for non-200 status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("simulation failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse response
	var simResponse struct {
		SimulationId string `json:"simulationId"`
		Success      bool   `json:"success"`
		Error        string `json:"error"`
	}

	if err := json.Unmarshal(body, &simResponse); err != nil {
		return "", fmt.Errorf("failed to parse simulation response: %v", err)
	}

	// Check for errors
	if !simResponse.Success {
		return "", fmt.Errorf("simulation failed: %s", simResponse.Error)
	}

	return simResponse.SimulationId, nil
}

func (c *Client) GetCallTrace(chainID uint64, simulationID string) (*CallTrace, error) {
	url := fmt.Sprintf("%s/call_trace", c.BaseURL)

	reqBody := map[string]interface{}{
		"networkId": chainID,
		"txId": map[string]interface{}{
			"simulationId": simulationID,
		},
		"withInternalCalls": true,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal call trace request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create call trace request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call trace request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read call trace response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("call trace request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var callTrace CallTrace
	if err := json.Unmarshal(body, &callTrace); err != nil {
		return nil, fmt.Errorf("failed to parse call trace response: %v", err)
	}

	return &callTrace, nil
}

func (c *Client) GetStateDiff(chainID uint64, simulationID string) (*StateDiff, error) {
	url := fmt.Sprintf("%s/state_diff", c.BaseURL)

	reqBody := map[string]interface{}{
		"networkId": chainID,
		"txId": map[string]interface{}{
			"simulationId": simulationID,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal state diff request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create state diff request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("state diff request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read state diff response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("state diff request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var stateDiff StateDiff
	if err := json.Unmarshal(body, &stateDiff); err != nil {
		return nil, fmt.Errorf("failed to parse state diff response: %v", err)
	}

	return &stateDiff, nil
}

// TODO MASON: wtf is going on here
func ExtractErrorFromCallTrace(trace *CallTrace) string {
	if trace == nil || trace.Error == "" {
		return ""
	}

	errorMsg := trace.Error

	if strings.Contains(errorMsg, "execution reverted") {
		if trace.Output != "" && strings.HasPrefix(trace.Output, "0x08c379a0") {
			if len(trace.Output) >= 138 { // 2 + 8 + 64 + 64 (0x + selector + offset + length)
				// Extract the string length (convert hex to decimal)
				lenHex := trace.Output[74:138]
				lenVal, ok := new(big.Int).SetString(lenHex, 16)
				if ok {
					msgLen := lenVal.Int64() * 2 // Each byte is 2 hex chars
					if len(trace.Output) >= 138+int(msgLen) {
						msgHex := trace.Output[138 : 138+int(msgLen)]
						// Convert hex to ASCII
						bytes, err := hexutil.Decode("0x" + msgHex)
						if err == nil {
							return string(bytes)
						}
					}
				}
			}
		}

		// Check recursive calls for revert reasons
		if trace.Calls != nil && len(trace.Calls) > 0 {
			for _, call := range trace.Calls {
				subCall := CallTrace(call)
				if reason := ExtractErrorFromCallTrace(&subCall); reason != "" {
					return reason
				}
			}
		}
	}

	// Common error patterns
	errorPatterns := map[string]string{
		"execution reverted":                 "Transaction reverted",
		"out of gas":                         "Transaction ran out of gas",
		"invalid opcode":                     "Contract execution error: invalid opcode",
		"invalid jump destination":           "Contract execution error: invalid jump destination",
		"stack underflow":                    "Contract execution error: stack underflow",
		"stack overflow":                     "Contract execution error: stack overflow",
		"invalid instruction":                "Contract execution error: invalid instruction",
		"insufficient balance":               "Insufficient balance for transfer",
		"exceeded the maximum amount of gas": "Exceeded gas limit",
	}

	// Try to match against known error patterns
	for pattern, humanMsg := range errorPatterns {
		if strings.Contains(errorMsg, pattern) {
			return humanMsg
		}
	}

	return errorMsg
}

// func SummarizeStateChanges(diff *StateDiff) []StateChangeSummary {
// 	if diff == nil || len(diff.StateChanges) == 0 {
// 		return nil
// 	}

// 	var summaries []StateChangeSummary

// 	for _, change := range diff.StateChanges {
// 		address := change.Address

// 		// Balance changes
// 		if change.BalanceChanges != nil && len(change.BalanceChanges) > 0 {
// 			for _, bal := range change.BalanceChanges {
// 				beforeBal, _ := new(big.Int).SetString(bal.PrevBalance, 10)
// 				afterBal, _ := new(big.Int).SetString(bal.Balance, 10)

// 				description := "Balance changed"
// 				valueChange := fmt.Sprintf("%s -> %s", bal.PrevBalance, bal.Balance)

// 				// If we can parse both values, provide more descriptive change
// 				if beforeBal != nil && afterBal != nil {
// 					delta := new(big.Int).Sub(afterBal, beforeBal)

// 					if delta.Sign() > 0 {
// 						description = "Balance increased"
// 						valueChange = fmt.Sprintf("+%s wei", delta.String())
// 					} else {
// 						description = "Balance decreased"
// 						valueChange = fmt.Sprintf("-%s wei", new(big.Int).Abs(delta).String())
// 					}

// 					// Check if this is a significant ETH transfer
// 					oneEth := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
// 					thousandthEth := new(big.Int).Div(oneEth, big.NewInt(1000)) // 0.001 ETH

// 					if delta.Abs(delta).Cmp(thousandthEth) > 0 {
// 						ethValue := utils.WeiToEth(delta.Abs(delta))
// 						sign := ""
// 						if delta.Sign() < 0 {
// 							sign = "-"
// 						} else {
// 							sign = "+"
// 						}
// 						valueChange = fmt.Sprintf("%s%s ETH", sign, ethValue)
// 					}
// 				}

// 				summaries = append(summaries, StateChangeSummary{
// 					Address:     address,
// 					Description: description,
// 					Type:        "balance",
// 					ValueChange: valueChange,
// 				})
// 			}
// 		}

// 		// Storage changes
// 		if change.StorageChanges != nil && len(change.StorageChanges) > 0 {
// 			description := fmt.Sprintf("Storage updated (%d slot changes)", len(change.StorageChanges))

// 			// Try to detect specific storage patterns
// 			if len(change.StorageChanges) == 1 {
// 				description = "Single storage slot updated"
// 			} else if len(change.StorageChanges) > 10 {
// 				description = fmt.Sprintf("Major storage update (%d slots)", len(change.StorageChanges))
// 			}

// 			summaries = append(summaries, StateChangeSummary{
// 				Address:     address,
// 				Description: description,
// 				Type:        "storage",
// 				ValueChange: fmt.Sprintf("%d storage slots modified", len(change.StorageChanges)),
// 			})
// 		}

// 		// Nonce changes
// 		if change.NonceChanges != nil && len(change.NonceChanges) > 0 {
// 			for _, nonce := range change.NonceChanges {
// 				description := "Nonce changed"
// 				if nonce.Nonce > nonce.PrevNonce {
// 					description = fmt.Sprintf("Nonce incremented by %d", nonce.Nonce-nonce.PrevNonce)
// 				}

// 				summaries = append(summaries, StateChangeSummary{
// 					Address:     address,
// 					Description: description,
// 					Type:        "nonce",
// 					ValueChange: fmt.Sprintf("%d -> %d", nonce.PrevNonce, nonce.Nonce),
// 				})
// 			}
// 		}

// 		// Code changes
// 		if change.CodeChanges != nil && len(change.CodeChanges) > 0 {
// 			for _, code := range change.CodeChanges {
// 				description := "Contract code changed"

// 				// Check for contract creation vs modification
// 				if code.PrevCode == "" && code.Code != "" {
// 					description = "New contract deployed"
// 					codeBytes, _ := hexutil.Decode(code.Code)
// 					summaries = append(summaries, StateChangeSummary{
// 						Address:     address,
// 						Description: description,
// 						Type:        "code",
// 						ValueChange: fmt.Sprintf("Code size: %d bytes", len(codeBytes)),
// 					})
// 				} else if code.PrevCode != "" && code.Code == "" {
// 					description = "Contract self-destructed"
// 					summaries = append(summaries, StateChangeSummary{
// 						Address:     address,
// 						Description: description,
// 						Type:        "code",
// 					})
// 				} else {
// 					description = "Contract code updated"
// 					prevCodeBytes, _ := hexutil.Decode(code.PrevCode)
// 					codeBytes, _ := hexutil.Decode(code.Code)
// 					summaries = append(summaries, StateChangeSummary{
// 						Address:     address,
// 						Description: description,
// 						Type:        "code",
// 						ValueChange: fmt.Sprintf("Code size changed: %d -> %d bytes",
// 							len(prevCodeBytes), len(codeBytes)),
// 					})
// 				}
// 			}
// 		}
// 	}

// 	return summaries
// }
