package actions

import (
	"fmt"
	"solver/bindings/aave_v3_pool"
	"solver/bindings/erc_20"
	"solver/internal/actions"
	aave_utils "solver/internal/actions/aave_v3/utils"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type RepayRequest struct {
	Token  string `json:"token"`
	Amount string `json:"amount"`
}

func Repay(lookup *actions.SchemaLookup[RepayRequest]) ([]signature.Plug, error) {
	tokenIn, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}
	amountIn, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert repayment amount to uint: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(references.Networks[lookup.ChainId].References["aave_v3"]["pool"]),
		amountIn,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	repayCalldata, err := poolAbi.Pack("repay",
		tokenIn,
		amountIn,
		aave_utils.InterestRateMode,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   *tokenIn,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[lookup.ChainId].References["aave_v3"]["pool"]),
		Data: repayCalldata,
	}}, nil
}
