package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"solver/internal/avs/config"
	"solver/internal/avs/routes"
	"solver/internal/avs/streams"

	"github.com/gin-gonic/gin"
)

func info() {
	_, address, err := config.GetAccount()
	if err != nil {
		log.Fatalf("Failed to get operator address: %v", err)
	}
	fmt.Printf("Operator Address: %s\n", address)
	fmt.Printf("Chain ID: %d\n", config.ChainId)
	fmt.Printf("Environment: %s\n", func() string {
		if config.Production {
			return "production"
		}
		return "development"
	}())
}

func help() {
	fmt.Println("Circuit AVS - Attestation Virtual Service for Plug")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("  start      Start the AVS service (default)")
	fmt.Println("  version    Display version information")
	fmt.Println("  info       Display operator information")
	fmt.Println("  help       Display this help message")
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			fmt.Println("Circuit AVS v0.1.0")
			return
		case "info":
			info()
			return
		case "help":
			help()
			return
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go streams.HandleStream(ctx)
	
	router := gin.Default()

	router.POST("task/validate", routes.ValidateTask)

	router.GET("status", routes.Status)
	router.GET("metrics", routes.Metrics)

	log.Printf("Circuit AVS started on port %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
