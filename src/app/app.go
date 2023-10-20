package app

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/api"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/database"
)

type App struct {
	echo        *echo.Echo
	config      *config.Config
	controllers *api.Controllers
}

func (a *App) Start() {
	addr := a.config.GetAddress()
	a.echo.Logger.Fatal(a.echo.Start(addr))
}

func newApp(e *echo.Echo, cfg *config.Config, controllers *api.Controllers) *App {
	a := &App{echo: e, config: cfg, controllers: controllers}
	// a.migrationUp()
	return a
}

func (a *App) migrationUp() {
	migration, err := database.NewMigration(a.config.Database)
	if err != nil {
		log.Fatal(err)
	}

	if migration.Up(); err != nil {
		log.Fatal(err)
	}
}
