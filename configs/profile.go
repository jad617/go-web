package configs

import (
	"log"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func GetActiveProfile() {
	_, goFileExecuted, _, ok := runtime.Caller(0)
	if !ok {
		log.Println("function runtime.Caller() failed to get the current execution filename")
	}

	projectRootDir := path.Join(path.Dir(goFileExecuted), "..")

	err := os.Chdir(projectRootDir)
	if err != nil {
		panic(err)
	}

	switch os.Getenv("GO_ENV") {
	case "local":
		err := godotenv.Load(projectRootDir + "/.env-local")
		if err != nil {
			log.Fatal(err)
		}
	case "pipeline":
		err := godotenv.Load(projectRootDir + "/.env-pipeline")
		if err != nil {
			log.Fatal(err)
		}
	}
}
