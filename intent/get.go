package intent

import (
	"encoding/json"
	"net/http"
	"solver/types"
	"solver/utils"
)

type GetIntentRequest struct {
	ChainId  int    `json:"chainId"`
	Protocol string `json:"protocol"`
	Type   string `json:"type"`
}

type GetIntentResponse struct {
	Plugs     types.Intent `json:"plugs"`
	Signature string       `json:"signature"`
}

func (i GetIntentRequest) Validate() error {
	if !utils.IsSupportedChain(i.ChainId) {
		return utils.ErrInvalidChainId("chainId", i.ChainId)
	}

	return nil
}

func Get(w http.ResponseWriter, r *http.Request) {
	var intentRequest GetIntentRequest
	if err := json.NewDecoder(r.Body).Decode(&intentRequest); err != nil {
		utils.Error(w, utils.ServerError{Message: "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	if err := intentRequest.Validate(); err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	// TODO: Determine the proper input values.
}

