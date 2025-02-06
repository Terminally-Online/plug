package plug

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/helpers/zerion"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type PlugOptionsProvider struct{}

func (p *PlugOptionsProvider) GetOptions(chainId uint64, from common.Address, action string) (map[int]actions.Options, error) {
	transferOptions, err := GetTransferOptions(chainId, from)
	if err != nil {
		return nil, err
	}

	switch action {
	case actions.ActionTransfer:
		return map[int]actions.Options{
			1: {Simple: transferOptions},
		}, nil
	default:
		return nil, nil
	}
}

func GetTransferOptions(chainId uint64, from common.Address) ([]actions.Option, error) {
	positions, err := zerion.GetFungiblePositions([]string{"base"}, from, from)
	if err != nil {
		return nil, err
	}
	var options []actions.Option
	for _, position := range positions {
		var implementation string
		var icon string
		for _, chainImplementation := range position.Attributes.FungibleInfo.Implementations {
			implementation = fmt.Sprintf("%s:%d", chainImplementation.Address, chainImplementation.Decimals)
			icon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, chainImplementation.Address)
		}
		if position.Attributes.FungibleInfo.Icon != nil {
			icon = position.Attributes.FungibleInfo.Icon.URL
		}
		quantity := utils.FormatNumber(position.Attributes.Quantity.Float, "")
		value := utils.FormatNumber(position.Attributes.Value, "$")
		options = append(options, actions.Option{
			Label: strings.ToUpper(position.Attributes.FungibleInfo.Symbol),
			Name:  position.Attributes.FungibleInfo.Name,
			Value: implementation,
			Icon:  icon,
			Info: actions.OptionInfo{
				Label: quantity,
				Value: value,
			},
		})
	}

	return options, nil
}
