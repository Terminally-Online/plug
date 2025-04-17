package actions

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/internal/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/accounts/abi"
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
	if d.Response == nil {
		return nil, nil
	}

	// If we have custom arguments, return slices for those
	if d.Response.Arguments != nil {
		slices := make([]coil.Slice, 0, len(*d.Response.Arguments))

		for i, arg := range *d.Response.Arguments {
			typeId := coil.TypeIDStatic
			var sliceType string

			// Type conversion logic
			switch arg.Type.T {
			case abi.UintTy:
				sliceType = fmt.Sprintf("uint%d", arg.Type.Size)
			case abi.IntTy:
				sliceType = fmt.Sprintf("int%d", arg.Type.Size)
			case abi.AddressTy:
				sliceType = "address"
			case abi.BoolTy:
				sliceType = "bool"
			case abi.BytesTy:
				typeId = coil.TypeIDString
				sliceType = "bytes"
			case abi.StringTy:
				typeId = coil.TypeIDString
				sliceType = "string"
			case abi.ArrayTy:
				typeId = coil.TypeIDArray
				sliceType = arg.Type.String()
			default:
				sliceType = arg.Type.String()
			}

			// Use the provided name or generate a default one
			name := arg.Name
			if name == "" {
				name = fmt.Sprintf("arg%d", i)
			}

			slices = append(slices, coil.Slice{
				Name:   &name,
				Index:  uint8(i),
				Start:  big.NewInt(int64(i * int(coil.WordSize))),
				Length: big.NewInt(int64(coil.WordSize)), // Default static size
				Type:   sliceType,
				TypeId: &typeId,
			})
		}

		return slices, nil
	}

	// For regular contract ABIs with Metadata
	if d.Response.Metadata != nil && d.Response.FunctionName != "" {
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

	return nil, nil
}

func (d *ActionDefinition[T]) GetCoilSlice(name string) (*coil.Slice, error) {
	if d.Response == nil {
		return nil, nil
	}

	// If we have custom arguments, look for the named argument
	if d.Response.Arguments != nil {
		for i, arg := range *d.Response.Arguments {
			if arg.Name == name {
				typeId := coil.TypeIDStatic
				var sliceType string

				// Type conversion logic
				switch arg.Type.T {
				case abi.UintTy:
					sliceType = fmt.Sprintf("uint%d", arg.Type.Size)
				case abi.IntTy:
					sliceType = fmt.Sprintf("int%d", arg.Type.Size)
				case abi.AddressTy:
					sliceType = "address"
				case abi.BoolTy:
					sliceType = "bool"
				case abi.BytesTy:
					typeId = coil.TypeIDString
					sliceType = "bytes"
				case abi.StringTy:
					typeId = coil.TypeIDString
					sliceType = "string"
				case abi.ArrayTy:
					typeId = coil.TypeIDArray
					sliceType = arg.Type.String()
				default:
					sliceType = arg.Type.String()
				}

				return &coil.Slice{
					Name:   &name,
					Index:  uint8(i),
					Start:  big.NewInt(int64(i * int(coil.WordSize))),
					Length: big.NewInt(int64(coil.WordSize)),
					Type:   sliceType,
					TypeId: &typeId,
				}, nil
			}
		}

		return nil, fmt.Errorf("no argument found with name: %s", name)
	}

	// For regular contract ABIs with Metadata
	if d.Response.Metadata != nil && d.Response.FunctionName != "" {
		abi, err := d.Response.Metadata.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to get ABI: %w", err)
		}

		slices, err := coil.GetCoilSlices(abi, d.Response.FunctionName, &name, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to find coils: %w", err)
		}

		if len(slices) == 0 {
			return nil, fmt.Errorf("no slices found for name: %s", name)
		}

		return &slices[0], nil
	}

	return nil, fmt.Errorf("neither Arguments nor Metadata+FunctionName provided")
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

func NewActionDefinitionWithCoils[T any](
	sentence string,
	action ActionFunc[T],
	options ActionOptionsFunc[T],
	isUserSpecific bool,
	isSearchable bool,
	response *ActionOnchainFunctionResponse,
	explicitCoilKeys map[string]string,
) ActionDefinitionInterface {
	return &ActionDefinitionWithCoils[T]{
		ActionDefinition: ActionDefinition[T]{
			Sentence:       sentence,
			Handler:        action,
			Options:        options,
			IsUserSpecific: isUserSpecific,
			IsSearchable:   isSearchable,
			Response:       response,
		},
		explicitCoilKeys: explicitCoilKeys,
	}
}

// ActionDefinitionWithCoils extends ActionDefinition with explicit coil information
type ActionDefinitionWithCoils[T any] struct {
	ActionDefinition[T]
	explicitCoilKeys map[string]string
}

// Override GetCoilKeys to use the explicit coil mapping
func (d *ActionDefinitionWithCoils[T]) GetCoilKeys() (map[string]string, error) {
	if d.explicitCoilKeys != nil {
		return d.explicitCoilKeys, nil
	}
	return d.ActionDefinition.GetCoilKeys()
}
