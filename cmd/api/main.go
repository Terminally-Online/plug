package main

import (
	"log"
	"net/http"
	"solver/router"
)

func main() {
	router := router.SetupRouter()
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
