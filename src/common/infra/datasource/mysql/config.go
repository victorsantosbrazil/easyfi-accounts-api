package mysql

import (
	"fmt"
)

type (
	Config struct {
		Host      string           `yaml:"host"`
		Port      int              `yaml:"port"`
		User      string           `yaml:"user"`
		Password  string           `yaml:"password"`
		Database  string           `yaml:"database"`
		Migration *MigrationConfig `yaml:"migration"`
	}

	MigrationConfig struct {
		Source string `yaml:"source"`
	}
)

func (c *Config) GetUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
}
