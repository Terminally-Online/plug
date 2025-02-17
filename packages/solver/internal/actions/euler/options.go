package euler

import (
	"fmt"
	"math"
	"math/big"
	"solver/bindings/euler_account_lens"
	"solver/bindings/euler_vault_lens"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type EulerOptionsProvider struct{}

func (p *EulerOptionsProvider) GetOptions(chainId uint64, address common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	switch action {
	case ActionEarn, ActionDepositCollateral, ActionWithdraw:
		vaults, err := GetVerifiedVaults(chainId)
		if err != nil {
			return nil, err
		}
		supplyTokenOptions, _, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId, vaults)
		if err != nil {
			return nil, err
		}
		addressPositions, err := GetAddressPositions(chainId, address)
		if err != nil {
			fmt.Printf("error getting address positions: %v\n", err)
			return nil, err
		}
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
			3: {Simple: addressPositions},
		}, nil

	case ActionBorrow, ActionRepay:
		vaults, err := GetVerifiedVaults(chainId)
		if err != nil {
			return nil, err
		}
		borrowTokenOptions, _, borrowTokenToVaultOptions, err := GetBorrowTokenToVaultOptions(chainId, vaults)
		if err != nil {
			return nil, err
		}
		addressPositions, err := GetAddressPositions(chainId, address)
		if err != nil {
			fmt.Printf("error getting address positions: %v\n", err)
			return nil, err
		}
		return map[int]actions.Options{
			1: {Simple: borrowTokenOptions},
			2: {Complex: borrowTokenToVaultOptions},
			3: {Simple: addressPositions},
		}, nil

	case ConstraintHealthFactor, ConstraintTimeToLiq:
		vaults, err := GetVerifiedVaults(chainId)
		if err != nil {
			return nil, err
		}
		_, borrowVaultOptions, _, err := GetBorrowTokenToVaultOptions(chainId, vaults)
		if err != nil {
			return nil, err
		}
		addressPositions, err := GetAddressPositions(chainId, address)
		if err != nil {
			fmt.Printf("error getting address positions: %v\n", err)
			return nil, err
		}
		return map[int]actions.Options{
			0: {Simple: borrowVaultOptions},
			1: {Simple: addressPositions},
			2: {Simple: actions.BaseThresholdFields},
		}, nil

	case ConstraintAPY:
		vaults, err := GetVerifiedVaults(chainId)
		if err != nil {
			return nil, err
		}
		_, supplyVaultOptions, _, err := GetSupplyTokenToVaultOptions(chainId, vaults)
		if err != nil {
			return nil, err
		}
		_, borrowVaultOptions, _, err := GetBorrowTokenToVaultOptions(chainId, vaults)
		if err != nil {
			return nil, err
		}
		return map[int]actions.Options{
			0: {Simple: actions.BaseLendActionTypeFields},
			1: {Complex: map[string][]actions.Option{
				"-1": borrowVaultOptions,
				"1":  supplyVaultOptions,
			}},
			2: {Simple: actions.BaseThresholdFields},
		}, nil

	default:
		return nil, fmt.Errorf("unsupported action for options: %s", action)
	}
}

func GetSupplyTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) ([]actions.Option, []actions.Option, map[string][]actions.Option, error) {
	type result struct {
		tokenOptions        []actions.Option
		vaultOptions        []actions.Option
		tokenToVaultOptions map[string][]actions.Option
	}

	cacheKey := fmt.Sprintf("euler:supplyTokens:%d", chainId)
	res, err := utils.WithCache(cacheKey, []time.Duration{5 * time.Minute}, true, func() (result, error) {
		seenToken := make(map[string]bool)
		tokenOptions := make([]actions.Option, 0)
		vaultOptions := make([]actions.Option, 0)
		tokenToVaultOptions := make(map[string][]actions.Option)

		for _, vault := range vaults {
			tokenAddress := fmt.Sprintf("%s:%d", vault.Asset, vault.AssetDecimals)
			if !seenToken[tokenAddress] {
				tokenOptions = append(tokenOptions, actions.Option{
					Label: vault.AssetSymbol,
					Name:  vault.AssetName,
					Value: tokenAddress,
					Icon:  actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
				})
				seenToken[tokenAddress] = true
			}

			var supplyApy string
			if len(vault.IrmInfo.InterestRateInfo) > 0 {
				supplyApyFloat := utils.UintToFloat(vault.IrmInfo.InterestRateInfo[0].SupplyAPY, 25)
				supplyApy = fmt.Sprintf("%.2f%%", supplyApyFloat)
			} else {
				supplyApy = "0.0%"
			}

			vaultOption := actions.Option{
				Label: vault.VaultSymbol,
				Name:  vault.VaultName,
				Value: vault.Vault.String(),
				Icon:  actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
				Info: actions.OptionInfo{
					Label: "Supply APY",
					Value: supplyApy,
				},
			}
			vaultOptions = append(vaultOptions, vaultOption)
			tokenToVaultOptions[tokenAddress] = append(tokenToVaultOptions[tokenAddress], vaultOption)
		}

		return result{
			tokenOptions:        tokenOptions,
			vaultOptions:        vaultOptions,
			tokenToVaultOptions: tokenToVaultOptions,
		}, nil
	})

	if err != nil {
		return nil, nil, nil, err
	}

	return res.tokenOptions, res.vaultOptions, res.tokenToVaultOptions, nil
}

func GetBorrowTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) ([]actions.Option, []actions.Option, map[string][]actions.Option, error) {
	type result struct {
		tokenOptions        []actions.Option
		vaultOptions        []actions.Option
		tokenToVaultOptions map[string][]actions.Option
	}

	cacheKey := fmt.Sprintf("euler:borrowTokens:%d", chainId)
	res, err := utils.WithCache(cacheKey, []time.Duration{5 * time.Minute}, true, func() (result, error) {
		seenToken := make(map[string]bool)
		tokenOptions := make([]actions.Option, 0)
		vaultOptions := make([]actions.Option, 0)
		tokenToVaultOptions := make(map[string][]actions.Option)

		for _, vault := range vaults {
			tokenAddress := fmt.Sprintf("%s:%d", vault.Asset, vault.AssetDecimals)
			if !seenToken[tokenAddress] {
				tokenOptions = append(tokenOptions, actions.Option{
					Label: vault.AssetSymbol,
					Name:  vault.AssetName,
					Value: tokenAddress,
					Icon:  actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
				})
				seenToken[tokenAddress] = true
			}

			var borrowApy string
			if len(vault.IrmInfo.InterestRateInfo) > 0 {
				borrowApyFloat := utils.UintToFloat(vault.IrmInfo.InterestRateInfo[0].BorrowAPY, 25)
				borrowApy = fmt.Sprintf("%.2f%%", borrowApyFloat)
			} else {
				borrowApy = "0.0%"
			}

			vaultOption := actions.Option{
				Label: vault.VaultSymbol,
				Name:  vault.VaultName,
				Value: vault.Vault.String(),
				Icon:  actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
				Info: actions.OptionInfo{
					Label: "Borrow APY",
					Value: borrowApy,
				},
			}
			vaultOptions = append(vaultOptions, vaultOption)
			tokenToVaultOptions[tokenAddress] = append(tokenToVaultOptions[tokenAddress], vaultOption)
		}

		return result{
			tokenOptions:        tokenOptions,
			vaultOptions:        vaultOptions,
			tokenToVaultOptions: tokenToVaultOptions,
		}, nil
	})

	if err != nil {
		return nil, nil, nil, err
	}

	return res.tokenOptions, res.vaultOptions, res.tokenToVaultOptions, nil
}

func GetAddressPositions(chainId uint64, address common.Address) ([]actions.Option, error) {
	if address == utils.ZeroAddress {
		return nil, nil
	}

	cacheKey := fmt.Sprintf("euler:positions:%d:%s", chainId, address.String())
	return utils.WithCache(cacheKey, []time.Duration{5 * time.Minute}, true, func() ([]actions.Option, error) {
		accountLensAbi, err := euler_account_lens.EulerAccountLensMetaData.GetAbi()
		if err != nil {
			return nil, utils.ErrABI("EulerAccountLens")
		}

		accountLensAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["account_lens"])
		evcAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["evc"])

		calls := make([]utils.MulticallCalldata, 256)
		for i := 0; i < 256; i++ {
			subAccountAddress := GetSubAccountAddress(address, uint8(i))
			calls[i] = utils.MulticallCalldata{
				Target: accountLensAddr,
				Method: "getAccountEnabledVaultsInfo",
				Args:   []interface{}{evcAddr, subAccountAddress},
				ABI:    accountLensAbi,
				OutputType: &struct {
					VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
				}{},
			}
		}

		multicallAddress := common.HexToAddress(references.Networks[chainId].References["multicall"]["primary"])
		results, err := utils.ExecuteMulticall(chainId, multicallAddress, calls)
		if err != nil {
			return nil, fmt.Errorf("multicall failed: %w", err)
		}

		options := make([]actions.Option, 0)
		// Create a map to store options by index to maintain order
		optionsByIndex := make(map[int]actions.Option)

		consecutiveEmptyAccounts := 0
		var firstEmptyAccount *actions.Option

		// We want to search through the results and only stop when we find 3 consecutive empty accounts to account for times when the user had a position in an earlier account but withdrew it while the later ones still have it
		for i, result := range results {
			accountInfo := result.(*struct {
				VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
			})
			subAccountAddress := GetSubAccountAddress(address, uint8(i))

			if len(accountInfo.VaultAccountInfo) == 0 {
				consecutiveEmptyAccounts++
				if consecutiveEmptyAccounts >= 3 {
					break
				}

				if firstEmptyAccount == nil {
					firstEmptyAccount = &actions.Option{
						Label: fmt.Sprintf("Account #%d", i+1),
						Name:  utils.FormatAddress(subAccountAddress),
						Value: fmt.Sprintf("%d", i),
						Info: actions.OptionInfo{
							Label: "No Asset Defined",
							Value: "$0.00",
						},
						Icon: actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "new account")},
					}
				}
				continue
			}

			consecutiveEmptyAccounts = 0

			// Collect vaults that need price info
			type VaultWithAccount struct {
				vaultAccountInfo  euler_account_lens.VaultAccountInfo
				accountIndex      int
				subAccountAddress common.Address
			}
			var failedVaults []VaultWithAccount

			// First pass - collect failed vaults and handle successful ones
			for _, vaultAccountInfo := range accountInfo.VaultAccountInfo {
				if vaultAccountInfo.LiquidityInfo.QueryFailure {
					failedVaults = append(failedVaults, VaultWithAccount{
						vaultAccountInfo:  vaultAccountInfo,
						accountIndex:      i,
						subAccountAddress: subAccountAddress,
					})
					continue
				}

				netValue := utils.UintToFloat(new(big.Int).Sub(vaultAccountInfo.LiquidityInfo.CollateralValueRaw, vaultAccountInfo.LiquidityInfo.LiabilityValue), 18)

				// Store successful vault in the map
				optionsByIndex[i] = actions.Option{
					Label: fmt.Sprintf("Account #%d", i+1),
					Name:  utils.FormatAddress(vaultAccountInfo.Account),
					Value: fmt.Sprintf("%d", i),
					Info: actions.OptionInfo{
						Label: vaultAccountInfo.Asset.Hex(),
						Value: fmt.Sprintf("$%.2f", netValue),
					},
					Icon: actions.OptionIcon{
						Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vaultAccountInfo.Asset.Hex())),
					},
				}
			}

			// If we have failed vaults, handle them in batch
			if len(failedVaults) > 0 {
				// Process results and create options
				for _, failedVault := range failedVaults {
					// Skip if we already have a successful vault for this index
					if _, exists := optionsByIndex[failedVault.accountIndex]; exists {
						continue
					}

					price, err := GetVaultPrice(failedVault.vaultAccountInfo.Vault.String(), chainId)
					if err != nil {
						fmt.Printf("error getting vault price: %v\n", err)
						continue
					}

					// For failed vaults, we need to:
					// 1. Convert assets to float using vault's asset decimals
					// 2. Convert borrowed to float using vault's asset decimals
					// 3. Calculate net value using the price
					decimals := uint8(price.vault.AssetDecimals.Uint64())
					fmt.Printf("Failed vault - Raw values: Assets: %v, Borrowed: %v, Decimals: %v, Price: %v\n",
						failedVault.vaultAccountInfo.Assets,
						failedVault.vaultAccountInfo.Borrowed,
						decimals,
						price.price)

					assetValue := utils.UintToFloat(failedVault.vaultAccountInfo.Assets, decimals)
					borrowedValue := utils.UintToFloat(failedVault.vaultAccountInfo.Borrowed, decimals)
					netValue := (assetValue - borrowedValue) * price.price * math.Pow10(int(decimals))
					fmt.Printf("Failed vault - Converted values: AssetValue: %v, BorrowedValue: %v, NetValue: %v\n",
						assetValue,
						borrowedValue,
						netValue)

					optionsByIndex[failedVault.accountIndex] = actions.Option{
						Label: fmt.Sprintf("Account #%d", failedVault.accountIndex+1),
						Name:  utils.FormatAddress(failedVault.vaultAccountInfo.Account),
						Value: fmt.Sprintf("%d", failedVault.accountIndex),
						Info: actions.OptionInfo{
							Label: failedVault.vaultAccountInfo.Asset.Hex(),
							Value: fmt.Sprintf("$%.2f", netValue),
						},
						Icon: actions.OptionIcon{
							Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(failedVault.vaultAccountInfo.Asset.Hex())),
						},
					}
				}
			}
		}

		fmt.Printf("optionsByIndex: %v\n", optionsByIndex)

		// Convert map to slice in order
		maxIndex := -1
		for index := range optionsByIndex {
			if index > maxIndex {
				maxIndex = index
			}
		}

		for i := 0; i <= maxIndex; i++ {
			if option, exists := optionsByIndex[i]; exists {
				options = append(options, option)
			}
		}

		if firstEmptyAccount != nil {
			options = append(options, *firstEmptyAccount)
		}

		return options, nil
	})
}
