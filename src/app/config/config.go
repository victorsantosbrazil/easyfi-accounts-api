package config

import (
	"fmt"

	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/infra/datasource/postgresql"
)

type Config struct {
	Port       int
	DataSource *postgresql.Config
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}
