package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Config represents the configuration structure
type Config struct {
	Servers []ServerConfig `json:"servers"`
}

// ServerConfig represents the configuration for a single server
type ServerConfig struct {
	Result       []ResultItem `json:"result"`
	Address      string       `json:"address"`
	ClientSecret string       `json:"client_secret"`
}

// ResultItem represents an item in the result
type ResultItem struct {
	Host       proto.Host `json:"host"`
	ID         uint64      `json:"id"`
	LastActive uint64      `json:"last_active"`
	Name       string      `json:"name"`
	Status     proto.State `json:"status"`
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, fmt.Errorf("failed to decode config file: %v", err)
	}

	return config, nil
}
