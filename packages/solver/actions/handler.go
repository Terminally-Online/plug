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

type Protocol struct {
	Name            string
	Icon            string
	Tags            []string
	Chains          []int
	Schemas         map[types.Action]types.ChainSchema
	OptionsProvider types.OptionsProvider
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

func NewBaseHandler(name string, icon string, tags []string, chains []int, sentences map[types.Action]string, optionsProvider types.OptionsProvider) *BaseHandler {
	// Pre-allocate schemas map with exact size
	schemas := make(map[types.Action]types.ChainSchema, len(sentences))

	// Pre-allocate actions slice for pre-warming
	actions := make([]types.Action, 0, len(sentences))

	// Initialize schemas and collect actions
	for action, sentence := range sentences {
		schemas[action] = types.ChainSchema{
			Schema: types.Schema{
				Sentence: sentence,
			},
		}
		actions = append(actions, action)
	}

	// Create cached options provider
	cachedProvider := types.NewCachedOptionsProvider(optionsProvider)

	// Create the handler
	handler := &BaseHandler{
		protocol: Protocol{
			Name:            name,
			Icon:            icon,
			Tags:            tags,
			Chains:          chains,
			Schemas:         schemas,
			OptionsProvider: cachedProvider,
		},
	}

	// Start pre-warming cache in background
	go func() {
		for _, chainId := range chains {
			cachedProvider.PreWarmCache(chainId, actions)
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
