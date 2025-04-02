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
	GetCoilSlice(string) (*coil.Slice, error)
	GetCoilSlices() ([]coil.Slice, error)
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
			ChainId:                  lookup.ChainId,
			Client:                   lookup.Client,
			From:                     lookup.From,
			Search:                   lookup.Search,
			PreviousActionDefinition: lookup.PreviousActionDefinition,
		}

		inputsJSON, err := json.Marshal(lookup.Inputs)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal inputs: %w", err)
		}

		var typedInputs T
		if err := json.Unmarshal(inputsJSON, &typedInputs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal inputs to target type: %w (input: %s)", err, string(inputsJSON))
		}

		typedLookup.Inputs = &typedInputs
		return d.Handler(typedLookup)
	}
}

func (d *ActionDefinition[T]) GetOptions() ActionOptionsFunc[any] {
	if d.Options == nil {
		return nil
	}
	return func(lookup *SchemaLookup[any]) (map[int]Options, error) {
		return d.Options(&SchemaLookup[T]{
			ChainId: lookup.ChainId,
			Client:  lookup.Client,
			From:    lookup.From,
			Search:  lookup.Search,
		})
	}
}

func (d *ActionDefinition[T]) GetCoilSlices() ([]coil.Slice, error) {
	if d.Response == nil || d.Response.Metadata == nil || d.Response.FunctionName == "" {
		return nil, nil
	}

	abi, err := d.Response.Metadata.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ABI: %w", err)
	}

	slices, err := coil.GetCoilSlices(abi, d.Response.FunctionName, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find coils: %w", err)
	}

	return slices, nil
}

func (d *ActionDefinition[T]) GetCoilSlice(name string) (*coil.Slice, error) {
	if d.Response == nil || d.Response.Metadata == nil || d.Response.FunctionName == "" {
		return nil, nil
	}

	abi, err := d.Response.Metadata.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ABI: %w", err)
	}

	fmt.Printf("getting coil slice for %s, slice name %s\n", d.Response.FunctionName, name)
	fmt.Printf("coil slice definition %+v\n", d)
	slices, err := coil.GetCoilSlices(abi, d.Response.FunctionName, &name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find coils: %w", err)
	}

	if len(slices) == 0 {
		return nil, fmt.Errorf("no slices found for name: %s", name)
	}

	return &slices[0], nil
}

func (d *ActionDefinition[T]) GetCoilKeys() (map[string]string, error) {
	slices, err := d.GetCoilSlices()
	if err != nil {
		return nil, err
	}

	coilKeys := make(map[string]string)
	for _, slice := range slices {
		name := *slice.Name
		coilKeys[name] = slice.Type
	}

	return coilKeys, nil
}
