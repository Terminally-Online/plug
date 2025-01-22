package nouns

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/bindings/nouns_art"
	"solver/bindings/nouns_auction_house"
	"solver/bindings/nouns_token"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionBid(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
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

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(common.HexToAddress(utils.Mainnet.References["nouns"]["auction_house"]), provider)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
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

	return []*types.Transaction{{
		To:    utils.Mainnet.References["nouns"]["auction_house"],
		Data:  "0x" + common.Bytes2Hex(bidCalldata),
		Value: *new(big.Int).Mul(amount, big.NewInt(1e18)),
	}}, nil
}

func HandleActionIncreaseBid(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
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

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(common.HexToAddress(utils.Mainnet.References["nouns"]["auction_house"]), provider)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
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

	return []*types.Transaction{{
		To:    utils.Mainnet.References["nouns"]["auction_house"],
		Data:  "0x" + common.Bytes2Hex(bidCalldata),
		Value: *bid,
	}}, nil
}

func HandleConstraintHasTrait(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
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

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(common.HexToAddress(utils.Mainnet.References["nouns"]["auction_house"]), provider)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
	if err != nil {
		return nil, err
	}
	nounsToken, err := nouns_token.NewNounsToken(common.HexToAddress(utils.Mainnet.References["nouns"]["token"]), provider)
	if err != nil {
		return nil, err
	}

	seeds, err := nounsToken.Seeds(utils.BuildCallOpts(params.From, big.NewInt(0)), auction.NounId)
	if err != nil {
		return nil, err
	}

	art, err := nouns_art.NewNounsArt(common.HexToAddress(utils.Mainnet.References["nouns"]["art"]), provider)
	if err != nil {
		return nil, err
	}

	var value string
	switch inputs.TraitType {
	case "background":
		value, err = art.Backgrounds(utils.BuildCallOpts(params.From, big.NewInt(0)), seeds.Background)
		if err != nil {
			return nil, err
		}
	case "body":
		bodyValue, err := art.Bodies(utils.BuildCallOpts(params.From, big.NewInt(0)), seeds.Body)
		if err != nil {
			return nil, err
		}
		value = string(bodyValue)
	case "accessory":
		accessoryValue, err := art.Accessories(utils.BuildCallOpts(params.From, big.NewInt(0)), seeds.Accessory)
		if err != nil {
			return nil, err
		}
		value = string(accessoryValue)
	case "head":
		headValue, err := art.Heads(utils.BuildCallOpts(params.From, big.NewInt(0)), seeds.Head)
		if err != nil {
			return nil, err
		}
		value = string(headValue)
	case "glasses":
		glassesValue, err := art.Glasses(utils.BuildCallOpts(params.From, big.NewInt(0)), seeds.Glasses)
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

func HandleConstraintIsTokenId(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Id *big.Int `json:"id"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(common.HexToAddress(utils.Mainnet.References["nouns"]["auction_house"]), provider)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
	if err != nil {
		return nil, err
	}

	if auction.NounId.Cmp(inputs.Id) != 0 {
		return nil, fmt.Errorf("noun id does not match")
	}

	return nil, nil
}

func HandleConstraintCurrentBidWithinRange(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
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

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(common.HexToAddress(utils.Mainnet.References["nouns"]["auction_house"]), provider)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
	if err != nil {
		return nil, err
	}

	if auction.Amount.Cmp(min) < 0 || auction.Amount.Cmp(max) > 0 {
		return nil, fmt.Errorf("current bid is not within range")
	}

	return nil, nil
}
