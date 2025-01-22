package plug

import "solver/types"

type PlugOptionsProvider struct{}

func (p *PlugOptionsProvider) GetOptions(chainId int, action string) (map[int]types.SchemaOptions, error) {
	return nil, nil
}
