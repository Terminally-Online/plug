package assert

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type AssertOptionsProvider struct{}

func (a *AssertOptionsProvider) GetOptions(chainId uint64, _ common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	booleanOptions := []actions.Option{
		{Label: "True", Value: "true"},
		{Label: "False", Value: "false"},
	}

	switch action {
	case AssertTrue:
		return map[int]actions.Options{
			0: {Simple: booleanOptions},
		}, nil
	case AssertFalse:
		return map[int]actions.Options{
			0: {Simple: booleanOptions},
		}, nil
	default:
		return nil, nil
	}
}
