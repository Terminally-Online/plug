package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_router"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type BorrowRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Target string `json:"target"`
}

func Borrow(lookup *actions.SchemaLookup[BorrowRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert borrow amount to uint: %w", err)
	}

	market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	morphoAbi, err := morpho_router.MorphoRouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho abi: %w", err)
	}
	borrowCalldata, err := morphoAbi.Pack(
		"borrow",
		market.Params,
		amount,
		big.NewInt(0),
		lookup.From,
		lookup.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack borrow calldata: %w", err)
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
		Data: borrowCalldata,
	}}, nil
}
