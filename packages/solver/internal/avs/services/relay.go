package services

import (
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

func RelayTask(response types.ExecutionResponse) {
	performerPrivateKey, performerAddress, err := config.GetAccount()
	if err != nil {
		log.Println("error occurred while getting account", err)
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
		log.Println("error occurred while encoding", err)
	}
	hash := crypto.Keccak256Hash(packed)
	signature, err := crypto.Sign(hash.Bytes(), performerPrivateKey)
	if err != nil {
		log.Println("error occurred while signing", err)
	}
	signature[64] += 27
	serialized := hexutil.Encode(signature)

	nodeUrl, err := config.GetNodeUrl(config.ChainId)
	if err != nil {
		log.Println("error occurred while calculating rpc url", err)
	}
	client, err := rpc.Dial(nodeUrl)
	if err != nil {
		log.Println("error occurred while connecting to RPC client", err)
	}
	defer client.Close()

	var result any
	if err := client.Call(
		&result, "sendTask",
		response.ProofOfTask,
		response.Data,
		response.TaskDefinitionId,
		performerAddress,
		serialized,
	); err != nil {
		log.Println("error occurred while sending task", err)
	}
}
