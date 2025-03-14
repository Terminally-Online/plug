package assert

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

func AssertOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	booleanOptions := []actions.Option{
		{Label: "True", Value: "true"},
		{Label: "False", Value: "false"},
	}

	return map[int]actions.Options{
		1: {Simple: booleanOptions},
	}, nil
}
