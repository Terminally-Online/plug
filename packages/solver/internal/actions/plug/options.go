package plug

import (
	"solver/internal/actions"
	// "solver/internal/helpers/zerion"

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
	// var options []actions.Option
	//
	// chainIds := []uint64{chainId}
	// holdings, err := zerion.GetFungibleHoldings(chainIds, from)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
