package engine

import (
	"time"
)

// Collector interface
type Collector interface {
	GetKey() string
	GetCollectionStream(stream chan<- Collection) error
}

// Executor interface
type Executor interface {
	GetKey() string
	Execute(execution interface{}) error
}

// Process interface
type Process interface {
	GetKey() string
	SyncState() error
	ProcessCollection(key string, collection interface{}) (ExecutionResult, error)
}

// Collection represents the data collected
type Collection struct {
	Key   string
	Data  interface{}
}

// ExecutionResult represents the result of processing a collection
type ExecutionResult struct {
	Key       string
	Execution interface{}
}

// Network represents the network of collectors and executors
type Network struct {
	Collectors []Collector
	Executors  []Executor
}

// Engine structure
type Engine struct {
	Network   Network
	Processes map[string]Process
	Stream    chan Collection
	Retries   int
	Delay     time.Duration
}

