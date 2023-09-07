package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DB struct {
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

type Config struct {
	Database DB `mapstructure:"db"`
}

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./config/")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("couldn't load config: %s", err)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("couldn't read config: %s", err)
	}

	return &c, nil
}
