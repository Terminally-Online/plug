package engine

import (
	"fmt"
	"time"
	"sync"
)

// NewEngine creates a new Engine instance
func NewEngine(network Network, processes map[string]Process, retries int, delay time.Duration) *Engine {
	e := &Engine{
		Network:   network,
		Processes: processes,
		Stream:    make(chan Collection),
		Retries:   retries,
		Delay:     delay,
	}

	fmt.Printf("Engine initialized with %d collectors, %d executors, and %d processes.\n",
		len(network.Collectors), len(network.Executors), len(processes))

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
	var wg sync.WaitGroup

	// Start executors
	executorChannels := make(map[string]chan interface{})
	for _, executor := range e.Network.Executors {
		executorChan := make(chan interface{})
		executorChannels[executor.GetKey()] = executorChan
		wg.Add(1)
		go func(exec Executor, ch <-chan interface{}) {
			defer wg.Done()
			for action := range ch {
				if err := exec.Execute(action); err != nil {
					fmt.Printf("Error executing action for %s: %v\n", exec.GetKey(), err)
				}
			}
		}(executor, executorChan)
	}

	// Start processes
	for name, process := range e.Processes {
		fmt.Printf("Running Process: %s\n", name)

		if err := process.SyncState(); err != nil {
			fmt.Printf("Error syncing state for process %s: %v\n", name, err)
		}

		wg.Add(1)
		go func(proc Process) {
			defer wg.Done()
			for collection := range e.Stream {
				execution, err := proc.ProcessCollection(collection.Key, collection.Data)
				if err != nil {
					fmt.Printf("Error processing collection for %s: %v\n", proc.GetKey(), err)
					continue
				}
				if execution.Key != "" {
					if ch, ok := executorChannels[execution.Key]; ok {
						ch <- execution.Execution
					} else {
						fmt.Printf("No executor found for key: %s\n", execution.Key)
					}
				}
			}
		}(process)
	}

	// Start collectors
	for _, collector := range e.Network.Collectors {
		wg.Add(1)
		go func(coll Collector) {
			defer wg.Done()
			if err := coll.GetCollectionStream(e.Stream); err != nil {
				fmt.Printf("Error in collector %s: %v\n", coll.GetKey(), err)
			}
		}(collector)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return nil
}
