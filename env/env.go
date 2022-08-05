package env

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	errEmptyEnvKey = fmt.Errorf("function got envKey: \"\" --> Was expecting a non-empty string")
	errEmptyValues = fmt.Errorf("got envValue: \"\" and defaultValue: \"\". Was expecting at least envKey to have a non-empty string value or defaultValue to be a valid non-empty string parameter")
)

func GetEnv(envKey string, defaultValue string) (string, error) {
	var envValue string

	if envKey == "" {
		return "", errEmptyEnvKey
	}

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
		return "", errEmptyValues
	}

	log.Printf("Value for %v has been defined as: %v\n", envKey, envValue)

	return envValue, nil
}
