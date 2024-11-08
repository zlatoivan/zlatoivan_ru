package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Config - основной конфиг
type Config struct {
	HTTPServer `yaml:"http_server"`
	Postgres   `yaml:"postgres"`
}

// HTTPServer - конфиг http-сервера
type HTTPServer struct {
	Port              string        `yaml:"port"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
}

// Postgres - конфиг postgres
type Postgres struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getConfigPath() (string, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return "", fmt.Errorf("CONFIG_PATH is not set")
	}

	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("config file %s not found", configPath)
	}

	return configPath, nil
}

func parseConfig(configPath string) (*Config, error) {
	var cfg Config
	configFileData, err := os.ReadFile(filepath.Clean(configPath))
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(%s): %w", configPath, err)
	}

	err = yaml.Unmarshal(configFileData, &cfg)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal(%s): %w", configPath, err)
	}

	return &cfg, nil
}

// New - создание конфига по файлу конфигурации
func New() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("getConfigPath: %w", err)
	}

	cfg, err := parseConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("parseConfig: %w", err)
	}

	return cfg, nil
}
