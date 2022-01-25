package job

import (
	"github.com/robfig/cron/v3"
)

func StartCronJobs() {

	c := cron.New()
	// Every hour on the half hour
	c.AddFunc("0 30 * * * *", func() {
		trainTask()
	})
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	// c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

}
