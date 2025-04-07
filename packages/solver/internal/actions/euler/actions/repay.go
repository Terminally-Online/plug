package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_evault_implementation"
	"solver/internal/actions"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type RepayRequest struct {
	Amount          coil.CoilInput[string, *big.Int] `json:"amount"`
	Token           string                           `json:"token"`
	Vault           string                           `json:"vault"`
	SubAccountIndex uint8                            `json:"sub-account"`
}

var RepayFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     euler_evault_implementation.EulerEvaultImplementationMetaData,
	FunctionName: "repay",
}

func Repay(lookup *actions.SchemaLookup[RepayRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	var approveUpdates []coil.Update
	approveAmount, approveUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&actions.Erc20ApprovalFunc,
		"_value",
		approveUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	approveCalldata, err := actions.Erc20ApprovalFunc.GetCalldata(
		common.HexToAddress(lookup.Inputs.Vault),
		approveAmount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	var repayUpdates []coil.Update
	amount, repayUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&RepayFunc,
		"amount",
		repayUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	repayCalldata, err := RepayFunc.GetCalldata(
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	repayCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		common.HexToAddress(lookup.Inputs.Vault),
		subAccountAddress,
		big.NewInt(0),
		repayCalldata,
		repayUpdates,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap repay call: %w", err)
	}

	return []signature.Plug{{
		To:      common.HexToAddress(lookup.Inputs.Token),
		Data:    approveCalldata,
		Updates: approveUpdates,
	}, repayCall}, nil
}
