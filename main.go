package main

import (
	"fmt"
	"os"
	"strconv"

	"market.ir/internal/server"
)

func main() {
	r := server.InitServer()
	var PORT int
	portEnvValue, err := strconv.Atoi(os.Getenv("PORT"))
	PORT = portEnvValue
	if err != nil {
		PORT = 8080
	}
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(PORT)))
}
