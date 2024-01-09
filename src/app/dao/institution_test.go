package dao

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os/exec"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/datasource"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/testing/integration"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/config"
)

const countQueryRegex = `SELECT COUNT\(id\) FROM ` + _TABLE_NAME
const getPageQueryRegex = `SELECT id, name FROM ` + _TABLE_NAME + ` ORDER BY name ASC LIMIT [0-9]+ OFFSET [0-9]+`

func setupTestEnvironment() (cfg *config.Config, tearDown func() error) {
	cfg, err := config.ReadProfileConfig("test")
	if err != nil {
		log.Fatal(err)
	}

	dsConfig, _ := cfg.DataSources.Mysql.Get("db")

	mysqlConfig := integration.MysqlConfig{
		RootPassword: "root",
		Database:     dsConfig.Database,
	}

	tearDownMysql := integration.RunMysql(mysqlConfig)

	err = setupMysqlDB(dsConfig)
	if err != nil {
		tearDownMysql()
		log.Fatal(err)
	}

	return cfg, tearDownMysql

}

func setupMysqlDB(dsConfig *datasource.MysqlDataSourceConfig) error {
	cmd := exec.Command(
		"bash",
		"./../../../dev/scripts/mysql/setup-database.sh",
		"-u", dsConfig.User,
		"-p", dsConfig.Password,
		"-d", dsConfig.Database,
	)
	return cmd.Run()
}

func TestCount(t *testing.T) {

	t.Run("should count elements inserted", func(t *testing.T) {
		cfg, tearDownTestEnvironment := setupTestEnvironment()
		defer tearDownTestEnvironment()

		dataSourcesCfg := cfg.DataSources
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
		cfg, tearDownTestEnvironment := setupTestEnvironment()
		defer tearDownTestEnvironment()

		dataSourcesCfg := cfg.DataSources
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
				First:         false,
				Last:          true,
			},
			Items: []InstitutionData{ // will return institutions sorted by name
				institutions[0], institutions[2],
			},
		}

		assert.Equal(t, actualPage, expectedPage)
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
