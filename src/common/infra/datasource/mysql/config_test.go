package mysql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMysqlDataSourceGetUrl(t *testing.T) {
	config := Config{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "testuser",
		Password: "testpasswd",
		Database: "testdb",
	}
	expected := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	actual := config.GetUrl()
	assert.Equal(t, expected, actual)
}
