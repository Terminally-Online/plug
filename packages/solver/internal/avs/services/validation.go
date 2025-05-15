package services

import (
	"context"
	"fmt"
	"log"
	"solver/bindings/plug_router"
	"solver/internal/avs/config"
	"solver/internal/bindings/references"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func Validate(proofOfTask string, data string) (bool, error) {
	nodeUrl, err := config.GetNodeUrl(config.ChainId)
	if err != nil {
		log.Println("error occurred while calculating rpc url", err)
	}
	dial, err := rpc.Dial(nodeUrl)
	if err != nil {
		log.Println("error occurred while connecting to RPC client", err)
	}
	defer dial.Close()

	client := ethclient.NewClient(dial)

	transactionHash := common.HexToHash(proofOfTask)
	receipt, err := client.TransactionReceipt(context.Background(), transactionHash)
	if err != nil {
		return false, err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return false, fmt.Errorf("transaction failed onchain")
	}

	if (receipt.ContractAddress == common.Address{}) {
		_, isPending, err := client.TransactionByHash(context.Background(), transactionHash)
		if err != nil {
			return false, err
		}
		if isPending {
			return false, fmt.Errorf("transaction is still pending")
		}
	}

	routerAddress := common.HexToAddress(references.Plug["router"])
	filterer, err := plug_router.NewPlugRouterFilterer(
		routerAddress,
		client,
	)
	if err != nil {
		return false, err
	}

	var found bool
	for _, log := range receipt.Logs {
		if log.Address != routerAddress {
			continue
		}

		plugResults, err := filterer.ParsePlugResult(*log)
		if err != nil || plugResults == nil {
			continue
		}

		if common.HexToHash(data) == plugResults.PlugsHash {
			found = true
			break
		}
	}

	return found, nil
}
