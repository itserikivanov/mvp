package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Health check endpoint
//
//  @Summary      Health check
//  @Description  Health check
//  @Tags         health
//  @Accept       json
//  @Produce      json
//  @Success      200	{object}	main.Response{}
//  @Router       /health [get]
func (h *handler) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Data: "OK"})
}
