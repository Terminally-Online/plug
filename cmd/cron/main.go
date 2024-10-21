package main

import (
	"time"
	scheduler "github.com/robfig/cron"
    "solver/cron"
)

func main() { 
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

    cronJob.Start()

    // NOTE: Keeps the process running until it is killed.
    select {}
}
