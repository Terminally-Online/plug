package cron

// * * * * * * <command to execute>
// | | | | | |
// | | | | | day of the week (0–7) (Sunday to Saturday, or use names; 7 is Sunday on some systems)
// | | | | month (1–12)
// | | | day of the month (1–31)
// | | hour (0–23)
// | minute (0–59)
// second (0–59)

type CronJob struct {
	Schedule string
	Job      func()
}

var CronJobs = []CronJob{
	{"0 0 0 * * *", AnonymousUsers},
	{"0 */5 * * * *", CollectibleMetadata},
	{"0 */10 * * * *", Simulations},
}
