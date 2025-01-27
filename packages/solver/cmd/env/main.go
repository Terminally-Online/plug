package main

import (
	"fmt"
	"os"
	"solver/internal/env"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [encrypt|decrypt]")
		return
	}
	command := os.Args[1]

	abort := env.Init()
	if abort {
		fmt.Println("Skipping encryption and decryption.")
		return
	}

	var err error
	switch command {
	case "encrypt":
		err = env.EncryptEnvFile()
	case "decrypt":
		err = env.DecryptEnvFile()
	default:
		fmt.Println("Usage: go run main.go [encrypt|decrypt]")
		return
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
