package processes

import (
	"fmt"
	"solver/engine"
)

type ExampleProcess struct {
	key string
}

func (p *ExampleProcess) GetKey() string {
	return p.key
}

func (p *ExampleProcess) SyncState() error {
	fmt.Printf("Syncing state for process %s\n", p.key)
	return nil
}

func (p *ExampleProcess) ProcessCollection(key string, collection interface{}) (engine.ExecutionResult, error) {
	fmt.Printf("Processing collection from %s: %v\n", key, collection)
	return engine.ExecutionResult{
		Key:       "ExampleExecutor",
		Execution: fmt.Sprintf("Processed %v", collection),
	}, nil
}

func NewExampleProcess(key string) (*ExampleProcess) {
	return &ExampleProcess{key: key}
}
