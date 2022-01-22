package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"market.ir/internal/db"
	"market.ir/internal/router"
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

	db.New()

	r := router.RegisterRouter()
	// r.Run(fmt.Sprintf(":%s", PORT))
	log.Fatal(r.Run(fmt.Sprintf(":%s", PORT)))

}
