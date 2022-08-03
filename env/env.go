package env

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetEnv(envKey string, defaultValue string) (string, error) {
	var envValue string

	upperEnv := os.Getenv(strings.ToUpper(envKey))
	lowerEnv := os.Getenv(strings.ToLower(envKey))

	switch {
	case upperEnv != "":
		envValue = upperEnv
	case lowerEnv != "":
		envValue = lowerEnv
	case defaultValue != "":
		log.Println(envKey, "is not defined. Default value", defaultValue, "has been used instead")
		return defaultValue, nil
	default:
		return "", fmt.Errorf("got envKey: %v and defaultValue: %v. Was expecting at least 1 valid string parameter to be provided", envKey, defaultValue)
	}

	log.Printf("Value for %v has been defined as: %v\n", envKey, envValue)

	return envValue, nil
}
