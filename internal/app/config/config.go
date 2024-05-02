package config

import (
	"flag"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	BindAddr               string        `toml:"bind_addr"`
	LogLevel               string        `toml:"log_level"`
	DatabaseUrl            string        `toml:"database_url"`
	AccessTokenPrivateKey  string        `toml:"access_token"`
	AccessTokenPublicKey   string        `toml:"access_token_public"`
	RefreshTokenPrivateKey string        `toml:"refresh_token"`
	RefreshTokenPublicKey  string        `toml:"refresh_token_public"`
	AccessTokenExpiresIn   time.Duration `toml:"access_token_expires_in"`
	RefreshTokenExpiresIn  time.Duration `toml:"refresh_token_expires_in"`
	AccessTokenMaxAge      int           `toml:"access_token_max_age"`
	RefreshTokenMaxAge     int           `toml:"refresh_token_max_age"`
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "internal/app/config/config.toml", "path to config file")
}

func NewConfig() *TomlConfig {
	flag.Parse()

	config := &TomlConfig{}
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
