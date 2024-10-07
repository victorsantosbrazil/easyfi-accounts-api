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
	mysql "github.com/victorsantosbrazil/easyfi-accounts-api/src/common/infra/datasource/mysql"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/testing/integration"
)

const countQueryRegex = `SELECT COUNT\(id\) FROM ` + _TABLE_NAME
const getPageQueryRegex = `SELECT id, name FROM ` + _TABLE_NAME + ` ORDER BY name ASC LIMIT [0-9]+ OFFSET [0-9]+`

func setupTestEnvironment() (dsConfig *mysql.Config, tearDown func() error) {
	dsConfig = &mysql.Config{
		Host:     "localhost",
		User:     "testuser",
		Password: "testpasswd",
		Database: "institution",
		Migration: &mysql.MigrationConfig{
			Source: "file://../../../../database/institution/schema",
		},
	}

	mysqlConfig := integration.MysqlConfig{
		RootPassword: "root",
		User:         dsConfig.User,
		Password:     dsConfig.Password,
		Database:     dsConfig.Database,
	}

	mysqlResource := integration.RunMysql(mysqlConfig)

	port, err := strconv.Atoi(mysqlResource.GetPort("3306/tcp"))
	if err != nil {
		mysqlResource.Stop()
		log.Fatal(err)
	}

	dsConfig.Port = port

	err = setupMysqlDB(dsConfig)
	if err != nil {
		mysqlResource.Stop()
		log.Fatal(err)
	}

	return dsConfig, mysqlResource.Stop

}

func setupMysqlDB(dsConfig *mysql.Config) error {
	m, err := mysql.NewMigration(dsConfig)
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
		assert.IsType(t, mysql.ScanRowError{}, actualErr)
	})

}

func populateDatabase(cfg *mysql.Config, institutions []InstitutionData) error {
	db, err := sql.Open("mysql", cfg.GetUrl())
	if err != nil {
		return err
	}
	defer db.Close()

	// Insert each institution into the database
	for _, institution := range institutions {
		_, err := db.Exec("INSERT INTO institution (id, name) VALUES (?, ?)", institution.Id, institution.Name)
		if err != nil {
			return err
		}
	}

	return nil

}
