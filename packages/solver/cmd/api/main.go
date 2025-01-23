package main

import (
	"log"
	"net/http"
	"solver/internal/api"
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

	// TODO: (#370) The cron jobs and api router need to share the same instance of the 
	//       solver so that if it is killed with the kill switch we need to halt the api
	//       and the cron jobs so that we have complete coverage.
	cronJob := scheduler.New()
	for _, job := range cron.CronJobs {
		err = cronJob.AddFunc(job.Schedule, job.Job)
		if err != nil {
			log.Fatal(err)
		}
	}

	go cronJob.Start()
	log.Printf("Started %d cron jobs...", len(cron.CronJobs))

	router := api.SetupRouter()

	log.Println("Started server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
