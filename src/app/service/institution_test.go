package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/domain/entity"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/infra/dao"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/model/pagination"
)

func TestGetPage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	institutionDAO := dao.NewMockInstitutionDAO(mockCtrl)
	institutionService := NewInstitutionService(institutionDAO)

	t.Run("returns page of banks", func(t *testing.T) {
		ctx := context.Background()
		pageParams := pagination.PageParams{
			Page: 1,
		}

		pageInstitutionData := dao.PageInstitutionData{
			Pagination: pagination.Pagination{
				Page: pageParams.Page,
			},
			Items: []dao.InstitutionData{
				{Id: 1, Name: "Brazil Bank"},
			},
		}

		institutionDAO.EXPECT().GetPage(ctx, pageParams).Return(pageInstitutionData, nil)

		expected := pagination.MapPage(pageInstitutionData, func(data dao.InstitutionData) entity.Institution {
			return entity.Institution{
				Id:   data.Id,
				Name: data.Name,
			}
		})

		actual, err := institutionService.GetPage(ctx, pageParams)

		if assert.NoError(t, err) {
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("returns error when fails to get page from dao", func(t *testing.T) {
		ctx := context.Background()
		pageParams := pagination.PageParams{
			Page: 1,
		}

		expectedErr := errors.New("error")
		institutionDAO.EXPECT().GetPage(ctx, pageParams).Return(dao.PageInstitutionData{}, expectedErr)

		_, actualErr := institutionService.GetPage(ctx, pageParams)

		assert.ErrorIs(t, actualErr, expectedErr)
	})

}
