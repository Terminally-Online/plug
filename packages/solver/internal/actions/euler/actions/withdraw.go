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

type WithdrawRequest struct {
	Amount          coil.CoilInput[string, *big.Int] `json:"amount"`
	Token           string                           `json:"token"`
	Vault           string                           `json:"vault"`
	SubAccountIndex uint8                            `json:"sub-account" default:"0"` // default to 0 for withdraw from non-collateral vaults
}

var WithdrawFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     euler_evault_implementation.EulerEvaultImplementationMetaData,
	FunctionName: "withdraw",
}

func HandleWithdraw(lookup *actions.SchemaLookup[WithdrawRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	var updates []coil.Update
	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&WithdrawFunc,
		"amount",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	calldata, err := WithdrawFunc.GetCalldata(
		amount,
		lookup.From,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	wrappedCalldata, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		common.HexToAddress(lookup.Inputs.Vault),
		subAccountAddress,
		big.NewInt(0),
		calldata,
		updates,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap withdraw collateral call: %w", err)
	}

	return []signature.Plug{wrappedCalldata}, nil
}
