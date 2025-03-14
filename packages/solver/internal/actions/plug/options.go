package plug

import (
	"solver/internal/actions"
	"solver/internal/actions/options"

	"github.com/ethereum/go-ethereum/common"
)

func TransferOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	fungiblesIndex := 1
	fungiblesOptions, err := options.GetFungiblesAndFungiblesHeldOptions(chainId, from, search[fungiblesIndex])
	if err != nil {
		return nil, err
	}

	recipientIndex := 2
	recipientOptions, err := options.GetAddressOptions(chainId, from, search[recipientIndex])
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		fungiblesIndex: {Simple: fungiblesOptions},
		recipientIndex: {Simple: recipientOptions},
	}, nil
}

func SwapOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	fungiblesOutIndex := 1
	fungiblesOutOptions, err := options.GetFungiblesAndFungiblesHeldOptions(chainId, from, search[fungiblesOutIndex])
	if err != nil {
		return nil, err
	}

	fungiblesInIndex := 2
	fungiblesInOptions, err := options.GetFungiblesAndFungiblesHeldOptions(chainId, from, search[fungiblesInIndex])
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		fungiblesOutIndex: {Simple: fungiblesOutOptions},
		fungiblesInIndex:  {Simple: fungiblesInOptions},
	}, nil
}

func PriceOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	fungiblesIndex := 0
	fungiblesOptions, err := options.GetFungiblesAndFungiblesHeldOptions(chainId, from, search[fungiblesIndex])
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		fungiblesIndex: {Simple: fungiblesOptions},
		1:              {Simple: actions.BaseThresholdFields},
	}, nil
}

func BalanceOptions(chainId uint64, from common.Address, search map[int]string, _ string) (map[int]actions.Options, error) {
	fungiblesIndex := 0
	fungiblesOptions, err := options.GetFungiblesAndFungiblesHeldOptions(chainId, from, search[fungiblesIndex])
	if err != nil {
		return nil, err
	}

	addressIndex := 1
	addressOptions, err := options.GetAddressOptions(chainId, from, search[addressIndex])
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		fungiblesIndex: {Simple: fungiblesOptions},
		addressIndex:   {Simple: addressOptions},
		2:              {Simple: actions.BaseThresholdFields},
	}, nil
}
