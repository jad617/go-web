package env

import (
	"log"
	"os"
	"strings"
)

func GetEnv(envKey string, defaultValue string) string {
	var envValue string

	upperEnv := os.Getenv(strings.ToUpper(envKey))
	lowerEnv := os.Getenv(strings.ToLower(envKey))

	switch {
	case upperEnv != "":
		envValue = upperEnv
	case lowerEnv != "":
		envValue = lowerEnv
	default:
		log.Println(envKey, "is not defined. Default value", defaultValue, "has been used instead")

		return defaultValue
	}

	log.Printf("Value for %v has been defined as: %v\n", envKey, envValue)

	return envValue
}
