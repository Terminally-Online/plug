package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/solver/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ActionFunc[T any] func(lookup *SchemaLookup[T]) ([]signature.Plug, error)
type ActionOptionsFunc[T any] func(lookup *SchemaLookup[T]) (map[int]Options, error)

type ActionOnchainFunctionResponse struct {
	Metadata     *bind.MetaData
	FunctionName string
}

func (r *ActionOnchainFunctionResponse) GetCalldata(inputs ...any) ([]byte, error) {
	abi, err := r.Metadata.GetAbi()
	if err != nil {
		return nil, err
	}

	calldata, err := abi.Pack(r.FunctionName, inputs...)
	if err != nil {
		return nil, err
	}

	return calldata, nil
}

// ActionInterface is used for type assertions in Protocol
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

// TODO: Although we are going to want all of this data to properly build
//       the transaction we should minify the return data to only the name
//       and type because the solver will have the positions when the
//       transactions is actually being built. Example:
//       .
//		 "coils": [
// 		     {
// 		         "start": 0,
// 		         "slice": {
// 		             "name": "balance",
// 		             "index": 0,
// 		             "start": 0,
// 		             "length": 32,
// 		             "type": "uint256",
// 		             "typeId": 0
// 		         }
// 		     }
// 		 ]
//       .
// 		 "coils": {
// 		     "balance": "uint256",
// 		 }
//       - CHANCE

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
