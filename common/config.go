/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
package common

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

var config *Configuration = nil
var configPath string = "./common/local.yaml"
var configTestPath string = "../common/local.yaml"

// Returns the local yaml configuration data
func GetConfig() *Configuration {
	if config == nil {
		yamlFile, err := ioutil.ReadFile(GetConfigPath())

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

func GetConfigPath() string {
	env := GetEnvironment()

	if env == "test" {
		return configTestPath
	}

	return configPath
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

	if env == "test" {
		return &configEnv.Test
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
