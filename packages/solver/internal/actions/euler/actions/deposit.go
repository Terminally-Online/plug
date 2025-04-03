package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_evault_implementation"
	"solver/internal/actions"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type DepositCollateralRequest struct {
	Amount          coil.CoilInput[string, *big.Int] `json:"amount"`
	Token           string                           `json:"token"`
	Vault           common.Address                   `json:"vault"`
	SubAccountIndex uint8                            `json:"sub-account"`
}

var DepositCollateralFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     euler_evault_implementation.EulerEvaultImplementationMetaData,
	FunctionName: "deposit",
}

var EnableCollateralFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     euler_evault_implementation.EulerEvaultImplementationMetaData,
	FunctionName: "enableCollateral",
}

func DepositCollateral(lookup *actions.SchemaLookup[DepositCollateralRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
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

	approveCalldata, err := actions.Erc20ApprovalFunc.GetCalldata(
		lookup.Inputs.Vault,
		approveAmount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	var depositUpdates []coil.Update
	depositAmount, depositUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&DepositCollateralFunc,
		"amount",
		depositUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	depositCalldata, err := DepositCollateralFunc.GetCalldata(
		depositAmount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		lookup.Inputs.Vault,
		subAccountAddress,
		big.NewInt(0),
		depositCalldata,
		depositUpdates,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	enableCollateralCalldata, err := EnableCollateralFunc.GetCalldata(
		subAccountAddress,
		lookup.Inputs.Vault,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	enableCollateralCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["evc"]),
		subAccountAddress,
		big.NewInt(0),
		enableCollateralCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap enable collateral call: %w", err)
	}

	return []signature.Plug{{
		To:      *token,
		Data:    approveCalldata,
		Updates: approveUpdates,
	}, depositCall, enableCollateralCall}, nil
}
