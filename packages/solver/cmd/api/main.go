package main

import (
	"log"
	"net/http"
	"solver/internal/api"
	"solver/internal/cron"
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

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	cronJob := scheduler.NewWithLocation(location)

	for _, job := range cron.CronJobs {
		err = cronJob.AddFunc(job.Schedule, job.Job)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Starting %d cron jobs...", len(cron.CronJobs))
	go cronJob.Start()
	log.Printf("Started %d cron jobs...", len(cron.CronJobs))

	router := api.SetupRouter()

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
