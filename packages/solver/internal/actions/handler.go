package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"
	"strings"

	// "strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type HandlerParams struct {
	Provider *ethclient.Client
	ChainId  uint64
	From     string
}

type BaseProtocolHandler interface {
	GetIcon() string
	GetTags() []string
	GetActions() []string
	GetChains() []*references.Network
	GetSchema(chainId string, from common.Address, search map[int]string, action string) (*ChainSchema, error)
	GetSchemas() map[string]ChainSchema
	GetTransaction(action string, rawInputs json.RawMessage, params HandlerParams) ([]signature.Plug, error)
}

type TransactionHandler func(rawInputs json.RawMessage, params HandlerParams) ([]signature.Plug, error)
type OptionsHandler func(chainId uint64) (map[int]Options, error)

type Protocol struct {
	Name            string
	Icon            string
	Tags            []string
	Chains          []*references.Network
	OptionsProvider OptionsProvider
	Schemas         map[string]ChainSchema
	txHandlers      map[string]TransactionHandler
	optHandlers     map[string]OptionsHandler
}

type BaseHandler struct {
	protocol Protocol
}

var (
	errUnsupportedAction = "unsupported action: %s"
	errInvalidChainID    = "invalid chain id: %s"
	errFailedOptions     = "failed to get options: %w"
)

type ActionDefinition struct {
	Type           string `default:"action,omitempty"`
	Sentence       string
	Handler        TransactionHandler
	IsUserSpecific bool
	IsSearchable   bool
}

func NewBaseHandler(
	name string,
	icon string,
	tags []string,
	chains []*references.Network,
	actionDefinitions map[string]ActionDefinition,
	optionsProvider OptionsProvider,
) *BaseHandler {
	sentences := make(map[string]string, len(actionDefinitions))
	transactions := make(map[string]TransactionHandler, len(actionDefinitions))
	for action, def := range actionDefinitions {
		sentences[action] = def.Sentence
		transactions[action] = def.Handler
	}
	getOptionsFor := func(action string) OptionsHandler {
		return func(chainId uint64) (map[int]Options, error) {
			return optionsProvider.GetOptions(chainId, common.Address(utils.ZeroAddress), map[int]string{}, action)
		}
	}
	optHandlers := make(map[string]OptionsHandler, len(actionDefinitions))
	for action := range actionDefinitions {
		optHandlers[action] = getOptionsFor(action)
	}
	schemas := make(map[string]ChainSchema, len(actionDefinitions))
	for action, def := range actionDefinitions {
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
			},
		}
	}
	cachedProvider := NewCachedOptionsProvider(optionsProvider)
	handler := &BaseHandler{
		protocol: Protocol{
			Name:            name,
			Icon:            icon,
			Tags:            tags,
			Chains:          chains,
			Schemas:         schemas,
			OptionsProvider: cachedProvider,
			txHandlers:      transactions,
			optHandlers:     optHandlers,
		},
	}

	actions := make([]string, 3, len(actionDefinitions))
	for action := range actionDefinitions {
		actions = append(actions, action)
	}
	go func() {
		for _, chain := range chains {
			for _, chainId := range chain.ChainIds {
				cachedProvider.PreWarmCache(chainId, common.Address(utils.ZeroAddress), actions)
			}
		}
	}()

	return handler
}

func (h *BaseHandler) GetName() string {
	return h.protocol.Name
}

func (h *BaseHandler) GetIcon() string {
	return h.protocol.Icon
}

func (h *BaseHandler) GetTags() []string {
	return h.protocol.Tags
}

func (h *BaseHandler) GetChains() []*references.Network {
	return h.protocol.Chains
}

func (h *BaseHandler) GetActions() []string {
	actions := make([]string, 0, len(h.protocol.Schemas))
	for action := range h.protocol.Schemas {
		actions = append(actions, action)
	}
	return actions
}

func (h *BaseHandler) GetSchemas() map[string]ChainSchema {
	return h.protocol.Schemas
}

func (h *BaseHandler) GetSchema(chainId string, from common.Address, search map[int]string, action string) (*ChainSchema, error) {
	chainSchema, exists := h.protocol.Schemas[action]
	if !exists {
		return nil, fmt.Errorf(errUnsupportedAction, action)
	}

	if h.protocol.OptionsProvider != nil {
		chainIdInt, err := strconv.ParseUint(chainId, 10, 64)
		if err != nil {
			return nil, fmt.Errorf(errInvalidChainID, chainId)
		}

		supported := false
		for _, supportedChain := range h.protocol.Chains {
			for _, supportedChainId := range supportedChain.ChainIds {
				if chainIdInt == supportedChainId {
					supported = true
					break
				}
			}
		}
		if !supported {
			return nil, fmt.Errorf("chain not supported: %d", chainIdInt)
		}

		if !h.protocol.Schemas[action].Schema.IsUserSpecific {
			from = utils.ZeroAddress
		}

		inputs, err := h.protocol.OptionsProvider.GetOptions(chainIdInt, from, search, action)
		if err != nil {
			return nil, fmt.Errorf(errFailedOptions, err)
		}

		for index := range inputs {
			if search[index] != "" {
				temp := inputs[index]
				var simple []Option
				searchTerm := strings.ToLower(search[index])
				
				for _, option := range temp.Simple {
					if strings.Contains(strings.ToLower(option.Label), searchTerm) ||
						strings.Contains(strings.ToLower(option.Name), searchTerm) ||
						strings.Contains(strings.ToLower(option.Value), searchTerm) {
						simple = append(simple, option)
					}
				}
				
				temp.Simple = simple
				inputs[index] = temp
			}
		}

		chainSchema.Schema.Options = inputs
	}

	return &chainSchema, nil
}

func (h *BaseHandler) GetTransaction(
	action string,
	rawInputs json.RawMessage,
	params HandlerParams,
) ([]signature.Plug, error) {
	handler, exists := h.protocol.txHandlers[action]
	if !exists {
		return nil, fmt.Errorf(errUnsupportedAction, action)
	}
	return handler(rawInputs, params)
}
