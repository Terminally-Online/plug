package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"solver/internal/api"
	"solver/internal/cache"
	"solver/internal/cron"
	"solver/internal/redis"
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
	useCache := flag.Bool("cache", true, "Enable or disable caching (default: true)")
	flag.Parse()
	if !*useCache {
		cache.Period = 1 * time.Nanosecond
		cache.UseStale = false
		cache.StaleBuffer = 0
	} else {
		if _, err := redis.CacheRedis.Ping(context.Background()).Result(); err != nil {
			log.Fatal(err)
		}
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(utils.ErrEnvironmentNotInitialized(err.Error()).Error())
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
	log.Fatal(http.ListenAndServe(":8080", router))
}
