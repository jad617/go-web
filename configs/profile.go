package configs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

var (
	err              error
	errVoidEnvValue  = fmt.Errorf("function GetActiveProfile() failed with error: environment variable GO_ENV is empty or set to an unrecognized value")
	errRuntimeCaller = fmt.Errorf("function GetActiveProfile() failed while calling function runtime.Caller(), error: failed to get the current execution filename")
	errOsChdir       = fmt.Errorf("function GetActiveProfile() failed while calling function os.Chdir(), with error: %w", err)
)

func GetActiveProfile() error {
	_, goFileExecuted, _, ok := runtime.Caller(0)
	if !ok {
		return errRuntimeCaller
	}

	projectRootDir := path.Join(path.Dir(goFileExecuted), "..")

	err := os.Chdir(projectRootDir)
	if err != nil {
		return errOsChdir
	}

	activeProfile := strings.ToLower(os.Getenv("GO_ENV"))

	switch activeProfile {
	case "local":
		err := godotenv.Overload(projectRootDir + "/.env-local")
		if err != nil {
			return fmt.Errorf("function GetActiveProfile() failed while calling other function, with error: %w", err)
		}

		return nil

	case "pipeline":
		err := godotenv.Overload(projectRootDir + "/.env-pipeline")
		if err != nil {
			return fmt.Errorf("function GetActiveProfile() failed while calling other function, with error: %w", err)
		}

		return nil

	default:
		return errVoidEnvValue
	}
}
