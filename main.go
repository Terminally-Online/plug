package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"solver/engine"
	"solver/engine/collectors"
	"solver/engine/executors"
	"solver/engine/processes"
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

	nodeURL := fmt.Sprintf("wss://eth-mainnet.g.alchemy.com/v2/%v", alchemyAPIKey)
	blockCollector, err := collectors.NewBlockCollector("BlockCollector", nodeURL)
	if err != nil { 
		log.Fatalf("Failed to create BlockCollector: %v", err)
	}

	network := engine.Network{
		Collectors: []engine.Collector{
			blockCollector,
		},
		Executors: []engine.Executor{
			executors.NewExampleExecutor("ExampleExecutor"),
		},
	}

	processes := map[string]engine.Process{
		"ExampleProcess": processes.NewExampleProcess("ExampleProcess"),
	}

	eng := engine.NewEngine(network, processes, 3, time.Second)

	fmt.Println("Starting the engine...")
	if err := eng.Run(); err != nil {
		fmt.Printf("Engine error: %v\n", err)
		eng.Restart(err)
	}
}
