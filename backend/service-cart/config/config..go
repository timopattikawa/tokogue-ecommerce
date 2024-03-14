package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func NewConfigInitiate() *Config {
	file, err := os.ReadFile("/config/config.yml")
	if err != nil {
		log.Fatalf("Error readfile : %v", err.Error())
	}

	var config = Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error unmarshal file : %v", err.Error())
	}

	return &config
}
