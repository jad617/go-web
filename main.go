package main

import (
	"go-web/configs"
	"go-web/handlers"
	"log"
)

var err error

func main() {
	configs.GetActiveProfile()
	r := handlers.Handler()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
