package actions

import (
	"fmt"
	"solver/bindings/erc_20"
	"solver/bindings/yearn_v3_gauge"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"solver/internal/actions/yearn_v3/types"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type StakeRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
}

var StakeFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     yearn_v3_gauge.YearnV3GaugeMetaData,
	FunctionName: "deposit",
}

func Stake(lookup *actions.SchemaLookup[StakeRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert stake amount to uint: %w", err)
	}

	vaults, err := reads.GetVaults(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}
	var targetVault *types.YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, lookup.Inputs.Token) {
			targetVault = &vault
			break
		}
	}
	if targetVault == nil || !targetVault.Staking.Available {
		return nil, fmt.Errorf("staking not available for vault: %s", lookup.Inputs.Token)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(targetVault.Staking.Address), amount)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	stakeCalldata, err := StakeFunc.GetCalldata(amount, lookup.From)
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
