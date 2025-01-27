package plug

import (
	"solver/internal/actions"
)

type PlugOptionsProvider struct{}

func (p *PlugOptionsProvider) GetOptions(chainId uint64, action string) (map[int]actions.Options, error) {
	return nil, nil
}
