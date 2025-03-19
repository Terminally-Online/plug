package coil

import (
	"encoding/json"
	"strings"
)

var (
	CoilPrefix = "<-{"
	CoilSuffix = "}"
)

type CoilInputInterface interface {
	GetIsLinked() bool
}

type CoilInput[T any, R any] struct {
	CoilInputInterface
	raw     string
	decoded T
}

func (c *CoilInput[T, R]) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return json.Unmarshal(data, &c.decoded)
	}
	c.raw = raw
	trimmed := strings.TrimPrefix(strings.TrimSuffix(raw, CoilSuffix), CoilPrefix)
	return json.Unmarshal([]byte(trimmed), &c.decoded)
}

func (c *CoilInput[T, R]) GetIsLinked() bool {
	return strings.HasPrefix(c.raw, CoilPrefix)
}

func (c *CoilInput[T, R]) GetValue() T {
	return c.decoded
}

func (c *CoilInput[T, R]) Get(valueFunc func() (R, error)) (R, error) {
	response, err := valueFunc()

	return response, err
}

type ValueFunc[R any] func() (R, error)

type FunctionResponseInterface interface { 
    GetUpdate(string) (*Update, error) 
}

func (c *CoilInput[T, R]) GetAndUpdate(valueFunc ValueFunc[R], coilFunc FunctionResponseInterface, param string, updates []Update) (R, []Update, error) {
	response, err := valueFunc()
	if err != nil {
		return response, nil, err

	}

	if update, err := coilFunc.GetUpdate(param); update != nil {
		updates = append(updates, *update)
		return response, updates, err
	} else if err != nil {
		return response, nil, err
	}

	return response, nil, err
}
