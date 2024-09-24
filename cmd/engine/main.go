package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"solver/engine"
	"solver/engine/collectors"
	"solver/engine/executors"
	"solver/engine/processes"
	"solver/utils"
	"time"
)

func main() {
	ethClient, err := utils.GetProvider(1)
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
