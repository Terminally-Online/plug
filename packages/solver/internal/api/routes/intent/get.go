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
	"solver/internal/solver"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
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
					Type:     chainSchema.Schema.Type,
					Sentence: chainSchema.Schema.Sentence,
					Coils:    chainSchema.Schema.Coils,
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
				Type:     chainSchema.Schema.Type,
				Sentence: chainSchema.Schema.Sentence,
				Coils:    chainSchema.Schema.Coils,
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

func GetRequest(w http.ResponseWriter, r *http.Request, c *redis.Client, s *solver.Solver) {
	var params SchemaQueryParams
	if err := Decoder.Decode(&params, r.URL.Query()); err != nil {
		utils.MakeHttpError(w, fmt.Sprintf("invalid parameters: %v", err), http.StatusBadRequest)
		return
	}

	protocol, exists := s.Protocols[params.Protocol]

	var result map[string]actions.ProtocolSchema
	var err error
	switch {
	case params.Protocol == "":
		result, err = GetAllSchemas(s, params.ChainId)
	case exists && params.Action == "":
		result, err = GetProtocolSchema(&protocol, params.Protocol, params.ChainId)
	case exists && params.Action != "":
		result, err = GetActionSchema(&protocol, params.Protocol, params.Action, params.ChainId, params.From, params.Search)
	default:
		utils.MakeHttpError(w, "invalid protocol", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func Get(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(GetRequest, GetContext, cache.Redis, s)
}
