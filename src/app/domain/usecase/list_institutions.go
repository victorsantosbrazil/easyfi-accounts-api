//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package usecase

import (
	"context"

	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/entity"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/repository"
)

type ListInstitutionsUseCaseResponse pagination.Page[ListInstitutionsUseCaseResponseItem]

type ListInstitutionsUseCaseResponseItem struct {
	CountryId int    `json:"countryId"`
	Name      string `json:"name"`
}

type ListInstitutionsUseCase interface {
	Run(ctx context.Context, pageParams pagination.PageParams) ListInstitutionsUseCaseResponse
}

type listInstitutionsUseCaseImpl struct {
	institutionRepository repository.InstitutionRepository
}

func (u *listInstitutionsUseCaseImpl) Run(ctx context.Context, pageParams pagination.PageParams) ListInstitutionsUseCaseResponse {
	pageInstitutions := u.institutionRepository.GetPage(ctx, pageParams)
	page := pagination.MapPage((pagination.Page[entity.Institution])(pageInstitutions), func(institution entity.Institution) ListInstitutionsUseCaseResponseItem {
		return ListInstitutionsUseCaseResponseItem{
			CountryId: institution.CountryId,
			Name:      institution.Name,
		}
	})

	return ListInstitutionsUseCaseResponse(page)
}

func NewListInstitutionsUseCase(institutionRepository repository.InstitutionRepository) ListInstitutionsUseCase {
	return &listInstitutionsUseCaseImpl{
		institutionRepository: institutionRepository,
	}
}
