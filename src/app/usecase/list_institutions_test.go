package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/domain/entity"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/service"
	pagination "github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/model/pagination"
)

func TestRun(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	institutionService := service.NewMockInstitutionService(mockCtrl)
	usecase := NewListInstitutionsUseCase(institutionService)

	t.Run("returns page of banks", func(t *testing.T) {
		ctx := context.Background()
		pageRequest := pagination.PageParams{Page: 1}

		institutions := []entity.Institution{
			{Id: 1, Name: "Brazil bank"},
			{Id: 2, Name: "Bank of America"},
		}

		pageInstitutions := service.PageInstitution{
			Pagination: pagination.Pagination{Page: 1},
			Items:      institutions,
		}

		institutionService.EXPECT().GetPage(ctx, pageRequest).Return(
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

		institutionService.EXPECT().GetPage(ctx, pageRequest).Return(
			service.PageInstitution{}, expectedErr)

		_, actualErr := usecase.Run(ctx, pageRequest)

		assert.ErrorIs(t, actualErr, expectedErr)
	})
}
