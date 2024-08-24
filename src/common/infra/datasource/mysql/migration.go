package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlmigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	migrate *migrate.Migrate
}

func (m *Migration) Up() error {
	return m.migrate.Up()
}

func NewMigration(dsConfig *Config) (*Migration, error) {
	db, err := sql.Open("mysql", dsConfig.GetUrl())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	driver, err := mysqlmigrate.WithInstance(db, &mysqlmigrate.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(dsConfig.Migration.Source, dsConfig.Database, driver)
	if err != nil {
		return nil, err
	}

	return &Migration{
		migrate: m,
	}, nil
}
