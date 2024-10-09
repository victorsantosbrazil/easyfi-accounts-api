package postgresql

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Migration struct {
	migrate *migrate.Migrate
}

func (m *Migration) Up() error {
	return m.migrate.Up()
}

func NewMigration(dsConfig *Config) (*Migration, error) {
	db, err := sql.Open("postgres", dsConfig.GetUrl())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(dsConfig.Migration.Source, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return &Migration{
		migrate: m,
	}, nil
}
