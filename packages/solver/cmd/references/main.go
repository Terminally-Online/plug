package main

import (
	"fmt"
	"log"
	"os"
	"solver/internal/references"
	"solver/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
	}

	fmt.Println("Starting contract reference generation...")
	if err := references.GenerateReferences(); err != nil {
		fmt.Printf("Error generating bindings: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All references generated successfully")
}
