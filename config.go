package main

import (
	"os/user"
	"strings"

	"github.com/spf13/viper"
)

var configKeys = []string{"owner", "sourceroot", "sourceprovider"}

// Config regroups all project config.
type Config struct {
	DefaultOwner      string `json:"owner" mapstructure:"owner"`
	PreferredProvider string `json:"sourceprovider" mapstructure:"sourceprovider"`
	RootPath          string `json:"sourceroot" mapstructure:"sourceroot"`
}

func getConfig() (*Config, error) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("purposed")
	viper.AutomaticEnv()

	for _, k := range configKeys {
		if err := viper.BindEnv(k); err != nil {
			return nil, err
		}
	}
	viper.SetDefault("sourceprovider", "github.com")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	if config.DefaultOwner == "" {
		config.DefaultOwner = currentUser.Name
	}

	if config.RootPath == "" {
		config.RootPath = currentUser.HomeDir
	}

	return &config, nil
}
