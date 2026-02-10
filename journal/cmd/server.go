package main

import (
	"log"
	"journal/internal/config"
	"journal/internal/app"
	"github.com/labstack/echo/v4"
)

func main() {
	config, error := config.Load()
	if error != nil {
		log.Fatal(error)
	}

	e := echo.New()

	app.New(e, config)

	e.Logger.Fatal(e.Start(":8000"))
}