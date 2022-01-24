package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"market.ir/internal/db"
	"market.ir/internal/job"
	"market.ir/internal/router"
)

func main() {
	err := godotenv.Load(".development.env")
	if err != nil {
		log.Fatal("Error loading \".development.env\" file")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	db.New()
	job.StartCronJobs()

	r := router.RegisterRouter()
	log.Fatal(r.Run(fmt.Sprintf(":%s", PORT)))

}
