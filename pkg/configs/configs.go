package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Driver   string `yaml:"driver"`
}

type Configs struct {
	Database database `yaml:"database"`
}

func NewConfigs() *Configs {
	file, err := os.Open("environment.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg Configs
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Failed to decode YAML file: %v", err)
	}

	return &cfg
}
