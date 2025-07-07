// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: July 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

package server

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port    string `envconfig:"APP_PORT" default:"8080"`
	Env     string `envconfig:"APP_ENV" default:"development"`
	LogPath string `envconfig:"APP_LOG_PATH" default:"./logs/app.log"`
	AppName string `envconfig:"APP_NAME" default:"gotth"`
	Version string `envconfig:"APP_VERSION" default:"v1.0.0"`
}

// Parses the `Config` struct based on environment variables.
func loadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Loads the config file strictly. If an error occurs trying to create the config
// the process will error out
func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
