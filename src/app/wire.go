//go:build wireinject
// +build wireinject

//go install github.com/google/wire/cmd/wire@latest
//go:generate wire

package app

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/api"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/dao"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/repository"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/usecase"
)

var (
	DAOSet        = wire.NewSet(dao.NewInstitutionDAO)
	RepositorySet = wire.NewSet(repository.NewInstitutionRepository)
	UseCaseSet    = wire.NewSet(usecase.NewListInstitutionsUseCase)
	ControllerSet = wire.NewSet(api.NewV1Group, api.NewInstitutionController)
)

func NewApp() (*App, error) {

	wire.Build(
		config.ReadConfig,
		echo.New,
		DAOSet,
		RepositorySet,
		UseCaseSet,
		ControllerSet,
		api.NewControllers,
		newApp,
	)

	return &App{}, nil
}
