package main

import (
	"log"
	"net/http"
	"solver/internal/actions"
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

	provider := actions.NewCachedOptionsProvider(&actions.DefaultOptionsProvider{})
	actions.SetCachedOptionsProvider(provider)

	// TODO: Would be nice if we did not define the list here. Honestly, not even really
	//       sure this is doing what I expected when I wrote it. Either way, we should 
	//       just have a rolling job that updates the global cache of all protocols and
	//       all actions that we currently support. The exception being those that are
	//       user specific because the cache will never even be hit for those.
	actionsList := []string{
		actions.ActionDeposit,
		actions.ActionBorrow,
		actions.ActionRedeem,
		actions.ActionRedeemMax,
		actions.ActionWithdraw,
		actions.ActionWithdrawMax,
		actions.ActionRepay,
		actions.ActionHarvest,
		actions.ActionTransfer,
		actions.ActionTransferFrom,
		actions.ActionApprove,
		actions.ActionSwap,
		actions.ActionRoute,
		actions.ActionStake,
		actions.ActionStakeMax,
		actions.ActionBuy,
		actions.ActionBid,
		actions.ActionRenew,
	}
	provider.PreWarmCache(8453, utils.ZeroAddress, actionsList)

	var CronJobs = []struct {
		Schedule string
		Job      func()
	}{
		{"0 0 0 * * *", cron.AnonymousUsers},                                                      // At the start of every day
		{"0 */5 * * * *", cron.CollectibleMetadata},                                               // Every 5 minutes
		{"0 */5 * * * *", func() { provider.PreWarmCache(8453, utils.ZeroAddress, actionsList) }}, // Every 5 minutes
		{"0 */1 * * * *", func() { cron.Simulations(s.Solver) }},                                  // Every 1 minute
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
