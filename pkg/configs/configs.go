package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	AppTransactionsBaseCurrency string
)

type Server struct {
	Port int `yaml:"port"`
}

type ExternalKeys struct {
	OpenExchangeRates string `yaml:"openExchangeRates"`
	Stripe            string `yaml:"stripe"`
}

type Currency struct {
	BaseCurrency    string `yaml:"baseCurrency"`
	CacheRates      bool   `yaml:"cacheRates"`
	CacheTTLSeconds int    `yaml:"cacheTTLSeconds"`
}

type Redis struct {
	Addr     string `yaml:"address"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Driver   string `yaml:"driver"`
}

type Configs struct {
	Database     Database     `yaml:"database"`
	Redis        Redis        `yaml:"redis"`
	Currency     Currency     `yaml:"currency"`
	ExternalKeys ExternalKeys `yaml:"externalKeys"`
	Server       Server       `yaml:"server"`
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
	AppTransactionsBaseCurrency = cfg.Currency.BaseCurrency

	return &cfg
}
