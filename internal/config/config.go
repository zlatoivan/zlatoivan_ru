package config

import (
	"os"
)

type Config struct {
	Server `yaml:"server"`
	Pg     `yaml:"postgres"`
}

type Server struct {
	HTTPPort string `yaml:"http_port" env-default:"7070"`
}

type Pg struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	DBname   string `yaml:"dbname" env-default:"postgres"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
}

func New() (Config, error) {
	_ = os.Getenv("CONFIG_PATH")
	//if configPath == "" {
	//	return Config{}, fmt.Errorf("CONFIG_PATH is not set")
	//}
	//
	//_, err := os.Stat(configPath)
	//if os.IsNotExist(err) {
	//	return Config{}, fmt.Errorf("config file %s not found", configPath)
	//}
	//
	//var cfg Config
	//
	//err = cleanenv.ReadConfig(configPath, &cfg)
	//if err != nil {
	//	return Config{}, fmt.Errorf("cleanenv.ReadConfig: %w", err)
	//}

	cfg := Config{
		Server: Server{
			HTTPPort: "7070",
		},
	}

	return cfg, nil
}
