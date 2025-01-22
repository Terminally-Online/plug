package yearn_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"solver/bindings/erc_20"
	"solver/bindings/yearn_v3_gauge"
	"solver/bindings/yearn_v3_pool"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionDeposit(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
		Vault  string   `json:"vault"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}
	if inputs.Amount.Sign() <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	vaults, err := GetVaults()
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
	if !strings.EqualFold(inputs.Token, targetVault.Token.Address) {
		return nil, fmt.Errorf("asset %s cannot be used in vault: %s", inputs.Token, inputs.Vault)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(inputs.Vault), inputs.Amount)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	vaultAbi, err := yearn_v3_pool.YearnV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Pool")
	}
	depositCalldata, err := vaultAbi.Pack("deposit", inputs.Amount, common.HexToAddress(params.From))
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(inputs.Token),
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(inputs.Vault),
		Data: depositCalldata,
	}}, nil
}

func HandleActionWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
		Vault  string   `json:"vault"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}
	if inputs.Amount.Sign() <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	vaults, err := GetVaults()
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
	if !strings.EqualFold(inputs.Token, targetVault.Token.Address) {
		return nil, fmt.Errorf("asset %s cannot be used in vault: %s", inputs.Token, inputs.Vault)
	}

	vaultAbi, err := yearn_v3_pool.YearnV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Pool")
	}
	depositCalldata, err := vaultAbi.Pack("deposit", inputs.Amount, common.HexToAddress(params.From))
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
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stake inputs: %v", err)
	}

	vaults, err := GetVaults()
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}
	if inputs.Amount.Sign() <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
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
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(targetVault.Staking.Address), inputs.Amount)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	gaugeAbi, err := yearn_v3_gauge.YearnV3GaugeMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Gauge")
	}
	stakeCalldata, err := gaugeAbi.Pack("deposit", inputs.Amount, common.HexToAddress(params.From))
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(inputs.Token),
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

	vaults, err := GetVaults()
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}

	var targetVault *YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, inputs.TokenIn) {
			targetVault = &vault
			break
		}
	}

	if targetVault == nil || !targetVault.Staking.Available {
		return nil, fmt.Errorf("staking not available for vault: %s", inputs.TokenIn)
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}

	erc20Contract, err := erc_20.NewErc20(common.HexToAddress(inputs.TokenIn), provider)
	if err != nil {
		return nil, fmt.Errorf("failed to create ERC20 contract instance: %v", err)
	}

	callOpts := utils.BuildCallOpts(params.From, big.NewInt(0))

	balance, err := erc20Contract.BalanceOf(callOpts, common.HexToAddress(params.From))
	if err != nil {
		return nil, fmt.Errorf("failed to get vault token balance: %v", err)
	}

	if balance.Sign() <= 0 {
		return nil, fmt.Errorf("no vault tokens available to stake")
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
		To:   common.HexToAddress(inputs.TokenIn),
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(targetVault.Staking.Address),
		Data: stakeCalldata,
	}}, nil
}

func HandleActionRedeem(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
		Token  string   `json:"token"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal redeem inputs: %v", err)
	}
	if inputs.Amount.Sign() <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0")
	}

	vaults, err := GetVaults()
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
	if targetVault == nil {
		return nil, fmt.Errorf("redeem not available for vault: %s", inputs.Token)
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}

	gaugeContract, err := yearn_v3_gauge.NewYearnV3Gauge(common.HexToAddress(targetVault.Staking.Address), provider)
	if err != nil {
		return nil, fmt.Errorf("failed to create gauge contract instance: %v", err)
	}

	callOpts := utils.BuildCallOpts(params.From, big.NewInt(0))
	maxRedeem, err := gaugeContract.MaxRedeem(callOpts, common.HexToAddress(params.From))
	if err != nil {
		return nil, fmt.Errorf("failed to get max redeemable amount: %v", err)
	}

	if maxRedeem.Cmp(inputs.Amount) < 0 {
		return nil, fmt.Errorf("insufficient staked balance")
	}

	gaugeAbi, err := yearn_v3_gauge.YearnV3GaugeMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("YearnV3Gauge")
	}

	calldata, err := gaugeAbi.Pack("redeem", inputs.Amount, common.HexToAddress(params.From), common.HexToAddress(params.From))
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

	vaults, err := GetVaults()
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
	if targetVault == nil {
		return nil, fmt.Errorf("redeem not available for vault: %s", inputs.Token)
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}

	gaugeContract, err := yearn_v3_gauge.NewYearnV3Gauge(common.HexToAddress(targetVault.Staking.Address), provider)
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
		Vault     string     `json:"token"`
		Operator  int        `json:"operator"`
		Threshold *big.Float `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs")
	}

	vaults, err := GetVaults()
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
		if rateFloat.Cmp(inputs.Threshold) >= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not less than threshold %.2f%%", rateFloat, inputs.Threshold)
		}
	case 1:
		if rateFloat.Cmp(inputs.Threshold) <= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not greater than threshold %.2f%%", rateFloat, inputs.Threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}
