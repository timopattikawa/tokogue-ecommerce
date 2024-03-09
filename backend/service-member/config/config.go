package config

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func NewConfig() *Config {
	f, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatal("Fail to open config file")
	}

	decoder := yaml.NewDecoder(f)
	var config = &Config{}

	if err := decoder.Decode(config); err != nil {
		log.Fatal("Fail to decode config file")
	}

	return config
}
