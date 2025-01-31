package euler

import (
	"solver/internal/actions"
)

type EulerOptionsProvider struct{}

func (p *EulerOptionsProvider) GetOptions(chainId uint64, action string) (map[int]actions.Options, error) {
	switch action {
	default:
		return nil, nil
	}
}