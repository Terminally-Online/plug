package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GenerateBindings is the main function to generate bindings for all contracts
func GenerateBindings() error {
	// Ensure the required directories exist
	if err := ensureDir("abigenBindings/bin"); err != nil {
		return err
	}
	if err := ensureDir("abigenBindings/abi"); err != nil {
		return err
	}
	if err := ensureDir("bindings"); err != nil {
		return err
	}

	// Get all contract files from the contracts directory
	contracts, err := filepath.Glob("contracts/*.sol")
	if err != nil {
		return fmt.Errorf("error reading contracts: %v", err)
	}

	for _, contract := range contracts {
		if err := generateBinding(contract); err != nil {
			return err
		}
	}

	return nil
}

// RunBindingGeneration is a wrapper function that can be called from main or other packages
func RunBindingGeneration() {
	if err := GenerateBindings(); err != nil {
		fmt.Printf("Error generating bindings: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All bindings generated successfully")
}

func generateBinding(contractPath string) error {
	contractName := strings.TrimSuffix(filepath.Base(contractPath), ".sol")
	lowerContractName := strings.ToLower(contractName)
	outputDirectory := fmt.Sprintf("bindings/%s", lowerContractName)

	if err := ensureDir(outputDirectory); err != nil {
		return err
	}
	
	cmd := exec.Command("abigen",
		"--bin", fmt.Sprintf("abigenBindings/bin/%s.bin", contractName),
		"--abi", fmt.Sprintf("abigenBindings/abi/%s.abi", contractName),
		"--pkg", lowerContractName,
		"--out", fmt.Sprintf("%s/%s.go", outputDirectory, lowerContractName))

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error generating binding for %s: %v\n%s", contractName, err, string(output))
	}
	
	fmt.Printf("Successfully generated binding for %s\n", contractName)
	return nil
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %v", dirName, err)
	}
	return nil
}
