package references

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"solver/internal/utils"
	"strings"
	"time"
)

var (
	retries          = 3
	escapeResults    = []string{"Contract source code not verified"}
	terminateResults = []string{"Invalid API Key"}
	sleepResults     = []string{"Rate limit"}

	Referral = "0x62180042606624f02d8a130da8a3171e9b33894d"

	Mainnet = &Network{
		ChainIds: []uint64{1, 31337},
		Explorer: "https://api.etherscan.io/api",
		References: map[string]map[string]string{
			"aave_v3": {
				"pool":                     "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
				"ui_pool_data_provider":    "0x3F78BBD206e4D3c504Eb854232EdA7e47E9Fd8FC",
				"ui_pool_address_provider": "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e",
			},
			"ens": {
				"registrar_controller": "0x253553366Da8546fC250F225fe3d25d0C782303b",
				"base_registrar":       "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85",
			},
			"morpho": {
				"router":      "0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb",
				"distributor": "0x330eefa8a787552DC5cAd3C3cA644844B1E61Ddb",
				"bundler":     "0x6566194141eefa99Af43Bb5Aa71460Ca2Dc90245",
				"vault":       "0xfbDEE8670b273E12b019210426E70091464b02Ab",
			},
			"multicall": {
				"primary": "0xcA11bde05977b3631167028862bE2a173976CA11",
			},
			"nouns": {
				"auction_house": "0x830BD73E4184ceF73443C15111a1DF14e495C706",
				"token":         "0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03",
				"art":           "0x6544bC8A0dE6ECe429F14840BA74611cA5098A92",
			},
			"weth": {
				"address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			},
			"yearn_v3": {
				"registry": "0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038",
				"pool":     "0x1ab62413e0cf2eBEb73da7D40C70E7202ae14467",
				"router":   "0x1112dbCF805682e828606f74AB717abf4b4FD8DE",
				"gauge":    "0x7Fd8Af959B54A677a1D8F92265Bd0714274C56a3",
			},
		},
	}

	Base = &Network{
		ChainIds: []uint64{8453},
		Explorer: "https://api.basescan.org/api",
		References: map[string]map[string]string{
			"aave_v3": {
				"pool":                     "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5",
				"ui_pool_data_provider":    "0x68100bD5345eA474D93577127C11F39FF8463e93",
				"ui_pool_address_provider": "0xe20fCBdBfFC4Dd138cE8b2E6FBb6CB49777ad64D",
			},
			"euler": {
				"balanceTracker": "0x029fDEe85BEdB0553D6fdc538546586641DD7438",
				"eVaultFactory": "0x7F321498A801A191a93C840750ed637149dDf8D0",
				"eVaultImplementation": "0x30a9A9654804F1e5b3291a86E83EdeD7cF281618",
				"eulerEarnFactory": "0x72bbDB652F2AEC9056115644EfCcDd1986F51f15",
				"eulerEarnImplementation": "0x6104c0F2a7750F1b143DAB49752e19DA43dec34A",
				"evc": "0x5301c7dD20bD945D2013b48ed0DEE3A284ca8989",
			},
			"morpho": {
				"router":      "0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb",
				"distributor": "0x5400dbb270c956e8985184335a1c62aca6ce1333",
				"bundler":     "0x6BFd8137e702540E7A42B74178A4a49Ba43920C4",
				"vault":       "0xF540D790413FCFAedAC93518Ae99EdDacE82cb78",
			},
			"multicall": {
				"primary": "0xcA11bde05977b3631167028862bE2a173976CA11",
			},
			"weth": {
				"address": "0x4200000000000000000000000000000000000006",
			},
			"yearn_v3": {
				"registry": "0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038",
				"pool":     "0x1ab62413e0cf2eBEb73da7D40C70E7202ae14467",
				"router":   "0x1112dbCF805682e828606f74AB717abf4b4FD8DE",
				"gauge":    "", // Not available on Base?
			},
		},
	}

	Networks = map[uint64]*Network{1: Mainnet, 31337: Mainnet, 8453: Base}
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
			return utils.ErrExplorer(result)
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
	processed := make(map[string]bool)

	for _, network := range Networks {
		for folderName, contracts := range network.References {
			for contractName, address := range contracts {
				key := fmt.Sprintf("%s:%s", folderName, contractName)

				if processed[key] {
					fmt.Printf("Skipping %s/%s on %s - already processed\n",
						folderName, contractName, network.Explorer)
					continue
				}

				if err := GenerateReference(network.Explorer, folderName, contractName, address, retries); err != nil {
					return err
				}

				processed[key] = true
			}
		}
	}

	return nil
}
