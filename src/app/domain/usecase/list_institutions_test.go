package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	pagination "github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/entity"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/repository"
)

func TestRun(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	institutionRepository := repository.NewMockInstitutionRepository(mockCtrl)
	usecase := NewListInstitutionsUseCase(institutionRepository)

	ctx := context.Background()
	pageRequest := pagination.PageParams{Page: 1}

	institutions := []entity.Institution{
		{Id: 1, Name: "Brazil bank"},
		{Id: 2, Name: "Bank of America"},
	}

	pageInstitutions := repository.PageInstitution{
		Pagination: pagination.Pagination{Page: 1},
		Items:      institutions,
	}

	institutionRepository.EXPECT().GetPage(ctx, pageRequest).Return(
		pageInstitutions)

	expectedPagination := pageInstitutions.Pagination
	expectedItems := make([]ListInstitutionsUseCaseResponseItem, len(institutions))

	for i, institution := range institutions {
		expectedItems[i] = ListInstitutionsUseCaseResponseItem{
			Id:   institution.Id,
			Name: institution.Name,
		}
	}

	expected := ListInstitutionsUseCaseResponse{Pagination: expectedPagination, Items: expectedItems}
	actual := usecase.Run(ctx, pageRequest)

	assert.Equal(t, expected, actual)
}
