package executors

import (
	"fmt"
)

type ExampleExecutor struct {
	key string
}

func (e *ExampleExecutor) GetKey() string {
	return e.key
}

func (e *ExampleExecutor) Execute(execution interface{}) error {
	fmt.Printf("Executing %v for %s\n", execution, e.key)
	return nil
}

func NewExampleExecutor(key string) (*ExampleExecutor) {
	return &ExampleExecutor{key: key}
}
