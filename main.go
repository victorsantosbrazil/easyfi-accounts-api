package main

import (
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app"
)

func main() {

	a, err := app.NewApp()

	if err != nil {
		log.Fatal(err)
	}

	a.Start()

}
