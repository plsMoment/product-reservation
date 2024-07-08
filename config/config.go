package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "./config/config.yaml"

type Config struct {
	Server ServerConfig `yaml:"server"`
	Db     DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Addr string `yaml:"addr"`
}

type DBConfig struct {
	Addr     string `yaml:"addr"`
	Name     string
	Username string
	Password string
	SSLMode  string `yaml:"ssl_mode"`
}

func ParseConfig() (*Config, error) {
	var config Config
	config.Db.Name = os.Getenv("POSTGRES_DB")
	config.Db.Username = os.Getenv("POSTGRES_USER")
	config.Db.Password = os.Getenv("POSTGRES_PASSWORD")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("yaml unmarshal failed: %w", err)
	}

	return &config, nil
}
