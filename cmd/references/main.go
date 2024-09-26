package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"solver/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
	}

	fmt.Println("Starting contract reference generation...")
	if err := utils.GenerateReferences(); err != nil {
		fmt.Printf("Error generating bindings: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All references generated successfully")
}
