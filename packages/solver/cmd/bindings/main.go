package main

import (
	"fmt"
	"os"
	"solver/utils"
)

func main() {
	fmt.Println("Starting contract binding generation...")
	if err := utils.GenerateBindings(); err != nil {
		fmt.Printf("Error generating bindings: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All bindings generated successfully")
}
