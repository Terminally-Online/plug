package actions

import (
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type BalanceRequest struct {
	Token   string         `json:"token"`
	Address common.Address `json:"address"`
}

var BalanaceFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     erc_20.Erc20MetaData,
	FunctionName: "balanceOf",
}


func Balance(lookup *actions.SchemaLookup[BalanceRequest]) ([]signature.Plug, error) {
	token, _, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, err
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	balanceCalldata, err := erc20Abi.Pack("balanceOf", lookup.Inputs.Address)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:   *token,
		Data: balanceCalldata,
	}}, nil
}
