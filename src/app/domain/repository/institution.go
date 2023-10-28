//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE.go -package=$GOPACKAGE

package repository

import (
	"context"

	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/entity"
)

type PageInstitution pagination.Page[entity.Institution]

type InstitutionRepository interface {
	GetPage(ctx context.Context, pageParams pagination.PageParams) PageInstitution
}

type institutionRepositoryImpl struct{}
