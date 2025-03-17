package yearn_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"solver/bindings/erc_20"
	"solver/bindings/yearn_v3_gauge"
	"solver/bindings/yearn_v3_pool"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionDeposit(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Vault  string `json:"vault"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	vaults, err := GetVaults(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}

	var targetVault *YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(strings.ToLower(vault.Address), strings.ToLower(inputs.Vault)) {
			targetVault = &vault
			break
		}
	}
	if targetVault == nil {
		return nil, fmt.Errorf("deposit not available for vault: %s", inputs.Vault)
	}
	if !strings.EqualFold(token.Hex(), targetVault.Token.Address) {
		return nil, fmt.Errorf("asset %s cannot be used in vault: %s", token, inputs.Vault)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(inputs.Vault), amount)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	vaultAbi, err := yearn_v3_pool.YearnV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Pool")
	}
	depositCalldata, err := vaultAbi.Pack("deposit", amount, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(inputs.Vault),
		Data: depositCalldata,
	}}, nil
}

func HandleActionWithdraw(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Vault  string `json:"vault"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	vaults, err := GetVaults(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}
	var targetVault *YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, inputs.Vault) {
			targetVault = &vault
			break
		}
	}
	if targetVault == nil {
		return nil, fmt.Errorf("deposit not available for vault: %s", inputs.Vault)
	}
	if !strings.EqualFold(token.Hex(), targetVault.Token.Address) {
		return nil, fmt.Errorf("asset %s cannot be used in vault: %s", token, inputs.Vault)
	}

	vaultAbi, err := yearn_v3_pool.YearnV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Pool")
	}
	depositCalldata, err := vaultAbi.Pack("deposit", amount, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(inputs.Vault),
		Data: depositCalldata,
	}}, nil
}

func HandleActionStake(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stake inputs: %v", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert stake amount to uint: %w", err)
	}

	vaults, err := GetVaults(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}
	var targetVault *YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, inputs.Token) {
			targetVault = &vault
			break
		}
	}
	if targetVault == nil || !targetVault.Staking.Available {
		return nil, fmt.Errorf("staking not available for vault: %s", inputs.Token)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(targetVault.Staking.Address), amount)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	gaugeAbi, err := yearn_v3_gauge.YearnV3GaugeMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Gauge")
	}
	stakeCalldata, err := gaugeAbi.Pack("deposit", amount, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(targetVault.Staking.Address),
		Data: stakeCalldata,
	}}, nil
}

func HandleActionRedeem(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal redeem inputs: %v", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert redeem amount to uint: %w", err)
	}

	vaults, err := GetVaults(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}

	var targetVault *YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, token.Hex()) {
			targetVault = &vault
			break
		}
	}
	if targetVault == nil {
		return nil, fmt.Errorf("redeem not available for vault: %s", token)
	}

	gaugeContract, err := yearn_v3_gauge.NewYearnV3Gauge(common.HexToAddress(targetVault.Staking.Address), lookup.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gauge contract instance: %v", err)
	}

	maxRedeem, err := gaugeContract.MaxRedeem(lookup.Client.ReadOptions(lookup.From), lookup.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get max redeemable amount: %v", err)
	}

	if maxRedeem.Cmp(amount) < 0 {
		return nil, fmt.Errorf("insufficient staked balance")
	}

	gaugeAbi, err := yearn_v3_gauge.YearnV3GaugeMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Gauge")
	}

	calldata, err := gaugeAbi.Pack("redeem", amount, lookup.From, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(targetVault.Staking.Address),
		Data: calldata,
	}}, nil
}

func HandleConstraintAPY(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Vault     string `json:"token"`
		Operator  int    `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs")
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold float: %w", err)
	}

	vaults, err := GetVaults(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}

	var targetVault *YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, inputs.Vault) {
			targetVault = &vault
			break
		}
	}
	if targetVault == nil {
		return nil, fmt.Errorf("cannot find data for vault: %s", inputs.Vault)
	}

	rateFloat := new(big.Float).SetInt(new(big.Int).Add(
		new(big.Int).SetInt64(int64(targetVault.APR.ForwardAPR.NetAPR*100)),
		new(big.Int).SetInt64(int64(targetVault.Extra.StakingRewardsAPR*100)),
	))

	switch inputs.Operator {
	case -1:
		if rateFloat.Cmp(big.NewFloat(thresholdFloat)) >= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not less than threshold %.2f%%", rateFloat, thresholdFloat)
		}
	case 1:
		if rateFloat.Cmp(big.NewFloat(thresholdFloat)) <= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not greater than threshold %.2f%%", rateFloat, thresholdFloat)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}
