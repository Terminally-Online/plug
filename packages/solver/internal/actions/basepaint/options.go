package basepaint

import (
	"solver/internal/actions"
	"solver/internal/actions/options"

	"github.com/ethereum/go-ethereum/common"
)

func MintLatestOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	return nil, nil
}

func TransferOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	recipientIndex := 1
	recipientOptions, err := options.GetAddressOptions(chainId, from, search[recipientIndex])
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		recipientIndex: {Simple: recipientOptions},
	}, nil
}
