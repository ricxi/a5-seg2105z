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
// and port for ListenAndServe to use to launch the server
func GetConfig(configPath string) string {

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var c Config
	err = json.Unmarshal(configFile, &c)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
