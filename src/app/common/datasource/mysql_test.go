package datasource

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMysqlDataSourceConfig(t *testing.T) {
	t.Run("should return config", func(t *testing.T) {
		expected := &MysqlDataSourceConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "testuser",
			Password: "testpasswd",
			Database: "testdb",
		}

		dssConfig := MysqlDataSourcesConfig{
			"db": expected,
		}

		result, err := dssConfig.Get("db")
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("should throw error when datasource does not exists by key", func(t *testing.T) {
		name := "db"
		dssConfig := MysqlDataSourcesConfig{}
		_, err := dssConfig.Get(name)
		assert.EqualError(t, err, ErrDataSourceNotFound(name).Error())
	})
}

func TestMysqlDataSourceGetUrl(t *testing.T) {
	config := MysqlDataSourceConfig{
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
