package main

import (
	"context"
	"log"
	"net/http"
	"solver/internal/api"
	"solver/internal/cache"
	"solver/internal/cron"
	"solver/internal/solver"
	"solver/internal/utils"
	"time"

	"github.com/joho/godotenv"
	scheduler "github.com/robfig/cron"
)

var (
	Solver   = solver.New()
	Schedule = scheduler.New()
	CronJobs = []struct {
		Schedule string
		Job      func()
	}{
		{"0 */1 * * * *", func() { cron.Simulations(Solver) }},
		{"0 */15 * * * *", func() { cron.IntentCleanup(time.Minute * 15) }},
	}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
	}

	if _, err := cache.Redis.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}

	for _, job := range CronJobs {
		err = Schedule.AddFunc(job.Schedule, job.Job)
		if err != nil {
			log.Fatal(err)
		}
	}

	go Schedule.Start()
	log.Printf("Started %d cron jobs...", len(CronJobs))

	router := api.SetupRouter(Solver)

	log.Println("Started server on http://localhost:8080")
	log.Println("OpenAPI specification available at: http://localhost:8080/openapi.json")
	log.Println("API Documentation UI available at: http://localhost:8080/docs")
	log.Fatal(http.ListenAndServe(":8080", router))
}
