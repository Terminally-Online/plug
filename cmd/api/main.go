package main

import (
	"log"
	"net/http"
	"solver/router"
	"solver/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
	}

	router := router.SetupRouter()
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
