package main

import (
	"log"
	"net/http"
	"solver/internal/actions"
	"solver/internal/api"
	"solver/internal/cron"
	"solver/internal/solver"
	"solver/internal/utils"
	"time"

	"github.com/joho/godotenv"
	scheduler "github.com/robfig/cron"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
	}

	s := solver.New()

	provider := actions.NewCachedOptionsProvider(&actions.DefaultOptionsProvider{})
	actions.SetCachedOptionsProvider(provider)

	var CronJobs = []struct {
		Schedule string
		Job      func()
	}{
		{"0 */1 * * * *", func() { cron.Simulations(s) }},
		{"0 */15 * * * *", func() { cron.IntentCleanup(time.Minute * 15) }},
	}

	schedule := scheduler.New()
	for _, job := range CronJobs {
		err = schedule.AddFunc(job.Schedule, job.Job)
		if err != nil {
			log.Fatal(err)
		}
	}

	go schedule.Start()
	log.Printf("Started %d cron jobs...", len(CronJobs))

	router := api.SetupRouter(s)

	log.Println("Started server on http://localhost:8080")
	log.Println("OpenAPI specification available at: http://localhost:8080/openapi.json")
	log.Println("API Documentation UI available at: http://localhost:8080/docs")
	log.Fatal(http.ListenAndServe(":8080", router))
}
