package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/nouns_auction_house"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type BidRequest struct {
	Amount string `json:"amount"`
}

var BidFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     nouns_auction_house.NounsAuctionHouseMetaData,
	FunctionName: "createBid",
}

func Bid(lookup *actions.SchemaLookup[BidRequest]) ([]signature.Plug, error) {
	amount, err := utils.StringToUint(lookup.Inputs.Amount, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert big amount to uint: %w", err)
	}

	client, err := client.New(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		client,
	)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(client.ReadOptions(lookup.From))
	if err != nil {
		return nil, err
	}

	bidCalldata, err := BidFunc.GetCalldata(auction.NounId, lookup.From)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:    common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		Data:  bidCalldata,
		Value: new(big.Int).Mul(amount, big.NewInt(1e18)),
	}}, nil
}
