package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type handler struct{}

type Response struct {
	Status string
}

func startServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handler{}
	e.GET("/health", h.healthCheck)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
