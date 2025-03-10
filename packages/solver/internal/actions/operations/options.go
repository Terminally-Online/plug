package operations

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type OperationsOptionsProvider struct{}

func (p *OperationsOptionsProvider) GetOptions(chainId uint64, from common.Address, search map[int]string, action string) (map[int]actions.Options, error) {
	switch action {
	}
	return nil, nil
}
