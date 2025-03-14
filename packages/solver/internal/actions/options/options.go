package options

import (
	"fmt"
	"slices"
	"solver/internal/actions"
	"solver/internal/client"
	"solver/internal/helpers/zerion"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"
)

var (
	ensSuffix = ".eth"
)

func getENSAddress(name string) (common.Address, error) {
	client, err := client.New(1)
	if err != nil {
		return common.Address{}, err
	}

	address, err := ens.Resolve(client, name)
	if err != nil {
		return common.Address{}, err
	}

	return address, nil
}

func getENSName(address common.Address) (*string, error) {
	client, err := client.New(1)
	if err != nil {
		return nil, err
	}

	name, err := ens.ReverseResolve(client, address)
	if err != nil {
		return nil, err
	}

	return &name, nil
}

func GetAddressOptions(chainId uint64, from common.Address, search string) ([]actions.Option, error) {
	var options []actions.Option

	if search != "" && common.IsHexAddress(search) {
		address := common.HexToAddress(search)
		ensName, err := getENSName(address)
		if err != nil {
			return nil, err
		}
		name := address.Hex()
		if ensName != nil {
			name = *ensName
		}
		options = append(options, actions.Option{
			Label: utils.FormatAddress(address),
			Name:  name,
			Value: address.Hex(),
		})
	}

	// NOTE: It may look like we are always checking the ENS data even when an
	//       address is provided and that would be correct -- There are hundreds of
	//       ENS minted and renewed that use their address as their ENS.
	if search != "" && !strings.HasSuffix(search, ensSuffix) {
		// NOTE: This is only here for optimistic suffixing so that we can start polling
		//       results for a name before they have technically finished. If we wanted to
		//       save some RPC calls we could only search when the search already ends with
		//       .eth, but we are in the business of making things feel magic.
		transformed := strings.TrimSuffix(search, ".")
		transformed = strings.TrimSuffix(transformed, "eth")
		transformed = fmt.Sprintf("%s%s", transformed, ensSuffix)

		if address, err := getENSAddress(transformed); err == nil && address != (common.Address{}) {
			options = append(options, actions.Option{
				Label: transformed,
				Name:  utils.FormatAddress(address),
				Value: address.Hex(),
				Icon: &actions.OptionIcon{
					Default: fmt.Sprintf("https://cdn.onplug.io/ens/%s", address.Hex()),
				},
			})
		}
	}

	options = append(options, actions.Option{
		Label: utils.FormatAddress(from),
		Name:  utils.FormatAddress(from),
		Value: from.Hex(),
	})

	return options, nil
}

func GetFungiblesHeldOptions(chainId uint64, from common.Address, search string) ([]actions.Option, error) {
	positions, err := zerion.GetFungiblePositions([]string{"base"}, from, from, search)
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

func GetFungiblesOptions(chainId uint64, from common.Address, search string) ([]actions.Option, error) {
	fungibles, err := zerion.GetFungibles(search, []string{"base"})
	if err != nil {
		return nil, err
	}

	chainName := client.GetChainName(chainId)
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
			defaultIcon = fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, address)
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

func GetFungiblesAndFungiblesHeldOptions(chainId uint64, from common.Address, search string) ([]actions.Option, error) {
	fungiblesHeldOptions, err := GetFungiblesHeldOptions(chainId, from, search)
	if err != nil {
		return nil, err
	}

	fungiblesOptions, err := GetFungiblesOptions(chainId, from, search)
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
