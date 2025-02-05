package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"

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
	GetChains() []uint64
	GetSchema(chainId string, from common.Address, action string) (*ChainSchema, error)
	GetSchemas() map[string]ChainSchema
	GetTransaction(action string, rawInputs json.RawMessage, params HandlerParams) ([]signature.Plug, error)
}

type TransactionHandler func(rawInputs json.RawMessage, params HandlerParams) ([]signature.Plug, error)
type OptionsHandler func(chainId uint64) (map[int]Options, error)

type Protocol struct {
	Name            string
	Icon            string
	Tags            []string
	Chains          []uint64
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
	chains []uint64,
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
			return optionsProvider.GetOptions(chainId, common.Address(utils.NativeTokenAddress), action)
		}
	}

	optHandlers := make(map[string]OptionsHandler, len(actionDefinitions))
	for action := range actionDefinitions {
		optHandlers[action] = getOptionsFor(action)
	}

	cachedProvider := NewCachedOptionsProvider(optionsProvider)

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
				Sentence: def.Sentence,
			},
		}
	}

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

	actions := make([]string, 0, len(actionDefinitions))
	for action := range actionDefinitions {
		actions = append(actions, action)
	}
	go func() {
		for _, chainId := range chains {
			cachedProvider.PreWarmCache(chainId, common.Address(utils.NativeTokenAddress), actions)
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

func (h *BaseHandler) GetChains() []uint64 {
	return h.protocol.Chains
}

func (h *BaseHandler) GetActions() []string {
	// Pre-allocate slice with exact capacity needed
	actions := make([]string, 0, len(h.protocol.Schemas))
	for action := range h.protocol.Schemas {
		actions = append(actions, action)
	}
	return actions
}

func (h *BaseHandler) GetSchemas() map[string]ChainSchema {
	return h.protocol.Schemas
}

func (h *BaseHandler) GetSchema(chainId string, from common.Address, action string) (*ChainSchema, error) {
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
		for _, supportedChainId := range h.protocol.Chains {
			if supportedChainId == chainIdInt {
				supported = true
				break
			}
		}
		if !supported {
			return nil, fmt.Errorf(errUnsupportedChain, chainIdInt)
		}

		options, err := h.protocol.OptionsProvider.GetOptions(chainIdInt, from, action)
		if err != nil {
			return nil, fmt.Errorf(errFailedOptions, err)
		}
		chainSchema.Schema.Options = options
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
