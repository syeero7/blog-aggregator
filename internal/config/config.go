package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFilepath, err := getConfigFilepath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(configFilepath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(name string) error {
	oldConfig, err := Read()
	if err != nil {
		return err
	}

	oldConfig.CurrentUserName = name
	newConfig, err := json.MarshalIndent(oldConfig, "", " ")
	if err != nil {
		return err
	}

	configFilepath, err := getConfigFilepath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFilepath, newConfig, 0o666); err != nil {
		return err
	}

	return nil
}

func getConfigFilepath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}
