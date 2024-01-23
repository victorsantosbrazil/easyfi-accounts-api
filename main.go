package main

import (
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/victorsantosbrazil/financial-institutions-api/docs/swagger"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app"
)

// @title Financial Institutions API
// @version 0.1.0
// @description API for managing financial institutions
func main() {

	a, err := app.NewApp()

	if err != nil {
		log.Fatal(err)
	}

	a.Start()

}
