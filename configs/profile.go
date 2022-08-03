package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetActiveProfile() {
	getWd, _ := os.Getwd()
	baseDir := filepath.Dir(getWd)

	switch os.Getenv("GO_ENV") {
	case "local":
		err := godotenv.Load(baseDir + "/.env-local")
		if err != nil {
			log.Fatal(err)
		}
	case "pipeline":
		err := godotenv.Load(baseDir + "/.env-pipeline")
		if err != nil {
			log.Fatal(err)
		}
	}
}
