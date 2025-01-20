package actions

import (
	"encoding/json"
	"fmt"
	"solver/types"
	"strconv"

	"github.com/ethereum/go-ethereum/ethclient"
)

type HandlerParams struct {
	Provider *ethclient.Client
	ChainId  int
	From     string
}

type BaseProtocolHandler interface {
	GetIcon() string
	GetTags() []string
	GetActions() []types.Action
	GetChains() []int
	GetSchema(chainId string, action types.Action) (*types.ChainSchema, error)
	GetSchemas() map[types.Action]types.ChainSchema
	GetTransaction(action types.Action, rawInputs json.RawMessage, params HandlerParams) ([]*types.Transaction, error)
}

// Handler function types
type TransactionHandler func(rawInputs json.RawMessage, params HandlerParams) ([]*types.Transaction, error)
type OptionsHandler func(chainId int) (map[int]types.SchemaOptions, error)

type Protocol struct {
	Name            string
	Icon            string
	Tags            []string
	Chains          []int
	Schemas         map[types.Action]types.ChainSchema
	OptionsProvider types.OptionsProvider
	// Add maps for handlers
	txHandlers  map[types.Action]TransactionHandler
	optHandlers map[types.Action]OptionsHandler
}

type BaseHandler struct {
	protocol Protocol
}

// Common error messages
var (
	errUnsupportedAction = "unsupported action: %s"
	errInvalidChainID    = "invalid chain id: %s"
	errUnsupportedChain  = "unsupported chain id: %d"
	errFailedOptions     = "failed to get options: %w"
)

type ActionDefinition struct {
	Type     string `default:"action"`
	Sentence string
	Handler  TransactionHandler
}

func NewBaseHandler(
	name string,
	icon string,
	tags []string,
	chains []int,
	actionDefinitions map[types.Action]ActionDefinition,
	optionsProvider types.OptionsProvider,
) *BaseHandler {
	// Extract sentences and handlers from definitions
	sentences := make(map[types.Action]string, len(actionDefinitions))
	transactions := make(map[types.Action]TransactionHandler, len(actionDefinitions))
	for action, def := range actionDefinitions {
		sentences[action] = def.Sentence
		transactions[action] = def.Handler
	}

	// Create options handlers for each action
	getOptionsFor := func(action types.Action) OptionsHandler {
		return func(chainId int) (map[int]types.SchemaOptions, error) {
			return optionsProvider.GetOptions(chainId, action)
		}
	}

	optHandlers := make(map[types.Action]OptionsHandler, len(actionDefinitions))
	for action := range actionDefinitions {
		optHandlers[action] = getOptionsFor(action)
	}

	// Create cached provider
	cachedProvider := types.NewCachedOptionsProvider(optionsProvider)

	// Initialize schemas
	schemas := make(map[types.Action]types.ChainSchema, len(actionDefinitions))
	for action, def := range actionDefinitions {
		schemas[action] = types.ChainSchema{
			Schema: types.Schema{
				Type: func() string {
					if def.Type == "" {
						return types.TypeAction
					}
					return def.Type
				}(),
				Sentence: def.Sentence,
			},
		}
	}

	// Create the handler
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

	// Start pre-warming cache in background
	actions := make([]types.Action, 0, len(actionDefinitions))
	for action := range actionDefinitions {
		actions = append(actions, action)
	}
	go func() {
		for _, chainId := range chains {
			cachedProvider.PreWarmCache(chainId, actions)
		}
	}()

	return handler
}

// Options provider that uses the handler map
type handlerOptionsProvider struct {
	handlers map[types.Action]OptionsHandler
}

func (p *handlerOptionsProvider) GetOptions(chainId int, action types.Action) (map[int]types.SchemaOptions, error) {
	handler, exists := p.handlers[action]
	if !exists {
		return nil, fmt.Errorf(errUnsupportedAction, action)
	}
	return handler(chainId)
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

func (h *BaseHandler) GetChains() []int {
	return h.protocol.Chains
}

func (h *BaseHandler) GetActions() []types.Action {
	// Pre-allocate slice with exact capacity needed
	actions := make([]types.Action, 0, len(h.protocol.Schemas))
	for action := range h.protocol.Schemas {
		actions = append(actions, action)
	}
	return actions
}

func (h *BaseHandler) GetSchemas() map[types.Action]types.ChainSchema {
	return h.protocol.Schemas
}

func (h *BaseHandler) GetSchema(chainId string, action types.Action) (*types.ChainSchema, error) {
	chainSchema, exists := h.protocol.Schemas[action]
	if !exists {
		return nil, fmt.Errorf(errUnsupportedAction, action)
	}

	if h.protocol.OptionsProvider != nil {
		chainIdInt, err := strconv.Atoi(chainId)
		if err != nil {
			return nil, fmt.Errorf(errInvalidChainID, chainId)
		}

		// Check if this chain is supported
		supported := false
		for _, supportedChainId := range h.protocol.Chains {
			if supportedChainId == chainIdInt {
				supported = true
				break
			}
		}
		if !supported {
			return nil, fmt.Errorf(errUnsupportedChain, chainIdInt)
		}

		options, err := h.protocol.OptionsProvider.GetOptions(chainIdInt, action)
		if err != nil {
			return nil, fmt.Errorf(errFailedOptions, err)
		}
		chainSchema.Schema.Options = options
	}

	return &chainSchema, nil
}

// Update GetTransaction to use the handler map
func (h *BaseHandler) GetTransaction(action types.Action, rawInputs json.RawMessage, params HandlerParams) ([]*types.Transaction, error) {
	handler, exists := h.protocol.txHandlers[action]
	if !exists {
		return nil, fmt.Errorf(errUnsupportedAction, action)
	}
	return handler(rawInputs, params)
}
