package intent

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"solver/solver"
	"solver/types"
	"solver/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type IntentRequest struct {
	ChainId int               `json:"chainId"`
	From    string            `json:"from"`
	Inputs  []json.RawMessage `json:"inputs"`
}

type Handler struct {
	solver *solver.Solver
}

func NewHandler(solver *solver.Solver) *Handler {
	return &Handler{
		solver: solver,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	protocol := r.URL.Query().Get("protocol")
	action := types.Action(r.URL.Query().Get("action"))

	// Case 1: No protocol - return all schemas for all protocols without fields included.
	if protocol == "" {
		allSchemas := make(map[types.Protocol]types.ProtocolSchema)

		for protocol, handler := range h.solver.GetProtocols() {
			protocolSchema := types.ProtocolSchema{
				Metadata: types.ProtocolMetadata{
					Icon: handler.GetIcon(),
					Tags: handler.GetTags(),
				},
				Schema: make(map[types.Action]types.Schema),
			}

			for _, supportedAction := range handler.GetActions() {
				schema, err := handler.GetSchema(supportedAction)
				if err != nil {
					utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
					return
				}

				schemaCopy := types.Schema{
					Sentence: schema.Sentence,
				}
				protocolSchema.Schema[supportedAction] = schemaCopy
			}
			allSchemas[protocol] = protocolSchema
		}

		if err := json.NewEncoder(w).Encode(allSchemas); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	handler, exists := h.solver.GetProtocolHandler(types.Protocol(protocol))
	if !exists {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported protocol: %s", protocol), http.StatusBadRequest)
		return
	}

	// Case 2: Protocol only - return all schemas for that protocol with fields included.
	if action == "" {
		protocolSchema := types.ProtocolSchema{
			Metadata: types.ProtocolMetadata{
				Icon: handler.GetIcon(),
				Tags: handler.GetTags(),
			},
			Schema: make(map[types.Action]types.Schema),
		}

		for _, supportedAction := range handler.GetActions() {
			schema, err := handler.GetSchema(supportedAction)
			if err != nil {
				utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
				return
			}
			protocolSchema.Schema[supportedAction] = *schema
		}

		response := map[types.Protocol]types.ProtocolSchema{
			types.Protocol(protocol): protocolSchema,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Case 3: Protocol and action - return specific schema
	schema, err := handler.GetSchema(action)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	protocolSchema := types.ProtocolSchema{
		Metadata: types.ProtocolMetadata{
			Icon: handler.GetIcon(),
			Tags: handler.GetTags(),
		},
		Schema: map[types.Action]types.Schema{
			action: *schema,
		},
	}

	response := map[types.Protocol]types.ProtocolSchema{
		types.Protocol(protocol): protocolSchema,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var req IntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validateRequest(&req); err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	transactionsBatch := make([]*types.Transaction, 0)
	for _, inputs := range req.Inputs {
		transactions, err := h.solver.GetTransaction(inputs, req.ChainId, req.From)
		if err != nil {
			utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
			return
		}

		transactionsBatch = append(transactionsBatch, transactions...)
	}

	if len(transactionsBatch) == 0 { 
		utils.MakeHttpError(w, "has no transactions to execute", http.StatusBadRequest)
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

	// TODO: Implement the EIP-712 signing schema.
	privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
	if err != nil {
		utils.Error(w, err, http.StatusInternalServerError)
		return
	}

	plugsHash := crypto.Keccak256Hash([]byte(req.From), []byte(salt), []byte(solver))
	signature, err := crypto.Sign(plugsHash.Bytes(), privateKey)
	if err != nil {
		utils.Error(w, err, http.StatusInternalServerError)
		return
	}

	response := types.Plugs{
		Plug: types.Plug{
			Socket: req.From,
			Plugs:  transactionsBatch,
			Solver: "0x" + common.Bytes2Hex(solver),
			Salt:   "0x" + common.Bytes2Hex(salt),
		},
		Signature: "0x" + common.Bytes2Hex(signature),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) validateRequest(req *IntentRequest) error {
	if req.ChainId == 0 {
		return fmt.Errorf("chainId is required")
	}

	if req.From == "" {
		return fmt.Errorf("from address is required")
	}

	return nil
}
