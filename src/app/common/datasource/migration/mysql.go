package migration

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/datasource"
)

type MysqlMigration struct {
	migrate *migrate.Migrate
}

func (m *MysqlMigration) Up() error {
	return m.migrate.Up()
}

func NewMysqlMigration(dsConfig *datasource.MysqlDataSourceConfig) (*MysqlMigration, error) {
	db, err := sql.Open("mysql", dsConfig.GetUrl())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(dsConfig.Migration.Source, dsConfig.Database, driver)
	if err != nil {
		return nil, err
	}

	return &MysqlMigration{
		migrate: m,
	}, nil
}
