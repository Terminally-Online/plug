package options

import (
	"fmt"
	"slices"
	"solver/internal/actions"
	"solver/internal/client"
	"solver/internal/helpers/zerion"
	"solver/internal/utils"
	"strings"
)

func GetFungiblesHeldOptions[T any](lookup *actions.SchemaLookup[T], index int) ([]actions.Option, error) {
	positions, err := zerion.GetFungiblePositions([]string{"base"}, lookup.From, lookup.From, lookup.Search[index])
	if err != nil {
		return nil, err
	}
	chainName := client.GetChainName(lookup.ChainId)
	var options []actions.Option
	for _, position := range positions {
		var option string
		var defaultIcon string
		var secondaryIcon string
		for _, chainImplementation := range position.Attributes.FungibleInfo.Implementations {
			address := chainImplementation.Address
			if address == "" {
				address = utils.NativeTokenAddress.Hex()
			}
			option = fmt.Sprintf("%s:%d:%d", address, chainImplementation.Decimals, 20)
			defaultIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", lookup.ChainId, address)
			secondaryIcon = fmt.Sprintf("https://cdn.onplug.io/blockchain/%s.png", chainName)
		}
		if position.Attributes.FungibleInfo.Icon != nil {
			defaultIcon = position.Attributes.FungibleInfo.Icon.URL
		}
		quantity := utils.FormatNumber(position.Attributes.Quantity.Float, "")
		value := utils.FormatNumber(position.Attributes.Value, "$")
		options = append(options, actions.Option{
			Label: strings.ToUpper(position.Attributes.FungibleInfo.Symbol),
			Name:  position.Attributes.FungibleInfo.Name,
			Value: option,
			Icon:  &actions.OptionIcon{Default: defaultIcon, Secondary: secondaryIcon},
			Info:  &actions.OptionInfo{Label: quantity, Value: value},
		})
	}

	return options, nil
}

func GetFungiblesOptions[T any](lookup *actions.SchemaLookup[T], index int) ([]actions.Option, error) {
	fungibles, err := zerion.GetFungibles(lookup.Search[index], []string{"base"})
	if err != nil {
		return nil, err
	}

	chainName := client.GetChainName(lookup.ChainId)
	var options []actions.Option
	for _, fungible := range fungibles {
		var option string
		var defaultIcon string
		var secondaryIcon string
		for _, chainImplementation := range fungible.Implementations {
			address := chainImplementation.Address
			if address == "" {
				address = utils.NativeTokenAddress.Hex()
			}
			option = fmt.Sprintf("%s:%d:%d", address, chainImplementation.Decimals, 20)
			defaultIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", lookup.ChainId, address)
			secondaryIcon = fmt.Sprintf("https://cdn.onplug.io/blockchain/%s.png", chainName)
		}

		options = append(options, actions.Option{
			Label: strings.ToUpper(fungible.Symbol),
			Name:  fungible.Name,
			Value: option,
			Icon:  &actions.OptionIcon{Default: defaultIcon, Secondary: secondaryIcon},
			Info:  &actions.OptionInfo{Label: "0.00", Value: "0.00"},
		})
	}

	return options, nil
}

func GetFungiblesAndFungiblesHeldOptions[T any](lookup *actions.SchemaLookup[T], index int) ([]actions.Option, error) {
	fungiblesHeldOptions, err := GetFungiblesHeldOptions(lookup, index)
	if err != nil {
		return nil, err
	}

	fungiblesOptions, err := GetFungiblesOptions(lookup, index)
	if err != nil {
		return nil, err
	}

	heldOptionsMap := make(map[string]actions.Option, len(fungiblesHeldOptions))
	for _, option := range fungiblesHeldOptions {
		heldOptionsMap[option.Value] = option
	}

	options := make([]actions.Option, 0, len(fungiblesHeldOptions)+len(fungiblesOptions))
	options = append(options, fungiblesHeldOptions...)

	return append(options, slices.DeleteFunc(fungiblesOptions, func(o actions.Option) bool {
		_, exists := heldOptionsMap[o.Value]
		return exists
	})...), nil
}
