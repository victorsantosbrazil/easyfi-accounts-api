//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE.go -package=$GOPACKAGE

package service

import (
	"context"

	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/domain/entity"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/infra/dao"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/model/pagination"
)

type PageInstitution = pagination.Page[entity.Institution]

type InstitutionService interface {
	GetPage(ctx context.Context, pageParams pagination.PageParams) (PageInstitution, error)
}

type institutionServiceImpl struct {
	institutionDAO dao.InstitutionDAO
}

func (r *institutionServiceImpl) GetPage(ctx context.Context, pageParams pagination.PageParams) (PageInstitution, error) {
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

func NewInstitutionService(institutionDAO dao.InstitutionDAO) InstitutionService {
	return &institutionServiceImpl{
		institutionDAO: institutionDAO,
	}
}
