//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire
//go:generate wire ./...

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/main/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/main/controllers"
	v1controllers "github.com/victorsantosbrazil/financial-institutions-api/src/main/controllers/v1"
)

var V1ControllersSet = wire.NewSet(v1controllers.NewV1Group, v1controllers.NewInstitutionsController)

func NewApp() (*App, error) {

	wire.Build(
		config.ReadConfig,
		echo.New,
		V1ControllersSet,
		controllers.NewControllerManager,
		newApp,
	)

	return &App{}, nil
}
