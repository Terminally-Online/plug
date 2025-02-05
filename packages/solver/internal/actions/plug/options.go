package plug

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type PlugOptionsProvider struct{}

func (p *PlugOptionsProvider) GetOptions(chainId uint64, _ common.Address, action string) (map[int]actions.Options, error) {
	return nil, nil
}
