package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/morpho_router"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type RepayRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Target string `json:"target"`
}

func Repay(lookup *actions.SchemaLookup[RepayRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert repay amount to uint: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(
		references.Networks[lookup.ChainId].References["morpho"]["router"]),
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}
	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	repayCalldata, err := morphoAbi.Pack(
		"repay",
		market.Params,
		amount,
		big.NewInt(0),
		lookup.From,
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack repay calldata: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
		Data: repayCalldata,
	}}, nil
}
