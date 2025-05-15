package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_vault"
	"solver/internal/actions"
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

var EarnFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     morpho_vault.MorphoVaultMetaData,
	FunctionName: "deposit",
}

func Earn(lookup *actions.SchemaLookup[EarnRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	approveAmount, approveUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&actions.Erc20ApprovalFunc,
		"_value",
		nil,
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
		return nil, fmt.Errorf("failed to pack approve calldata: %w", err)
	}

	depositAmount, depositUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&EarnFunc,
		"assets",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	depositCalldata, err := EarnFunc.GetCalldata(
		depositAmount,
		lookup.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack deposit calldata: %w", err)
	}

	return []signature.Plug{{
		To:      *token,
		Data:    approveCalldata,
		Updates: approveUpdates,
	}, {
		To:      common.HexToAddress(lookup.Inputs.Vault),
		Data:    depositCalldata,
		Updates: depositUpdates,
	}}, nil
}
