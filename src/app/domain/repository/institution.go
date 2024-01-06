//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE.go -package=$GOPACKAGE

package repository

import (
	"context"

	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/dao"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/entity"
)

type PageInstitution = pagination.Page[entity.Institution]

type InstitutionRepository interface {
	GetPage(ctx context.Context, pageParams pagination.PageParams) (PageInstitution, error)
}

type institutionRepositoryImpl struct {
	institutionDAO dao.InstitutionDAO
}

func (r *institutionRepositoryImpl) GetPage(ctx context.Context, pageParams pagination.PageParams) (PageInstitution, error) {
	pageData, err := r.institutionDAO.GetPage(ctx, pageParams)

	if err != nil {
		return PageInstitution{}, err
	}

	page := pagination.MapPage(pageData, func(data dao.InstitutionData) entity.Institution {
		return entity.Institution{
			Id:   data.Id,
			Name: data.Name,
		}
	})
	return PageInstitution(page), nil

}

func NewInstitutionRepository(institutionDAO dao.InstitutionDAO) InstitutionRepository {
	return &institutionRepositoryImpl{
		institutionDAO: institutionDAO,
	}
}
