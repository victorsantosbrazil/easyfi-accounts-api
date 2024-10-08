//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE.go -package=$GOPACKAGE

package dao

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/model/pagination"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/infra/datasource/postgresql"
)

const _TABLE_NAME = "institution"

type InstitutionData struct {
	Id   int
	Name string
}

type PageInstitutionData = pagination.Page[InstitutionData]

type InstitutionDAO interface {
	Count(ctx context.Context) (int, error)
	GetPage(ctx context.Context, pageParams pagination.PageParams) (PageInstitutionData, error)
}

type institutionDAOImpl struct {
	db *sql.DB
}

func (d *institutionDAOImpl) Count(ctx context.Context) (int, error) {
	var totalElements int
	countQuery := "SELECT COUNT(id) FROM " + _TABLE_NAME
	err := d.db.QueryRowContext(ctx, countQuery).Scan(&totalElements)
	return totalElements, err
}

func (d *institutionDAOImpl) GetPage(ctx context.Context, pageParams pagination.PageParams) (PageInstitutionData, error) {
	offset := (pageParams.Page - 1) * pageParams.Size
	query := fmt.Sprintf("SELECT id, name FROM %s ORDER BY name ASC LIMIT %d OFFSET %d", _TABLE_NAME, pageParams.Size, offset)

	totalElements, err := d.Count(ctx)
	if err != nil {
		return PageInstitutionData{}, err
	}

	rows, err := d.db.QueryContext(ctx, query)
	if err != nil {
		return PageInstitutionData{}, err
	}
	defer rows.Close()

	institutions := make([]InstitutionData, 0)

	for rows.Next() {
		var institution InstitutionData
		err := rows.Scan(&institution.Id, &institution.Name)
		if err != nil {
			return PageInstitutionData{}, postgresql.NewScanRowError(err.Error())
		}
		institutions = append(institutions, institution)
	}

	var totalPages = (totalElements + pageParams.Size - 1) / pageParams.Size

	return PageInstitutionData{
		Pagination: pagination.Pagination{
			Page:          pageParams.Page,
			Size:          pageParams.Size,
			TotalPages:    totalPages,
			TotalElements: totalElements,
		},
		Items: institutions,
	}, nil
}

func NewInstitutionDAO(config *postgresql.Config) (InstitutionDAO, error) {

	db, err := sql.Open("postgres", config.GetUrl())
	if err != nil {
		return nil, fmt.Errorf("error opening connection to database db: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database db: %s", err)
	}

	return &institutionDAOImpl{
		db: db,
	}, nil

}
