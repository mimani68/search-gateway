package job

import (
	"fmt"

	"market.ir/pkg/sdk"
)

func StartCronJobs() {

	trainTask()

}

func trainTask() bool {
	keies := []string{
		"فروش",
		"قیمت",
		"سود",
	}
	for _, key := range keies {
		result := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", key)
		fmt.Println(result)
	}
	return true
}
