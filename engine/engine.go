package engine

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// NewEngine creates a new Engine instance
func NewEngine(networks map[string]Network, retries int, delay time.Duration) *Engine {
	e := &Engine{
		Networks: networks,
		Stream:   make(chan Collection),
		Retries:  retries,
		Delay:    delay,
	}

	totalCollectors := 0
	totalExecutors := 0
	totalProcesses := 0
	for networkName, network := range networks {
		totalCollectors += len(network.Collectors)
		totalExecutors += len(network.Executors)
		totalProcesses += len(network.Processes)
		fmt.Printf("Engine network %s with %d collectors, %d executors, and %d processes.\n",
			networkName, len(network.Collectors), len(network.Executors), len(network.Processes))
	}

	fmt.Printf("Engine initialized with %d networks, %d collectors, %d executors, and %d processes.\n",
		len(networks), totalCollectors, totalExecutors, totalProcesses)

	return e
}

// Restart handles engine restart on error
func (e *Engine) Restart(err error) {
	fmt.Printf("Engine failed while running: %v\n", err)

	if e.Retries == 0 {
		fmt.Println("Engine has no more retries left. Exiting process.")
		return
	}

	fmt.Printf("Engine will retry in %v seconds: %d attempts left.\n", e.Delay.Seconds(), e.Retries)

	e.Retries--

	time.AfterFunc(e.Delay, func() {
		err := e.Run()
		if err != nil {
			e.Restart(err)
		}
	})
}

// Run starts the engine
func (e *Engine) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Manage the state of all the networks we are running on
	for networkName, network := range e.Networks {
		// Start executors
		executorChannels := make(map[string]chan interface{})
		for _, executor := range network.Executors {
			executorChan := make(chan interface{})
			executorChannels[executor.GetKey()] = executorChan
			wg.Add(1)
			go func(net string, exec Executor, ch <-chan interface{}) {
				defer wg.Done()
				for action := range ch {
					if err := exec.Execute(action); err != nil {
						fmt.Printf("Error executing action for %s on network %s: %v\n", exec.GetKey(), net, err)
					}
				}
			}(networkName, executor, executorChan)
		}

		// Start processes
		for _, process := range network.Processes {
			fmt.Printf("Running Process: %s on network %s\n", process.GetKey(), networkName)

			if err := process.SyncState(); err != nil {
				fmt.Printf("Error syncing state for process %s on network %s: %v\n", process.GetKey(), networkName, err)
			}

			wg.Add(1)
			go func(net string, proc Process) {
				defer wg.Done()
				for collection := range e.Stream {
					if collection.NetworkName != net { continue }
					execution, err := proc.ProcessCollection(collection.Key, collection.Data)
					if err != nil {
						fmt.Printf("Error processing collection for %s on network %s: %v\n", proc.GetKey(), net, err)
						continue
					}
					if execution.Key != "" {
						if ch, ok := executorChannels[execution.Key]; ok {
							ch <- execution.Execution
						} else {
							fmt.Printf("No executor found for key: %s on network %s\n", execution.Key, net)
						}
					}
				}
			}(networkName, process)
		}

		// Start collectors
		for _, collector := range network.Collectors {
			wg.Add(1)
			go func(net string, coll Collector) {
				defer wg.Done()
				if err := coll.GetCollectionStream(ctx, net, e.Stream); err != nil {
					fmt.Printf("Error in collector %s on network %s: %v\n", coll.GetKey(), net, err)
				}
			}(networkName, collector)
		}
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return nil
}
