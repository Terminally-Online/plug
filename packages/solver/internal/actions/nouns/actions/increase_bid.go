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

type IncreaseBidRequest struct {
	Percent string `json:"percent"`
}

func IncreaseBid(lookup *actions.SchemaLookup[IncreaseBidRequest]) ([]signature.Plug, error) {
	percent, err := utils.StringToUint(lookup.Inputs.Percent, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to convert percent to uint: %w", err)
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

	bid := new(big.Int).Mul(
		auction.Amount,
		new(big.Int).Add(
			big.NewInt(1),
			new(big.Int).Div(percent, big.NewInt(100)),
		),
	)

	auctionHouseAbi, err := nouns_auction_house.NounsAuctionHouseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	bidCalldata, err := auctionHouseAbi.Pack("createBid", auction.NounId)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:    common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		Data:  bidCalldata,
		Value: bid,
	}}, nil
}
