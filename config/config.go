package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port string `json:"port"`

	DBType     string `json:"db_type"`
	DBAddress  string `json:"db_address"`
	DBProtocal string `json:"db_protocol"`
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

var Config *Configuration

func LoadConfig() error {
	file, err := os.Open("config/config.json")

	if err != nil {
		return err
	}

	Config = new(Configuration)

	err = json.NewDecoder(file).Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
