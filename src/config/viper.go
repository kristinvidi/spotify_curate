package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppEnv struct {
	Env string `mapstructure:"env"`
}

type AppClientInformation struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectURI  string `mapstructure:"redirect_uri"`
	State        string `mapstructure:"state"`
}

type Authentication struct {
	Scope         string `mapstructure:"scope"`
	GrantType     string `mapstructure:"grant_type"`
	ContentType   string `mapstructure:"content_type"`
	Authorization string `mapstructure:"authorization"`
}

type DB struct {
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

type Config struct {
	AppEnv         AppEnv               `mapstructure:"app_env"`
	AppClientInfo  AppClientInformation `mapstructure:"app_client_information"`
	Authentication Authentication       `mapstructure:"authentication"`
	Database       DB                   `mapstructure:"db"`
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
