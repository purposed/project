package main

import (
	"os/user"
	"path"

	"github.com/spf13/viper"
)

var _curUsr, _ = user.Current()
var configDir = path.Join(_curUsr.HomeDir, ".config", "dalloriam")

var defaultConfig = Config{
	DefaultOwner:      "dalloriam",
	PreferredProvider: "github.com",
}

// Config regroups all project config.
type Config struct {
	DefaultOwner      string `json:"default_owner" mapstructure:"default_owner"`
	PreferredProvider string `json:"preferred_provider" mapstructure:"preferred_provider"`
	RootPath          string `json:"root_path" mapstructure:"root_path"`
}

func loadConfigFromDisk() (*Config, error) {
	viper.SetConfigName("project")
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfig() *Config {
	cfg, err := loadConfigFromDisk()
	if err != nil {
		return &defaultConfig
	}

	if cfg.RootPath == "" {
		cfg.RootPath = _curUsr.HomeDir
	}

	return cfg
}
