package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"

	"github.com/victorsantosbrazil/financial-institutions-api/src/config"
	"github.com/victorsantosbrazil/financial-institutions-api/src/utils"
)

func main() {
	e := echo.New()

	cfg, err := config.ReadConfig()

	if err != nil {
		e.Logger.Fatal(err)
	}

	migration, err := utils.NewMigration(cfg.Database)
	if err != nil {
		e.Logger.Fatal(err)
	}

	migration.Up()

	e.Logger.Fatal(e.Start(cfg.GetAddress()))
}
