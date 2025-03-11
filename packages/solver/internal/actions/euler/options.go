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

type EulerOptionsProvider struct{}

func (p *EulerOptionsProvider) GetOptions(chainId uint64, address common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, err
	}

	switch action {
	case ActionEarn, ActionWithdraw:
		supplyTokenOptions, _, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId, vaults)
		if err != nil {
			return nil, err
		}
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
		}, nil
	case ActionDepositCollateral, ActionWithdrawCollateral:
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
		addressPositions, err := GetAddressPositions(chainId, address)
		if err != nil {
			fmt.Printf("error getting address positions: %v\n", err)
			return nil, err
		}
		return map[int]actions.Options{
			0: {Simple: addressPositions},
		}, nil

	case ConstraintAPY:
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

		calls := make([]client.MulticallCalldata, 256)
		for i := 0; i < 256; i++ {
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

		// First pass - collect and organize vault data by subaccount
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

			// Process vaults for this subaccount
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

		// Second pass - create options from the organized data
		options := make([]actions.Option, 0)
		optionsByIndex := make(map[int]actions.Option)

		// Create options for active accounts
		for i := 0; i <= maxIndex; i++ {
			if accountVaults, exists := vaultsBySubaccount[i]; exists {
				var debtIcon string
				var netValue float64
				// var healthFactor float64
				var totalCollateralValue float64

				// find the largest value collateral vault
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
					// healthFactor = utils.UintToFloat(accountVaults.debtVault.LiquidityInfo.HealthFactor, 18)
				} else {
					netValue = totalCollateralValue
					debtIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "debt")
					// healthFactor = -1
				}

				collateralIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(largestCollateralVault.vault.Asset.Hex()))
				allIcon := fmt.Sprintf("%s%%7C%s", collateralIcon, debtIcon)
				fmt.Printf("allIcon: %s\n", allIcon)

				optionsByIndex[i] = actions.Option{
					Label: fmt.Sprintf("Account #%d", i),
					Name:  utils.FormatAddress(accountAddress),
					Value: fmt.Sprintf("%d", i),
					Info: actions.OptionInfo{
						Label: "Health Factor: 1.2",
						Value: fmt.Sprintf("$%.2f", netValue),
					},
					Icon: actions.OptionIcon{Default: allIcon},
				}
			} else {
				// Create empty account option
				option := createEmptyAccountOption(i, chainId, address)
				optionsByIndex[i] = option
			}
		}

		// Add options in order
		for i := 0; i <= maxIndex; i++ {
			if option, exists := optionsByIndex[i]; exists {
				options = append(options, option)
			}
		}

		// Add one more empty account at the end
		nextIndex := maxIndex + 1
		options = append(options, createEmptyAccountOption(nextIndex, chainId, address))

		return options, nil
	})
}

// Helper function to create an empty account option
func createEmptyAccountOption(index int, chainId uint64, address common.Address) actions.Option {
	collateralIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "collateral")
	debtIcon := fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "debt")

	return actions.Option{
		Label: fmt.Sprintf("Account #%d", index),
		Name:  utils.FormatAddress(GetSubAccountAddress(address, uint8(index))),
		Value: fmt.Sprintf("%d", index),
		Info: actions.OptionInfo{
			Label: "Health Factor: -",
			Value: "$0.00",
		},
		Icon: actions.OptionIcon{Default: fmt.Sprintf("%s%%7C%s", collateralIcon, debtIcon)},
	}
}
