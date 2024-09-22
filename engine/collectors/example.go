package collectors

import (
	"fmt"
	"time"
	"solver/engine"
)

type ExampleCollector struct {
	key string
}

func (c *ExampleCollector) GetKey() string {
	return c.key
}

func (c *ExampleCollector) GetCollectionStream(stream chan<- engine.Collection) error {
	// Simulating collection of data
	go func() {
		for i := 0; ; i++ {
			stream <- engine.Collection{
				Key:  c.key,
				Data: fmt.Sprintf("Data %d from %s", i, c.key),
			}
			time.Sleep(time.Second)
		}
	}()
	return nil
}

func NewExampleCollector(key string) (*ExampleCollector) {
	return &ExampleCollector{key: key}
}
