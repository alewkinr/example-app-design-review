package config

import (
	"os"
)

// Config — application configuration
type Config struct {
	Host string
	Port string

	*Log
}

// MustNewConfig — constructor for configuration struct, or panic if error
func MustNewConfig() *Config {
	cfg := &Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		Log: &Log{
			Level: os.Getenv("LOG.LEVEL"),
		},
	}

	if cfg.Host == "" || cfg.Port == "" {
		panic("⛔️HOST and PORT environment variables are required")
	}

	if cfg.Log.Level == "" {
		cfg.Log.Level = "info"
	}

	return cfg
}
