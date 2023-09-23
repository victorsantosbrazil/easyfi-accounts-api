package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/main/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/main/controllers"
	"github.com/victorsantosbrazil/financial-institutions-api/src/main/utils"
)

type App struct {
	echo              *echo.Echo
	config            *config.Config
	controllerManager *controllers.ControllerManager
}

func (a *App) Start() {
	addr := a.config.GetAddress()
	a.echo.Logger.Fatal(a.echo.Start(addr))
}

func newApp(e *echo.Echo, cfg *config.Config, controllerManager *controllers.ControllerManager) *App {
	a := &App{echo: e, config: cfg, controllerManager: controllerManager}
	// a.migrationUp()
	return a
}

func (a *App) migrationUp() {
	migration, err := utils.NewMigration(a.config.Database)
	if err != nil {
		log.Fatal(err)
	}

	if migration.Up(); err != nil {
		log.Fatal(err)
	}
}
