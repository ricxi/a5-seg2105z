// Author: Richard Xiong

package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config struct is used to load information from a config file
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// GetConfig receives the path to the config file
// as an argument and returns a string of the host
// and port for ListenAndServe to use
func GetConfig(configPath string) string {

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var config Config
	if json.Unmarshal(configFile, &config); err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%s", config.Host, config.Port)
}
