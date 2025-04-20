package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Config represents the application configuration stored in the JSON file
type Config struct {
	DBURL           string `json:"db_connection_string"` // Match JSON key
	CurrentUserName string `json:"current_user_name,omitempty"`
}

// Read loads the configuration from the config file in the user's home directory
func Read() (Config, error) {
	var cfg Config

	configPath, err := getConfigFilePath()
	if err != nil {
		return cfg, fmt.Errorf("error getting config file path: %w", err)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return cfg, fmt.Errorf("error reading config file: %w", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("error parsing config file: %w", err)
	}

	return cfg, nil
}

// SetUser updates the current user in the configuration and writes it to disk
func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

// getConfigFilePath returns the absolute path to the config file
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}

	return filepath.Join(homeDir, configFileName), nil
}

// write saves the configuration to the config file
func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting config file path: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}
