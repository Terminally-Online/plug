package services

import (
	"fmt"
	"log"
	"math/big"
	"solver/internal/avs/config"
	"solver/internal/avs/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

func RelayTask(response types.ExecutionResponse) error {
	performerPrivateKey, performerAddress, err := config.GetAccount()
	if err != nil {
		log.Println("Error occurred while getting account:", err)
		return fmt.Errorf("failed to get account: %w", err)
	}
	
	arguments := abi.Arguments{
		{Type: abi.Type{T: abi.StringTy}},
		{Type: abi.Type{T: abi.BytesTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy}},
	}
	packed, err := arguments.Pack(
		response.ProofOfTask,
		response.Data,
		common.HexToAddress(performerAddress),
		big.NewInt(int64(response.TaskDefinitionId)),
	)
	if err != nil {
		log.Println("Error occurred while encoding:", err)
		return fmt.Errorf("failed to encode parameters: %w", err)
	}
	
	hash := crypto.Keccak256Hash(packed)
	signature, err := crypto.Sign(hash.Bytes(), performerPrivateKey)
	if err != nil {
		log.Println("Error occurred while signing:", err)
		return fmt.Errorf("failed to sign message: %w", err)
	}
	
	signature[64] += 27
	serialized := hexutil.Encode(signature)

	nodeUrl, err := config.GetNodeUrl(config.ChainId)
	if err != nil {
		log.Println("Error occurred while getting RPC URL:", err)
		return fmt.Errorf("failed to get RPC URL: %w", err)
	}
	
	client, err := rpc.Dial(nodeUrl)
	if err != nil {
		log.Println("Error occurred while connecting to RPC client:", err)
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer client.Close()

	log.Printf("Relaying task with proof: %s and task definition ID: %d", 
		response.ProofOfTask, response.TaskDefinitionId)
	
	var result any
	if err := client.Call(
		&result, "sendTask",
		response.ProofOfTask,
		response.Data,
		response.TaskDefinitionId,
		performerAddress,
		serialized,
	); err != nil {
		log.Println("Error occurred while sending task:", err)
		return fmt.Errorf("failed to send task: %w", err)
	}
	
	log.Printf("Successfully relayed task to attestation center")
	return nil
}
