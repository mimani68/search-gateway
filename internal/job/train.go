package job

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"market.ir/pkg/sdk"
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

func trainTask() {
	keies := []string{
		"فروش",
		"قیمت",
		"سود",
	}
	for _, key := range keies {
		result := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", key)
		// result := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", key)
		// result := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", key)
		// result := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", key)
		// result := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", key)
		fmt.Println(result)
	}
}
