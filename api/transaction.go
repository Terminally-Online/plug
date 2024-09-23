package api

import (
	"encoding/json"
	"net/http"
	"solver/intent"
)

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var action intent.Action
	if err := json.NewDecoder(r.Body).Decode(&action); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	inputs, err := intent.ParseAction(action)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := inputs.Build()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
