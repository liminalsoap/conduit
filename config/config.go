package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Postgresql
	Logger
	Http
}

type Postgresql struct {
	Url string
}

type Logger struct {
	Level string
}

type Http struct {
	Port   string
	Secret string
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, err
}
