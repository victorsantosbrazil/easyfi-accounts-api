package config

import (
	cmnconfig "github.com/victorsantosbrazil/financial-institutions-api/src/app/common/config"
)

func ReadConfig() (*Config, error) {
	var cfg Config

	readCfg := cmnconfig.ReadConfigOptions{
		ConfigName: "config",
		ConfigType: "yaml",
		ConfigPath: ".",
	}

	err := cmnconfig.ReadConfig(&cfg, readCfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
