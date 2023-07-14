package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port int
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func ReadConfig() (cfg *Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
