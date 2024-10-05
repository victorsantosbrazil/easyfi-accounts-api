package config

import (
	"fmt"

	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/infra/datasource/mysql"
)

type Config struct {
	Port       int
	DataSource *mysql.Config
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}
