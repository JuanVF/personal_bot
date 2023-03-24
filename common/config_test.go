package common

import (
	"testing"
)

func TestEnv_env_should_be_set_to_test(t *testing.T) {
	tName := "TestEnv - it can detect the test environment"
	env := GetEnvironment()

	expected := "test"

	if env != expected {
		logger.TestError(tName, expected, env, t)
	}
}
