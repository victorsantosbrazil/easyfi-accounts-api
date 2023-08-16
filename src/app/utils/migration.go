package utils

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
)

const _MIGRATIONS_DIR = "database/migrations"

type Migration interface {
	Up() error
}

type migrationImpl struct {
	migrate *migrate.Migrate
}

func NewMigration(databaseConfig config.DatabaseConfig) (Migration, error) {

	db, err := sql.Open(databaseConfig.Driver, databaseConfig.GetUrl())
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", _MIGRATIONS_DIR), databaseConfig.Driver, driver)
	if err != nil {
		return nil, err
	}

	return &migrationImpl{m}, nil

}

func (m *migrationImpl) Up() error {
	return m.migrate.Up()
}
