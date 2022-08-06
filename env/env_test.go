package env_test

import (
	"errors"
	"go-web/env"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

type testEnv struct {
	testName      string
	envKey        string
	envValue      string
	defaultValue  string
	expectedValue string
	expectedError error
}

var errTest = errors.New("")

func TestGetEnv(t *testing.T) {
	testEnvValues := []testEnv{
		{testName: "1 envValue takes precedence", envKey: "env1", envValue: "value1", defaultValue: "notEmptyDefault1", expectedValue: "value1", expectedError: nil},
		{testName: "1.1 envKey capitalized (subtest)", envKey: "ENV4", envValue: "value4", defaultValue: "notEmptyDefault4", expectedValue: "value4", expectedError: nil},
		{testName: "2 only envValue is defined", envKey: "env2", envValue: "value2", defaultValue: "", expectedValue: "value2", expectedError: nil},
		{testName: "2.1 envKey capitialized (subtest)", envKey: "ENV5", envValue: "value5", defaultValue: "", expectedValue: "value5", expectedError: nil},
		{testName: "3 only defaultValue is defined", envKey: "env3", envValue: "", defaultValue: "notEmptyDefault3", expectedValue: "notEmptyDefault3", expectedError: nil},
		{testName: "3.1 envKey capitialized (subtest)", envKey: "ENV6", envValue: "", defaultValue: "notEmptyDefault6", expectedValue: "notEmptyDefault6", expectedError: nil},
		{testName: "4 is Error when only envKey defined", envKey: "env11", envValue: "", defaultValue: "", expectedError: errTest},
		{testName: "4.1 is Error envKey capitalized (subtest)", envKey: "ENV12", envValue: "", defaultValue: "", expectedError: errTest},
		{testName: "5 is Error when all fields are empty", envKey: "", envValue: "", defaultValue: "", expectedError: errTest},
		{testName: "6 is Error when only envValue is defined", envKey: "", envValue: "value2", defaultValue: "", expectedError: errTest},
		{testName: "7 is Error when only envValue & defaultValue are defined", envKey: "", envValue: "value3", defaultValue: "notEmptyDefault3", expectedError: errTest},
		{testName: "8 is Error when only defaultValue is defined", envKey: "", envValue: "", defaultValue: "notEmptyDefault4", expectedError: errTest},
	}

	for _, envTest := range testEnvValues {
		t.Run(envTest.testName, func(t *testing.T) {
			var (
				errIsNil           bool // Default value is false
				expectedErrorIsNil bool // Default value is false
			)

			// Set the ENV var if envKey is defined
			if envTest.envKey != "" {
				t.Setenv(envTest.envKey, envTest.envValue)
			}

			// Fetch envValue will be returned from the function if defined
			// If not, defaultValue will be instead
			got, err := env.GetEnv(envTest.envKey, envTest.defaultValue)
			want := envTest.expectedValue

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}

			// Define err into "true" value if nil
			if err == nil {
				errIsNil = true
			}

			// Define envTest.expectedError into "true" value if nil
			if envTest.expectedError == nil {
				expectedErrorIsNil = true
			}

			// Both bool vars should have the same value
			// If not, it will FAIL
			if !assert.Equal(t, expectedErrorIsNil, errIsNil) {
				t.Errorf("Error not catched, error msg is: %v", err)
			}
		})
	}
}
