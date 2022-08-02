package main

import (
	"go-web/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var err error

func init() {
	// For local development we load the .env file
	if os.Getenv("GO_ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	r := handlers.Handler()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
