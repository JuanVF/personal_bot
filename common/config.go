package common

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var config *Configuration = nil
var configPath string = "./common/local.yaml"

// Returns the local yaml configuration data
func GetConfig() *Configuration {
	if config == nil {
		yamlFile, err := ioutil.ReadFile(configPath)

		if err != nil {
			GetLogger().Error("Config", fmt.Sprintf("yamlFile.Get err   #%v ", err))
		}

		configEnv := &EnvironmentConfig{}

		err = yaml.Unmarshal(yamlFile, configEnv)

		if err != nil {
			GetLogger().Error("Config", fmt.Sprintf("Unmarshal: %v", err))

			panic(err)
		}

		config = GetConfigByEnvironment(configEnv)
	}

	return config
}

// It will return the configuration based on the current environment
func GetConfigByEnvironment(configEnv *EnvironmentConfig) *Configuration {
	env := GetEnvironment()

	if env == "development" {
		return &configEnv.Development
	}

	if env == "container" {
		return &configEnv.Container
	}

	return nil
}

// Returns the environment of the app
func GetEnvironment() string {
	env := os.Getenv("ENVIRONMENT")

	if env == "" {
		return "development"
	}

	return env
}
