//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package usecase

import (
	"context"

	"github.com/victorsantosbrazil/financial-institutions-api/src/common/app/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/domain/entity"
	"github.com/victorsantosbrazil/financial-institutions-api/src/domain/repository"
)

type ListInstitutionsUseCaseResponse pagination.Page[ListInstitutionsUseCaseResponseItem]

type ListInstitutionsUseCaseResponseItem struct {
	Id   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Brazil Bank"`
}

type ListInstitutionsUseCase interface {
	Run(ctx context.Context, pageParams pagination.PageParams) (ListInstitutionsUseCaseResponse, error)
}

type listInstitutionsUseCaseImpl struct {
	institutionRepository repository.InstitutionRepository
}

func (u *listInstitutionsUseCaseImpl) Run(ctx context.Context, pageParams pagination.PageParams) (ListInstitutionsUseCaseResponse, error) {
	pageInstitutions, err := u.institutionRepository.GetPage(ctx, pageParams)

	if err != nil {
		return ListInstitutionsUseCaseResponse{}, err
	}

	page := pagination.MapPage(pageInstitutions, func(institution entity.Institution) ListInstitutionsUseCaseResponseItem {
		return ListInstitutionsUseCaseResponseItem{
			Id:   institution.Id,
			Name: institution.Name,
		}
	})

	return ListInstitutionsUseCaseResponse(page), nil
}

func NewListInstitutionsUseCase(institutionRepository repository.InstitutionRepository) ListInstitutionsUseCase {
	return &listInstitutionsUseCaseImpl{
		institutionRepository: institutionRepository,
	}
}
