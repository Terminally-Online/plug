package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/nouns_auction_house"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type IsTokenIdRequest struct {
	Id *big.Int `json:"id"`
}

func IsTokenId(lookup *actions.SchemaLookup[IsTokenIdRequest]) ([]signature.Plug, error) {
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		lookup.Client,
	)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(lookup.Client.ReadOptions(lookup.From))
	if err != nil {
		return nil, err
	}

	if auction.NounId.Cmp(lookup.Inputs.Id) != 0 {
		return nil, fmt.Errorf("noun id does not match")
	}

	return nil, nil
}
