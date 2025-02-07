package solver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
	"solver/internal/utils"

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

func (h *Handler) GetIntent(w http.ResponseWriter, r *http.Request) {
	protocol := r.URL.Query().Get("protocol")
	action := r.URL.Query().Get("action")
	chainId := r.URL.Query().Get("chainId")
	from := r.URL.Query().Get("from")

	if protocol == "" {
		allSchemas := make(map[string]actions.ProtocolSchema)

		for protocol, handler := range h.Solver.GetProtocols() {
			if chainId != "" {
				supportsChain := false
				chainLoop:
				for _, chain := range handler.GetChains() {
					for _, _chainId := range chain.ChainIds {
						if chainId == fmt.Sprint(_chainId) {
							supportsChain = true
							break chainLoop
						}
					}
				}
				if !supportsChain {
					continue
				}
			}

			var chains []*references.Network
			for _, chain := range handler.GetChains() {
				chain.References = nil
				chains = append(chains, chain)
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

		if err := json.NewEncoder(w).Encode(allSchemas); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	handler, exists := h.Solver.GetProtocolHandler(protocol)
	if !exists {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported protocol: %s", protocol), http.StatusBadRequest)
		return
	}

	if action == "" {
		protocolSchema := actions.ProtocolSchema{
			Metadata: actions.ProtocolMetadata{
				Icon:   handler.GetIcon(),
				Tags:   handler.GetTags(),
				Chains: handler.GetChains(),
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

	chainSchema, err := handler.GetSchema(chainId, common.HexToAddress(from), action)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	protocolSchema := actions.ProtocolSchema{
		Metadata: actions.ProtocolMetadata{
			Icon:   handler.GetIcon(),
			Tags:   handler.GetTags(),
			Chains: handler.GetChains(),
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

func (h *Handler) PostIntent(w http.ResponseWriter, r *http.Request) {
	var req IntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	transactionsBatch := make([]signature.Plug, 0)
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
				transactionsBatch = []signature.Plug{transaction}
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

	message, err := h.Solver.GetPlugs(req.ChainId, req.From, transactionsBatch)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	simulationRequest, simulationResponse, err := h.Solver.GetSimulation("1", req.ChainId, message)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type IntentResponse struct {
		Message     signature.LivePlugs           `json:"message"`
		Transaction simulation.SimulationRequest  `json:"transaction"`
		Simulation  simulation.SimulationResponse `json:"simulation"`
	}
	response := IntentResponse{
		Message:     *message,
		Transaction: simulationRequest,
		Simulation:  simulationResponse,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}
