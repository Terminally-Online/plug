package actions

import (
	"solver/bindings/nouns_auction_house"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type CurrentBidRequest struct{}

var CurrentBidFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     nouns_auction_house.NounsAuctionHouseMetaData,
	FunctionName: "auction",
}

func CurrentAuction(lookup *actions.SchemaLookup[CurrentBidRequest]) ([]signature.Plug, error) {
	calldata, err := CurrentBidFunc.GetCalldata(lookup.Inputs)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		Data: calldata,
	}}, nil
}
