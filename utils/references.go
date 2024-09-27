package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	retries          = 3
	escapeResults    = []string{"Contract source code not verified"}
	terminateResults = []string{"Invalid API Key"}
	sleepResults     = []string{"Rate limit"}
)

type Network struct {
	Explorer   string
	References map[string]map[string]string
}

type AdditionalSource struct {
	Filename   string
	SourceCode string
}

var (
	Referral = "0x62180042606624f02d8a130da8a3171e9b33894d"

	Mainnet = &Network{
		Explorer: "https://api.etherscan.io/api",
		References: map[string]map[string]string{
			"aave_v2": {
				"pool": "0x02d84abd89ee9db409572f19b6e1596c301f3c81",
			},
			"aave_v3": {
				"pool": "0x34339f94350ec5274ea44d0c37dae9e968c44081",
			},
		},
	}
)

func GenerateReference(explorer string, folderName string, contractName string, address string, retries int) error {
	url := fmt.Sprintf("%s?module=contract&action=getsourcecode&address=%s&apiKey=%s", explorer, address, os.Getenv("ETHERSCAN_API_KEY"))
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var response struct {
		Message string            `json:"message"`
		Result  []json.RawMessage `json:"result"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	if response.Message == "NOTOK" {
		result := string(response.Result[0])
		shouldRetry := true
		for _, termination := range terminateResults {
			if strings.Contains(strings.ToLower(result), strings.ToLower(termination)) {
				shouldRetry = false
				break
			}
		}
		shouldRetry = shouldRetry && retries > 0

		fmt.Printf("Failed to get source for %s/%s: %s.\n", folderName, contractName, result)
		if !shouldRetry {
			return ErrExplorerFailed(result)
		}
		fmt.Printf("Retrying %d more times before aborting.\n", retries)

		shouldSleep := false
		for _, sleepResult := range sleepResults {
			if strings.Contains(strings.ToLower(result), strings.ToLower(sleepResult)) {
				shouldSleep = true
				break
			}
		}
		if shouldSleep {
			fmt.Println("Sleeping for 5 seconds before retrying.")
			time.Sleep(5 * time.Second)
		}
		return GenerateReference(explorer, folderName, contractName, address, retries-1)
	}

	var explorerResponse struct {
		ABI          string `json:"ABI"`
		ContractName string `json:"ContractName"`
	}
	err = json.Unmarshal(response.Result[0], &explorerResponse)
	if err != nil {
		return err
	}

	for _, escapeResult := range escapeResults {
		if strings.Contains(explorerResponse.ABI, escapeResult) {
			fmt.Printf("Source code not verified for %s/%s.\n", folderName, contractName)
			return nil
		}
	}

	var abiJSON interface{}
	err = json.Unmarshal([]byte(explorerResponse.ABI), &abiJSON)
	if err != nil {
		return fmt.Errorf("failed to parse ABI JSON: %v", err)
	}

	formattedABI, err := json.MarshalIndent(abiJSON, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format ABI JSON: %v", err)
	}

	fileName := fmt.Sprintf("%s.json", contractName)
	referencePath := filepath.Join(".", "abis", folderName, fileName)

	if err := os.MkdirAll(filepath.Dir(referencePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	err = os.WriteFile(referencePath, formattedABI, 0644)
	if err != nil {
		return fmt.Errorf("failed to write ABI file: %v", err)
	}

	fmt.Printf("Generated ABI file for %s/%s at %s\n", folderName, contractName, referencePath)
	return nil
}

func GenerateReferences() error {
	for folderName, contracts := range Mainnet.References {
		for contractName, address := range contracts {
			if err := GenerateReference(Mainnet.Explorer, folderName, contractName, address, retries); err != nil {
				return err
			}
		}
	}
	return nil
}
