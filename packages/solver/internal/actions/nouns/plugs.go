package nouns

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/nouns_art"
	"solver/bindings/nouns_auction_house"
	"solver/bindings/nouns_token"
	"solver/bindings/plug_router"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionBid(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Amount string `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	amount, err := utils.StringToUint(inputs.Amount, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert big amount to uint: %w", err)
	}

	client, err := client.New(params.ChainId)
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

	auction, err := auctionHouse.Auction(client.ReadOptions(params.From))
	if err != nil {
		return nil, err
	}

	auctionHouseAbi, err := nouns_auction_house.NounsAuctionHouseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	bidCalldata, err := auctionHouseAbi.Pack("createBid", auction.NounId)
	if err != nil {
		return nil, err
	}

	return []plug_router.PlugTypesLibPlug{{
		To:    common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		Data:  bidCalldata,
		Value: new(big.Int).Mul(amount, big.NewInt(1e18)),
	}}, nil
}

func HandleActionIncreaseBid(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Percent string `json:"percent"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	percent, err := utils.StringToUint(inputs.Percent, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to convert percent to uint: %w", err)
	}

	client, err := client.New(params.ChainId)
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

	auction, err := auctionHouse.Auction(client.ReadOptions(params.From))
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

	return []plug_router.PlugTypesLibPlug{{
		To:    common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		Data:  bidCalldata,
		Value: bid,
	}}, nil
}

func HandleConstraintHasTrait(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		TraitType string `json:"traitType"`
		Trait     string `json:"trait"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	if inputs.TraitType == "" || inputs.Trait == "" {
		return nil, fmt.Errorf("trait type and trait must be provided")
	}

	client, err := client.New(params.ChainId)
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

	auction, err := auctionHouse.Auction(client.ReadOptions(params.From))
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

	seeds, err := nounsToken.Seeds(client.ReadOptions(params.From), auction.NounId)
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
	switch inputs.TraitType {
	case "background":
		value, err = art.Backgrounds(client.ReadOptions(params.From), seeds.Background)
		if err != nil {
			return nil, err
		}
	case "body":
		bodyValue, err := art.Bodies(client.ReadOptions(params.From), seeds.Body)
		if err != nil {
			return nil, err
		}
		value = string(bodyValue)
	case "accessory":
		accessoryValue, err := art.Accessories(client.ReadOptions(params.From), seeds.Accessory)
		if err != nil {
			return nil, err
		}
		value = string(accessoryValue)
	case "head":
		headValue, err := art.Heads(client.ReadOptions(params.From), seeds.Head)
		if err != nil {
			return nil, err
		}
		value = string(headValue)
	case "glasses":
		glassesValue, err := art.Glasses(client.ReadOptions(params.From), seeds.Glasses)
		if err != nil {
			return nil, err
		}
		value = string(glassesValue)
	default:
		return nil, fmt.Errorf("invalid trait type")
	}

	if value != inputs.Trait {
		return nil, fmt.Errorf("trait does not match")
	}

	return nil, nil
}

func HandleConstraintIsTokenId(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Id *big.Int `json:"id"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		params.Client,
	)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(params.Client.ReadOptions(params.From))
	if err != nil {
		return nil, err
	}

	if auction.NounId.Cmp(inputs.Id) != 0 {
		return nil, fmt.Errorf("noun id does not match")
	}

	return nil, nil
}

func HandleConstraintCurrentBidWithinRange(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Min string `json:"min"`
		Max string `json:"max"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	min, err := utils.StringToUint(inputs.Min, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert min to uint: %w", err)
	}
	max, err := utils.StringToUint(inputs.Max, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert max to uint: %w", err)
	}

	if min.Cmp(max) >= 0 {
		return nil, fmt.Errorf("min must be less than max")
	}

	client, err := client.New(params.ChainId)
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

	auction, err := auctionHouse.Auction(client.ReadOptions(params.From))
	if err != nil {
		return nil, err
	}

	if auction.Amount.Cmp(min) < 0 || auction.Amount.Cmp(max) > 0 {
		return nil, fmt.Errorf("current bid is not within range")
	}

	return nil, nil
}
