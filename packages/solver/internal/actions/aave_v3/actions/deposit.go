package actions

import (
	"fmt"
	"solver/bindings/aave_v3_pool"
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type DepositRequest struct {
		Token  string `json:"token"`
		Amount string `json:"amount"`
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

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(references.Networks[lookup.ChainId].References["aave_v3"]["pool"]),
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	depositCalldata, err := poolAbi.Pack("deposit",
		token,
		amount,
		lookup.From,
		uint16(0),
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[lookup.ChainId].References["aave_v3"]["pool"]),
		Data: depositCalldata,
	}}, nil
}
