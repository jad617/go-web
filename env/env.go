package env

import (
	"log"
	"os"
	"strings"
)

func GetEnv(envVar string, defaultValue string) string {
	env := ""

	upperEnvVar := strings.ToUpper(envVar)
	lowerEnvVar := strings.ToLower(envVar)

	upperEnv := os.Getenv(upperEnvVar)
	lowerEnv := os.Getenv(lowerEnvVar)

	if upperEnv != "" {
		env = upperEnv
	} else if lowerEnv != "" {
		env = lowerEnv
	} else {
		log.Println("Default value", defaultValue, "has been used")
		return defaultValue
	}

	log.Printf("Value for %v has been defined as: %v\n", envVar, env)

	return env
}
