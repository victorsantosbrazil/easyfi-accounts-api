package config

import (
	"fmt"

	"github.com/victorsantosbrazil/financial-institutions-api/src/common/infra/datasource"
)

type Config struct {
	Port        int
	DataSources *datasource.DataSourcesConfig
}

func (c *Config) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}
