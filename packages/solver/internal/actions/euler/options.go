package euler

import (
	"fmt"
	"solver/bindings/euler_account_lens"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type EulerOptionsProvider struct{}

func (p *EulerOptionsProvider) GetOptions(chainId uint64, address common.Address, action string) (map[int]actions.Options, error) {
	supplyTokenOptions, supplyVaultOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}

	borrowTokenOptions, borrowVaultOptions, borrowTokenToVaultOptions, err := GetBorrowTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}

	addressPositions, err := GetAddressPositions(chainId, address)
	if err != nil {
		fmt.Printf("error getting address positions: %v\n", err)
		return nil, err
	}
	fmt.Printf("address positions: %v\n", addressPositions)

	switch action {
		case ActionEarn:
			return map[int]actions.Options{
				1: {Simple: supplyTokenOptions},
				2: {Complex: supplyTokenToVaultOptions},
			}, nil
		case ActionDepositCollateral:
			return map[int]actions.Options{
				1: {Simple: supplyTokenOptions},
				2: {Complex: supplyTokenToVaultOptions},
			}, nil
		case ActionWithdraw:
			return map[int]actions.Options{
				1: {Simple: supplyTokenOptions},
				2: {Complex: supplyTokenToVaultOptions},
			}, nil
		case ActionBorrow:
			return map[int]actions.Options{
				1: {Simple: borrowTokenOptions},
				2: {Complex: borrowTokenToVaultOptions},
			}, nil
		case ActionRepay:
			return map[int]actions.Options{
				1: {Simple: borrowTokenOptions},
				2: {Complex: borrowTokenToVaultOptions},
			}, nil
		case ConstraintHealthFactor:
			return map[int]actions.Options{
				1: {Simple: actions.BaseThresholdFields},
			}, nil
		case ConstraintAPY:
			return map[int]actions.Options{
				0: {Simple: actions.BaseLendActionTypeFields},
				1: {Complex: map[string][]actions.Option{ 
					"-1": borrowVaultOptions,
					"1": supplyVaultOptions,
				}},
				2: {Simple: actions.BaseThresholdFields},
			}, nil
		case ConstraintTimeToLiq:
			return map[int]actions.Options{
				1: {Simple: actions.BaseThresholdFields},
			}, nil
	default:
		return nil, fmt.Errorf("unsupported action for options: %s", action)
	}
}

func GetSupplyTokenToVaultOptions(chainId uint64) ([]actions.Option, []actions.Option, map[string][]actions.Option, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, nil, nil, err
	}

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
				Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String())),
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
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String())),
			Info: actions.OptionInfo{
				Label: "Supply APY",
				Value: supplyApy,
			},
		}
		vaultOptions = append(vaultOptions, vaultOption)
		tokenToVaultOptions[tokenAddress] = append(tokenToVaultOptions[tokenAddress], vaultOption)
	}

	return tokenOptions, vaultOptions, tokenToVaultOptions, nil
}

func GetBorrowTokenToVaultOptions(chainId uint64) ([]actions.Option, []actions.Option, map[string][]actions.Option, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, nil, nil, err
	}

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
				Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String())),
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
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String())),
			Info: actions.OptionInfo{
				Label: "Borrow APY",
				Value: borrowApy,
			},
		}
		vaultOptions = append(vaultOptions, vaultOption)
		tokenToVaultOptions[tokenAddress] = append(tokenToVaultOptions[tokenAddress], vaultOption)
	}

	return tokenOptions, vaultOptions, tokenToVaultOptions, nil
}

func GetAddressPositions(chainId uint64, address common.Address) ([]actions.Option, error) {
	fmt.Printf("address: %s\n", address.String())
	if (address == utils.ZeroAddress) {
		return nil, nil
	}

	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	accountLens, err := euler_account_lens.NewEulerAccountLens(
		common.HexToAddress(references.Networks[chainId].References["euler"]["account_lens"]),
		provider,
	)
	if err != nil {
		return nil, utils.ErrABI("EulerAccountLens")
	}

	accountEnabledVaults, err := accountLens.GetAccountEnabledVaultsInfo(
		nil,
		common.HexToAddress(references.Networks[chainId].References["euler"]["account_lens"]),
		address,
	)
	if err != nil {
		return nil, err
	}
	fmt.Printf("account enabled vaults: %v\n", accountEnabledVaults)

	for idx, vault := range accountEnabledVaults.VaultAccountInfo {
		fmt.Printf("account %d (%s): %v\n", idx, vault.Account, vault.LiquidityInfo)
	}

	return nil, nil
}