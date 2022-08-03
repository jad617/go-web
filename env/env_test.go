package env

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// SETUP
// Importantly you need to call Run() once you've done what you need
func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

type testEnv struct {
	envKey       string
	value        string
	defaultValue string
}

var testEnvValues = []testEnv{
	{envKey: "env1", value: "value1", defaultValue: "notEmptyDefault1"},
	{envKey: "env2", value: "value2", defaultValue: ""},
	{envKey: "", value: "", defaultValue: ""},
}

func TestGetEnv(t *testing.T) {
	for _, env := range testEnvValues {

		if env.envKey != "" {
			t.Setenv(env.envKey, env.value)
		}

		got := GetEnv(env.envKey, env.defaultValue)
		want := env.value

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

var testEnvDefaults = []testEnv{
	{envKey: "env1", value: "", defaultValue: "notEmptyDefault1"},
	{envKey: "", value: "", defaultValue: ""},
}

func TestGetEnvDefault(t *testing.T) {
	for _, env := range testEnvDefaults {

		if env.envKey != "" {
			t.Setenv(env.envKey, env.value)
		}

		got := GetEnv(env.envKey, env.defaultValue)
		want := env.defaultValue

		if got != want {
			t.Logf("got %v, wanted %v", got, want)
		}
	}
}
