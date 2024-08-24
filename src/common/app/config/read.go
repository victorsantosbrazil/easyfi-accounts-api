package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ReadConfigOptions struct {
	ConfigName string
	ConfigType string
	ConfigPath string
}

func ReadConfig(i interface{}, options ReadConfigOptions) error {
	viper.SetConfigType(options.ConfigType)
	viper.AddConfigPath(options.ConfigPath)

	err := readDefaultConfig(options)
	if err != nil {
		return err
	}

	profiles := viper.GetStringSlice("profiles")
	err = readActiveProfilesConfig(profiles, options)
	if err != nil {
		return err
	}

	if err := viper.Unmarshal(i); err != nil {
		return err
	}

	return nil
}

func readDefaultConfig(options ReadConfigOptions) error {
	viper.SetConfigName(options.ConfigName)

	if err := viper.MergeInConfig(); err != nil {
		return err
	}
	return nil
}

func readActiveProfilesConfig(activeProfiles []string, options ReadConfigOptions) error {
	for _, activeProfile := range activeProfiles {
		configName := fmt.Sprintf("%s-%s", options.ConfigName, activeProfile)
		viper.SetConfigName(configName)
		if err := viper.MergeInConfig(); err != nil {
			return err
		}
	}
	return nil
}
