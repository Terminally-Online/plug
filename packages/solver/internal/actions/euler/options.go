package euler

import (
	"fmt"
	"math"
	"math/big"
	"solver/bindings/euler_account_lens"
	"solver/bindings/euler_vault_lens"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/utils"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func GetSupplyTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) (tokenOptions []actions.Option, vaultOptions []actions.Option, tokenToVaultOptions map[string][]actions.Option, err error) {
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
					Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
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
				Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
				Info:  &actions.OptionInfo{Label: "Supply APY", Value: supplyApy},
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

func SupplyTokenToVaultOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, err
	}

	supplyTokenOptions, _, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId, vaults)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyTokenOptions},
		2: {Complex: supplyTokenToVaultOptions},
	}, nil
}

func createEmptyAccountOption(index int, chainId uint64, address common.Address) actions.Option {
	collateralIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "collateral")
	debtIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "debt")

	return actions.Option{
		Label: fmt.Sprintf("Account #%d", index),
		Name:  utils.FormatAddress(GetSubAccountAddress(address, uint8(index))),
		Value: fmt.Sprintf("%d", index),
		Info:  &actions.OptionInfo{Label: "Health Factor: -", Value: "$0.00"},
		Icon:  &actions.OptionIcon{Default: fmt.Sprintf("%s%%7C%s", collateralIcon, debtIcon)},
	}
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

		calls := make([]client.MulticallCalldata, 256)
		for i := range 256 {
			subAccountAddress := GetSubAccountAddress(address, uint8(i))
			calls[i] = client.MulticallCalldata{
				Target: accountLensAddr,
				Method: "getAccountEnabledVaultsInfo",
				Args:   []interface{}{evcAddr, subAccountAddress},
				ABI:    accountLensAbi,
				OutputType: &struct {
					VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
				}{},
			}
		}

		client, err := client.New(chainId)
		if err != nil {
			return nil, fmt.Errorf("multicall failed: %w", err)
		}
		results, err := client.Multicall(calls)
		if err != nil {
			return nil, fmt.Errorf("multicall failed: %w", err)
		}

		vaultsBySubaccount := make(map[int]struct {
			debtVault        *euler_account_lens.VaultAccountInfo
			collateralVaults []euler_account_lens.VaultAccountInfo
		})
		consecutiveEmptyAccounts := 0
		maxIndex := -1

		for i, result := range results {
			accountInfo := result.(*struct {
				VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
			})

			if len(accountInfo.VaultAccountInfo) == 0 {
				consecutiveEmptyAccounts++
				if consecutiveEmptyAccounts >= 3 {
					break
				}
				continue
			}

			consecutiveEmptyAccounts = 0
			maxIndex = i

			for _, vaultAccountInfo := range accountInfo.VaultAccountInfo {
				if vaultAccountInfo.LiquidityInfo.QueryFailure {
					// This is a collateral vault
					temp := vaultsBySubaccount[i]
					temp.collateralVaults = append(temp.collateralVaults, vaultAccountInfo)
					vaultsBySubaccount[i] = temp
				} else {
					// This is a debt vault
					temp := vaultsBySubaccount[i]
					temp.debtVault = &vaultAccountInfo
					vaultsBySubaccount[i] = temp
				}
			}
		}

		options := make([]actions.Option, 0)
		optionsByIndex := make(map[int]actions.Option)

		for i := 0; i <= maxIndex; i++ {
			if accountVaults, exists := vaultsBySubaccount[i]; exists {
				var debtIcon string
				var netValue float64
				var healthFactor string
				var totalCollateralValue float64

				var largestCollateralVault *struct {
					vault euler_account_lens.VaultAccountInfo
					value float64
				}
				for _, collateralVault := range accountVaults.collateralVaults {
					vaultPrice, err := GetVaultPrice(collateralVault.Vault.String(), chainId)
					if err != nil {
						fmt.Printf("error getting vault price: %v\n", err)
						continue
					}

					decimals := uint8(vaultPrice.vault.AssetDecimals.Uint64())
					assetValue := utils.UintToFloat(collateralVault.Assets, decimals)
					netValue := assetValue * vaultPrice.price * math.Pow10(int(decimals))
					totalCollateralValue += netValue
					if largestCollateralVault == nil || netValue > largestCollateralVault.value {
						largestCollateralVault = &struct {
							vault euler_account_lens.VaultAccountInfo
							value float64
						}{
							vault: collateralVault,
							value: netValue,
						}
					}
				}

				accountAddress := GetSubAccountAddress(address, uint8(i))
				if accountVaults.debtVault != nil {
					netValue = utils.UintToFloat(new(big.Int).Sub(accountVaults.debtVault.LiquidityInfo.CollateralValueRaw, accountVaults.debtVault.LiquidityInfo.LiabilityValue), 18)
					debtIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(accountVaults.debtVault.Asset.Hex()))

					totalValueBorrowed := accountVaults.debtVault.LiquidityInfo.LiabilityValue
					lltvValueCollateral := accountVaults.debtVault.LiquidityInfo.CollateralValueLiquidation

					healthFactorFloat := new(big.Float).SetInt(lltvValueCollateral)
					healthFactorFloat.Quo(healthFactorFloat, new(big.Float).SetInt(totalValueBorrowed))
					healthFactorF64, _ := healthFactorFloat.Float64()
					healthFactor = fmt.Sprintf("Health Factor: %.2f", healthFactorF64)
				} else {
					netValue = totalCollateralValue
					debtIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "debt")
					healthFactor = "Health Factor: -"
				}

				collateralIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(largestCollateralVault.vault.Asset.Hex()))
				allIcon := fmt.Sprintf("%s%%7C%s", collateralIcon, debtIcon)
				fmt.Printf("allIcon: %s\n", allIcon)

				optionsByIndex[i] = actions.Option{
					Label: fmt.Sprintf("Account #%d", i),
					Name:  utils.FormatAddress(accountAddress),
					Value: fmt.Sprintf("%d", i),
					Info:  &actions.OptionInfo{Label: healthFactor, Value: fmt.Sprintf("$%.2f", netValue)},
					Icon:  &actions.OptionIcon{Default: allIcon},
				}
			} else {
				option := createEmptyAccountOption(i, chainId, address)
				optionsByIndex[i] = option
			}
		}

		for i := 0; i <= maxIndex; i++ {
			if option, exists := optionsByIndex[i]; exists {
				options = append(options, option)
			}
		}

		nextIndex := maxIndex + 1
		options = append(options, createEmptyAccountOption(nextIndex, chainId, address))

		return options, nil
	})
}

func PositionOptions(chainId uint64, from common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	addressPositions, err := GetAddressPositions(chainId, from)
	if err != nil {
		fmt.Printf("error getting address positions: %v\n", err)
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: addressPositions},
	}, nil
}

func SupplyTokenToVaultToPositionsOptions(chainId uint64, from common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, err
	}

	supplyTokenOptions, _, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId, vaults)
	if err != nil {
		return nil, err
	}
	addressPositions, err := GetAddressPositions(chainId, from)
	if err != nil {
		fmt.Printf("error getting address positions: %v\n", err)
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyTokenOptions},
		2: {Complex: supplyTokenToVaultOptions},
		3: {Simple: addressPositions},
	}, nil
}

func GetBorrowTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) (tokenOptions []actions.Option, vaultOptions []actions.Option, tokenToVaultOptions map[string][]actions.Option, err error) {
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
					Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
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
				Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
				Info:  &actions.OptionInfo{Label: "Borrow APY", Value: borrowApy},
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

func BorrowTokenToVaultOptions(chainId uint64, from common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, err
	}

	borrowTokenOptions, _, borrowTokenToVaultOptions, err := GetBorrowTokenToVaultOptions(chainId, vaults)
	if err != nil {
		return nil, err
	}
	addressPositions, err := GetAddressPositions(chainId, from)
	if err != nil {
		fmt.Printf("error getting address positions: %v\n", err)
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: borrowTokenOptions},
		2: {Complex: borrowTokenToVaultOptions},
		3: {Simple: addressPositions},
	}, nil
}

func GetTokenOptions(chainId uint64, from common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
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
}
