package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"market.ir/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := server.InitServer()
	var PORT int
	portEnvValue, err := strconv.Atoi(os.Getenv("PORT"))
	PORT = portEnvValue
	if err != nil {
		PORT = 8080
	}
	r.Run(fmt.Sprintf(":%s", strconv.Itoa(PORT)))
}
