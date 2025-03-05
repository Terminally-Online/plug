package database

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type DatabaseOptionsProvider struct{}

func (d *DatabaseOptionsProvider) GetOptions(chainId uint64, _ common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	// NOTE: For the database operations, we don't need predefined options as the inputs are free-form strings
	//       and addresses. However, in the future, we might want to provide common keys as suggestions
	//       so I am leaving it here for now. If you are ever not sure of what to work on this
	//       is a nice random ticket to pick up. - CHANCE
	switch action {
	case SetValue:
		return nil, nil
	case GetValue:
		return nil, nil
	case RemoveValue:
		return nil, nil
	default:
		return nil, nil
	}
}
