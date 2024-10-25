package intent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/utils"
)

type GetIntentRequest struct {
	ChainId  int    `json:"chainId"`
	Protocol string `json:"protocol"`
	Type     string `json:"type"`
}

func (i GetIntentRequest) Validate() error {
	if !utils.IsSupportedChain(i.ChainId) {
		return utils.ErrInvalidChainId("chainId", i.ChainId)
	}

	return nil
}

func Get(w http.ResponseWriter, r *http.Request) {
	var request GetIntentRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.Error(w, utils.ServerError{Message: "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	inputs, err := GetActionInputs(request.Type)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	provider, err := utils.GetProvider(request.ChainId)
	if err != nil {
		utils.Error(w, fmt.Errorf("failed to connect to Ethereum node: %v", err), http.StatusBadRequest)
		return
	}

	schema, err := inputs.Get(provider, request.ChainId)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(schema); err != nil {
		utils.Error(w, err, http.StatusInternalServerError)
		return
	}
}
