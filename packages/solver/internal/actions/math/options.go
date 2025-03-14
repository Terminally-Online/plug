package math

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

var (
	operationOptions = map[int]actions.Options{
		1: {
			Simple: []actions.Option{
				{
					Name:  "Add",
					Label: "+",
					Value: "+",
				},
				{
					Name:  "Subtract",
					Label: "-",
					Value: "-",
				},
				{
					Name:  "Multiply",
					Label: "*",
					Value: "*",
				},
				{
					Name:  "Divide",
					Label: "÷",
					Value: "÷",
				},
				{
					Name:  "Modulo",
					Label: "%",
					Value: "%",
				},
			},
		},
	}
)

func CalculateOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	return operationOptions, nil
}
