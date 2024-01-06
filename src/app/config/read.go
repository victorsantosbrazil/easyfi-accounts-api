package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/util/path"
)

func ReadConfig() (cfg *Config, err error) {

	if err := loadConfig(); err != nil {
		return nil, err
	}

	profile := fmt.Sprintf("%v", viper.Get("profile"))
	if profile != "" {
		if err := loadProfileConfig(profile); err != nil {
			return nil, err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func ReadProfileConfig(profile string) (cfg *Config, err error) {
	if err := loadProfileConfig(profile); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadConfig() error {
	return loadConfigFromFile("config.yaml")
}

func loadProfileConfig(profile string) error {
	filename := fmt.Sprintf("config-%s.yaml", profile)
	return loadConfigFromFile(filename)
}

func loadConfigFromFile(fileName string) error {
	rootDir, err := path.FindProjectRootPath()
	if err != nil {
		return fmt.Errorf("could not find project root path: %s", err)
	}

	pathConfigFile := fmt.Sprintf("%s/%s", rootDir, fileName)

	viper.SetConfigFile(pathConfigFile)
	viper.SetConfigType("yaml")

	if err := viper.MergeInConfig(); err != nil {
		return fmt.Errorf("error loading environment variables: %s", err)
	}

	return nil
}
