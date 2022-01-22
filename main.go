package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"market.ir/internal/db"
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
	db.Client()
	r.Run(fmt.Sprintf("0.0.0.0:%s", PORT))

}
