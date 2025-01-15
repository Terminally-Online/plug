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
			"weth": {
				"address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			},
			"aave_v3": {
				"pool":                  "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
				"ui_pool_data_provider": "0x194324C9Af7f56E22F1614dD82E18621cb9238E7",
			},
			"yearn_v3": {
				"registry": "0xff31A1B020c868F6eA3f61Eb953344920EeCA3af",
				"pool":     "0x1ab62413e0cf2eBEb73da7D40C70E7202ae14467",
				"router":   "0x1112dbCF805682e828606f74AB717abf4b4FD8DE",
				"gauge":    "0x7Fd8Af959B54A677a1D8F92265Bd0714274C56a3",
			},
			"multicall": {
				"primary": "0xcA11bde05977b3631167028862bE2a173976CA11",
			},
			"ens": {
				"registrar_controller": "0x253553366Da8546fC250F225fe3d25d0C782303b",
				"base_registrar":       "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85",
			},
			"nouns": {
				"auction_house": "0x830BD73E4184ceF73443C15111a1DF14e495C706",
				"token":         "0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03",
				"art":           "0x6544bC8A0dE6ECe429F14840BA74611cA5098A92",
			},
			"morpho": {
				"router":      "0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb",
				"distributor": "0x330eefa8a787552DC5cAd3C3cA644844B1E61Ddb",
				"bundler":     "0x4095F064B8d3c3548A3bebfd0Bbfd04750E30077",
				"vault":       "0xfbDEE8670b273E12b019210426E70091464b02Ab",
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
		ABI            string `json:"ABI"`
		ContractName   string `json:"ContractName"`
		Implementation string `json:"Implementation"`
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

	// If there is an implement address, we have found a proxy contract and need to
	// pull in the ABI from the implementation contract.
	if explorerResponse.Implementation != "" && explorerResponse.Implementation != address {
		return GenerateReference(explorer, folderName, contractName, explorerResponse.Implementation, retries)
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
	time.Sleep(time.Millisecond * 500)
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
