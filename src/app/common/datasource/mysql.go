package datasource

import (
	"errors"
	"fmt"
)

type (
	MysqlDataSourcesConfig map[string]*MysqlDataSourceConfig

	MysqlDataSourceConfig struct {
		Host      string                          `yaml:"host"`
		Port      int                             `yaml:"port"`
		User      string                          `yaml:"user"`
		Password  string                          `yaml:"password"`
		Database  string                          `yaml:"database"`
		Migration *MysqlDataSourceMigrationConfig `yaml:"migration"`
	}

	MysqlDataSourceMigrationConfig struct {
		Source string `yaml:"source"`
	}
)

func ErrDataSourceNotFound(key string) error {
	return errors.New("data source not found for key " + key)
}

func (c MysqlDataSourcesConfig) Get(key string) (*MysqlDataSourceConfig, error) {
	ds, ok := c[key]
	if !ok {
		return nil, ErrDataSourceNotFound(key)
	}
	return ds, nil
}

func (c *MysqlDataSourceConfig) GetUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
}
