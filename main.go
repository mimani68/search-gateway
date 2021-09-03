package main

import (
	"fmt"
	"os"
	"strconv"

	"market.ir/internal/router"
)

func main() {
	r := router.SetupRouter()
	var PORT int
	portEnvValue, err := strconv.Atoi(os.Getenv("PORT"))
	PORT = portEnvValue
	if err != nil {
		PORT = 8080
	}
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(PORT)))
}
