package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/victorsantosbrazil/financial-institutions-api/docs/swagger"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/log"
)

// @title Financial Institutions API
// @version 0.1.0
// @description API for managing financial institutions
func main() {
	logger := log.NewLogger()

	a, err := app.NewApp()
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = a.Start()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
