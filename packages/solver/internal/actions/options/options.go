package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/client"
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
