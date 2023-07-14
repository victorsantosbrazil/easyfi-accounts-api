package main

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/config"
)

func main() {
	e := echo.New()

	cfg, err := config.ReadConfig()

	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(cfg.GetAddress()))
}
