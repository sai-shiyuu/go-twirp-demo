package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Dev  FullConfig
	Prod FullConfig
}

type FullConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var Configs *FullConfig

func LoadConfig(path string) (*Config, error) {
	c := &Config{}
	if _, err := toml.DecodeFile(path, &c); err != nil {
		return nil, err
	}
	return c, nil
}
