package database

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type DatabaseOptionsProvider struct{}

func (d *DatabaseOptionsProvider) GetOptions(chainId uint64, _ common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	// For the database operations, we don't need predefined options as the inputs are free-form strings
	// and addresses. However, in the future, we might want to provide common keys as suggestions.

	switch action {
	case SetValue:
		// No predefined options for key or value fields
		return nil, nil
	case GetValue:
		// No predefined options for key or address fields
		return nil, nil
	case RemoveValue:
		// No predefined options for key field
		return nil, nil
	default:
		return nil, nil
	}
}