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
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/datasource"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/datasource/migration"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/testing/integration"
)

const countQueryRegex = `SELECT COUNT\(id\) FROM ` + _TABLE_NAME
const getPageQueryRegex = `SELECT id, name FROM ` + _TABLE_NAME + ` ORDER BY name ASC LIMIT [0-9]+ OFFSET [0-9]+`

func setupTestEnvironment() (datasourcesConfig *datasource.DataSourcesConfig, tearDown func() error) {
	dbConfig := &datasource.MysqlDataSourceConfig{
		Host:     "localhost",
		User:     "testuser",
		Password: "testpasswd",
		Database: "institution",
		Migration: &datasource.MysqlDataSourceMigrationConfig{
			Source: "file://../../../database/institution/schema",
		},
	}

	datasourcesConfig = &datasource.DataSourcesConfig{
		Mysql: &datasource.MysqlDataSourcesConfig{
			"db": dbConfig,
		},
	}

	mysqlConfig := integration.MysqlConfig{
		RootPassword: "root",
		User:         dbConfig.User,
		Password:     dbConfig.Password,
		Database:     dbConfig.Database,
	}

	mysqlResource := integration.RunMysql(mysqlConfig)

	port, err := strconv.Atoi(mysqlResource.GetPort("3306/tcp"))
	if err != nil {
		mysqlResource.Stop()
		log.Fatal(err)
	}

	dbConfig.Port = port

	err = setupMysqlDB(dbConfig)
	if err != nil {
		mysqlResource.Stop()
		log.Fatal(err)
	}

	return datasourcesConfig, mysqlResource.Stop

}

func setupMysqlDB(dsConfig *datasource.MysqlDataSourceConfig) error {
	m, err := migration.NewMysqlMigration(dsConfig)
	if err != nil {
		return err
	}
	return m.Up()
}

func TestCount(t *testing.T) {

	t.Run("should count elements inserted", func(t *testing.T) {
		dataSourcesCfg, tearDownTestEnvironment := setupTestEnvironment()
		defer tearDownTestEnvironment()

		dbConfig, err := dataSourcesCfg.Mysql.Get("db")
		if err != nil {
			t.Fatal(err)
		}

		institutions := []InstitutionData{
			{Id: 1, Name: "JPMorgan Chase & Co."},
			{Id: 2, Name: "Bank of America"},
			{Id: 3, Name: "Wells Fargo & Co."},
			{Id: 4, Name: "Citigroup Inc."},
			{Id: 5, Name: "HSBC Holdings plc"},
		}

		err = populateDatabase(dbConfig, institutions)
		if err != nil {
			t.Fatal(err)
		}

		institutionsDAO, err := NewInstitutionDAO(dataSourcesCfg)
		if err != nil {
			t.Fatal(err)
		}

		actualCount, err := institutionsDAO.Count(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, len(institutions), actualCount)
	})

	t.Run("when scanning query result fails then return error", func(t *testing.T) {
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
		dataSourcesCfg, tearDownTestEnvironment := setupTestEnvironment()
		defer tearDownTestEnvironment()

		dbConfig, err := dataSourcesCfg.Mysql.Get("db")

		if err != nil {
			t.Fatal(err)
		}

		institutions := []InstitutionData{
			{Id: 1, Name: "JPMorgan Chase & Co."},
			{Id: 2, Name: "Bank of America"},
			{Id: 3, Name: "Wells Fargo & Co."},
			{Id: 4, Name: "Citigroup Inc."},
			{Id: 5, Name: "HSBC Holdings plc"},
		}

		err = populateDatabase(dbConfig, institutions)
		if err != nil {
			t.Fatal(err)
		}

		institutionsDAO, err := NewInstitutionDAO(dataSourcesCfg)
		if err != nil {
			t.Fatal(err)
		}

		pageParams := pagination.PageParams{Size: 3, Page: 1}
		actualPage, err := institutionsDAO.GetPage(context.Background(), pageParams)
		if err != nil {
			t.Fatalf("got err when get page: %s", err)
		}

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

	t.Run("should return error when row scanning fails", func(t *testing.T) {
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
		assert.IsType(t, datasource.ScanRowError{}, actualErr)
	})

}

func populateDatabase(cfg *datasource.MysqlDataSourceConfig, institutions []InstitutionData) error {
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
