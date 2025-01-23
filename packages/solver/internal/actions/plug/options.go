package plug

import (
	"solver/internal/actions"
)

type PlugOptionsProvider struct{}

func (p *PlugOptionsProvider) GetOptions(chainId int, action string) (map[int]actions.Options, error) {
	return nil, nil
}
