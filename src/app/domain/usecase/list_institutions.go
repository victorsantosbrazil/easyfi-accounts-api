package usecase

import (
	"context"

	cmnmodel "github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model"
)

type ListInstitutionsUseCaseResponse struct {
	Pagination cmnmodel.Pagination                   `json:"pagination"`
	Items      []ListInstitutionsUseCaseResponseItem `json:"items"`
}

type ListInstitutionsUseCaseResponseItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ListInstitutionsUseCase interface {
	Run(ctx context.Context, pageRequest cmnmodel.PageRequest) ListInstitutionsUseCaseResponse
}

func NewListInstitutionsUseCase() ListInstitutionsUseCase {
	return nil
}
