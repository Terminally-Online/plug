package intent

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"solver/types"
	"solver/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PostIntentRequest struct {
	ChainId int            `json:"chainId"`
	From    string         `json:"from"`
	Solver  *string        `json:"solver"`
	Actions []types.Action `json:"actions"`
}

type PostIntentResponse struct {
	Plugs     types.Intent         `json:"plugs"`
	Signature string        `json:"signature"`
}

func (i PostIntentRequest) Validate() error {
	if !utils.IsSupportedChain(i.ChainId) {
		return utils.ErrInvalidChainId("chainId", i.ChainId)
	}

	if !utils.IsAddress(i.From) {
		return utils.ErrInvalidAddress("from", i.From)
	}

	if len(i.Actions) < utils.MinActions || len(i.Actions) > utils.MaxActions {
		return utils.ErrInvalidArrayLength("actions", &utils.MinActions, &utils.MaxActions)
	}

	return nil
}

func getIntent(w http.ResponseWriter, provider *ethclient.Client, intentRequest PostIntentRequest) (*types.Intent, error) {
	var transactions []types.Transaction
	for _, action := range intentRequest.Actions {
		inputs, err := ParseAction(action)
		if err != nil {
			return nil, utils.ErrTransactionFailed(err.Error())
		}

		deposit, err := inputs.Build(provider, intentRequest.ChainId, intentRequest.From)
		if err != nil {
			return nil, utils.ErrTransactionFailed(err.Error())
		}

		for _, tx := range deposit {
			transactions = append(transactions, *tx)
		}
	}

	plugs := types.Intent{
		Socket:       intentRequest.From,
		Transactions: transactions,
	}

	// Generate the encoded solver value so that the smart contract can decode it.
	// Note: Used in Solidity with:
	// 		body: `(uint48 expiration, address solver))`
	// 		encode: `abi.encode(uint48(0), msg.sender)`
	// 		decode: `abi.decode(data, (uint48, address))`
	if intentRequest.Solver != nil {
		solverArguments := abi.Arguments{
			{Type: abi.Type{T: abi.UintTy, Size: 48}},
			{Type: abi.Type{T: abi.AddressTy}},
		}
		expiration := big.NewInt(0).Add(big.NewInt(time.Now().Unix()), big.NewInt(300))

		solver, err := solverArguments.Pack(expiration, common.HexToAddress(*intentRequest.Solver))
		if err != nil {
			return nil, utils.ErrTransactionFailed(err.Error())
		}

		plugs.Solver = "0x" + common.Bytes2Hex(solver)
	}

	saltArguments := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 96}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
	}
	salt, err := saltArguments.Pack(
		big.NewInt(time.Now().Unix()),
		common.HexToAddress(intentRequest.From),
		common.HexToAddress(os.Getenv("ONE_CLICKER_ADDRESS")),
		common.HexToAddress(os.Getenv("IMPLEMENTATION_ADDRESS")),
	)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
	}
	plugs.Salt = "0x" + common.Bytes2Hex(salt)

	return &plugs, nil
}

func Post(w http.ResponseWriter, r *http.Request) {
	var intentRequest PostIntentRequest
	if err := json.NewDecoder(r.Body).Decode(&intentRequest); err != nil {
		utils.Error(w, utils.ServerError{Message: "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	if err := intentRequest.Validate(); err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	provider, err := utils.GetProvider(intentRequest.ChainId)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	plugs, err := getIntent(w, provider, intentRequest)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	// Generate the signature that is used to pass through the Router and
	// pass the Socket-sided signature verification so that payment approval
	// can take place without being frontrun.
	var signature []byte
	if intentRequest.Solver != nil {
		// TODO: Implement the EIP-712 signing schema.
		privateKey, err := crypto.HexToECDSA(os.Getenv("SOLVER_PRIVATE_KEY"))
		if err != nil {
			utils.Error(w, err, http.StatusInternalServerError)
			return
		}

		plugsHash := crypto.Keccak256Hash([]byte(plugs.Socket), []byte(plugs.Salt), []byte(plugs.Solver))
		signature, err = crypto.Sign(plugsHash.Bytes(), privateKey)
		if err != nil {
			utils.Error(w, err, http.StatusInternalServerError)
			return
		}
	}

	intentResponse := PostIntentResponse{
		Plugs:     *plugs,
		Signature: "0x" + common.Bytes2Hex(signature),
	}

	if err := json.NewEncoder(w).Encode(intentResponse); err != nil {
		utils.Error(w, err, http.StatusInternalServerError)
		return
	}
}
