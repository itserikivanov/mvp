package main

import (
	"os"

	_ "github.com/itserikivanov/mvp/server/src/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type handler struct{}

type Response struct {
	Data interface{}
	Errors []string
}

func startServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handler{}
	
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/health", h.healthCheck)
	e.POST("/api/recipe", h.generateRecipe)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
