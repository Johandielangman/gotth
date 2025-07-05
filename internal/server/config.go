package server

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port string `envconfig:"APP_PORT" default:"8080"`
	Env  string `envconfig:"APP_ENV" default:"development"`
}

func loadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
