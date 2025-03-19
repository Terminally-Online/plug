package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_distributor"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

func ClaimRewards(lookup *actions.SchemaLookup[any]) ([]signature.Plug, error) {
	distributions, err := reads.GetDistributions(lookup.From, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get distributions: %w", err)
	}
	if len(distributions) == 0 {
		return nil, fmt.Errorf("no active distributions found for address: %s", lookup.From)
	}

	distributorAbi, err := morpho_distributor.MorphoDistributorMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho distributor abi: %w", err)
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
		claimCalldata, err := distributorAbi.Pack(
			"claim",
			lookup.From,
			common.HexToAddress(distribution.Asset.Address),
			claimable,
			proof,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to pack claim calldata: %w", err)
		}
		claims = append(claims, signature.Plug{
			To:   common.HexToAddress(distribution.Distributor.Address),
			Data: claimCalldata,
		})
	}

	return claims, nil
}
