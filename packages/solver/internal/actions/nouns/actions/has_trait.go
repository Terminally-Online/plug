package actions

import (
	"fmt"
	"solver/bindings/nouns_art"
	"solver/bindings/nouns_auction_house"
	"solver/bindings/nouns_token"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type HasTraitRequest struct {
	TraitType string `json:"traitType"`
	Trait     string `json:"trait"`
}

func HasTrait(lookup *actions.SchemaLookup[HasTraitRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.TraitType == "" || lookup.Inputs.Trait == "" {
		return nil, fmt.Errorf("trait type and trait must be provided")
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
	nounsToken, err := nouns_token.NewNounsToken(
		common.HexToAddress(references.Mainnet.References["nouns"]["token"]),
		client,
	)
	if err != nil {
		return nil, err
	}

	seeds, err := nounsToken.Seeds(client.ReadOptions(lookup.From), auction.NounId)
	if err != nil {
		return nil, err
	}

	art, err := nouns_art.NewNounsArt(
		common.HexToAddress(references.Mainnet.References["nouns"]["art"]),
		client,
	)
	if err != nil {
		return nil, err
	}

	var value string
	switch lookup.Inputs.TraitType {
	case "background":
		value, err = art.Backgrounds(client.ReadOptions(lookup.From), seeds.Background)
		if err != nil {
			return nil, err
		}
	case "body":
		bodyValue, err := art.Bodies(client.ReadOptions(lookup.From), seeds.Body)
		if err != nil {
			return nil, err
		}
		value = string(bodyValue)
	case "accessory":
		accessoryValue, err := art.Accessories(client.ReadOptions(lookup.From), seeds.Accessory)
		if err != nil {
			return nil, err
		}
		value = string(accessoryValue)
	case "head":
		headValue, err := art.Heads(client.ReadOptions(lookup.From), seeds.Head)
		if err != nil {
			return nil, err
		}
		value = string(headValue)
	case "glasses":
		glassesValue, err := art.Glasses(client.ReadOptions(lookup.From), seeds.Glasses)
		if err != nil {
			return nil, err
		}
		value = string(glassesValue)
	default:
		return nil, fmt.Errorf("invalid trait type")
	}

	if value != lookup.Inputs.Trait {
		return nil, fmt.Errorf("trait does not match")
	}

	return nil, nil
}
