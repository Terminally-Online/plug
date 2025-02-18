package solver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/bindings/plug_router"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type IntentRequest struct {
	Id      string            `json:"id"`
	ChainId uint64            `json:"chainId"`
	From    string            `json:"from"`
	Inputs  []json.RawMessage `json:"inputs"`
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

func (h *Handler) GetPlug(w http.ResponseWriter, r *http.Request) {
	var req IntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	transactionsBatch := make([]plug_router.PlugTypesLibPlug, 0)
	var breakOuter bool
	for _, inputs := range req.Inputs {
		transactions, err := h.Solver.GetTransaction(inputs, req.ChainId, req.From)
		if err != nil {
			utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// NOTE: Some plug actions have exclusive transactions that need to be run alone
		//       before the rest of the Plug can run. For this, we will just break out
		//       of the loop and execute any solo transactions that are needed for
		//       the rest of the batch to run in sequence.
		for _, transaction := range transactions {
			if transaction.Exclusive {
				// NOTE: Set the field to false to avoid tarnishing the response shape.
				transaction.Exclusive = false
				transactionsBatch = []plug_router.PlugTypesLibPlug{transaction}
				breakOuter = true
				break
			}
		}

		if breakOuter {
			break
		}

		transactionsBatch = append(transactionsBatch, transactions...)
	}

	if len(transactionsBatch) == 0 {
		utils.MakeHttpError(w, "has no transactions to execute", http.StatusBadRequest)
		return
	}

	type IntentResponse struct {
		Transactions []plug_router.PlugTypesLibPlug `json:"transactions"`
		Plug         *simulation.SimulationRequest  `json:"plug,omitempty"`
		Simulation   *simulation.SimulationResponse `json:"simulation,omitempty"`
	}

	message, err := signature.GetLivePlugs(req.ChainId, req.From, transactionsBatch)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	simulationRequest, simulationResponse, err := h.Solver.Simulator.GetSimulationResponse("1", req.ChainId, message)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := IntentResponse{
		Transactions: transactionsBatch,
		Plug:         simulationRequest,
		Simulation:   simulationResponse,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
