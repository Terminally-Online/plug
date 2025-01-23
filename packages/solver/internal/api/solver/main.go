package solver

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"solver/internal/actions"
	"solver/internal/solver"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Handler struct {
	Solver *solver.Solver
}

func New() *Handler {
	return &Handler{
		Solver: solver.New(),
	}
}

func (h *Handler) GetIntent(w http.ResponseWriter, r *http.Request) {
	protocol := r.URL.Query().Get("protocol")
	action := r.URL.Query().Get("action")
	chainId := r.URL.Query().Get("chainId")

	// Case 1: No protocol - return all schemas for all protocols without options.
	if protocol == "" {
		allSchemas := make(map[string]actions.ProtocolSchema)

		for protocol, handler := range h.Solver.GetProtocols() {
			protocolSchema := actions.ProtocolSchema{
				Metadata: actions.ProtocolMetadata{
					Icon:   handler.GetIcon(),
					Tags:   handler.GetTags(),
					Chains: handler.GetChains(),
				},
				Schema: make(map[string]actions.Schema),
			}

			// Get all schemas without options
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

	// Case 2: Protocol only - return all schemas for that protocol without options.
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

	// Case 3: Protocol and action - return specific schema
	chainSchema, err := handler.GetSchema(chainId, action)
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

	// Generate the encoded solver value so that the smart contract can decode it.
	// Note: Used in Solidity with:
	// 		body: `(uint48 expiration, address solver))`
	// 		encode: `abi.encode(uint48(0), msg.sender)`
	// 		decode: `abi.decode(data, (uint48, address))`
	solverArguments := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 48}},
		{Type: abi.Type{T: abi.AddressTy}},
	}
	expiration := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(300))

	solver, err := solverArguments.Pack(expiration, common.HexToAddress(os.Getenv("SOLVER_ADDRESS")))
	if err != nil {
		utils.MakeHttpError(w, "failed to pack solver: "+err.Error(), http.StatusInternalServerError)
		return
	}

	saltArguments := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 96}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
	}
	salt, err := saltArguments.Pack(
		big.NewInt(time.Now().Unix()),
		common.HexToAddress(req.From),
		common.HexToAddress(os.Getenv("ONE_CLICKER_ADDRESS")),
		common.HexToAddress(os.Getenv("IMPLEMENTATION_ADDRESS")),
	)
	if err != nil {
		utils.MakeHttpError(w, "failed to pack salt: "+err.Error(), http.StatusInternalServerError)
		return
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	plugsHash := crypto.Keccak256Hash([]byte(req.From), []byte(salt), []byte(solver))
	signedPlugsHash, err := crypto.Sign(plugsHash.Bytes(), privateKey)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := signature.LivePlugs{
		Plugs: signature.Plugs{
			Socket: common.HexToAddress(req.From),
			Plugs:  transactionsBatch,
			Solver: solver,
			Salt:   salt,
		},
		Signature: signedPlugsHash,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) GetKill(w http.ResponseWriter, r *http.Request) {
	response := KillResponse{Killed: h.Solver.IsKilled}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) PostKill(w http.ResponseWriter, r *http.Request) {
	h.Solver.IsKilled = !h.Solver.IsKilled

	response := KillResponse{Killed: h.Solver.IsKilled}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

