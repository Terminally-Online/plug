package main

import (
	"log"
	"net/http"
	"solver/internal/api"
	"solver/internal/api/solver"
	"solver/internal/cron"
	"solver/internal/utils"

	"github.com/joho/godotenv"
	scheduler "github.com/robfig/cron"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
	}

	s := solver.New()

	var CronJobs = []struct {
		Schedule string
		Job      func()
	}{
		{"0 0 0 * * *", cron.AnonymousUsers},                     // At the start of every day
		{"0 */5 * * * *", cron.CollectibleMetadata},              // Every 5 minutes
		{"0 */1 * * * *", func() { cron.Simulations(s.Solver) }}, // Every 1 minute
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
	log.Fatal(http.ListenAndServe(":8080", router))
}
