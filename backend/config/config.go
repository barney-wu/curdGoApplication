package config

import (
	"encoding/json"
	"log"
	"os"
)

//Configuration is the output of func LoadConfig
type Configuration struct {
	Host       string
	HTTPPort   string
	DbHost     string
	DbName     string
	DbUsername string
	DbPassword string
}

//LoadConfig decode configuration in json
func LoadConfig(file string) Configuration {
	var config Configuration
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	return config
}
