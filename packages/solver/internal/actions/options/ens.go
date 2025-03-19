package options

import (
	"solver/internal/client"

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
