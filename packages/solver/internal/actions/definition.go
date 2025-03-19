package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/coil"
	"solver/internal/solver/signature"
)

type ActionFunc[T any] func(lookup *SchemaLookup[T]) ([]signature.Plug, error)
type ActionOptionsFunc[T any] func(lookup *SchemaLookup[T]) (map[int]Options, error)


type ActionDefinitionInterface interface {
	GetType() string
	GetSentence() string
	GetIsUserSpecific() bool
	GetHandler() ActionFunc[any]
	GetOptions() ActionOptionsFunc[any]
	GetCoils() ([]coil.Update, error)
	GetCoilKeys() (map[string]string, error)
}

type ActionDefinition[T any] struct {
	ActionDefinitionInterface
	Type           string `default:"action,omitempty"`
	Sentence       string
	Handler        ActionFunc[T]
	Options        ActionOptionsFunc[T]
	IsUserSpecific bool
	IsSearchable   bool
	Response       *ActionOnchainFunctionResponse
}

func NewActionDefinition[T any](
	sentence string,
	action ActionFunc[T],
	options ActionOptionsFunc[T],
	isUserSpecific bool,
	isSearchable bool,
	response *ActionOnchainFunctionResponse,
) ActionDefinitionInterface {
	return &ActionDefinition[T]{
		Sentence:       sentence,
		Handler:        action,
		Options:        options,
		IsUserSpecific: isUserSpecific,
		IsSearchable:   isSearchable,
		Response:       response,
	}
}

func (d *ActionDefinition[T]) GetType() string {
	return d.Type
}

func (d *ActionDefinition[T]) GetSentence() string {
	return d.Sentence
}

func (d *ActionDefinition[T]) GetIsUserSpecific() bool {
	return d.IsUserSpecific
}

func (d *ActionDefinition[T]) GetHandler() ActionFunc[any] {
	return func(lookup *SchemaLookup[any]) ([]signature.Plug, error) {
		typedLookup := &SchemaLookup[T]{
			ChainId: lookup.ChainId,
			Client:  lookup.Client,
			From:    lookup.From,
			Search:  lookup.Search,
		}

		inputsJSON, err := json.Marshal(lookup.Inputs)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal inputs: %w", err)
		}

		var typedInputs T
		if err := json.Unmarshal(inputsJSON, &typedInputs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal inputs to target type: %w", err)
		}

		typedLookup.Inputs = &typedInputs
		return d.Handler(typedLookup)
	}
}

func (d *ActionDefinition[T]) GetOptions() ActionOptionsFunc[any] {
	return func(lookup *SchemaLookup[any]) (map[int]Options, error) {
		return d.Options(&SchemaLookup[T]{
			ChainId: lookup.ChainId,
			Client:  lookup.Client,
			From:    lookup.From,
			Search:  lookup.Search,
		})
	}
}

func (d *ActionDefinition[T]) GetCoils() ([]coil.Update, error) {
	if d.Response == nil || d.Response.Metadata == nil || d.Response.FunctionName == "" {
		return nil, nil
	}

	abi, err := d.Response.Metadata.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ABI: %w", err)
	}

	coils, err := coil.FindCoils(abi, d.Response.FunctionName, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find coils: %w", err)
	}

	return coils, nil
}

func (d *ActionDefinition[T]) GetCoilKeys() (map[string]string, error) {
	coils, err := d.GetCoils()
	if err != nil {
		return nil, err
	}

	coilKeys := make(map[string]string)
	for _, coil := range coils {
		name := *coil.Slice.Name
		coilKeys[name] = coil.Slice.Type
	}

	return coilKeys, nil
}
