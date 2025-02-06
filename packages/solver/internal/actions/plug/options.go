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
	case actions.ConstraintPrice:
		return map[int]actions.Options{
			0: {Simple: transferOptions},             // Token selection
			1: {Simple: actions.BaseThresholdFields}, // Comparison operators
		}, nil
	case actions.ConstraintBalance:
		return map[int]actions.Options{
			0: {Simple: transferOptions},             // Token selection
			2: {Simple: actions.BaseThresholdFields}, // Comparison operators
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
	chainName, err := utils.GetChainName(chainId)
	if err != nil {
		return nil, err
	}
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
			defaultIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, address)
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
			Icon:  actions.OptionIcon{Default: defaultIcon, Secondary: secondaryIcon},
			Info:  actions.OptionInfo{Label: quantity, Value: value},
		})
	}

	return options, nil
}
