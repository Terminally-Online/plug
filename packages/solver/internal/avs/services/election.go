package services

import (
	"context"
	"solver/bindings/othentic_attestation"
	"solver/internal/avs/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getActiveOperators() ([]othentic_attestation.Struct0, error) {
	client, err := ethclient.Dial(config.Node_8453)
	if err != nil {
		return nil, err
	}

	attestationCenter, err := othentic_attestation.NewOthenticAttestation(
		config.AttestationCenter,
		client,
	)
	if err != nil {
		return nil, err
	}

	_, performerAddress, err := config.GetAccount()
	if err != nil {
		return nil, err
	}

	return attestationCenter.GetActiveOperatorsDetails(&bind.CallOpts{
		From:    common.HexToAddress(performerAddress),
		Pending: true,
		Context: context.Background(),
	})
}

func ElectRoundRobin(blockNumber uint64) (*common.Address, error) {
	operators, err := getActiveOperators()
	if err != nil {
		return nil, err
	}

	index := blockNumber % uint64(len(operators))

	return &operators[index].Operator, nil
}
