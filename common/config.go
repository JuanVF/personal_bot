package common

import (
	"fmt"
	"io/ioutil"

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

		config = &Configuration{}

		err = yaml.Unmarshal(yamlFile, config)

		if err != nil {
			GetLogger().Error("Config", fmt.Sprintf("Unmarshal: %v", err))
		}
	}

	return config
}
