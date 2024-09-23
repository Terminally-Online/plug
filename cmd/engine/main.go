package main

import (
	"fmt"
	"log"
	"os"
	"solver/engine"
	"solver/engine/collectors"
	"solver/engine/executors"
	"solver/engine/processes"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	alchemyAPIKey := os.Getenv("ALCHEMY_API_KEY")
	if alchemyAPIKey == "" {
		log.Fatal("ALCHEMY_API_KEY environment variable not set")
	}
	ethClient, err := ethclient.Dial(fmt.Sprintf("wss://eth-mainnet.g.alchemy.com/v2/%v", alchemyAPIKey))
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	transferEventSignature := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

	blockCollector := collectors.NewBlockCollector("BlockCollector", ethClient)
	eventCollector := collectors.NewEventCollector(
		"EventCollector",
		ethClient,
		[]collectors.EventFilter{
			{
				Topics: []common.Hash{transferEventSignature},
			},
		},
	)

	networks := make(map[string]engine.Network)
	networks["ethereum"] = engine.Network{
		Collectors: []engine.Collector{
			blockCollector,
			eventCollector,
		},
		Executors: []engine.Executor{
			executors.NewExampleExecutor("ExampleExecutor"),
		},
		Processes: []engine.Process{
			processes.NewExampleProcess("ExampleProcess"),
		},
	}

	eng := engine.NewEngine(networks, 3, time.Second)

	fmt.Println("Starting the engine...")
	if err := eng.Run(); err != nil {
		fmt.Printf("Engine error: %v\n", err)
		eng.Restart(err)
	}
}
