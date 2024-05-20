package config

import (
    "github.com/kelseyhightower/envconfig"
    "log"
)

type Config struct {
    ConfluenceAPIURL string `envconfig:"CONFLUENCE_API_URL"`
    ConfluenceUsername string `envconfig:"CONFLUENCE_USERNAME"`
    ConfluenceAPIToken string `envconfig:"CONFLUENCE_API_TOKEN"`
}

func LoadConfig() *Config {
    var cfg Config
    if err := envconfig.Process("", &cfg); err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
    return &cfg
}