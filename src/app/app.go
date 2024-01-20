package app

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/api"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/datasource/migration"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
)

type App struct {
	echo        *echo.Echo
	config      *config.Config
	controllers *api.Controllers
}

func (a *App) Start() {
	a.setupDatabase()
	addr := a.config.GetAddress()
	a.echo.Logger.Fatal(a.echo.Start(addr))
}

func (a *App) setupDatabase() {
	dbConfig, err := a.config.DataSources.Mysql.Get("db")
	if err != nil {
		a.echo.Logger.Fatal(err)
	}

	m, err := migration.NewMysqlMigration(dbConfig)
	if err != nil {
		a.echo.Logger.Fatal(err)
	}

	m.Up()
}

func newApp(e *echo.Echo, cfg *config.Config, controllers *api.Controllers) *App {
	return &App{echo: e, config: cfg, controllers: controllers}
}
