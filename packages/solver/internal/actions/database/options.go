package database

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type DatabaseOptionsProvider struct{}

func (d *DatabaseOptionsProvider) GetOptions(chainId uint64, _ common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	return nil, nil
}
