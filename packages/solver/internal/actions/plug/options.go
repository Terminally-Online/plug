package plug

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/options"
	"solver/internal/client"
	"solver/internal/helpers/zerion"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func TransferOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	recipientIndex := 2
	recipientOptions, err := options.GetAddressOptions(chainId, from, search[recipientIndex])
	if err != nil {
		return nil, err
	}
	transferOptions, err := GetTransferOptions(chainId, from)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1:              {Simple: transferOptions},
		recipientIndex: {Simple: recipientOptions},
	}, nil
}

func SwapOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	transferOptions, err := GetTransferOptions(chainId, from)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: transferOptions},
		2: {Simple: transferOptions},
	}, nil
}

func PriceOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	transferOptions, err := GetTransferOptions(chainId, from)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: transferOptions},
		1: {Simple: actions.BaseThresholdFields},
	}, nil
}

func BalanceOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	addressOptions, err := options.GetAddressOptions(chainId, from, search[1])
	if err != nil {
		return nil, err
	}
	transferOptions, err := GetTransferOptions(chainId, from)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: transferOptions},
		1: {Simple: addressOptions},
		2: {Simple: actions.BaseThresholdFields},
	}, nil
}

func GetTransferOptions(chainId uint64, from common.Address) ([]actions.Option, error) {
	positions, err := zerion.GetFungiblePositions([]string{"base"}, from, from)
	if err != nil {
		return nil, err
	}
	chainName := client.GetChainName(chainId)
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
			Icon:  &actions.OptionIcon{Default: defaultIcon, Secondary: secondaryIcon},
			Info:  &actions.OptionInfo{Label: quantity, Value: value},
		})
	}

	return options, nil
}
