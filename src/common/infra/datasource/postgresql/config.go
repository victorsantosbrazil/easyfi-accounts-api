package postgresql

import (
	"fmt"
	"net/url"
)

type (
	Config struct {
		Host      string           `yaml:"host"`
		Port      int              `yaml:"port"`
		User      string           `yaml:"user"`
		Password  string           `yaml:"password"`
		Database  string           `yaml:"database"`
		SslMode   string           `yaml:"sslmode"`
		Migration *MigrationConfig `yaml:"migration"`
	}

	MigrationConfig struct {
		Source string `yaml:"source"`
	}
)

func (c *Config) GetUrl() string {
	// Construct the base connection string
	baseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		url.PathEscape(c.User),
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.Database)

	// Add query parameters
	query := url.Values{}

	// Use the specified sslmode if set, otherwise use "require" as the default
	if c.SslMode != "" {
		query.Add("sslmode", c.SslMode)
	} else {
		query.Add("sslmode", "require")
	}

	// Construct the final URL
	finalUrl := fmt.Sprintf("%s?%s", baseUrl, query.Encode())

	return finalUrl
}
