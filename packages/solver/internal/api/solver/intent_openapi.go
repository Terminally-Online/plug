package solver

import (
	"net/http"
	"solver/internal/actions"
	"solver/internal/database/models"

	"github.com/swaggest/openapi-go"
)

// SchemaQueryParams defines the query parameters for the GET /solver endpoint
type SchemaQueryParams struct {
	ChainID  string `query:"chainId" description:"Chain ID to filter schemas by"`
	Protocol string `query:"protocol" description:"Protocol name to filter schemas by"`
	Action   string `query:"action" description:"Action name to filter schemas by"`
	From     string `query:"from" description:"Wallet address to generate schemas for"`
}

// SetupOpenAPIForGetSchema defines the OpenAPI documentation for the GET /solver endpoint
func SetupOpenAPIForGetSchema(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Solver")
	oc.SetSummary("Get Schema")
	oc.SetDescription("Retrieves available action schemas for protocols based on query parameters. Requires API key.")
	
	// Define query parameters
	oc.AddReqStructure(SchemaQueryParams{})
	
	// Define success response
	oc.AddRespStructure(map[string]actions.ProtocolSchema{}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "no protocols found on chainId"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	
	oc.AddRespStructure(
		map[string]string{"error": "unsupported protocol"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)

	return nil
}

// SetupOpenAPIForGetSolution defines the OpenAPI documentation for the POST /solver endpoint
func SetupOpenAPIForGetSolution(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Solver")
	oc.SetSummary("Solve Intent")
	oc.SetDescription("Processes an intent and returns a solution (transaction data). Requires API key.")
	
	// Define request body
	oc.AddReqStructure(models.Intent{})
	
	// Define success response
	// Using a generic map since the solution structure is complex and dynamic
	oc.AddRespStructure(map[string]interface{}{}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)
	
	oc.AddRespStructure(
		map[string]string{"error": "failed to initialize intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)
	
	oc.AddRespStructure(
		map[string]string{"error": "failed to solve intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}