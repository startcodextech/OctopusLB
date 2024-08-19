package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

const configFile = "config.json"

var (
	configInstance *Config
	configMutex    sync.Mutex
	once           sync.Once
)

type (
	Config struct {
		DHCP *DHCPConfig `json:"managers"`
	}
)

func SaveConfig() error {
	configMutex.Lock()
	defer configMutex.Unlock()

	data, err := json.MarshalIndent(configInstance, "", "  ")
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
	var err error
	once.Do(func() {
		configInstance, err = loadConfigFromFile()
		if err != nil {
			configInstance = &Config{}
		}
	})
	return configInstance, err
}

func loadConfigFromFile() (*Config, error) {
	configMutex.Lock()
	defer configMutex.Unlock()

	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config from load file: %w", err)
	}

	return &config, nil
}
