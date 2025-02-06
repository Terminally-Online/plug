package euler

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_account_lens"
	"solver/bindings/euler_vault_lens"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type EulerOptionsProvider struct{}

func (p *EulerOptionsProvider) GetOptions(chainId uint64, address common.Address, action string) (map[int]actions.Options, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, err
	}

	supplyTokenOptions, supplyVaultOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId, vaults)
	if err != nil {
		return nil, err
	}

	borrowTokenOptions, borrowVaultOptions, borrowTokenToVaultOptions, err := GetBorrowTokenToVaultOptions(chainId, vaults)
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
				3: {Simple: addressPositions},
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

func GetSupplyTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) ([]actions.Option, []actions.Option, map[string][]actions.Option, error) {
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

func GetBorrowTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) ([]actions.Option, []actions.Option, map[string][]actions.Option, error) {
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

	// Do we really have to make this call for all 256 sub accounts? Or do we need to lean on an indexer like euler does.
	options := make([]actions.Option, 0)
	for i := 0; i < 5; i++ {	
		subAccountAddress := GetSubAccountAddress(address, uint8(i))
		accountEnabledVaults, err := accountLens.GetAccountEnabledVaultsInfo(
			nil,
			common.HexToAddress(references.Networks[chainId].References["euler"]["evc"]),
			subAccountAddress,
		)
		if err != nil {
			fmt.Printf("error getting account enabled vaults contract instance: %v\n", err)
			return nil, err
		}

		for _, vault := range accountEnabledVaults.VaultAccountInfo {
			// account liquidity info returns a query failure if it's a borrow not a supply vault.
			if vault.LiquidityInfo.QueryFailure { continue };

			netValue := new(big.Int).Sub(vault.LiquidityInfo.CollateralValueRaw, vault.LiquidityInfo.LiabilityValue)
			accountOption := actions.Option{
				Label: fmt.Sprintf("Account %d", i),
				Name:  vault.Account.String(),
				Value: fmt.Sprintf("$%.2f", utils.UintToFloat(netValue, 18)),
				Info: actions.OptionInfo{
					Label: "Borrowing Value",
					Value: fmt.Sprintf("$%.2f", utils.UintToFloat(vault.LiquidityInfo.LiabilityValue, 18)),
				},
			}
			options = append(options, accountOption)

			fmt.Printf("account %d (%s) Vault (%s): \n%v\n\n", i, vault.Account, vault.Vault, vault.LiquidityInfo)
		}
	}

	return options, nil
}