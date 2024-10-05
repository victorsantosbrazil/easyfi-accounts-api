package config

import (
	cmnconfig "github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/config"
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
