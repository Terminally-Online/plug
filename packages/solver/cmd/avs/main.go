package main

import (
	"fmt"
	"log"
	"os"
	"solver/internal/avs/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	router := gin.Default()
	router.POST("p2p/message", routes.PerformTask)
	router.POST("task/validate", routes.ValidateTask)

	log.Printf("Started on %s", port)

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}
