package main

import (
	"github.com/labstack/echo/v4"
)

type handler struct{}

type Response struct {
	Data interface{}
	Errors []string
}

func defineRoutes(e *echo.Echo) {
	h := handler{}

	g := e.Group("/api")
	g.GET("/health", h.healthCheck)
	g.POST("/recipe", h.generateRecipe)
}