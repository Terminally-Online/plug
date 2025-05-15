package intent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"solver/internal/actions"
	"solver/internal/api/middleware"
	"solver/internal/api/routes"
	"solver/internal/cache"
	"solver/internal/redis"
	"solver/internal/solver"
	"solver/internal/utils"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/gorilla/schema"
	"github.com/swaggest/openapi-go"
)

var Decoder = func() *schema.Decoder {
	d := schema.NewDecoder()
	d.IgnoreUnknownKeys(true)
	d.SetAliasTag("query")
	return d
}()

type SearchQueryParam struct {
	Index int    `schema:"index" query:"index" description:"Index of the search parameter"`
	Value string `schema:"value" query:"value" description:"Value of the search parameter"`
}
type SchemaQueryParams struct {
	ChainId  uint64             `schema:"chainId" query:"chainId" description:"Chain ID to filter schemas by"`
	Protocol string             `schema:"protocol" query:"protocol" description:"Protocol name to filter schemas by"`
	Action   string             `schema:"action" query:"action" description:"Action name to filter schemas by"`
	From     common.Address     `schema:"from" query:"from" description:"Wallet address to generate schemas for"`
	Search   []SearchQueryParam `schema:"search" query:"search" description:"Search parameters to filter schemas by"`
}

func GetContext(oc openapi.OperationContext) error {
	oc.SetTags("Solver")
	oc.SetSummary("Get Schema")
	oc.SetDescription("Retrieves available action schemas for protocols based on query parameters. Requires API key.")

	oc.AddReqStructure(SchemaQueryParams{})
	oc.AddRespStructure(map[string]actions.ProtocolSchema{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "no protocols found on chainId"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	oc.AddRespStructure(
		map[string]string{"error": "unsupported protocol"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)

	return nil
}

func GetAllSchemas(s *solver.Solver, chainId uint64) (map[string]actions.ProtocolSchema, error) {
	all := make(map[string]actions.ProtocolSchema)
	for protocol, handler := range s.Protocols {
		var supportsChain bool
		for _, chain := range handler.Chains {
			if slices.Contains(chain.ChainIds, chainId) {
				supportsChain = true
				break
			}
		}
		if !supportsChain {
			continue
		}

		protocolSchema := actions.ProtocolSchema{
			Metadata: actions.ProtocolMetadata{
				Icon:   handler.Icon,
				Tags:   handler.Tags,
				Chains: handler.Chains,
			},
			Schema: make(map[string]actions.Schema),
		}

		schemas := handler.Schemas
		for supportedAction := range handler.Actions {
			if chainSchema, ok := schemas[supportedAction]; ok {
				protocolSchema.Schema[supportedAction] = actions.Schema{
					Sentence:   chainSchema.Schema.Sentence,
					Coils:      chainSchema.Schema.Coils,
					Properties: chainSchema.Schema.Properties,
				}
			}
		}

		all[protocol] = protocolSchema
	}

	if len(all) == 0 {
		return nil, fmt.Errorf("no protocols found")
	}

	return all, nil
}

func GetProtocolSchema(handler *actions.Protocol, protocol string, chainId uint64) (map[string]actions.ProtocolSchema, error) {
	protocolSchema := actions.ProtocolSchema{
		Metadata: actions.ProtocolMetadata{
			Icon:   handler.Icon,
			Tags:   handler.Tags,
			Chains: handler.Chains,
		},
		Schema: make(map[string]actions.Schema),
	}

	schemas := handler.Schemas
	for supportedAction := range handler.Actions {
		if chainSchema, ok := schemas[supportedAction]; ok {
			protocolSchema.Schema[supportedAction] = actions.Schema{
				Sentence:   chainSchema.Schema.Sentence,
				Coils:      chainSchema.Schema.Coils,
				Properties: chainSchema.Schema.Properties,
			}
		}
	}

	response := map[string]actions.ProtocolSchema{
		protocol: protocolSchema,
	}

	return response, nil
}

func GetActionSchema(handler *actions.Protocol, protocol string, action string, chainId uint64, from common.Address, searchParams []SearchQueryParam) (map[string]actions.ProtocolSchema, error) {
	response, err := GetProtocolSchema(handler, protocol, chainId)
	if err != nil {
		return nil, err
	}

	searchMap := make(map[int]string)
	for _, param := range searchParams {
		searchMap[param.Index] = param.Value
	}

	var chainSchema actions.ChainSchema
	var chainSchemaErr error

	middleware.TrackOptionsBuildTime(protocol, chainId, action, func() error {
		schemaResults, err := handler.GetSchema(chainId, from, searchMap, action)
		if err != nil {
			chainSchemaErr = err
			return err
		}
		if schemaResults != nil {
			chainSchema = *schemaResults
		} else {
			chainSchemaErr = fmt.Errorf("schema results is nil")
			return chainSchemaErr
		}
		return nil
	})

	if chainSchemaErr != nil {
		return nil, chainSchemaErr
	}

	response[protocol] = actions.ProtocolSchema{
		Metadata: response[protocol].Metadata,
		Schema: map[string]actions.Schema{
			action: chainSchema.Schema,
		},
	}

	return response, nil
}

func getCacheKey(params SchemaQueryParams) string {
	cacheKey := fmt.Sprintf("schema:%d:%s:%s:%s", params.ChainId, params.Protocol, params.Action, params.From.Hex())
	paramKeyParts := make([]string, len(params.Search))
	for i, param := range params.Search {
		paramKeyParts[i] = fmt.Sprintf("%d:%s", param.Index, param.Value)
	}

	sort.Strings(paramKeyParts)

	if len(paramKeyParts) != 0 {
		cacheKey = fmt.Sprintf("%s:%s", cacheKey, strings.Join(paramKeyParts, ":"))
	}

	return cacheKey
}

func GetRequest(s *solver.Solver, params SchemaQueryParams) (map[string]actions.ProtocolSchema, error) {
	var result map[string]actions.ProtocolSchema
	var err error

	switch {
	case params.Protocol == "":
		result, err = GetAllSchemas(s, params.ChainId)
	case params.Protocol != "" && params.Action == "":
		handler, exists := s.Protocols[params.Protocol]
		if !exists {
			return nil, fmt.Errorf("unsupported protocol")
		}
		result, err = GetProtocolSchema(&handler, params.Protocol, params.ChainId)
	case params.Protocol != "" && params.Action != "":
		handler, exists := s.Protocols[params.Protocol]
		if !exists {
			return nil, fmt.Errorf("unsupported protocol")
		}
		result, err = GetActionSchema(&handler, params.Protocol, params.Action, params.ChainId, params.From, params.Search)
	default:
		return nil, utils.ErrInvalidField("protocol", params.Protocol)
	}

	return result, err
}

func GetCachedRequest(w http.ResponseWriter, r *http.Request, c *redisv8.Client, s *solver.Solver) {
	var params SchemaQueryParams
	if err := Decoder.Decode(&params, r.URL.Query()); err != nil {
		utils.RespondWithError(w, utils.ErrInvalidParameters(err))
		return
	}

	result, err := cache.WithCache(getCacheKey(params), cache.WithOptions(
		cache.WithDuration(cache.Period),
		cache.WithStaleData(cache.UseStale),
		cache.WithStaleBuffer(cache.StaleBuffer),
	), func() (map[string]actions.ProtocolSchema, error) {
		return GetRequest(s, params)
	})
	if err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to get schema: "+err.Error()))
		return
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to encode response: "+err.Error()))
		return
	}
}

func Get(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(GetCachedRequest, GetContext, redis.CacheRedis, s)
}
