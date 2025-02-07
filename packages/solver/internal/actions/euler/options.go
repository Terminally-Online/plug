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

	switch action {
	case ActionEarn:
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
			3: {Simple: addressPositions},
		}, nil
	case ActionDepositCollateral:
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
			3: {Simple: addressPositions},
		}, nil
	case ActionWithdraw:
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
			3: {Simple: addressPositions},
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
			3: {Simple: addressPositions},
		}, nil
	case ConstraintHealthFactor:
		return map[int]actions.Options{
			0: {Simple: borrowVaultOptions},
			1: {Simple: addressPositions},
			2: {Simple: actions.BaseThresholdFields},
		}, nil
	case ConstraintAPY:
		return map[int]actions.Options{
			0: {Simple: actions.BaseLendActionTypeFields},
			1: {Complex: map[string][]actions.Option{
				"-1": borrowVaultOptions,
				"1":  supplyVaultOptions,
			}},
			2: {Simple: actions.BaseThresholdFields},
		}, nil
	case ConstraintTimeToLiq:
		return map[int]actions.Options{
			0: {Simple: borrowVaultOptions},
			1: {Simple: addressPositions},
			2: {Simple: actions.BaseThresholdFields},
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
	if address == utils.ZeroAddress {
		return nil, nil
	}

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
	for i, result := range results {
		accountInfo := result.(*struct {
			VaultAccountInfo []euler_account_lens.VaultAccountInfo `json:"vaultAccountInfo"`
		})
		subAccountAddress := GetSubAccountAddress(address, uint8(i))

		if len(accountInfo.VaultAccountInfo) == 0 {
			options = append(options, actions.Option{
				Label: fmt.Sprintf("Account %d", i),
				Name:  fmt.Sprintf("%s...%s", subAccountAddress.String()[:6], subAccountAddress.String()[len(subAccountAddress.String())-4:]),
				Value: fmt.Sprintf("%d", i),
				Info: actions.OptionInfo{
					Label: "Net Asset Value",
					Value: "$0.00",
				},
			})
			continue
		}

		for _, vault := range accountInfo.VaultAccountInfo {
			if vault.LiquidityInfo.QueryFailure {
				continue
			}

			netValue := new(big.Int).Sub(vault.LiquidityInfo.CollateralValueRaw, vault.LiquidityInfo.LiabilityValue)
			accountOption := actions.Option{
				Label: fmt.Sprintf("Account %d", i),
				Name:  fmt.Sprintf("%s...%s", vault.Account.String()[:6], vault.Account.String()[len(vault.Account.String())-4:]),
				Value: fmt.Sprintf("%d", i),
				Info: actions.OptionInfo{
					Label: "Net Asset Value",
					Value: fmt.Sprintf("$%.2f", utils.UintToFloat(netValue, 18)),
				},
			}
			options = append(options, accountOption)
		}
	}

	return options, nil
}
