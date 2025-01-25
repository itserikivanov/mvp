package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

type handler struct{}

type Response struct {
	Data interface{}
	Errors []string
}

// @title Swagger MVP API
// @version 1.0
// @description Go server for MVP app
// @BasePath /api
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
