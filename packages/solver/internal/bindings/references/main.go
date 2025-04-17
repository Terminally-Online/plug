package references

import (
	"encoding/json"
	"fmt"
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

	MultichainEtherscan = "https://api.etherscan.io/v2/api?chainid=%d"

	Referral = "0xE09C3f7a144D3b1Bef42499486cA5C36d20ec5d1"

	Multicall = map[string]string{
		"primary": "0xcA11bde05977b3631167028862bE2a173976CA11",
	}
	// NOTE: We are not actively verifying our contracts when they are deployed
	//       because we are still in development so to set the ABI you just toss it
	//       into the abi directory above.
	Plug = map[string]string{
		"router":   "0x0000000021EAfaa2A0ADeec53B7E25F662920212",
		"factory":  "0x0000000030c2d2825F563E2F7b78943B0Ea9D145",
		"socket":   "0x0000000011A65597897563205669f9c46dEEE244",
		"assert":   "0x0000000005d8F29675fC43df88588bD0D5c0DeC5",
		"boolean":  "0x00000000410986831F18E06d908bE25e5Fb949A0",
		"database": "0x0000000006d777c8390a5E84Ecb88A6556A1d3B5",
		"math":     "0x000000000c0352950e3aa28973824f4d01ccec4f",
		"evm":      "0x000000005DAf9C20a195d6f895c96A226e52Ed98",
	}

	Mainnet = &Network{
		Name: "Ethereum Mainnet",
		Icon: struct {
			Default string `json:"default"`
		}{
			Default: "https://cdn.onplug.io/blockchain/ethereum.png",
		},
		ChainIds: []uint64{1, 31337},
		References: map[string]map[string]string{
			"aave_v3": {
				"pool":                     "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
				"ui_pool_data_provider":    "0x3F78BBD206e4D3c504Eb854232EdA7e47E9Fd8FC",
				"ui_pool_address_provider": "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e",
			},
			"euler": {
				"evault_implementation":     "0x8Ff1C814719096b61aBf00Bb46EAd0c9A529Dd7D",
				"euler_earn_implementation": "0xBa42141648dFD74388f3541C1d80fa9387043Da9",
				"evc":                       "0x0C9a3dd6b8F28529d72d7f9cE918D493519EE383",
				"governed_perspective":      "0xC0121817FF224a018840e4D15a864747d36e6Eb2",
				"account_lens":              "0x88062031730136292902Cd4f6f07fDB224E60E9F",
				"utils_lens":                "0x1Ad6eDF948B037a9364607f9e3c1e640166Ee539",
				"vault_lens":                "0x75AAf54F12784935128306BEe2520de55890a29A",
			},
			"morpho": {
				"router":      "0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb",
				"distributor": "0x330eefa8a787552DC5cAd3C3cA644844B1E61Ddb",
				"bundler":     "0x6566194141eefa99Af43Bb5Aa71460Ca2Dc90245",
				"vault":       "0xfbDEE8670b273E12b019210426E70091464b02Ab",
			},
			"multicall": Multicall,
			"nouns": {
				"auction_house": "0x830BD73E4184ceF73443C15111a1DF14e495C706",
				"token":         "0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03",
				"art":           "0x6544bC8A0dE6ECe429F14840BA74611cA5098A92",
			},
			"plug": Plug,
			"weth": {
				"address": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
			},
			"yearn_v3": {
				"registry": "0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038",
				"pool":     "0x1ab62413e0cf2eBEb73da7D40C70E7202ae14467",
				"router":   "0x1112dbCF805682e828606f74AB717abf4b4FD8DE",
			},
		},
	}

	Base = &Network{
		Name: "base",
		Icon: struct {
			Default string `json:"default"`
		}{
			Default: "https://cdn.onplug.io/blockchain/base.png",
		},
		ChainIds: []uint64{8453},
		References: map[string]map[string]string{
			"aave_v3": {
				"pool":                     "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5",
				"ui_pool_data_provider":    "0x68100bD5345eA474D93577127C11F39FF8463e93",
				"ui_pool_address_provider": "0xe20fCBdBfFC4Dd138cE8b2E6FBb6CB49777ad64D",
			},
			"basepaint": {
				"referral": "0xaff1A9E200000061fC3283455d8B0C7e3e728161",
			},
			"euler": {
				"evault_implementation":     "0x30a9A9654804F1e5b3291a86E83EdeD7cF281618",
				"euler_earn_implementation": "0x6104c0F2a7750F1b143DAB49752e19DA43dec34A",
				"evc":                       "0x5301c7dD20bD945D2013b48ed0DEE3A284ca8989",
				"governed_perspective":      "0xafC8545c49DF2c8216305922D9753Bf60bf8c14A",
				"account_lens":              "0x40c1DbD5855bFbCDd3844C4327777FD1c5E039eb",
				"utils_lens":                "0x6E1033296eDbD7Ef23544E2A4Fa6E78e77D294E1",
				"vault_lens":                "0x26c577bF95d3c4AD8155834a0149D6BB76F2D090",
			},
			"morpho": {
				"router":      "0xBBBBBbbBBb9cC5e90e3b3Af64bdAF62C37EEFFCb",
				"distributor": "0x5400dbb270c956e8985184335a1c62aca6ce1333",
				"bundler":     "0x6BFd8137e702540E7A42B74178A4a49Ba43920C4",
				"vault":       "0xF540D790413FCFAedAC93518Ae99EdDacE82cb78",
			},
			"multicall": Multicall,
			"plug":      Plug,
			"weth": {
				"address": "0x4200000000000000000000000000000000000006",
			},
			"yearn_v3": {
				"registry": "0xd40ecF29e001c76Dcc4cC0D9cd50520CE845B038",
				"pool":     "0x1ab62413e0cf2eBEb73da7D40C70E7202ae14467",
				"router":   "0x1112dbCF805682e828606f74AB717abf4b4FD8DE",
			},
		},
	}

	Networks = map[uint64]*Network{1: Mainnet, 8453: Base}
)

func GenerateReference(explorer string, folderName string, contractName string, address string, retries int) error {
	fileName := fmt.Sprintf("%s.json", contractName)
	referencePath := filepath.Join(".", "abis", folderName, fileName)

	if _, err := os.Stat(referencePath); err == nil {
		fmt.Printf("ABI file already exists for %s/%s at %s, skipping generation\n", folderName, contractName, referencePath)
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(referencePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	response, err := utils.MakeHTTPRequest(
		fmt.Sprintf("%s&module=contract&action=getsourcecode&address=%s&apiKey=%s", explorer, address, os.Getenv("ETHERSCAN_API_KEY")),
		"GET",
		map[string]string{
			"accept":        "application/json",
			"authorization": fmt.Sprintf("Basic %v", os.Getenv("ZERION_KEY")),
		},
		nil,
		nil,
		struct {
			Message string            `json:"message"`
			Result  []json.RawMessage `json:"result"`
		}{},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Getting source for %s/%s\n", folderName, contractName)

	if response.Message == "NOTOK" {
		fmt.Printf("Failed to get source for %s/%s: %s.\n", folderName, contractName, response.Result[0])
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

	var abiJSON any
	err = json.Unmarshal([]byte(explorerResponse.ABI), &abiJSON)
	if err != nil {
		return fmt.Errorf("failed to parse ABI JSON: %v", err)
	}

	formattedABI, err := json.MarshalIndent(abiJSON, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format ABI JSON: %v", err)
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
		explorer := fmt.Sprintf(MultichainEtherscan, network.ChainIds[0])

		for folderName, contracts := range network.References {
			for contractName, address := range contracts {
				key := fmt.Sprintf("%s:%s", folderName, contractName)

				fmt.Printf("Generating reference for %s/%s on %s\n", folderName, contractName, explorer)

				if processed[key] {
					fmt.Printf("Skipping %s/%s on %s - already processed\n",
						folderName, contractName, explorer)
					continue
				}

				if err := GenerateReference(explorer, folderName, contractName, address, retries); err != nil {
					return err
				}

				processed[key] = true
			}
		}
	}

	return nil
}
