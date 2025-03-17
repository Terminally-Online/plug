package intent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/internal/actions"
	"solver/internal/api/routes"
	"solver/internal/bindings/references"
	"solver/internal/solver"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/swaggest/openapi-go"
)

type SchemaQueryParams struct {
	ChainID  string `query:"chainId" description:"Chain ID to filter schemas by"`
	Protocol string `query:"protocol" description:"Protocol name to filter schemas by"`
	Action   string `query:"action" description:"Action name to filter schemas by"`
	From     string `query:"from" description:"Wallet address to generate schemas for"`
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

func GetRequest(w http.ResponseWriter, r *http.Request, s *solver.Solver) {
	chainIdQueryParam := r.URL.Query().Get("chainId")
	protocol := r.URL.Query().Get("protocol")
	action := r.URL.Query().Get("action")
	from := r.URL.Query().Get("from")

	chainId, err := strconv.ParseUint(chainIdQueryParam, 10, 64)
	if err != nil {
		utils.MakeHttpError(w, fmt.Sprintf("invalid chain ID: %s", chainId), http.StatusBadRequest)
		return
	}

	searchParams := make(map[int]string)
	for key, values := range r.URL.Query() {
		if strings.HasPrefix(key, "search[") && strings.HasSuffix(key, "]") {
			index := strings.TrimPrefix(strings.TrimSuffix(key, "]"), "search[")
			if len(values) > 0 {
				indexInt, err := strconv.Atoi(index)
				if err != nil {
					continue
				}
				searchParams[indexInt] = values[0]
			}
		}
	}

	if protocol == "" {
		allSchemas := make(map[string]actions.ProtocolSchema)

		for protocol, handler := range s.Protocols {
			protocolChains := handler.Chains
			// TODO: Only show matching chain ids. I removed this when refactoring.
			chains := make([]*references.Network, len(protocolChains))
			for i, chain := range protocolChains {
				chainCopy := *chain
				chainCopy.References = nil
				chains[i] = &chainCopy
			}

			protocolSchema := actions.ProtocolSchema{
				Metadata: actions.ProtocolMetadata{
					Icon:   handler.Icon,
					Tags:   handler.Tags,
					Chains: chains,
				},
				Schema: make(map[string]actions.Schema),
			}

			schemas := handler.Schemas
			for supportedAction, _ := range handler.Actions {
				if chainSchema, ok := schemas[supportedAction]; ok {
					protocolSchema.Schema[supportedAction] = actions.Schema{
						Type:     chainSchema.Schema.Type,
						Sentence: chainSchema.Schema.Sentence,
						Coils:    chainSchema.Schema.Coils,
					}
				}
			}

			allSchemas[protocol] = protocolSchema
		}

		if len(allSchemas) == 0 {
			utils.MakeHttpError(w, fmt.Sprintf("no protocols found on chainId %s", chainId), http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(allSchemas); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	handler, exists := s.Protocols[protocol]
	if !exists {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported protocol: %s", protocol), http.StatusBadRequest)
		return
	}

	chains := make([]*references.Network, len(handler.Chains))
	for i, chain := range handler.Chains {
		chainCopy := *chain
		chainCopy.References = nil
		chains[i] = &chainCopy
	}

	if action == "" {
		protocolSchema := actions.ProtocolSchema{
			Metadata: actions.ProtocolMetadata{
				Icon:   handler.Icon,
				Tags:   handler.Tags,
				Chains: chains,
			},
			Schema: make(map[string]actions.Schema),
		}

		schemas := handler.Schemas
		for supportedAction, _ := range handler.Actions {
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

		if err := json.NewEncoder(w).Encode(response); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	chainSchema, err := handler.GetSchema(chainId, common.HexToAddress(from), searchParams, action)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	protocolSchema := actions.ProtocolSchema{
		Metadata: actions.ProtocolMetadata{
			Icon:   handler.Icon,
			Tags:   handler.Tags,
			Chains: chains,
		},
		Schema: map[string]actions.Schema{
			action: chainSchema.Schema,
		},
	}

	response := map[string]actions.ProtocolSchema{
		protocol: protocolSchema,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func Get(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(GetRequest, GetContext, s)
}
