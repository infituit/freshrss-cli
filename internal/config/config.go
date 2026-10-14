package config

import "os"

type Config struct {
	BaseURL     string
	Username    string
	APIPassword string
	APIKey      string
}

func MergeEnv(cfg Config) Config {
	if cfg.BaseURL == "" {
		cfg.BaseURL = os.Getenv("FRESHRSS_URL")
	}
	if cfg.Username == "" {
		cfg.Username = os.Getenv("FRESHRSS_USER")
	}
	if cfg.APIPassword == "" {
		cfg.APIPassword = os.Getenv("FRESHRSS_API_PASSWORD")
	}
	if cfg.APIKey == "" {
		cfg.APIKey = os.Getenv("FRESHRSS_API_KEY")
	}
	return cfg
}
