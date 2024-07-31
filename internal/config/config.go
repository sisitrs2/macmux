package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	CommandModeToggleKey string `json:"command_mode_toggle_key"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("Error opening config file: %w", err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("Error parsing config file: %w", err)
	}

	return &config, nil
}

// PrintConfig prints the current configuration to the console.
func (c *Config) PrintConfig() {
	fmt.Printf("Current Configuration:\n")
	fmt.Printf("Command Mode Toggle Key: %s\n", c.CommandModeToggleKey)
}
