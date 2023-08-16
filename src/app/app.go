package app

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/utils"
)

type App struct {
	echo   *echo.Echo
	config *config.Config
	logger echo.Logger
}

func (a *App) Start() {
	address := a.config.GetAddress()
	a.logger.Fatal(a.echo.Start(address))
}

func newApp(e *echo.Echo, cfg *config.Config) *App {
	a := &App{echo: e, config: cfg, logger: e.Logger}
	a.migrationUp()
	return a
}

func (a *App) migrationUp() {
	migration, err := utils.NewMigration(a.config.Database)
	if err != nil {
		a.logger.Fatal(err)
	}

	if migration.Up(); err != nil {
		a.logger.Fatal(err)
	}
}
