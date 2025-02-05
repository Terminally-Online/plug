package euler

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/utils"
	"strings"
)

type EulerOptionsProvider struct{}

func (p *EulerOptionsProvider) GetOptions(chainId uint64, action string) (map[int]actions.Options, error) {
	supplyTokenOptions, supplyVaultOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}

	borrowTokenOptions, borrowVaultOptions, borrowTokenToVaultOptions, err := GetBorrowTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}

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