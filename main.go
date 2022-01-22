package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"market.ir/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	r := server.InitServer()
	r.Run(fmt.Sprintf(":%s", PORT))
}
