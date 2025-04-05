package routes

import (
	"context"
	"solver/internal/avs/config"
	"solver/internal/avs/streams"

	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	_, address, _ := config.GetAccount()
	c.JSON(200, gin.H{
		"status":   "running",
		"address":  address,
		"version":  "0.1.0",
		"chain_id": config.ChainId,
	})
}

func Metrics(c *gin.Context) {
	ctx := context.Background()
	consumerGroup := "circuit-operators"

	pending, err := streams.GetPendingMessages(ctx, consumerGroup)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"pending_count": pending.Count,
		"consumers":     pending.Consumers,
	})
}
