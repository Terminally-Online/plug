package math

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type MathOptionsProvider struct{}

func (m *MathOptionsProvider) GetOptions(chainId uint64, _ common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	switch action {
	case Calculate:
		return map[int]actions.Options{
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
		}, nil
	}

	return nil, nil
}
