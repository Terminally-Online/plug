package services

import (
	"context"
	"solver/bindings/othentic_attestation"
	"solver/internal/avs/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getActiveOperators(client *ethclient.Client) ([]othentic_attestation.Struct0, error) {

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

func ElectRoundRobin() (*common.Address, error) {
	client, err := ethclient.Dial(config.Node_8453)
	if err != nil {
		return nil, err
	}

	operators, err := getActiveOperators(client)
	if err != nil {
		return nil, err
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	index := blockNumber % uint64(len(operators))

	return &operators[index].Operator, nil

}

func IsElectedLeader(operatorAddress string) bool {
	electedLeader, err := ElectRoundRobin()
	if err != nil {
		return false
	}

	return *electedLeader == common.HexToAddress(operatorAddress)
}
