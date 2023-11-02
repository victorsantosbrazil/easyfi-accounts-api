//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE.go -package=$GOPACKAGE

package dao

import (
	"context"

	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
)

type InstitutionData struct {
	Id   int
	Name string
}

type PageInstitutionData = pagination.Page[InstitutionData]

type InstitutionDAO interface {
	GetPage(ctx context.Context, pageParams pagination.PageParams) PageInstitutionData
}

func NewInstitutionDAO() InstitutionDAO {
	return nil
}
