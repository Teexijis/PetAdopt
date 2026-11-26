package config

import (
	"fmt"
	"os"
	"time"
	"go.yaml.in/yaml/v4"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	JWT JWTConfig `yaml:"-"`
	AccountServiceSettings AccountServiceSettings `yaml:"AccountServiceSettings"`
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
// из env
type JWTConfig struct {
    Secret string
    TTL time.Duration
}

type AccountServiceSettings struct {
    MinNameLen int `yaml:"min_name_len"`
    MaxNameLen int `yaml:"max_name_len"`
    MaxEmailLen int `yaml:"max_email_len"`
    MinPasswordLen int `yaml:"min_password_len"`
    MaxPasswordLen int `yaml:"max_password_len"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл конфигурации: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("не удалось разобрать файл yaml: %w", err)
	}

	config.Database.Username = os.Getenv("ACCOUNT_DB_USER")
	config.Database.Password = os.Getenv("ACCOUNT_DB_PASSWORD")
	if config.Database.Username == "" || config.Database.Password == "" {
		return nil, fmt.Errorf("не задано имя пользователя или пароль бд")
	}

    config.JWT.Secret = os.Getenv("JWT_SECRET")
    if config.JWT.Secret == "" {
        return nil, fmt.Errorf("не задан секретный ключ")
    }

    config.JWT.TTL, err = time.ParseDuration(os.Getenv("JWT_TTL"))
    if err != nil || config.JWT.TTL <= 0 {
        return nil, fmt.Errorf("некорректный срок действия TTL")
    }

	return &config, nil
}


