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
	"solver/internal/client"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionDeposit(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Vault  string `json:"vault"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
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

	vaults, err := GetVaults(params.ChainId)
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
	depositCalldata, err := vaultAbi.Pack("deposit", amount, common.HexToAddress(params.From))
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

func HandleActionWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Vault  string `json:"vault"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
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

	vaults, err := GetVaults(params.ChainId)
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
	depositCalldata, err := vaultAbi.Pack("deposit", amount, common.HexToAddress(params.From))
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(inputs.Vault),
		Data: depositCalldata,
	}}, nil
}

func HandleActionStake(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
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

	vaults, err := GetVaults(params.ChainId)
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
	stakeCalldata, err := gaugeAbi.Pack("deposit", amount, common.HexToAddress(params.From))
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

func HandleActionStakeMax(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		TokenIn string `json:"token"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stake max inputs: %v", err)
	}

	token, _, err := utils.ParseAddressAndDecimals(inputs.TokenIn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	vaults, err := GetVaults(params.ChainId)
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

	if targetVault == nil || !targetVault.Staking.Available {
		return nil, fmt.Errorf("staking not available for vault: %s", token)
	}

	client, err := client.New(params.ChainId)
	if err != nil {
		return nil, err
	}

	erc20Contract, err := erc_20.NewErc20(*token, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create ERC20 contract instance: %v", err)
	}

	callOpts := utils.BuildCallOpts(params.From, big.NewInt(0))

	balance, err := erc20Contract.BalanceOf(callOpts, common.HexToAddress(params.From))
	if err != nil {
		return nil, fmt.Errorf("failed to get vault token balance: %v", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(targetVault.Staking.Address), balance)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	gaugeAbi, err := yearn_v3_gauge.YearnV3GaugeMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Gauge")
	}
	stakeCalldata, err := gaugeAbi.Pack("deposit", balance, common.HexToAddress(params.From))
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

func HandleActionRedeem(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
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

	vaults, err := GetVaults(params.ChainId)
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

	client, err := client.New(params.ChainId)
	if err != nil {
		return nil, err
	}

	gaugeContract, err := yearn_v3_gauge.NewYearnV3Gauge(common.HexToAddress(targetVault.Staking.Address), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gauge contract instance: %v", err)
	}

	callOpts := utils.BuildCallOpts(params.From, big.NewInt(0))
	maxRedeem, err := gaugeContract.MaxRedeem(callOpts, common.HexToAddress(params.From))
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

	calldata, err := gaugeAbi.Pack("redeem", amount, common.HexToAddress(params.From), common.HexToAddress(params.From))
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(targetVault.Staking.Address),
		Data: calldata,
	}}, nil
}

func HandleActionRedeemMax(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal redeem max inputs: %v", err)
	}

	token, _, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	vaults, err := GetVaults(params.ChainId)
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

	client, err := client.New(params.ChainId)
	if err != nil {
		return nil, err
	}

	gaugeContract, err := yearn_v3_gauge.NewYearnV3Gauge(common.HexToAddress(targetVault.Staking.Address), client)
	if err != nil {
		return nil, fmt.Errorf("failed to create gauge contract instance: %v", err)
	}

	callOpts := utils.BuildCallOpts(params.From, big.NewInt(0))
	maxRedeem, err := gaugeContract.MaxRedeem(callOpts, common.HexToAddress(params.From))
	if err != nil {
		return nil, fmt.Errorf("failed to get max redeemable amount: %v", err)
	}

	if maxRedeem.Sign() <= 0 {
		return nil, fmt.Errorf("no staked tokens available to redeem")
	}

	gaugeAbi, err := yearn_v3_gauge.YearnV3GaugeMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Gauge")
	}

	calldata, err := gaugeAbi.Pack("redeem", maxRedeem, common.HexToAddress(params.From), common.HexToAddress(params.From))
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(targetVault.Staking.Address),
		Data: calldata,
	}}, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Vault     string `json:"token"`
		Operator  int    `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs")
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold float: %w", err)
	}

	vaults, err := GetVaults(params.ChainId)
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
