package cron

import (
    "fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func AnonymousUsers() {
	apiURL := os.Getenv("PLUG_APP_API_URL")
	if apiURL == "" {
		log.Println("PLUG_APP_API_URL environment variable is not set")
		return
	}

	apiKey := os.Getenv("PLUG_APP_API_KEY")
	if apiKey == "" {
		log.Println("PLUG_APP_API_KEY environment variable is not set")
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", apiURL, "jobs/anonymous"), nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request to clean anonymous users: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code when cleaning anonymous users: %d", resp.StatusCode)
		return
	}

	log.Println("Successfully cleaned anonymous users")
}
