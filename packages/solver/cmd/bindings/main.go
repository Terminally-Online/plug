package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %v", dirName, err)
	}
	return nil
}

func GenerateBindings() error {
	if err := ensureDir("bindings"); err != nil {
		return err
	}

	var abiFiles []string
	err := filepath.Walk("abis", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			abiFiles = append(abiFiles, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking through references directory: %v", err)
	}

	for _, abiPath := range abiFiles {
		contractName := strings.TrimSuffix(filepath.Base(abiPath), ".json")
		folderName := filepath.Base(filepath.Dir(abiPath))
		packageName := fmt.Sprintf("%s_%s", strings.ToLower(folderName), contractName)
		outputDirectory := filepath.Join("bindings", packageName)
		if err := ensureDir(outputDirectory); err != nil {
			return err
		}
		outputFile := filepath.Join(outputDirectory, packageName+".go")
		cmd := exec.Command("abigen",
			"--abi", abiPath,
			"--pkg", packageName,
			"--out", outputFile)

		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error generating binding for %s/%s: %v\n%s", folderName, contractName, err, string(output))
		}
		fmt.Printf("Successfully generated binding for %s/%s\n", folderName, contractName)
	}

	return nil
}

func main() {
	fmt.Println("Starting contract binding generation...")
	if err := GenerateBindings(); err != nil {
		fmt.Printf("Error generating bindings: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All bindings generated successfully")
}
