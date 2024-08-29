package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string         `yaml:"env" env-default:"local"`
	Server ServerConfig   `yaml:"server"`
	DB     DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type AppConfig struct {
	ExpiriesTime string
	AppSecret    string
	SrvPort      string
}

// MustLoad такое название согласованнон и означает что
// функция не может возвращать ошибку а сразу паникует
func MustLoad() *Config {
	var cfg Config
	path := "config/config.yml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found in path: " + path)
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
