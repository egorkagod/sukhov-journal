package main

import (
	"github.com/labstack/echo/v4"
	"journal/internal/app"
	"journal/internal/config"
	"log"
)

func main() {
	config, error := config.Load()
	if error != nil {
		log.Fatal(error)
	}

	e := echo.New()

	app.New(e, config)

	e.Logger.Fatal(e.Start(":8080"))
}
