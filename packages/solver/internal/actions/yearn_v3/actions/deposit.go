package actions

import (
	"fmt"
	"solver/bindings/erc_20"
	"solver/bindings/yearn_v3_pool"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

var ApproveFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "approve",
}

var DepositFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     yearn_v3_pool.YearnV3PoolMetaData,
	FunctionName: "deposit",
}

type DepositRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Vault  string `json:"vault"`
}

func Deposit(lookup *actions.SchemaLookup[DepositRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	approveCalldata, err := ApproveFunc.GetCalldata(lookup.Inputs.Vault, amount)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositCalldata, err := DepositFunc.GetCalldata(amount, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(lookup.Inputs.Vault),
		Data: depositCalldata,
	}}, nil
}
