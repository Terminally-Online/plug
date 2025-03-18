package actions

import (
	"fmt"
	"log"
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
		log.Println(action)
		log.Println(definition)

		schemas[action] = ChainSchema{
			Schema: Schema{
				Type: func() string {
					if definition.GetType() == "" {
						return TypeAction
					}
					return definition.GetType()
				}(),
				Sentence:       definition.GetSentence(),
				IsUserSpecific: definition.GetIsUserSpecific(),
			},
		}
	}

	p.Schemas = schemas

	return p
}

func (p *Protocol) GetSchema(chainId uint64, from common.Address, search map[int]string, action string) (*ChainSchema, error) {
	chainSchema, schemaExists := p.Schemas[action]
	actionDef, actionExists := p.Actions[action]
	if !schemaExists || !actionExists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}

	if actionDef.GetOptions() != nil {
		if !chainSchema.Schema.IsUserSpecific {
			from = utils.ZeroAddress
		}

		lookup, err := NewSchemaLookup[any](chainId, from, search, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create schema lookup: %w", err)
		}

		inputs, err := actionDef.GetOptions()(lookup)
		if err != nil {
			return nil, fmt.Errorf("failed to get options: %w", err)
		}

		chainSchema.Schema.Options = inputs
	}

	return &chainSchema, nil
}
