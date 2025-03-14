package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/solver/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"

	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type HandlerParams struct {
	Client  *client.Client
	ChainId uint64
	From    common.Address
}

func (p *HandlerParams) New(chainId uint64, from common.Address) (HandlerParams, error) {
	client, err := client.New(chainId)
	if err != nil {
		return HandlerParams{}, err
	}

	return HandlerParams{
		Client:  client,
		ChainId: chainId,
		From:    from,
	}, nil
}

type ActionFunc func(rawInputs json.RawMessage, params HandlerParams) ([]signature.Plug, error)
type OptionsFunc func(chainId uint64, from common.Address, search map[int]string, action string) (map[int]Options, error)
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

func (p *Protocol) GetSchema(chainId string, from common.Address, search map[int]string, action string) (*ChainSchema, error) {
	chainSchema, exists := p.Schemas[action]
	if !exists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}

	actionDefinition, exists := p.Actions[action]
	if !exists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}

	if actionDefinition.Options != nil {
		chainIdInt, err := strconv.ParseUint(chainId, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid chain ID: %s", chainId)
		}

		// NOTE: Override the address value so that we utilize the cache from the global state
		//       instead of using the "from" parameter as a key lookup even when provided if
		//       the action being queried against only supports global state.
		if !chainSchema.Schema.IsUserSpecific {
			from = utils.ZeroAddress
		}

		inputs, err := actionDefinition.Options(chainIdInt, from, search, action)
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
