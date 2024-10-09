//go:build wireinject
// +build wireinject

//go install github.com/google/wire/cmd/wire@latest
//go:generate wire

package app

import (
	"github.com/google/wire"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/api/rest"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/config"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/infra/dao"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/service"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/usecase"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/echo"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/infra/datasource/postgresql"
)

var (
	DAOSet        = wire.NewSet(dao.NewInstitutionDAO)
	ServiceSet    = wire.NewSet(service.NewInstitutionService)
	UseCaseSet    = wire.NewSet(usecase.NewListInstitutionsUseCase)
	ControllerSet = wire.NewSet(rest.NewV1Group, rest.NewInstitutionController)
)

func NewApp() (*App, error) {
	wire.Build(
		config.ReadConfig,
		GetDataSourceConfig,
		echo.New,
		DAOSet,
		ServiceSet,
		UseCaseSet,
		ControllerSet,
		rest.NewControllers,
		newApp,
	)

	return &App{}, nil
}

func GetDataSourceConfig(cfg *config.Config) *postgresql.Config {
	return cfg.DataSource
}
