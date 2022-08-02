package main

import (
	"go-web/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var err error

func main() {
	if os.Getenv("GO_ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}
	}

	r := handlers.Handler()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
