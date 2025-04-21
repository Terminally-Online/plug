package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/client"
	"solver/internal/helpers/zerion"
	"strings"
)

func GetCollectiblesHeldOptions[T any](lookup *actions.SchemaLookup[T], index int) ([]actions.Option, error) {
	collectibles, err := zerion.GetCollectibles([]string{"base"}, lookup.From, lookup.From, 100)
	if err != nil {
		return nil, err
	}

	chainName := client.GetChainName(lookup.ChainId)
	var options []actions.Option

	searchTerm := strings.ToLower(lookup.Search[index])

	for _, collectible := range collectibles {
		if collectible.Attributes.NFTInfo.Flags.IsSpam {
			continue
		}

		nftName := collectible.Attributes.NFTInfo.Name
		collectionName := collectible.Attributes.CollectionInfo.Name

		if searchTerm != "" &&
			!strings.Contains(strings.ToLower(nftName), searchTerm) &&
			!strings.Contains(strings.ToLower(collectionName), searchTerm) {
			continue
		}

		interfaceType := strings.ReplaceAll(collectible.Attributes.NFTInfo.Interface, "ERC", "")
		optionValue := fmt.Sprintf("%s:%s:%s", collectible.Attributes.NFTInfo.ContractAddress, collectible.Attributes.NFTInfo.TokenID, interfaceType)

		var defaultIcon string
		if collectible.Attributes.NFTInfo.Content.Detail != nil {
			defaultIcon = collectible.Attributes.NFTInfo.Content.Detail.URL
		} else if collectible.Attributes.NFTInfo.Content.Preview != nil {
			defaultIcon = collectible.Attributes.NFTInfo.Content.Preview.URL
		} else {
			defaultIcon = collectible.Attributes.NFTInfo.Content.Icon.URL
		}

		secondaryIcon := fmt.Sprintf("https://cdn.onplug.io/blockchain/%s.png", chainName)

		var quantity string
		if collectible.Attributes.Amount != "" {
			quantity = collectible.Attributes.Amount
		} else {
			quantity = "1"
		}

		options = append(options, actions.Option{
			Label: nftName,
			Name:  collectionName,
			Value: optionValue,
			Icon:  &actions.OptionIcon{Default: defaultIcon, Secondary: secondaryIcon},
			Info:  &actions.OptionInfo{Label: quantity, Value: fmt.Sprintf("#%s", collectible.Attributes.NFTInfo.TokenID)},
		})
	}

	return options, nil
}
