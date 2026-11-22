package config

import (
	"fmt"
	"os"
	"go.yaml.in/yaml/v4"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

// из env пароль и юзернэйм 
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int	`yaml:"port"`
	Username string `yaml:"-"`
	Password string `yaml:"-"`
	DBName   string `yaml:"name"`
	SSLMode  string `yaml:"ssl_mode"`
}


func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	config.Database.Username = os.Getenv("ACCOUNT_DB_USER")
	config.Database.Password = os.Getenv("ACCOUNT_DB_PASSWORD")
	if config.Database.Username == "" || config.Database.Password == "" {
		return nil, fmt.Errorf("database username or password is not set")
	}

	return &config, nil
}


