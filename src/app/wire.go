//go:build wireinject
// +build wireinject

//go install github.com/google/wire/cmd/wire@latest
//go:generate wire

package app

import (
	"github.com/google/wire"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/api"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/usecase"
	"github.com/victorsantosbrazil/financial-institutions-api/src/common/app/echo"
	"github.com/victorsantosbrazil/financial-institutions-api/src/common/infra/datasource"
	"github.com/victorsantosbrazil/financial-institutions-api/src/domain/repository"
	"github.com/victorsantosbrazil/financial-institutions-api/src/infra/dao"
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
		GetDataSourcesConfig,
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

func GetDataSourcesConfig(cfg *config.Config) *datasource.DataSourcesConfig {
	return cfg.DataSources
}
