package actions

import (
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

type ActionDefinition[T any] struct {
	Type           string `default:"action,omitempty"`
	Sentence       string
	Handler        ActionFunc[T]
	Options        ActionOptionsFunc[T]
	Response       ActionOnchainFunctionResponse
	IsUserSpecific bool
	IsSearchable   bool
}

func NewActionDefinition[T any](
	sentence string,
	handler ActionFunc[T],
	options ActionOptionsFunc[T],
	isUserSpecific bool,
	isSearchable bool,
) ActionDefinition[T] {
	return ActionDefinition[T]{
		Sentence:       sentence,
		Handler:        handler,
		Options:        options,
		IsUserSpecific: isUserSpecific,
		IsSearchable:   isSearchable,
	}
}

// func (a *ActionDefinition[T]) GetCoils() ([]coil.Update, error) {
// 	return nil, nil
// if a.Metadata == nil || a.FunctionName == "" {
// 	return []coil.Update{}, nil
// }

// abi, err := a.Metadata.GetAbi()
// if err != nil {
// 	return []coil.Update{}, fmt.Errorf("failed to get ABI: %w", err)
// }

// coils, err := coil.FindCoils(abi, a.FunctionName, nil, nil)
// if err != nil {
// 	return []coil.Update{}, fmt.Errorf("failed to find coils: %w", err)
// }

// return coils, nil
// }

func (d *ActionDefinition[T]) GetType() string {
	return d.Type
}

func (d *ActionDefinition[T]) GetSentence() string {
	return d.Sentence
}

func (d *ActionDefinition[T]) GetIsUserSpecific() bool {
	return d.IsUserSpecific
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
