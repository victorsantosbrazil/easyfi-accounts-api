package repository

import (
	"context"
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

	institutionDAO.EXPECT().GetPage(ctx, pageParams).Return(pageInstitutionData)

	expected := pagination.MapPage(pageInstitutionData, func(data dao.InstitutionData) entity.Institution {
		return entity.Institution{
			Id:   data.Id,
			Name: data.Name,
		}
	})

	actual := institutionRepository.GetPage(ctx, pageParams)

	assert.Equal(t, expected, actual)

}
