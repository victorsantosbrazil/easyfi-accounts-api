package app

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/api"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
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
	app := &App{echo: e, config: cfg, controllers: controllers}
	return app
}
