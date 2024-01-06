package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/dao"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/entity"
)

func TestGetPage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	institutionDAO := dao.NewMockInstitutionDAO(mockCtrl)
	institutionRepository := NewInstitutionRepository(institutionDAO)

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

		actual, _ := institutionRepository.GetPage(ctx, pageParams)

		assert.Equal(t, expected, actual)
	})

	t.Run("returns error when failed to load data", func(t *testing.T) {
		ctx := context.Background()
		pageParams := pagination.PageParams{
			Page: 1,
		}

		expectedErr := errors.New("error")
		institutionDAO.EXPECT().GetPage(ctx, pageParams).Return(dao.PageInstitutionData{}, expectedErr)

		_, actualErr := institutionRepository.GetPage(ctx, pageParams)

		assert.ErrorIs(t, expectedErr, actualErr)
	})

}
