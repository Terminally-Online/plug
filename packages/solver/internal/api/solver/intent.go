package solver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type IntentDomain struct {
	ChainId uint64 `json:"chainId"`
	From    string `json:"from"`
}

type IntentRequest struct {
	Id     string            `json:"id"`
	Inputs []json.RawMessage `json:"inputs"`
	IntentDomain
}

func (r *IntentRequest) Validate() error {
	if r.ChainId == 0 {
		return fmt.Errorf("'chainId' is required")
	}

	if r.From == "" {
		return fmt.Errorf("'from' is required")
	}

	return nil
}

func (req *IntentRequest) UnmarshalJSON(data []byte) error {
	type AuxIntentRequest IntentRequest
	aux := (*AuxIntentRequest)(req)

	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("invalid request body: %w", err)
	}

	if err := req.Validate(); err != nil {
		return err
	}

	return nil
}

func (h *Handler) GetSchema(w http.ResponseWriter, r *http.Request) {
	chainId := r.URL.Query().Get("chainId")
	protocol := r.URL.Query().Get("protocol")
	action := r.URL.Query().Get("action")
	from := r.URL.Query().Get("from")

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

		for protocol, handler := range h.Solver.Protocols {
			protocolChains, err := handler.GetChains(chainId)
			if err != nil {
				continue
			}
			chains := make([]*references.Network, len(protocolChains))
			for i, chain := range protocolChains {
				chainCopy := *chain
				chainCopy.References = nil
				chains[i] = &chainCopy
			}

			protocolSchema := actions.ProtocolSchema{
				Metadata: actions.ProtocolMetadata{
					Icon:   handler.GetIcon(),
					Tags:   handler.GetTags(),
					Chains: chains,
				},
				Schema: make(map[string]actions.Schema),
			}

			schemas := handler.GetSchemas()
			for _, supportedAction := range handler.GetActions() {
				if chainSchema, ok := schemas[supportedAction]; ok {
					protocolSchema.Schema[supportedAction] = actions.Schema{
						Type:     chainSchema.Schema.Type,
						Sentence: chainSchema.Schema.Sentence,
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

	handler, exists := h.Solver.Protocols[protocol]
	if !exists {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported protocol: %s", protocol), http.StatusBadRequest)
		return
	}

	protocolChains, err := handler.GetChains(chainId)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	chains := make([]*references.Network, len(protocolChains))
	for _, chain := range protocolChains {
		chainCopy := *chain
		chainCopy.References = nil
		chains = append(chains, &chainCopy)
	}

	if action == "" {
		protocolSchema := actions.ProtocolSchema{
			Metadata: actions.ProtocolMetadata{
				Icon:   handler.GetIcon(),
				Tags:   handler.GetTags(),
				Chains: chains,
			},
			Schema: make(map[string]actions.Schema),
		}

		schemas := handler.GetSchemas()
		for _, supportedAction := range handler.GetActions() {
			if chainSchema, ok := schemas[supportedAction]; ok {
				protocolSchema.Schema[supportedAction] = actions.Schema{
					Type:     chainSchema.Schema.Type,
					Sentence: chainSchema.Schema.Sentence,
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
			Icon:   handler.GetIcon(),
			Tags:   handler.GetTags(),
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

func (h *Handler) GetSolution(w http.ResponseWriter, r *http.Request) {
	var intent *models.Intent
	if err := json.NewDecoder(r.Body).Decode(&intent); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	intent.ApiKeyId = r.Header.Get("X-Api-Key-Id")
	var err error

	intent, err = intent.GetOrCreate(database.DB)
	if err != nil {
		utils.MakeHttpError(w, "failed to initialize intent: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if solution, err := h.Solver.Solve(intent); err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(solution); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
