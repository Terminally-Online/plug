package engine

import (
	"context"
	"time"
)

// Collector interface
type Collector interface {
	GetKey() string
	GetCollectionStream(ctx context.Context, networkName string, stream chan<- Collection) error
}

// Collection represents the data collected
type Collection struct {
	NetworkName string
	Key         string
	Data        interface{}
}

// Executor interface
type Executor interface {
	GetKey() string
	Execute(execution interface{}) error
}

// ExecutionResult represents the result of processing a collection
type ExecutionResult struct {
	Key       string
	Execution interface{}
}

// Process interface
type Process interface {
	GetKey() string
	SyncState() error
	ProcessCollection(key string, collection interface{}) (ExecutionResult, error)
}

// Network represents the network of collectors and executors
type Network struct {
	Collectors []Collector
	Executors  []Executor
	Processes  []Process
}

// Engine structure
type Engine struct {
	Networks map[string]Network
	Stream   chan Collection
	Retries  int
	Delay    time.Duration
}
