//go:build wireinject
// +build wireinject

//go install github.com/google/wire/cmd/wire@latest
//go:generate wire

package app

import (
	"github.com/google/wire"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/api"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/config"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/usecase"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/echo"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/infra/datasource/mysql"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/domain/repository"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/infra/dao"
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
		GetDataSourceConfig,
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

func GetDataSourceConfig(cfg *config.Config) *mysql.Config {
	return cfg.DataSource
}
