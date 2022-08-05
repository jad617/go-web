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
		{testName: "1. Testing if envValue takes precedence over defaultValue", envKey: "env1", envValue: "value1", defaultValue: "notEmptyDefault1", expectedValue: "value1", expectedError: nil},
		{testName: "2. Testing when defaultValue is not defined but envValue is", envKey: "env2", envValue: "value2", defaultValue: "", expectedValue: "value2", expectedError: nil},
		{testName: "3. Testing when defaultValue is defined but not envValue", envKey: "env3", envValue: "", defaultValue: "notEmptyDefault3", expectedValue: "notEmptyDefault3", expectedError: nil},
		{testName: "4. Testing same as Test 1 but with upper case envKey", envKey: "ENV4", envValue: "value4", defaultValue: "notEmptyDefault4", expectedValue: "value4", expectedError: nil},
		{testName: "5. Testing same as Test 2 but with upper case envKey", envKey: "ENV5", envValue: "value5", defaultValue: "", expectedValue: "value5", expectedError: nil},
		{testName: "6. Testing same as Test 3 but with upper case envKey", envKey: "ENV6", envValue: "", defaultValue: "notEmptyDefault6", expectedValue: "notEmptyDefault6", expectedError: nil},
		{testName: "7. Testing that Error should be catched when all fields are empty strings", envKey: "", envValue: "", defaultValue: "", expectedError: errTest},
		{testName: "8. Testing that Error should be catched when envKey is empty string even if envValue is defined", envKey: "", envValue: "value2", defaultValue: "", expectedError: errTest},
		{testName: "9. Testing that Error should be catched when envKey is empty string even if envValue and defaultValue are defined", envKey: "", envValue: "value3", defaultValue: "notEmptyDefault3", expectedError: errTest},
		{testName: "10. Testing that Error should be catched when envKey is empty string even if defaultValue is defined ", envKey: "", envValue: "", defaultValue: "notEmptyDefault4", expectedError: errTest},
		{testName: "11. Testing that Error envKey defined when all other fields are empty", envKey: "env11", envValue: "", defaultValue: "", expectedError: errTest},
		{testName: "12. Testing that Error upper envKey defined when all other fields are empty", envKey: "ENV12", envValue: "", defaultValue: "", expectedError: errTest},
	}

	for _, envTest := range testEnvValues {
		t.Run(envTest.testName, func(t *testing.T) {
			if envTest.envKey != "" {
				t.Setenv(envTest.envKey, envTest.envValue)
			}

			got, err := env.GetEnv(envTest.envKey, envTest.defaultValue)
			want := envTest.expectedValue

			var errIsNil bool           // Default value is false
			var expectedErrorIsNil bool // Default value is false

			if err == nil {
				errIsNil = true
			}

			if envTest.expectedError == nil {
				expectedErrorIsNil = true
			}

			if !assert.Equal(t, errIsNil, expectedErrorIsNil) {
				t.Errorf("Error not catched, error msg is: %v", err)
			}

			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		})
	}
}
