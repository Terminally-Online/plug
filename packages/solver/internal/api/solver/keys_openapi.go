package solver

import (
	"net/http"
	"solver/internal/database/models"

	"github.com/swaggest/openapi-go"
)

// ApiKeyResponse simplified model for OpenAPI documentation of API keys
type ApiKeyResponse struct {
	ID        string  `json:"id" description:"Unique identifier for the API key"`
	Key       string  `json:"key" description:"The API key value"`
	SocketID  *string `json:"socketId,omitempty" description:"ID of the associated socket"`
	RateLimit int     `json:"rateLimit" description:"Rate limit for the API key"`
	Role      string  `json:"role" description:"Role of the API key (user or admin)"`
}

// ApiKeyCreateRequest simplified model for creating an API key
type ApiKeyCreateRequest struct {
	SocketID  *string `json:"socketId,omitempty" description:"ID of the socket to associate with this key"`
	RateLimit int     `json:"rateLimit" description:"Rate limit for this API key"`
	Role      string  `json:"role" description:"Role for this API key (user or admin)"`
}

// SetupOpenAPIForCreateApiKey defines the OpenAPI documentation for the POST /api-key endpoint
func SetupOpenAPIForCreateApiKey(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("API Keys")
	oc.SetSummary("Create API Key")
	oc.SetDescription("Creates a new API key with the specified parameters. Requires admin privileges.")
	
	// Define request body
	oc.AddReqStructure(ApiKeyCreateRequest{})
	
	// Define success response
	oc.AddRespStructure(models.ApiKey{}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)
	
	oc.AddRespStructure(
		map[string]string{"error": "failed to save api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// SetupOpenAPIForReadApiKey defines the OpenAPI documentation for the GET /api-key/{id} endpoint
func SetupOpenAPIForReadApiKey(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("API Keys")
	oc.SetSummary("Get API Key")
	oc.SetDescription("Retrieves an API key by ID. Requires admin privileges.")
	
	// Define path parameter
	oc.AddReqStructure(struct {
		ID string `path:"id" description:"API key ID"`
	}{})
	
	// Define success response
	oc.AddRespStructure(models.ApiKey{}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "failed to find api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)

	return nil
}

// SetupOpenAPIForUpdateApiKey defines the OpenAPI documentation for the POST /api-key/{id} endpoint
func SetupOpenAPIForUpdateApiKey(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("API Keys")
	oc.SetSummary("Update API Key")
	oc.SetDescription("Updates an existing API key with the provided parameters. Requires admin privileges.")
	
	// Define path parameter
	oc.AddReqStructure(struct {
		ID string `path:"id" description:"API key ID"`
	}{})
	
	// Define request body
	oc.AddReqStructure(ApiKeyCreateRequest{})
	
	// Define success response
	oc.AddRespStructure(nil, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "failed to find api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)
	
	oc.AddRespStructure(
		map[string]string{"error": "failed to update api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// SetupOpenAPIForDeleteApiKey defines the OpenAPI documentation for the DELETE /api-key/{id} endpoint
func SetupOpenAPIForDeleteApiKey(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("API Keys")
	oc.SetSummary("Delete API Key")
	oc.SetDescription("Deletes an API key by ID. Requires admin privileges.")
	
	// Define path parameter
	oc.AddReqStructure(struct {
		ID string `path:"id" description:"API key ID"`
	}{})
	
	// Define success response
	oc.AddRespStructure(models.ApiKey{}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "failed to find api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	
	oc.AddRespStructure(
		map[string]string{"error": "failed to delete api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}