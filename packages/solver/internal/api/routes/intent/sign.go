package intent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/bindings/othentic_attestation"
	"solver/internal/api/routes"
	"solver/internal/client"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/redis"
	"solver/internal/solver"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/swaggest/openapi-go"
)

type SignatureRequest struct {
	IntentID        string `json:"intentId" binding:"required"`
	SignedMessage   string `json:"signedMessage" binding:"required"`
	OperatorAddress string `json:"operatorAddress" binding:"required"`
}

func SignContext(oc openapi.OperationContext) error {
	oc.SetTags("AVS")
	oc.SetSummary("Sign LivePlugs for AVS Operator")
	oc.SetDescription("Signs and returns LivePlugs for a specific AVS operator. Requires verification through a signed message and attestation center validation.")

	oc.AddReqStructure(SignatureRequest{})

	oc.AddRespStructure(map[string]signature.LivePlugs{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)
	oc.AddRespStructure(
		map[string]string{"error": "invalid signature"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusUnauthorized),
	)
	oc.AddRespStructure(
		map[string]string{"error": "operator not active"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusForbidden),
	)
	oc.AddRespStructure(
		map[string]string{"error": "intent not found"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)

	return nil
}

func verifyOperatorInAttestationCenter(client bind.ContractBackend, attestationCenterAddress common.Address, operatorAddress common.Address) (bool, error) {
	attestationCenter, err := othentic_attestation.NewOthenticAttestation(
		attestationCenterAddress,
		client,
	)
	if err != nil {
		return false, err
	}

	operators, err := attestationCenter.GetActiveOperatorsDetails(&bind.CallOpts{
		Pending: true,
	})
	if err != nil {
		return false, err
	}

	for _, op := range operators {
		if op.Operator == operatorAddress {
			return true, nil
		}
	}

	return false, nil
}

func verifySignature(message string, signedMessage string, operatorAddress string) (bool, error) {
	sig, err := hexutil.Decode(signedMessage)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %v", err)
	}

	if sig[64] >= 27 {
		sig[64] -= 27
	}

	messageHash := crypto.Keccak256Hash([]byte(message))

	pubKey, err := crypto.SigToPub(messageHash.Bytes(), sig)
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr == common.HexToAddress(operatorAddress), nil
}

func MakeHTTPError(w http.ResponseWriter, err string, code int) { 
}

func SignRequest(w http.ResponseWriter, r *http.Request, c *redisv8.Client, s *solver.Solver) {
	var req SignatureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.MakeHTTPError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("Request signed LivePlugs for intent %s", req.IntentID)
	valid, err := verifySignature(message, req.SignedMessage, req.OperatorAddress)
	if err != nil {
		utils.MakeHTTPError(w, "signature verification error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !valid {
		utils.MakeHTTPError(w, "invalid signature", http.StatusUnauthorized)
		return
	}

	// Get the intent from database
	var intent models.Intent
	result := database.DB.Where("id = ?", req.IntentID).First(&intent)
	if result.Error != nil {
		utils.MakeHTTPError(w, "intent not found: "+result.Error.Error(), http.StatusNotFound)
		return
	}

	rpcUrl, err := client.GetQuicknodeUrl(intent.ChainId)
	if err != nil {
		utils.MakeHTTPError(w, "failed to get RPC URL: "+err.Error(), http.StatusInternalServerError)
		return
	}
	ethClient, err := ethclient.Dial(rpcUrl)
	if err != nil {
		utils.MakeHTTPError(w, "failed to connect to blockchain: "+err.Error(), http.StatusInternalServerError)
		return
	}

	attestationCenterAddress := common.HexToAddress("0x62180042606624f02d8a130da8a3171e9b33894d") // Use your actual address

	isActive, err := verifyOperatorInAttestationCenter(ethClient, attestationCenterAddress, common.HexToAddress(req.OperatorAddress))
	if err != nil {
		utils.MakeHTTPError(w, "failed to verify operator: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !isActive {
		utils.MakeHTTPError(w, "operator is not active", http.StatusForbidden)
		return
	}

	solution, err := s.Solve(&intent, false, true)
	if err != nil {
		utils.MakeHTTPError(w, "failed to solve intent: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]any{
		"livePlugs": solution.LivePlugs,
	}); err != nil {
		utils.MakeHTTPError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

// Sign creates a new route handler for the LivePlugs signing endpoint
func Sign(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(SignRequest, SignContext, redis.CacheRedis, s)
}
