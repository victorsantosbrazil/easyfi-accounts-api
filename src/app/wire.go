//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
)

func NewApp() (*App, error) {
	wire.Build(
		echo.New,
		config.ReadConfig,
		newApp,
	)

	return &App{}, nil
}
