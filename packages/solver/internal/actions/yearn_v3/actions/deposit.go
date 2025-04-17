package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/yearn_v3_pool"
	"solver/internal/actions"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

var DepositFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     yearn_v3_pool.YearnV3PoolMetaData,
	FunctionName: "deposit",
}

type DepositRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Token  string                           `json:"token"`
	Vault  string                           `json:"vault"`
}

func Deposit(lookup *actions.SchemaLookup[DepositRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
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
		common.HexToAddress(lookup.Inputs.Vault),
		approvalAmount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	var depositUpdates []coil.Update
	depositAmount, depositUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&DepositFunc,
		"assets",
		depositUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	depositCalldata, err := DepositFunc.GetCalldata(
		depositAmount,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      *token,
		Data:    approveCalldata,
		Updates: approvalUpdates,
	}, {
		To:      common.HexToAddress(lookup.Inputs.Vault),
		Data:    depositCalldata,
		Updates: depositUpdates,
	}}, nil
}
