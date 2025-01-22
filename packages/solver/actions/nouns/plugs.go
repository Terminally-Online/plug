package nouns

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/bindings/nouns_art"
	"solver/bindings/nouns_auction_house"
	"solver/bindings/nouns_token"
	"solver/internal/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionBid(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount *big.Int `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	if inputs.Amount.Sign() <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		provider,
	)
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

	return []signature.Plug{{
		To:    common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		Data:  bidCalldata,
		Value: new(big.Int).Mul(inputs.Amount, big.NewInt(1e18)),
	}}, nil
}

func HandleActionIncreaseBid(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Percent *big.Int `json:"percent"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	if inputs.Percent.Sign() <= 0 {
		return nil, fmt.Errorf("percent must be greater than zero")
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		provider,
	)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
	if err != nil {
		return nil, err
	}

	bid := new(big.Int).Mul(
		auction.Amount,
		new(big.Int).Add(big.NewInt(1), inputs.Percent.Div(inputs.Percent, big.NewInt(100))),
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

func HandleConstraintHasTrait(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
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
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		provider,
	)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
	if err != nil {
		return nil, err
	}
	nounsToken, err := nouns_token.NewNounsToken(
		common.HexToAddress(references.Mainnet.References["nouns"]["token"]),
		provider,
	)
	if err != nil {
		return nil, err
	}

	seeds, err := nounsToken.Seeds(utils.BuildCallOpts(params.From, big.NewInt(0)), auction.NounId)
	if err != nil {
		return nil, err
	}

	art, err := nouns_art.NewNounsArt(
		common.HexToAddress(references.Mainnet.References["nouns"]["art"]),
		provider,
	)
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

func HandleConstraintIsTokenId(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
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
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		provider,
	)
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

func HandleConstraintCurrentBidWithinRange(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Min *big.Int `json:"min"`
		Max *big.Int `json:"max"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, err
	}

	if inputs.Min.Sign() <= 0 || inputs.Max.Sign() <= 0 {
		return nil, fmt.Errorf("min and max must be greater than zero")
	}

	if inputs.Min.Cmp(inputs.Max) >= 0 {
		return nil, fmt.Errorf("min must be less than max")
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, err
	}
	auctionHouse, err := nouns_auction_house.NewNounsAuctionHouse(
		common.HexToAddress(references.Mainnet.References["nouns"]["auction_house"]),
		provider,
	)
	if err != nil {
		return nil, err
	}

	auction, err := auctionHouse.Auction(utils.BuildCallOpts(params.From, big.NewInt(0)))
	if err != nil {
		return nil, err
	}

	if auction.Amount.Cmp(inputs.Min) < 0 || auction.Amount.Cmp(inputs.Max) > 0 {
		return nil, fmt.Errorf("current bid is not within range")
	}

	return nil, nil
}
