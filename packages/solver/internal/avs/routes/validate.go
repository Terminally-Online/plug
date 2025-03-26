package routes

import (
	"log"
	"net/http"
	"solver/internal/avs/services"

	"github.com/gin-gonic/gin"
)

func ValidateTask(c *gin.Context) {
	var request map[string]any
	var proofOfTask string
	var data string
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("error decoding json body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if val, ok := request["proofOfTask"].(string); ok {
		proofOfTask = val
	}
	log.Println("proofOfTask:", proofOfTask)

	if val, ok := request["data"].(string); ok {
		data = val
	}

	result, err := services.Validate(proofOfTask, data)
	if err != nil {
		log.Println("Validation error:", err)
		errorResponse := services.NewResponse(map[string]any{}, "Error during validation step")
		c.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	log.Printf("Vote: %s", func() string {
		if result {
			return "Approve"
		}
		return "Not Approved"
	}())

	response := services.NewResponse(map[string]any{
		"result": result,
	}, "Task validated successfully")
	c.JSON(http.StatusOK, response)
}
