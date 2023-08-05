package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

func (d *DatabaseConfig) GetUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.Username, d.Password, d.Host, d.Port, d.Name)
}

type Config struct {
	Port     int
	Database DatabaseConfig
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
