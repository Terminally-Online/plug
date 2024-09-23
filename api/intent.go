package api

import (
	"encoding/json"
	"net/http"
	"solver/intent"
	"solver/utils"
)

func CreateIntent(w http.ResponseWriter, r *http.Request) {
	var intentRequest intent.IntentRequest
	if err := json.NewDecoder(r.Body).Decode(&intentRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := intentRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var transactions []utils.Transaction
	for _, action := range intentRequest.Actions {
		inputs, err := intent.ParseAction(action)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		transaction, err := inputs.Build(intentRequest.From)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		transactions = append(transactions, *transaction)
	}

	intentResponse := intent.IntentResponse{
		Request:      intentRequest,
		Transactions: transactions,
	}

	if err := json.NewEncoder(w).Encode(intentResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
