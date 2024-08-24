package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Config struct {
	Port       int
	Datasource DatasourceConfig
}

type DatasourceConfig struct {
	Host string
	Port int
}

func TestReadConfig(t *testing.T) {

	t.Run("should read test config", func(t *testing.T) {
		cfg := &Config{}
		readCfgOptions := ReadConfigOptions{
			ConfigName: "config",
			ConfigType: "yaml",
			ConfigPath: ".",
		}

		err := ReadConfig(cfg, readCfgOptions)
		if assert.NoError(t, err) {
			assert.Equal(t, 8080, cfg.Port)
			assert.Equal(t, "testdb", cfg.Datasource.Host)
			assert.Equal(t, 5432, cfg.Datasource.Port)
		}
	})
}
