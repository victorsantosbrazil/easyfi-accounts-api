package app

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/api"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/datasource/migration"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/log"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
)

type App struct {
	echo        *echo.Echo
	logger      log.Logger
	config      *config.Config
	controllers *api.Controllers
}

func (a *App) Start() {
	a.setupDatabase()
	addr := a.config.GetAddress()
	err := a.echo.Start(addr)
	if err != nil {
		a.logger.Fatal(err.Error())
	}
}

func (a *App) setupDatabase() {
	dbConfig, err := a.config.DataSources.Mysql.Get("db")
	if err != nil {
		a.logger.Fatal("Fail setting up database: " + err.Error())
	}

	m, err := migration.NewMysqlMigration(dbConfig)
	if err != nil {
		a.logger.Fatal("Fail setting up database: " + err.Error())
	}

	m.Up()
}

func newApp(e *echo.Echo, cfg *config.Config, logger log.Logger, controllers *api.Controllers) *App {
	return &App{echo: e, config: cfg, logger: logger, controllers: controllers}
}
