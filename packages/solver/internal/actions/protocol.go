package actions

import (
	"fmt"
	"solver/internal/bindings/references"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type Protocol struct {
	Name    string
	Icon    string
	Tags    []string
	Chains  []*references.Network
	Actions map[string]ActionDefinitionInterface
	Schemas map[string]ChainSchema
}

func NewProtocol(p Protocol) Protocol {
	schemas := make(map[string]ChainSchema, len(p.Actions))
	for action, definition := range p.Actions {
		coilKeys, err := definition.GetCoilKeys()
		if err != nil {
			continue
		}

		schemas[action] = ChainSchema{
			Schema: Schema{
				Type: func() string {
					if definition.GetType() == "" {
						return TypeAction
					}
					return definition.GetType()
				}(),
				Sentence:   definition.GetSentence(),
				Properties: definition.GetProperties(),
				Coils:      coilKeys,
			},
		}
	}

	p.Schemas = schemas

	return p
}

func (p *Protocol) GetSchema(chainId uint64, from common.Address, search map[int]string, action string) (*ChainSchema, error) {
	schema, schemaExists := p.Schemas[action]
	definition, definitionExists := p.Actions[action]
	if !schemaExists || !definitionExists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}

	if definition.GetOptions() != nil {
		if definition.GetIsUserSpecific() && from == utils.ZeroAddress {
			return nil, utils.ErrMissingField("from")
		} else if !definition.GetIsUserSpecific() {
			from = utils.ZeroAddress
		}

		lookup, err := NewSchemaLookup[any](chainId, from, search, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create schema lookup: %w", err)
		}

		inputs, err := definition.GetOptions()(lookup)
		if err != nil {
			return nil, fmt.Errorf("failed to get options: %w", err)
		}

		schema.Schema.Options = inputs
	}

	return &schema, nil
}
