package main

import (
	"log"
	"solver/cron"
	"solver/utils"
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
		panic(err)
	}
	cronJob := scheduler.NewWithLocation(location)

	for _, job := range cron.CronJobs {
		err = cronJob.AddFunc(job.Schedule, job.Job)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("Running %d jobs...", len(cron.CronJobs))

	cronJob.Start()

	// NOTE: Keeps the process running until it is killed.
	select {}
}
