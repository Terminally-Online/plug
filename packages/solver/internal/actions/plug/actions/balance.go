package actions

import (
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type BalanceRequest struct {
	Token  string         `json:"token"`
	Holder common.Address `json:"holder"`
}

var BalanceFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "balanceOf",
}

func Balance(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	token, _, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, err
	}

	balanceCalldata, err := BalanceFunc.GetCalldata(lookup.Inputs.Holder)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:   *token,
		Data: balanceCalldata,
	}}, nil
}
