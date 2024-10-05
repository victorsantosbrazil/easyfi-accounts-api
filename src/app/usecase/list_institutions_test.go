package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	pagination "github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/model/pagination"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/domain/entity"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/domain/repository"
)

func TestRun(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	institutionRepository := repository.NewMockInstitutionRepository(mockCtrl)
	usecase := NewListInstitutionsUseCase(institutionRepository)

	t.Run("returns page of banks", func(t *testing.T) {
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
			pageInstitutions, nil)

		expectedPagination := pageInstitutions.Pagination
		expectedItems := make([]ListInstitutionsUseCaseResponseItem, len(institutions))

		for i, institution := range institutions {
			expectedItems[i] = ListInstitutionsUseCaseResponseItem{
				Id:   institution.Id,
				Name: institution.Name,
			}
		}

		actual, err := usecase.Run(ctx, pageRequest)
		expected := ListInstitutionsUseCaseResponse{Pagination: expectedPagination, Items: expectedItems}
		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("returns error when fails to get page of institutions", func(t *testing.T) {
		ctx := context.Background()
		pageRequest := pagination.PageParams{Page: 1}

		expectedErr := errors.New("error")

		institutionRepository.EXPECT().GetPage(ctx, pageRequest).Return(
			repository.PageInstitution{}, expectedErr)

		_, actualErr := usecase.Run(ctx, pageRequest)

		assert.ErrorIs(t, actualErr, expectedErr)
	})
}
