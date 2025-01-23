package cron

type CronJob struct {
	Schedule string
	Job      func()
}

var CronJobs = []CronJob{
	{"0 0 0 * * *", AnonymousUsers},        // At the start of every day
	{"0 */5 * * * *", CollectibleMetadata}, // Every 5 minutes
	{"0 */1 * * * *", Simulations},         // Every 1 minute
}
