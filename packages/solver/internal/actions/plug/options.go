package plug

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type PlugOptionsProvider struct{}

func (p *PlugOptionsProvider) GetOptions(chainId uint64, from common.Address, action string) (map[int]actions.Options, error) {
	switch action {
	case actions.ActionTransfer:
		return map[int]actions.Options{
			1: {Simple: []actions.Option{{
				Label: from.Hex(),
			}}},
		}, nil
	default:
		return nil, nil
	}
}
