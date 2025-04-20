package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_router"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	morpho_utils "solver/internal/actions/morpho/utils"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type HealthFactorRequest struct {
	Target string `json:"target"`
}

var HealthFactorFunc = actions.ActionOnchainFunctionResponse{
	Arguments: &abi.Arguments{
		{Name: "healthFactor", Type: abi.Type{T: abi.UintTy, Size: 256}},
	},
}

// NOTE: Morpho has health factor formatted as a WAD (18 decimals)
func HealthFactor(lookup *actions.SchemaLookup[HealthFactorRequest]) ([]signature.Plug, error) {
	market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	morpho, err := morpho_router.NewMorphoRouter(
		common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
		lookup.Client,
	)
	if err != nil {
		return nil, err
	}
	position, err := morpho.Position(
		lookup.Client.ReadOptions(lookup.From),
		[32]byte(common.Hex2BytesFixed(market.UniqueKey, 32)),
		lookup.From,
	)
	if err != nil {
		return nil, err
	}

	borrowAssets := morpho_utils.MulDivUp(
		position.BorrowShares,
		new(big.Int).Add(market.State.BorrowAssets, morpho_utils.VirtualAssets),
		new(big.Int).Add(market.State.BorrowShares, morpho_utils.VirtualShares),
	)
	maxBorrow := morpho_utils.WMulDown(
		morpho_utils.MulDivDown(position.Collateral, market.State.Price, morpho_utils.OraclePriceScale),
		market.LLTV,
	)

	healthFactor := morpho_utils.WMulDown(maxBorrow, borrowAssets)
	if borrowAssets.Cmp(big.NewInt(0)) != 1 {
		// NOTE: If the user has no borrow assets, the health factor is infinite 
		//       and we can't divide by zero so set it as max health.
		healthFactor = utils.Uint256Max
	}

	healthFactorCalldata, err := HealthFactorFunc.GetCalldata(healthFactor)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		Selector: signature.ForwardedCall,
		Data:     healthFactorCalldata,
	}}, nil
}
