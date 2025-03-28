package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MakeHTTPRequest[T any](fullUrl string, method string, headers map[string]string, parameters url.Values, body io.Reader, responseType T) (T, error) {
	client := http.Client{}
	u, err := url.Parse(fullUrl)
	if err != nil {
		return responseType, err
	}

	if method == "GET" {
		query := u.Query()
		for key, value := range parameters {
			query.Set(key, strings.Join(value, ","))
		}
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return responseType, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return responseType, err
	}
	if res == nil {
		return responseType, fmt.Errorf("Error: Calling %s returned an empt response", u.String())
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return responseType, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return responseType, fmt.Errorf("Error calling %s: %s", u.String(), string(responseData))
	}

	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return responseType, err
	}
	return responseObject, nil
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func RespondWithError(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	message := err.Error()
	errorType := "InternalError"

	if serverErr, ok := err.(SolverError); ok {
		statusCode = serverErr.GetStatusCode()
		message = serverErr.Message
		errorType = serverErr.Type
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(ErrorResponse{
		Error:   errorType,
		Message: message,
		Code:    statusCode,
	}); err != nil {
		http.Error(w, "Failed to encode error response", http.StatusInternalServerError)
	}
}
