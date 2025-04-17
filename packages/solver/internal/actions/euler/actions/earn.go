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

type EarnRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Token  string                           `json:"token"`
	Vault  string                           `json:"vault"`
}

var DepositFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     euler_evault_implementation.EulerEvaultImplementationMetaData,
	FunctionName: "deposit",
}

func Earn(lookup *actions.SchemaLookup[EarnRequest]) ([]signature.Plug, error) {
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
		common.HexToAddress(lookup.Inputs.Vault),
		approveAmount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositAmount, depositUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&DepositFunc,
		"amount",
		approveUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	depositCalldata, err := DepositFunc.GetCalldata(
		depositAmount,
		common.HexToAddress(lookup.Inputs.Vault),
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	wrappedDepositCall, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		common.HexToAddress(lookup.Inputs.Vault),
		lookup.From,
		big.NewInt(0),
		depositCalldata,
		depositUpdates,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	return []signature.Plug{{
		To:      *token,
		Data:    approveCalldata,
		Updates: approveUpdates,
	}, wrappedDepositCall}, nil
}
