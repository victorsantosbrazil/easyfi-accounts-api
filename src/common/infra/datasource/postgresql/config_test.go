package postgresql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresDataSourceGetUrl(t *testing.T) {
	config := Config{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "testuser",
		Password: "testpasswd",
		Database: "testdb",
		SslMode:  "disable",
	}
	expected := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.SslMode,
	)
	actual := config.GetUrl()
	assert.Equal(t, expected, actual)
}
