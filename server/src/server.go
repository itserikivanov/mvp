package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func startServer(port string) {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	configureSwagger(e)
	defineRoutes(e)
	
	e.Logger.Fatal(e.Start(":" + port))
}
