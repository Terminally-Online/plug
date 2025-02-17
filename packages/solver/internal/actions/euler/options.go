package euler

import (
	"fmt"
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
	res, err := utils.WithCache(cacheKey, []time.Duration{5 * time.Minute}, func() (result, error) {
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
	res, err := utils.WithCache[result](cacheKey, []time.Duration{5 * time.Minute}, func() (result, error) {
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
	return utils.WithCache(cacheKey, []time.Duration{5 * time.Minute}, func() ([]actions.Option, error) {
		accountLensAbi, err := euler_account_lens.EulerAccountLensMetaData.GetAbi()
		if err != nil {
			return nil, utils.ErrABI("EulerAccountLens")
		}

		accountLensAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["account_lens"])
		evcAddr := common.HexToAddress(references.Networks[chainId].References["euler"]["evc"])

		var nextVault *actions.Option

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

		options := make([]actions.Option, 0)
		for i, result := range results {
			accountInfo := result.(*struct {
				VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
			})
			subAccountAddress := GetSubAccountAddress(address, uint8(i))

			if len(accountInfo.VaultAccountInfo) == 0 {
				options = append(options, actions.Option{
					Label: fmt.Sprintf("Account %d", i+1),
					Name:  utils.FormatAddress(subAccountAddress),
					Value: fmt.Sprintf("%d", i),
					Info: actions.OptionInfo{
						Label: "No Asset Defined",
						Value: "$0.00",
					},
					Icon: actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, "new account")},
				})
				break
			}

			// TODO: This seems wrong in general because we have a for loop for the 256 calls. Then, each account may contain a set of vaults.
			//       Right now, an option is added for each vault. So if an account has multiple vaults, a user had multiple assets in some of
			//       those vaults, we would have more than 256 options? I am not understanding something here.
			//		 .
			//       Shouldn't we be returning one option for each account, not vault? Why are we adding options likes this?
			//       - CHANCE

			// NOTE: Only return accounts that have value plus one that does not have a value so that the user can
			//       choose to create a new position in a specific subaccount.
			for _, vault := range accountInfo.VaultAccountInfo {
				if vault.LiquidityInfo.QueryFailure {
					continue
				}

				netValue := utils.UintToFloat(new(big.Int).Sub(vault.LiquidityInfo.CollateralValueRaw, vault.LiquidityInfo.LiabilityValue), 18)
				accountOption := actions.Option{
					Label: fmt.Sprintf("Account #%d", i+1),
					Name:  utils.FormatAddress(vault.Account),
					Value: fmt.Sprintf("%d", i),
					Info: actions.OptionInfo{
						Label: vault.Asset.Hex(),
						Value: fmt.Sprintf("$%.2f", netValue),
					},
					Icon: actions.OptionIcon{
						Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.Hex())),
					},
				}

				if netValue == 0 && nextVault != nil {
					nextVault = &accountOption
					continue
				}

				options = append(options, accountOption)
			}
		}

		return options, nil
	})
}
