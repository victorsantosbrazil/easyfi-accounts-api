package dao

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/model/pagination"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/infra/datasource/postgresql"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/testing/integration"
)

const countQueryRegex = `SELECT COUNT\(id\) FROM ` + _TABLE_NAME
const getPageQueryRegex = `SELECT id, name FROM ` + _TABLE_NAME + ` ORDER BY name ASC LIMIT [0-9]+ OFFSET [0-9]+`

func setupTestEnvironment() (dsConfig *postgresql.Config, tearDown func() error) {
	dsConfig = &postgresql.Config{
		Host:     "localhost",
		User:     "testuser",
		Password: "testpasswd",
		Database: "account",
		SslMode:  "disable",
		Migration: &postgresql.MigrationConfig{
			Source: "file://../../../../database/schema",
		},
	}

	postgresqlConfig := integration.PostgresConfig{
		User:     dsConfig.User,
		Password: dsConfig.Password,
		Database: dsConfig.Database,
	}

	postgresqlResource := integration.RunPostgres(postgresqlConfig)

	port, err := strconv.Atoi(postgresqlResource.GetPort("5432/tcp"))
	if err != nil {
		postgresqlResource.Stop()
		log.Fatal(err)
	}

	dsConfig.Port = port

	err = setupDB(dsConfig)
	if err != nil {
		postgresqlResource.Stop()
		log.Fatal(err)
	}

	return dsConfig, postgresqlResource.Stop
}

func setupDB(dsConfig *postgresql.Config) error {
	m, err := postgresql.NewMigration(dsConfig)
	if err != nil {
		return err
	}
	return m.Up()
}

func TestCount(t *testing.T) {

	t.Run("should count elements inserted", func(t *testing.T) {
		dsConfig, tearDownTestEnvironment := setupTestEnvironment()
		defer tearDownTestEnvironment()

		institutions := []InstitutionData{
			{Id: 1, Name: "JPMorgan Chase & Co."},
			{Id: 2, Name: "Bank of America"},
			{Id: 3, Name: "Wells Fargo & Co."},
			{Id: 4, Name: "Citigroup Inc."},
			{Id: 5, Name: "HSBC Holdings plc"},
		}

		err := populateDatabase(dsConfig, institutions)
		if err != nil {
			t.Fatal(err)
		}

		institutionsDAO, err := NewInstitutionDAO(dsConfig)
		if err != nil {
			t.Fatal(err)
		}

		actualCount, err := institutionsDAO.Count(context.Background())
		if assert.NoError(t, err) {
			assert.Equal(t, len(institutions), actualCount)
		}
	})

	t.Run("should return error when row scan fails", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		institutionsDAO := &institutionDAOImpl{
			db: db,
		}

		expectedErr := errors.New("fail")
		countRows := sqlmock.NewRows([]string{"count(id)"}).AddRow(0).RowError(0, expectedErr)
		dbmock.ExpectQuery(countQueryRegex).WillReturnRows(countRows)

		_, actualErr := institutionsDAO.Count(context.Background())
		assert.ErrorIs(t, actualErr, expectedErr)
	})
}

func TestGetPage(t *testing.T) {
	t.Run("should return page with institutions", func(t *testing.T) {
		dsConfig, tearDownTestEnvironment := setupTestEnvironment()
		defer tearDownTestEnvironment()

		institutions := []InstitutionData{
			{Id: 1, Name: "JPMorgan Chase & Co."},
			{Id: 2, Name: "Bank of America"},
			{Id: 3, Name: "Wells Fargo & Co."},
			{Id: 4, Name: "Citigroup Inc."},
			{Id: 5, Name: "HSBC Holdings plc"},
		}

		err := populateDatabase(dsConfig, institutions)
		if err != nil {
			t.Fatal(err)
		}

		institutionsDAO, err := NewInstitutionDAO(dsConfig)
		if err != nil {
			t.Fatal(err)
		}

		pageParams := pagination.PageParams{Size: 3, Page: 1}
		actualPage, err := institutionsDAO.GetPage(context.Background(), pageParams)

		if assert.NoError(t, err) {
			expectedPage := PageInstitutionData{
				Pagination: pagination.Pagination{
					Page:          pageParams.Page,
					Size:          pageParams.Size,
					TotalPages:    2,
					TotalElements: len(institutions),
				},
				Items: []InstitutionData{ // will return institutions sorted by name
					institutions[1], institutions[3], institutions[4],
				},
			}

			assert.Equal(t, expectedPage, actualPage)
		}
	})

	t.Run("should return error when count fails", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		institutionsDAO := &institutionDAOImpl{
			db: db,
		}

		expectedErr := errors.New("fail")
		dbmock.ExpectQuery(countQueryRegex).WillReturnError(expectedErr)

		pageParams := pagination.PageParams{Size: 3, Page: 1}
		_, actualErr := institutionsDAO.GetPage(context.Background(), pageParams)
		assert.ErrorIs(t, actualErr, expectedErr)
	})

	t.Run("should return error when query fails", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		institutionsDAO := &institutionDAOImpl{
			db: db,
		}

		countRows := sqlmock.NewRows([]string{"count(id)"}).AddRow(20)
		dbmock.ExpectQuery(countQueryRegex).WillReturnRows(countRows)

		expectedErr := errors.New("fail")
		dbmock.ExpectQuery(getPageQueryRegex).WillReturnError(expectedErr)

		pageParams := pagination.PageParams{Size: 3, Page: 1}
		_, actualErr := institutionsDAO.GetPage(context.Background(), pageParams)
		assert.ErrorIs(t, actualErr, expectedErr)
	})

	t.Run("should return error when row scan fails", func(t *testing.T) {
		db, dbmock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		institutionsDAO := &institutionDAOImpl{
			db: db,
		}

		countRows := sqlmock.NewRows([]string{"count(id)"}).AddRow(20)
		dbmock.ExpectQuery(countQueryRegex).WillReturnRows(countRows)

		getPageRows := sqlmock.NewRows([]string{"id", "name", "unknown"}).AddRow(1, "Bank", "")
		dbmock.ExpectQuery(getPageQueryRegex).WillReturnRows(getPageRows)
		pageParams := pagination.PageParams{Size: 3, Page: 1}
		_, actualErr := institutionsDAO.GetPage(context.Background(), pageParams)
		assert.IsType(t, postgresql.ScanRowError{}, actualErr)
	})

}

func populateDatabase(cfg *postgresql.Config, institutions []InstitutionData) error {
	db, err := sql.Open("postgres", cfg.GetUrl())
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert each institution into the database
	for _, institution := range institutions {
		_, err := db.Exec("INSERT INTO institution (id, name) VALUES ($1, $2)", institution.Id, institution.Name)

		if err != nil {
			return err
		}
	}

	return nil

}
