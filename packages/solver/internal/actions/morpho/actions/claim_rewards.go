package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_distributor"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

var ClaimRewardsFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     morpho_distributor.MorphoDistributorMetaData,
	FunctionName: "claim",
}

func ClaimRewards(lookup *actions.SchemaLookup[any]) ([]signature.Plug, error) {
	// TODO MASON: this is the last build time read I need to figure out how to get rid of.. Would love some input here as I don't want to pack an arbitrary amount of distributions into the option value
	distributions, err := reads.GetDistributions(lookup.From, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get distributions: %w", err)
	}
	if len(distributions) == 0 {
		return nil, fmt.Errorf("no active distributions found for address: %s", lookup.From)
	}

	claims := []signature.Plug{}
	for _, distribution := range distributions {
		claimable, ok := new(big.Int).SetString(distribution.Claimable, 10)
		if !ok {
			return nil, fmt.Errorf("failed to parse claimable amount: %s", distribution.Claimable)
		}
		proof := make([][32]byte, len(distribution.Proof))
		for i, proofItem := range distribution.Proof {
			proof[i] = common.HexToHash(proofItem)
		}

		claimCalldata, err := ClaimRewardsFunc.GetCalldata(
			lookup.From,
			common.HexToAddress(distribution.Asset.Address),
			claimable,
			proof,
		)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}
		claims = append(claims, signature.Plug{
			To:   common.HexToAddress(distribution.Distributor.Address),
			Data: claimCalldata,
		})
	}

	return claims, nil
}
