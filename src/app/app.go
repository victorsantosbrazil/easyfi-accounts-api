package app

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/api"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/common/app/log"
	"github.com/victorsantosbrazil/financial-institutions-api/src/common/infra/datasource/migration"
)

type App struct {
	echo        *echo.Echo
	logger      log.Logger
	config      *config.Config
	controllers *api.Controllers
}

func (a *App) Start() error {
	err := a.setupDatabase()
	if err != nil {
		return err
	}

	addr := a.config.GetAddress()
	return a.echo.Start(addr)
}

func (a *App) setupDatabase() error {
	dbConfig, err := a.config.DataSources.Mysql.Get("db")
	if err != nil {
		return err
	}

	m, err := migration.NewMysqlMigration(dbConfig)
	if err != nil {
		return err
	}

	m.Up()
	return nil
}

func newApp(e *echo.Echo, cfg *config.Config, controllers *api.Controllers) *App {
	return &App{echo: e, config: cfg, controllers: controllers}
}
