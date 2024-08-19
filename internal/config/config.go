package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

const configFile = "config.json"

var configMutex sync.Mutex

type (
	Config struct {
		DHCP DHCPConfig `json:"dhcp"`
	}
)

func SaveConfig(config *Config) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func LoadConfig() (*Config, error) {
	configMutex.Lock()
	defer configMutex.Unlock()

	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{
				DHCP: DHCPConfig{},
			}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
