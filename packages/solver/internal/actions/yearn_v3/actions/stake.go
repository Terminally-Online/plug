package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/yearn_v3_gauge"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type StakeRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Token  string                           `json:"token"`
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

	targetVault, err := reads.GetVault(lookup.ChainId, token.String())
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %v", err)
	}
	if !targetVault.Staking.Available {
		return nil, fmt.Errorf("staking not available for vault: %s", lookup.Inputs.Token)
	}

	var approvalUpdates []coil.Update
	approvalAmount, approvalUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&actions.Erc20ApprovalFunc,
		"_value",
		approvalUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}
	approveCalldata, err := actions.Erc20ApprovalFunc.GetCalldata(
		common.HexToAddress(targetVault.Staking.Address),
		approvalAmount,
	)

	var stakeUpdates []coil.Update
	stakeAmount, stakeUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&StakeFunc,
		"_assets",
		stakeUpdates,
		lookup.PreviousActionDefinition,
	)

	stakeCalldata, err := StakeFunc.GetCalldata(stakeAmount, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      *token,
		Data:    approveCalldata,
		Updates: approvalUpdates,
	}, {
		To:      common.HexToAddress(targetVault.Staking.Address),
		Data:    stakeCalldata,
		Updates: stakeUpdates,
	}}, nil
}
