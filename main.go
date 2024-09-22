package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"solver/engine"
	"solver/engine/collectors"
	"solver/engine/executors"
	"solver/engine/processes"
	"time"
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
	alchemyURL := fmt.Sprintf("wss://eth-mainnet.g.alchemy.com/v2/%v", alchemyAPIKey)

	blockCollector, err := collectors.NewBlockCollector("BlockCollector", alchemyURL)
	if err != nil {
		log.Fatalf("Failed to create BlockCollector: %v", err)
	}

	networks := make(map[string]engine.Network)
	networks["ethereum"] = engine.Network{
		Collectors: []engine.Collector{
			blockCollector,
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
