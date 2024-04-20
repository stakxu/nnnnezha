// config/loader.go
package config

import (
	"encoding/json"
	"fmt"
	"os"
	
	"github.com/nezhahq/agent/proto"
)

type Config struct {
	Servers []ServerConfig `json:"servers"`
}

type ServerConfig struct {
	Result       []ResultItem `json:"result"`
	Address      string       `json:"address"`
	ClientSecret string       `json:"client_secret"`
}

type ResultItem struct {
	Host       proto.Host `json:"host"`
	ID         uint64      `json:"id"`
	LastActive uint64      `json:"last_active"`
	Name       string      `json:"name"`
	Status     proto.State `json:"status"`
}

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
