package call

import (
	"solver/internal/client"
)

type Caller struct{}

func New() Caller {
	return Caller{}
}

func (c *Caller) Call() error {
	return nil
}

func Write(chainId uint64) (error, error) {
	_, err := client.New(chainId)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
