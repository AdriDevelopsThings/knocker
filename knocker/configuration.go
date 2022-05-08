package knocker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PortConfig struct {
	ListenAddress string `json:"knock_listen_address"`
	OpenPort      string `json:"open_port"`
	TTL           int    `json:"ttl"`
}

var ports []PortConfig = make([]PortConfig, 0)

func ReadConfig() error {
	filepath := os.Getenv("KNOCKER_CONFIGURATION_FILE")
	if len(filepath) == 0 {
		filepath = "configuration.json"
	}
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Error ReadFile configuration file %q:  %v\n", filepath, err)
	}
	err = json.Unmarshal(yamlFile, &ports)

	if err != nil {
		return fmt.Errorf("Error while unmarshal configuration file %q:  %v\n", filepath, err)
	}
	return nil
}
