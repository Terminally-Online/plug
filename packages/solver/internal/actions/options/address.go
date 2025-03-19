package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func GetAddressOptions[T any](lookup *actions.SchemaLookup[T], index int) ([]actions.Option, error) {
	var options []actions.Option

	if lookup.Search[index] != "" && common.IsHexAddress(lookup.Search[index]) {
		address := common.HexToAddress(lookup.Search[index])
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
	if lookup.Search[index] != "" && !strings.HasSuffix(lookup.Search[index], ensSuffix) {
		// NOTE: This is only here for optimistic suffixing so that we can start polling
		//       results for a name before they have technically finished. If we wanted to
		//       save some RPC calls we could only search when the search already ends with
		//       .eth, but we are in the business of making things feel magic.
		transformed := strings.TrimSuffix(lookup.Search[index], ".")
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
		Label: utils.FormatAddress(lookup.From),
		Name:  utils.FormatAddress(lookup.From),
		Value: lookup.From.Hex(),
	})

	return options, nil
}
