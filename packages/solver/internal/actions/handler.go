package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/solver/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type SchemaLookup struct {
	ChainId uint64
	Client  *client.Client
	From    common.Address
	Search  map[int]string
}

func NewSchemaLookup(chainId uint64, from common.Address, search map[int]string) (*SchemaLookup, error) {
	client, err := client.New(chainId)
	if err != nil {
		return nil, err
	}

	return &SchemaLookup{
		Client:  client,
		ChainId: chainId,
		From:    from,
		Search:  search,
	}, nil
}

type ActionFunc func(lookup *SchemaLookup, raw json.RawMessage) ([]signature.Plug, error)
type OptionsFunc func(lookup *SchemaLookup) (map[int]Options, error)
type ActionDefinition struct {
	Type     string `default:"action,omitempty"`
	Sentence string
	Handler  ActionFunc
	Options  OptionsFunc

	Metadata     *bind.MetaData
	FunctionName string

	IsUserSpecific bool
	IsSearchable   bool
}

func (a *ActionDefinition) GetCoils() ([]coil.Update, error) {
	if a.Metadata == nil || a.FunctionName == "" {
		return []coil.Update{}, nil
	}

	abi, err := a.Metadata.GetAbi()
	if err != nil {
		return []coil.Update{}, fmt.Errorf("failed to get ABI: %w", err)
	}

	coils, err := coil.FindCoils(abi, a.FunctionName, nil, nil)
	if err != nil {
		return []coil.Update{}, fmt.Errorf("failed to find coils: %w", err)
	}

	return coils, nil
}

type Protocol struct {
	Name    string
	Icon    string
	Tags    []string
	Chains  []*references.Network
	Actions map[string]ActionDefinition

	OptionsProvider OptionsProvider
	Schemas         map[string]ChainSchema
}

func New(p Protocol) Protocol {
	schemas := make(map[string]ChainSchema, len(p.Actions))
	for action, def := range p.Actions {
		coils, err := def.GetCoils()
		if err != nil {
			log.Error("failed to get coils", "action", action, "error", err)
		}
		schemas[action] = ChainSchema{
			Schema: Schema{
				Type: func() string {
					if def.Type == "" {
						return TypeAction
					}
					return def.Type
				}(),
				Sentence:       def.Sentence,
				IsUserSpecific: def.IsUserSpecific,
				Coils:          coils,
			},
			LinkedInputs: coils,
		}
	}

	p.Schemas = schemas

	return p
}

func (p *Protocol) SupportsChain(chainId uint64) bool {
	for _, chain := range p.Chains {
		if slices.Contains(chain.ChainIds, chainId) {
			return true
		}
	}

	return false
}

func (p *Protocol) GetSchema(chainId uint64, from common.Address, search map[int]string, action string) (*ChainSchema, error) {
	chainSchema, schemaExists := p.Schemas[action]
	actionDefinition, actionExists := p.Actions[action]
	if !schemaExists || !actionExists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}

	if actionDefinition.Options != nil {
		// NOTE: Override the address value so that we utilize the cache from the global state
		//       instead of using the "from" parameter as a key lookup even when provided if
		//       the action being queried against only supports global state.
		if !chainSchema.Schema.IsUserSpecific {
			from = utils.ZeroAddress
		}

		lookup, err := NewSchemaLookup(chainId, from, search)
		if err != nil {
			return nil, fmt.Errorf("failed to create schema lookup: %w", err)
		}

		inputs, err := actionDefinition.Options(lookup)
		if err != nil {
			return nil, fmt.Errorf("failed to get options: %w", err)
		}

		chainSchema.Schema.Options = inputs
	}

	return &chainSchema, nil
}

// func (h *BaseHandler) GetTransaction(
// 	action string,
// 	rawInputs json.RawMessage,
// 	params HandlerParams,
// ) ([]signature.Plug, error) {
// 	return h.Protocol.Actions[action].Handler(rawInputs, params)
// }
